package cache01_go_cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

/*
1. **[patrickmn/go-cache](https://github.com/patrickmn/go-cache)**

  - 最经典的 Go 内存缓存库
  - 支持过期时间、定期清理
  - 适合中小规模的本地缓存场景
  - 类似 Java 的 `ConcurrentHashMap` + TTL

go get github.com/patrickmn/go-cache

详细文档：https://pkg.go.dev/github.com/patrickmn/go-cache
*/
func Test01(t *testing.T) {
	// Create a cache with a default expiration time of 5 minutes,
	// and which purges expired items every 10 minutes
	// defaultExpiration：key的过期时间，超过过期时间之后无法访问key
	// cleanupInterval：key的清理时间，每隔cleanupInterval这一段时间会把内存里面过期的key清理一次
	// 即：key 过期后立刻就无法访问（Get 返回不到），但真正从内存里删除要等到下一次 cleanupInterval 到来。
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	fooCache, fooOk := c.Get("foo")
	if fooOk {
		fooVal := fooCache.(string)
		fmt.Println("fooVal", fooVal)
	}

	c.Set("foo", "bar123456789", 1*time.Minute) // 替换该key原有的val
	fooCache2, fooOk2 := c.Get("foo")
	if fooOk2 {
		fooVal2 := fooCache2.(string)
		fmt.Println("fooVal2", fooVal2)
	}

	c.SetDefault("foo1", "aaa")
	if x, found := c.Get("foo1"); found {
		foo1 := x.(string)
		// ...
		fmt.Println("foo1", foo1)
	}

}

type CacheStruct struct {
	name    string
	Age     int
	addTime time.Time
}

// 缓存结构体
func Test2(t *testing.T) {
	cs := CacheStruct{
		name:    "这个是缓存测试对象",
		Age:     18,
		addTime: time.Now(),
	}

	c := cache.New(5*time.Second, 10*time.Second)

	c.SetDefault("cs", cs)
	c.SetDefault("cs2", &cs)

	if x, found := c.Get("cs"); found {
		fmt.Println("csVal", x)
	}
	if x, found := c.Get("cs2"); found {
		fmt.Println("cs2Val", x)
	}
	expiration, t2, b := c.GetWithExpiration("cs")
	fmt.Println("expiration ", expiration)
	fmt.Println("t2 ", t2)
	fmt.Println("b ", b)

	time.Sleep(6 * time.Second)
	// 当cache.New(5*time.Second, 10*time.Second)，那么5s后将无法访问到数据了

	x21, found21 := c.Get("cs")
	fmt.Println("x21", x21)
	fmt.Println("found21", found21)

	x22, found22 := c.Get("cs2")
	fmt.Println("x22", x22)
	fmt.Println("found22", found22)

	expiration1, t21, b1 := c.GetWithExpiration("cs")
	fmt.Println("expiration1 ", expiration1)
	fmt.Println("t21 ", t21)
	fmt.Println("b1 ", b1)

}
