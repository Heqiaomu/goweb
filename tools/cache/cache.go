package cache

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/songzhibin97/gkit/options"
	"sync"
)

type Cache struct {
	*local_cache.Cache
}

var cache *Cache
var cacheOnec sync.Once

func NewCache(op options.Option) {
	cacheOnec.Do(func() {
		ch := local_cache.NewCache(op)
		cache = &Cache{
			&ch,
		}
	})
}

func GetLocalCache() *Cache {
	return cache
}
