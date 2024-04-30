package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type StoreURLCache interface {
	StoreUrl(string, string)
}
type storeCache struct {
	newCache *cache.Cache
}

func (s *storeCache) StoreUrl(key, originalUrl string) {
	s.newCache.Set(key, originalUrl, -1)
}

func NewCache() StoreURLCache {
	return &storeCache{
		newCache: cache.New(0, 360*time.Minute),
	}
}
