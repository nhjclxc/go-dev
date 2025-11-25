package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"testing"
	"time"
)

// 使用 https://github.com/hashicorp/golang-lru 来实现lru缓存
// 这里先测试一波https://github.com/hashicorp/golang-lru的功能
// key = qname + qtype

// 使用新版本的包
// go get github.com/hashicorp/golang-lru/v2

type CacheItem struct {
	Msg      string
	ExpireAt time.Time
}

func Test11(t *testing.T) {

	// size缓存大小，可以缓存多少个item
	// ttl表示每隔多久执行一次内存淘汰，0表示不执行内存淘汰，由用户自己控制。time.Second*5表示5秒执行一次内存淘汰,
	cache := expirable.NewLRU[string, *CacheItem](128, nil, time.Second*4)

	now := time.Now()
	cache.Add("a", &CacheItem{
		Msg:      "a-val",
		ExpireAt: now.Add(2 * time.Second),
	})
	cache.Add("b", &CacheItem{
		Msg:      "b-val",
		ExpireAt: now.Add(3 * time.Second),
	})
	cache.Add("c", &CacheItem{
		Msg:      "c-val",
		ExpireAt: now.Add(51 * time.Second),
	})
	av, ok := cache.Get("a")
	fmt.Println(av, ok)
	bv, ok := cache.Get("b")
	fmt.Println(bv, ok)
	cv, ok := cache.Get("c")
	fmt.Println(cv, ok)

	time.Sleep(3 * time.Second)

	fmt.Println("----------------------------------------")
	av2, ok := cache.Get("a")
	fmt.Println(av2, ok)
	bv2, ok := cache.Get("b")
	fmt.Println(bv2, ok)
	cv2, ok := cache.Get("c")
	fmt.Println(cv2, ok)

}

var cache *expirable.LRU[string, *CacheItem]

func Test22(t *testing.T) {

	// size缓存大小，可以缓存多少个item
	// ttl设置为0的时候，自己在定义一个getCache方法用于业务内存淘汰的判断
	cache = expirable.NewLRU[string, *CacheItem](128, nil, 0)

	now := time.Now()
	cache.Add("c", &CacheItem{
		Msg:      "c-val",
		ExpireAt: now.Add(3 * time.Second),
	})
	cv, ok := cache.Get("c")
	fmt.Println(cv, ok)
	cv12, ok := GetCache("c")
	fmt.Println(cv12, ok)

	time.Sleep(5 * time.Second)

	fmt.Println("----------------------------------------")
	cv2, ok := cache.Get("c")
	fmt.Println(cv2, ok)

	cv22, ok := GetCache("c")
	fmt.Println(cv22, ok)
}

// Get 时自己判断
func GetCache(key string) (*CacheItem, bool) {
	item, ok := cache.Get(key)
	if !ok || item == nil {
		return nil, false
	}
	if time.Now().After(item.ExpireAt) {
		cache.Remove(key)
		return nil, false
	}
	return item, true
}

func Test33(t *testing.T) {

	cache := expirable.NewLRU[string, string](128, nil, 3*time.Second)

	cache.Add("a", "a-val")
	cache.Add("c", "c-val")
	av, ok := cache.Get("a")
	fmt.Println(av, ok)
	cv, ok := cache.Get("c")
	fmt.Println(cv, ok)

	time.Sleep(2 * time.Second)

	cache.Add("dd", "dd-val")

	time.Sleep(2 * time.Second)

	av2, ok := cache.Get("a")
	fmt.Println(av2, ok)
	cv2, ok := cache.Get("c")
	fmt.Println(cv2, ok)
	dd, ok := cache.Get("dd")
	fmt.Println(dd, ok)

}
