package algo01_memory_elimination_algo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
### 1. TTL（Time To Live，存活时间）
  - **核心思想**：缓存每条数据带一个过期时间，到期自动淘汰。
  - **优点**：实现简单，防止缓存无限增长。
  - **缺点**：不考虑访问频率或时间，可能提前删除热点数据。
  - **常用场景**：HTTP 缓存、Redis TTL。
*/

// ttlEntry 表示缓存中的一个元素
type ttlEntry struct {
	key        string
	value      any
	expiration int64
}

// TtlCache TTL 缓存
type TtlCache struct {
	cache map[string]*ttlEntry
	mu    sync.Mutex
}

func NewTtlCache() *TtlCache {
	ttlCache := &TtlCache{
		cache: make(map[string]*ttlEntry),
	}
	go ttlCache.run()
	return ttlCache
}

// run 每一秒执行一次，看看哪些缓存过期了
func (ttl *TtlCache) run() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 每隔一秒进行内存淘汰
	for range ticker.C {
		now := time.Now().Unix()
		ttl.mu.Lock()
		for k, v := range ttl.cache {
			if now > v.expiration {
				delete(ttl.cache, k)
			}
		}
		ttl.mu.Unlock()
	}
}

func (ttl *TtlCache) Put(key string, value any, expiration int64) {
	ttl.mu.Lock()
	defer ttl.mu.Unlock()

	if ttlE, ok := ttl.cache[key]; ok {
		ttlE.value = value
		ttlE.expiration = expiration + time.Now().Unix()
	} else {
		ttl.cache[key] = &ttlEntry{
			key:        key,
			value:      value,
			expiration: expiration + time.Now().Unix(),
		}
	}

}

func (ttl *TtlCache) Get(key string) (any, error) {
	ttl.mu.Lock()
	defer ttl.mu.Unlock()
	if ttlE, ok := ttl.cache[key]; ok {
		if time.Now().Unix() < ttlE.expiration {
			return ttlE.value, nil
		}
		// 已过期，删除
		delete(ttl.cache, key)
	}
	return nil, fmt.Errorf("key=[%s]对应的元素不存在！", key)
}

func (ttl *TtlCache) Print() {
	fmt.Printf("TTL:  \n")
	for k, v := range ttl.cache {
		fmt.Printf("\t %s:%v, %v \n", k, v.value, v.expiration)
	}
	fmt.Println()
}

func TestTTL(t *testing.T) {

	cache := NewTtlCache()
	cache.Put("a", 111, 3)
	cache.Put("b", 222, 5)

	for i := 0; i < 7; i++ {
		time.Sleep(1 * time.Second)
		v, ok := cache.Get("b")
		fmt.Printf("%v 秒后：a=%v, ok=%v\n", i+1, v, ok)
	}
}
