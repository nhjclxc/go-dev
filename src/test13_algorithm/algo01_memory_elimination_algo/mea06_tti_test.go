package algo01_memory_elimination_algo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
2. TTI（Time To Idle，空闲时间淘汰）
  - **核心思想**：数据在一定时间内未被访问，则淘汰。
  - **优点**：热点数据如果持续访问不会被删除。
  - **缺点**：需要记录最后访问时间。
  - **常用场景**：缓存中间件，如 Ehcache。
*/

type ttiEntry struct {
	key             string
	value           any
	visitExpiration time.Duration
	lastVisitTime   time.Time
}

type TtiCache struct {
	cache map[string]*ttiEntry
	mu    sync.Mutex
}

func NewTtiCache() *TtiCache {
	t := &TtiCache{
		cache: make(map[string]*ttiEntry),
	}

	go t.clearUp()
	return t
}

func (tti *TtiCache) clearUp() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 每隔一秒进行内存淘汰
	for range ticker.C {
		tti.mu.Lock()
		for key, ttiE := range tti.cache {
			if time.Since(ttiE.lastVisitTime) > ttiE.visitExpiration {
				delete(tti.cache, key)
			}
		}
		tti.mu.Unlock()
	}
}

func (tti *TtiCache) Put(key string, value any, expiration time.Duration) {
	tti.mu.Lock()
	defer tti.mu.Unlock()

	if ttiE, ok := tti.cache[key]; ok {
		ttiE.value = value
		ttiE.lastVisitTime = time.Now()
		ttiE.visitExpiration = expiration
	} else {
		tti.cache[key] = &ttiEntry{
			key:             key,
			value:           value,
			lastVisitTime:   time.Now(),
			visitExpiration: expiration,
		}
	}
}

func (tti *TtiCache) Get(key string) (any, error) {
	tti.mu.Lock()
	defer tti.mu.Unlock()
	if ttiE, ok := tti.cache[key]; ok {
		if time.Since(ttiE.lastVisitTime) < ttiE.visitExpiration {
			ttiE.lastVisitTime = time.Now() // 重置访问时间
			return ttiE.value, nil
		}
		// 已过期，删除
		delete(tti.cache, key)
	}
	return nil, fmt.Errorf("key=[%s]对应的元素不存在！", key)
}

func (tti *TtiCache) Print() {
	fmt.Printf("TTI:  \n")
	for k, v := range tti.cache {
		fmt.Printf("\t %s:%v, %v , %v \n", k, v.value, v.visitExpiration, v.lastVisitTime)
	}
	fmt.Println()
}

func TestTTI(t *testing.T) {

	cache := NewTtiCache()

	// 设置每个键的空闲时间为 3 秒
	cache.Put("a", "apple", 3*time.Second)
	cache.Put("b", "banana", 3*time.Second)

	for i := 1; i <= 7; i++ {
		time.Sleep(1 * time.Second)

		// 每 2 秒访问一次 a（应该一直不会过期）
		if i%2 == 0 {
			cache.Get("a")
			fmt.Printf("[%v秒] 访问 a，刷新空闲时间\n", i)
		}

		cache.Print()
	}
	/*
		TTI:
			 b:banana, 3s , 2025-10-23 12:18:25.057797 +0800 CST m=+0.000903459
			 a:apple, 3s , 2025-10-23 12:18:25.057797 +0800 CST m=+0.000903001

		[2秒] 访问 a，刷新空闲时间
		TTI:
			 a:apple, 3s , 2025-10-23 12:18:27.059706 +0800 CST m=+2.002819501
			 b:banana, 3s , 2025-10-23 12:18:25.057797 +0800 CST m=+0.000903459

		TTI:
			 a:apple, 3s , 2025-10-23 12:18:27.059706 +0800 CST m=+2.002819501

		[4秒] 访问 a，刷新空闲时间
		TTI:
			 a:apple, 3s , 2025-10-23 12:18:29.060317 +0800 CST m=+4.003436834

		TTI:
			 a:apple, 3s , 2025-10-23 12:18:29.060317 +0800 CST m=+4.003436834

		[6秒] 访问 a，刷新空闲时间
		TTI:
			 a:apple, 3s , 2025-10-23 12:18:31.060965 +0800 CST m=+6.004090917

		TTI:
			 a:apple, 3s , 2025-10-23 12:18:31.060965 +0800 CST m=+6.004090917

	*/
}
