package cache

import (
	"fmt"
	"github.com/allegro/bigcache/v2"
	"time"
)

var globalCache *bigcache.BigCache

func InitCache(config bigcache.Config) {
	var err error
	if globalCache, err = bigcache.NewBigCache(config); err != nil {
		fmt.Printf("init cache failed: %s\n", err)
		return
	}
}

func InitCacheWithEviction(eviction time.Duration) {
	InitCache(bigcache.DefaultConfig(eviction))
}

func InitCacheWithDefault() {
	InitCacheWithEviction(30 * time.Second)
}

func Set(key string, value []byte) error {
	if globalCache == nil {
		InitCacheWithDefault()
	}
	return globalCache.Set(key, value)
}

func Get(key string) ([]byte, error) {
	if globalCache == nil {
		InitCacheWithDefault()
	}
	return globalCache.Get(key)
}
