package algo01_memory_elimination_algo

import (
	"fmt"
	"testing"
	"time"
)

/*
2. LFU（Least Frequently Used，最不常用）
	核心思想：使用次数最少的数据被淘汰。
	优点：能够保留高访问频率的热点数据。
	缺点：频率统计可能导致“老旧热点”被误保留。
	常用场景：数据库缓存、对象缓存。

LFU 的核心思想是：
	淘汰使用频率（访问次数）最少的缓存项。
当缓存满了：
	删除访问次数最少的元素；
	如果多个元素的访问次数相同，则删除最久未使用的那个（通常用时间戳辅助）。
*/

// 淘汰最少使用的缓存数据，使用次数相同时淘汰时间戳大的

type lfuEntry struct {
	key       string
	value     any
	counter   int
	timestamp int64
}

type LfuCache struct {
	capacity int
	cache    map[string]*lfuEntry
}

func NewLfuCache(capacity int) *LfuCache {
	return &LfuCache{
		capacity: capacity,
		cache:    make(map[string]*lfuEntry),
	}
}

func (lfu *LfuCache) Put(key string, value any) {
	if entry, ok := lfu.cache[key]; ok {
		entry.value = value
		entry.counter += 1
		entry.timestamp = time.Now().Unix()
		return
	}

	// 放入新数据
	lfu.cache[key] = &lfuEntry{
		key:       key,
		value:     value,
		counter:   1,
		timestamp: time.Now().Unix(),
	}

	// 检查是否要淘汰内存
	if len(lfu.cache) > lfu.capacity {
		//minKey := ""
		//minCounter := 0
		//var minTimestamp int64 = 0
		//for k, v := range lfu.cache {
		//	//fmt.Printf("\t 内存淘汰 %s:%v,%v,%v,%v,%v,%v \n", k, v.value, v.counter, v.timestamp, minKey, minCounter, minTimestamp)
		//	// 淘汰最少使用的缓存数据，使用次数相同时淘汰时间戳大的
		//	if v.counter < minCounter || minCounter == 0 {
		//		minKey = k
		//		minCounter = v.counter
		//		continue
		//	}
		//	if v.counter == minCounter && v.timestamp < minTimestamp {
		//		minKey = k
		//		minCounter = v.counter
		//		continue
		//	}
		//}
		minKey := ""
		minCounter := int(^uint(0) >> 1)       // 最大 int
		minTimestamp := int64(^uint64(0) >> 1) // 最大 int64
		for k, v := range lfu.cache {
			// 淘汰最少使用的缓存数据，使用次数相同时淘汰时间戳大的
			if v.counter < minCounter || (v.counter == minCounter && v.timestamp < minTimestamp) {
				minKey = k
				minCounter = v.counter
				minTimestamp = v.timestamp
			}
		}
		fmt.Printf("minKey = [%s] \n", minKey)
		if minKey != "" {
			delete(lfu.cache, minKey)
		}
	}
}

func (lfu *LfuCache) Get(key string) (any, error) {
	if entry, ok := lfu.cache[key]; ok {
		entry.counter += 1
		entry.timestamp = time.Now().Unix()
		return entry, nil
	}
	return nil, fmt.Errorf("key=[%s]对应的元素不存在！", key)
}

func (lfu *LfuCache) Print() {
	fmt.Printf("LFU: %v \n", time.Now())
	for k, v := range lfu.cache {
		fmt.Printf("\t %s:%v,%v,%v \n", k, v.value, v.counter, v.timestamp)
	}
	fmt.Println()
}

func TestLfuCache(t *testing.T) {

	cache := NewLfuCache(3)

	cache.Put("a", 111)
	time.Sleep(1 * time.Second)
	cache.Put("b", 222)
	time.Sleep(1 * time.Second)
	cache.Put("c", 333)
	time.Sleep(1 * time.Second)
	cache.Print()

	cache.Get("a")
	time.Sleep(1 * time.Second)
	cache.Get("a")
	time.Sleep(1 * time.Second)
	cache.Get("b")
	time.Sleep(1 * time.Second)
	cache.Print()

	cache.Put("d", 555) // 应该淘汰 c（freq最低）
	cache.Print()
}
