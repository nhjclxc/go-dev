package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type HB struct {
	best   bool
	amount float64
	index  int
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// genHB 随机拆分红包
func genHB(size int, totalAmount float64) ([]*HB, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size 必须大于 0")
	}
	if totalAmount < float64(size)*0.01 {
		return nil, fmt.Errorf("总金额太小，无法保证每人至少 0.01")
	}

	// 将总金额转换为最小单位（分）, 避免小数的数据泄露
	totalCents := int(totalAmount * 100)
	minUnit := 1 // 1 分
	amounts := make([]int, size)

	// 每个人至少 1 分
	for i := 0; i < size; i++ {
		amounts[i] = minUnit
	}
	// 计算剩下的随机分配进而的大小 = 总的金额大小 - 人数*一分钱
	remaining := totalCents - size*minUnit

	// 随机分配剩余的分，不断的给所有人分一分钱一次，直至全部分完
	for remaining > 0 {
		idx := r.Intn(size)
		amounts[idx]++
		remaining--
	}

	// 转换回浮点数（元），并找最大红包
	hbs := make([]*HB, size)
	bestIndex := 0
	bestAmount := 0.0
	for i := 0; i < size; i++ {
		// 避免总数不对的方法1：计算分配金额的时候不计算保留两位小数
		amt := float64(amounts[i]) / 100.0 // 保留精度
		if amt > bestAmount {
			bestAmount = amt
			bestIndex = i
		}

		hbs[i] = &HB{
			amount: amt,
			index:  i,
		}
	}

	hbs[bestIndex].best = true
	return hbs, nil
}

func TestName(t *testing.T) {
	hb, err := genHB(5, 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0.0
	for _, h := range hb {
		fmt.Printf("%+v\n", h)
		total += h.amount
	}
	fmt.Printf("总金额: %.2f\n", total)
}
