package algo01_memory_elimination_algo

import (
	"fmt"
	"testing"
)

/*
### 3. FIFO（First In First Out，先进先出）

* **核心思想**：最先进入缓存的数据先淘汰。
* **优点**：实现简单，不需要复杂的数据结构。
* **缺点**：不考虑访问频率或时间，容易误删热点数据。
* **常用场景**：简单队列缓存。
*/

// fifoEntry 表示缓存中的一个元素
type fifoEntry struct {
	key   string
	value any
}

// FifoCache FIFO 缓存
type FifoCache struct {
	capacity int
	cache    map[string]*fifoEntry
	queue    []string // 存储插入顺序的 key
}

func NewFifoCache(capacity int) *FifoCache {
	return &FifoCache{
		capacity: capacity,
		cache:    make(map[string]*fifoEntry),
		queue:    make([]string, 0),
	}
}

func (fifo *FifoCache) Put(key string, value any) {
	if fifoE, ok := fifo.cache[key]; ok {
		fifoE.value = value
		// 把key在queue里面的顺序调一下
		fifo.queue = append(fifo.queue[1:], key)
		return
	}

	fifoE := fifoEntry{key: key, value: value}
	fifo.cache[key] = &fifoE
	fifo.queue = append(fifo.queue, key)

	// 内存淘汰
	if len(fifo.queue) > fifo.capacity {
		outKey := fifo.queue[0]
		delete(fifo.cache, outKey)
		fifo.queue = fifo.queue[1:]
		fmt.Printf("outKey = [%s] \n", outKey)
	}
}

func (fifo *FifoCache) Get(key string) (any, error) {
	if fifoE, ok := fifo.cache[key]; ok {
		return fifoE.value, nil
	}
	return nil, fmt.Errorf("key=[%s]对应的元素不存在！", key)
}

func (fifo *FifoCache) Print() {
	fmt.Printf("FIFO:  \n")
	for k, v := range fifo.cache {
		fmt.Printf("\t %s:%v \n", k, v.value)
	}
	fmt.Println()
}

func TestFIFO(t *testing.T) {

	cache := NewFifoCache(3)

	cache.Put("a", 111)
	cache.Put("b", 222)
	cache.Put("c", 333)
	cache.Print()

	cache.Put("d", 444) // 淘汰 a
	cache.Print()

	cache.Put("e", 555) // 淘汰 b
	cache.Print()

	// 获取测试
	v, err := cache.Get("c")
	fmt.Println("Get c =", v, "err =", err)
}
