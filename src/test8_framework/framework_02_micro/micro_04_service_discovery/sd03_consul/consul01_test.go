package sd03_consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"time"

	"testing"
)

// consul做key-val存储
func Test11(t *testing.T) {
	//put存数据，注意：consul的kv是不支持过期时间的
	put, err := ConsulClient.KV().Put(&api.KVPair{
		Key:   "key1",
		Value: []byte("consul-value"),
	}, nil)
	if err != nil {
		fmt.Println("kv opt err", err)
		return
	}
	fmt.Println(put.RequestTime, put.Warnings)

}

func Test12(t *testing.T) {
	// get取数据
	key := "key1"
	kvPair, queryMeta, err := ConsulClient.KV().Get(key, nil)
	if err != nil {
		fmt.Println("get opt err", err)
		return
	}
	if kvPair == nil {
		fmt.Printf(" key = %s not sxists", key)
		return
	}
	fmt.Println(kvPair.Key, string(kvPair.Value), queryMeta.LastIndex)

}

func Test13(t *testing.T) {
	// 监听kv的变化
	var lastIndex uint64 = 0

	key := "key1"
	for {
		// 核心是 Blocking Query（阻塞查询）
		pair, meta, err := ConsulClient.KV().Get(
			key,
			&api.QueryOptions{WaitIndex: lastIndex, WaitTime: 3 * time.Minute},
		)
		if err != nil {
			fmt.Printf("listen kv change err: %s \n", err)
			return
		}
		if err != nil {
			fmt.Println("KV watch error:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if meta.LastIndex == lastIndex {
			continue // 没有变化
		}

		lastIndex = meta.LastIndex
		fmt.Printf("KV %s changed: %v\n", key, pair)

	}
}
