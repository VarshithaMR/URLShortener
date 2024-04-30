package cache

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type StoreURLCache interface {
	StoreUrl(CacheItem, string)
}
type storeCache struct {
	newCache *cache.Cache
}

type CacheItem struct {
	ShortKey string
	Value    map[string]EndPoint
}

type EndPoint struct {
	Domain string
	Path   string
}

func (s *storeCache) StoreUrl(cacheItem CacheItem, shortVal string) {
	value, found := s.newCache.Get(cacheItem.ShortKey)
	if !found {
		fmt.Printf("Key-Value %s-%s stored in cache\n", cacheItem.ShortKey, shortVal)
		s.newCache.Set(cacheItem.ShortKey, cacheItem.Value, -1)
		return
	}
	cacheData := value.(map[string]EndPoint)
	cachePathKey, found := cacheData[shortVal]
	if found {
		fmt.Printf("URL found already %s%s", cachePathKey.Domain, cachePathKey.Path)
		fmt.Printf("URL - key %s, URL - val %s found already\n", cacheItem.ShortKey, shortVal)
		return
	}

	cacheData[shortVal] = EndPoint{
		Domain: cacheItem.Value[shortVal].Domain,
		Path:   cacheItem.Value[shortVal].Path,
	}
	fmt.Printf("Key-Value %s-%s stored in cache", cacheItem.ShortKey, shortVal)
	s.newCache.Set(cacheItem.ShortKey, cacheData, -1)
}

func NewCache() StoreURLCache {
	return &storeCache{
		newCache: cache.New(0, 360*time.Minute),
	}
}
