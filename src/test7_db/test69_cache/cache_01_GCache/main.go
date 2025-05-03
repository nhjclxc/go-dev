package main

import (
	"fmt"
	"github.com/bluele/gcache"
	"time"
)

// go get github.com/bluele/gcache
func main() {
	cache := gcache.New(100). // 最大缓存项数
					LRU().                        // 缓存策略
					Expiration(10 * time.Second). // 过期时间
					Build()

	_ = cache.Set("key", "value")

	val, err := cache.Get("key")
	if err == nil {
		fmt.Println(val) // 输出: value
	}
}
