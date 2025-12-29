package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"testing"
	"time"
)

// 回源服务器选择

type OriginServer struct {
	Address string
	Port    int
	Weight  int
	Backup  bool
}

// weightedRandom 在所有有权重，且backup=false的列表中取出一台服务器
// 将所有的weight相加，取一个随机数，随机数落到哪个区域就选哪一个服务器
func weightedRandom(weightServers []OriginServer) (*OriginServer, error) {
	if len(weightServers) == 0 {
		return nil, fmt.Errorf("no servers len(weightServers) == 0")
	}
	// 根据weight从大到小排序
	slices.SortStableFunc(weightServers, func(e1 OriginServer, e2 OriginServer) int {
		return cmp.Compare(e2.Weight, e1.Weight)
	})
	fmt.Printf("weightServers: %v\n", weightServers)

	totalWeight := 0
	for _, originServer := range weightServers {
		if originServer.Weight > 0 {
			totalWeight += originServer.Weight
		}
	}

	if totalWeight == 0 {
		return nil, fmt.Errorf("no servers totalWeight == 0")
	}
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	rnum := r.Intn(totalWeight)
	fmt.Printf("rnum: %d\n", rnum)

	current := 0
	for _, server := range weightServers {
		if server.Weight <= 0 {
			continue
		}
		current += server.Weight
		if rnum < current {
			return &server, nil
		}
	}

	// 兜底
	return &weightServers[len(weightServers)-1], nil
}

func selectServer(originServers []OriginServer) (*OriginServer, error) {
	if len(originServers) == 0 {
		return nil, fmt.Errorf("no servers len(originServers) == 0")
	}

	// 先返回所有权重且backup=false的服务器
	primaryServers := make([]OriginServer, 0)
	backupServers := make([]OriginServer, 0)
	for _, originServer := range originServers {
		if originServer.Backup {
			backupServers = append(backupServers, originServer)
		} else {
			primaryServers = append(primaryServers, originServer)
		}
	}

	// 现在主服务器里面选择
	if len(primaryServers) > 0 {
		return weightedRandom(primaryServers)
	}
	fmt.Printf("primaryServers is null.\n")
	// 在备用服务器里面选择
	if len(backupServers) > 0 {
		return weightedRandom(backupServers)
	}

	return nil, fmt.Errorf("no server found")
}

// 测试回源服务器加权随机访问
func Test111(t *testing.T) {
	// 这个对应的lua：src/test15_cdn/cdn05_openresty/resty03_gateway/lua/middleware/origin/origin_server_load_balance_rand.lua
	servers := []OriginServer{
		{"192.168.10.2", 80, 3, true},
		{"192.168.10.3", 80, 2, true},
		{"192.168.10.1", 80, 5, true},
		{"192.168.10.5", 80, 10, true},
	}

	s, err := selectServer(servers)
	if err != nil {
		panic(err)
	}

	println(s.Address, s.Port)
}

// weightedPoll 在所有有权重，且backup=false的列表中取出一台服务器
// 将所有的weight相加，轮询选择服务器

// map[ip]当前权重
// var originServerPollMap map[string]int
var originServerPollMap map[string]int = make(map[string]int)

func smoothWeightedRoundRobin(weightServers []OriginServer) (*OriginServer, error) {
	if len(weightServers) == 0 {
		return nil, fmt.Errorf("no servers")
	}

	totalWeight := 0

	// 计算 totalWeight，并初始化 currentWeight
	for _, s := range weightServers {
		if s.Weight > 0 {
			totalWeight += s.Weight
			if _, ok := originServerPollMap[s.Address]; !ok {
				originServerPollMap[s.Address] = 0
			}
		}
	}

	if totalWeight == 0 {
		return nil, fmt.Errorf("total weight is zero")
	}

	var selected *OriginServer

	// 核心算法
	for i := range weightServers {
		s := &weightServers[i]

		if s.Weight <= 0 {
			continue
		}

		originServerPollMap[s.Address] += s.Weight

		if selected == nil ||
			originServerPollMap[s.Address] > originServerPollMap[selected.Address] {
			selected = s
		}
	}

	if selected == nil {
		return nil, fmt.Errorf("no server selected")
	}

	// 关键步骤：平滑
	originServerPollMap[selected.Address] -= totalWeight

	return selected, nil
}

func selectServerPoll(originServers []OriginServer) (*OriginServer, error) {
	if len(originServers) == 0 {
		return nil, fmt.Errorf("no servers len(originServers) == 0")
	}

	// 先返回所有权重且backup=false的服务器
	primaryServers := make([]OriginServer, 0)
	backupServers := make([]OriginServer, 0)
	for _, originServer := range originServers {
		if originServer.Backup {
			backupServers = append(backupServers, originServer)
		} else {
			primaryServers = append(primaryServers, originServer)
		}
	}

	// 现在主服务器里面选择
	if len(primaryServers) > 0 {
		return smoothWeightedRoundRobin(primaryServers)
	}
	fmt.Printf("primaryServers is null.\n")
	// 在备用服务器里面选择
	if len(backupServers) > 0 {
		return smoothWeightedRoundRobin(backupServers)
	}

	return nil, fmt.Errorf("no server found")
}

// 测试回源服务器加权轮询访问
func Test222(t *testing.T) {
	// 这个对应的lua：src/test15_cdn/cdn05_openresty/resty03_gateway/lua/middleware/origin/origin_server_load_balance_poll.lua

	servers := []OriginServer{
		{"192.168.10.2", 80, 3, false},
		{"192.168.10.3", 80, 2, false},
		{"192.168.10.1", 80, 5, false},
		{"192.168.10.5", 80, 10, true},
	}
	fmt.Println(servers)
	m := make(map[string]int)
	for i := 0; i < 50; i++ {
		s, _ := selectServerPoll(servers)
		fmt.Printf("选择的ip：%s \n", s.Address)
		if _, ok := m[s.Address]; !ok {
			m[s.Address] = 1
		} else {
			m[s.Address] += 1
		}
		fmt.Println("==================")
	}
	fmt.Println("==================")
	fmt.Println("==================")
	for k, v := range m {
		fmt.Printf("<UNK>%s<UNK>%d\n", k, v)
	}
}
