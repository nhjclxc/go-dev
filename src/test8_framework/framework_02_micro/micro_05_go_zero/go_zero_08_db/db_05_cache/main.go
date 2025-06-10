package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
)


// https://go-zero.dev/docs/tasks/memory-cache

func main() {

	/*
	函数签名:
	    NewCache func(expire time.Duration, opts ...CacheOption) (*Cache, error)
	说明:
	    创建 cache 对象。
	入参:
	    1. expire: 过期时间
	    2. opts: 操作选项
			2.1 WithLimit: 设置 cache 存储数据数量上限
			2.2 WithName: 设置 cache 名称，输出日志时会打印
	返回值:
	    1. *Cache: cache对象
	    2. error: 创建结果
	 */

	newCache, err := collection.NewCache(time.Second*30,
		collection.WithName("any"), collection.WithLimit(50),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	newCache.Set("mykey", 666)

	get, b := newCache.Get("mykey")
	if b == false {
		fmt.Println("不存在key=mykey的高速缓存")
	} else {
		fmt.Println("key = mykey的高速缓存 = ", get)
	}

	// 函数签名:
	//    SetWithExpire func(key string, value interface{}, expire time.Duration)
	//说明:
	//    添加值到缓存, 同时指定过期时间
	newCache.SetWithExpire("mykey111", 123, time.Second * 3)

	get111, b111 := newCache.Get("mykey111")
	if b111 == false {
		fmt.Println("不存在key=mykey111 的高速缓存")
	} else {
		fmt.Println("key = mykey111 的高速缓存 = ", get111)
	}

	time.Sleep(time.Second * 5)

	get1112, b1112 := newCache.Get("mykey1112")
	if b1112 == false {
		fmt.Println("不存在key=mykey1112 的高速缓存")
	} else {
		fmt.Println("key = mykey1112 的高速缓存 = ", get1112)
	}

	// 函数签名:
	//    Take func(key string, fetch func() (interface{}, error)) (interface{}, error)
	//说明:
	//    获取缓存，如果缓存中存在，则返回缓存中的值，如果缓存不存在，则执行 fetch 函数的返回结果。
	//入参:
	//    1. key: key
	//    2. fetch: 自定义返回结果
	take, err := newCache.Take("mykey111", func() (any, error) {

		fmt.Println("缓存不存在执行fetch函数")

		// 模拟一些业务
		newCache.SetWithExpire("mykey111", 999, time.Second*3)

		return 999, nil
	})
	if err != nil {
		return
	}
	fmt.Println("take = ", take)




}
