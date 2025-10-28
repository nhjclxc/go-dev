package algo01_memory_elimination_algo

import (
	"container/list"
	"fmt"
	"testing"
	"time"
)

/*
1. LRU（Least Recently Used，最近最少使用）
核心思想：最近最少使用的数据最可能不再被访问，所以优先淘汰。
优点：符合“热点数据”原则。
缺点：需要维护访问顺序，开销略大。
常用场景：Redis、操作系统页面缓存。
*/

// 最近最久未使用

// LruCache Lur缓存对象
type LruCache struct {
	// 缓存大小
	capacity int
	// 使用hash来存储，便于取缓存数据
	cache map[string]*list.Element
	// 使用list双向链表来存储数据
	queue *list.List
}

// 定义每一个数据节点
type entry struct {
	key         string
	lastUseTime time.Time
	value       any
}

// NewLruCache 创建缓存对象
func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

// Put 向缓存中存入数据或修改数据
func (lru *LruCache) Put(key string, value any) {
	if ele, ok := lru.cache[key]; ok {
		// 修改
		e, ok2 := ele.Value.(*entry)
		if ok2 {
			e.lastUseTime = time.Now()
			e.value = value
		}
		lru.queue.MoveToFront(ele)
		return
	}

	// 新增
	e := entry{value: value, key: key, lastUseTime: time.Now()}
	ele := lru.queue.PushFront(&e)
	lru.cache[key] = ele

	// 判断新节点进来之后，是否超出容量，超出容量则删除队尾的一个元素
	if lru.queue.Len() > lru.capacity {
		tailElement := lru.queue.Back()
		if tailElement != nil {
			lru.queue.Remove(tailElement)
			delete(lru.cache, tailElement.Value.(*entry).key)
		}
	}
}

// Get 获取缓存数据
func (lru *LruCache) Get(key string) (any, error) {
	if ele, ok := lru.cache[key]; ok {
		// 使用了的话将当前这个元素移到队头
		lru.queue.MoveToFront(ele)
		e, ok2 := ele.Value.(*entry)
		if ok2 {
			e.lastUseTime = time.Now()
			return e.value, nil
		}
		return nil, fmt.Errorf("元素类型不匹配")
	}
	return nil, fmt.Errorf("key=[%s]所对应的元素不存在", key)
}

func (lru *LruCache) Print() {
	fmt.Printf("LRU: %v \n", time.Now())
	for ele := lru.queue.Front(); ele != nil; ele = ele.Next() {
		if ele.Value == nil {
			fmt.Println("ele.Value nil")
			continue
		}

		e, ok := ele.Value.(*entry)
		if !ok || e == nil {
			fmt.Println(" !ok || e == nil")
			continue
		}

		fmt.Printf("\t %s:%v,%v \n", e.key, e.value, e.lastUseTime)
	}
	fmt.Println()
}

func TestLruCache(t *testing.T) {

	cache := NewLruCache(3)
	cache.Put("a", 1)
	time.Sleep(1 * time.Second)
	cache.Put("b", 2)
	time.Sleep(1 * time.Second)
	cache.Put("c", 3)
	time.Sleep(1 * time.Second)
	cache.Print() // c b a

	cache.Get("a")
	cache.Print() // a c b

	cache.Put("d", 4) // 超过容量，删除尾部 b
	cache.Print()     // d a c
}
