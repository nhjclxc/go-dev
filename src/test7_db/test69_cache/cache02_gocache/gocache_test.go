package cache02_gocache

import (
	"fmt"
	"testing"
	"time"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
)

/*
2. **[eko/gocache](https://github.com/eko/gocache)**

  - 封装性更强，支持多种存储后端（内存、Redis、Memcached…）
  - 支持多层缓存（multi-level caching）
  - 在团队项目中更灵活（本地 + 分布式都能接入）

go get github.com/eko/gocache/lib/v4

https://github.com/eko/gocache?tab=readme-ov-file#installation

*/

func Test1(t *testing.T) {
	cache.new
	// 使用内存存储
	memoryStore := store.NewMemoryStore()
	cacheManager := cache.New[string](memoryStore)

	// 设置缓存
	_ = cacheManager.Set("foo", "bar", store.WithExpiration(5*time.Minute))

	// 获取缓存
	val, _ := cacheManager.Get("foo")
	fmt.Println("gocache:", val)
}
