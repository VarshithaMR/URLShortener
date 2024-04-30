package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type StoreURLCache interface {
	StoreUrl(CacheItem, string, string)
	GetFullUrl(string) (string, error)
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

type RedirectCacheItem struct {
	newUrl string
	pair   Pair
}

type Pair struct {
	shortKey string
	shortVal string
}

func (s *storeCache) StoreUrl(cacheItem CacheItem, shortVal string, newUrl string) {
	var redirectCacheVal RedirectCacheItem
	value, found := s.newCache.Get(cacheItem.ShortKey)
	if !found {
		fmt.Printf("Key-Value %s-%s stored in cache\n", cacheItem.ShortKey, shortVal)
		s.newCache.Set(cacheItem.ShortKey, cacheItem.Value, -1)

		fmt.Printf("NewUrl-shortKey-shortVal, %s-%s-%s stored in cache\n", newUrl, cacheItem.ShortKey, shortVal)
		redirectCacheVal = RedirectCacheItem{
			newUrl: newUrl,
			pair: Pair{
				shortKey: cacheItem.ShortKey,
				shortVal: shortVal,
			},
		}
		s.newCache.Set(newUrl, redirectCacheVal, -1)
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

	fmt.Printf("NewUrl-shortKey-shortVal, %s-%s-%s stored in cache\n", newUrl, cacheItem.ShortKey, shortVal)
	redirectCacheVal = RedirectCacheItem{
		newUrl: newUrl,
		pair: Pair{
			shortKey: cacheItem.ShortKey,
			shortVal: shortVal,
		},
	}
	s.newCache.Set(newUrl, redirectCacheVal, -1)
}

func (s *storeCache) GetFullUrl(url string) (string, error) {
	value, found := s.newCache.Get(url)
	if !found {
		fmt.Printf("This url doesnt exist in cache. Please request existing ones")
		return "", errors.New("short url doesnt exist")
	}

	pair := value.(Pair)
	value, found = s.newCache.Get(pair.shortKey)
	if !found {
		return "", errors.New("short key doesnt exist")
	}

	cacheData := value.(map[string]EndPoint)
	urlPart := cacheData[pair.shortVal]
	return urlPart.Domain + urlPart.Path, nil

}

func NewCache() StoreURLCache {
	return &storeCache{
		newCache: cache.New(0, 360*time.Minute),
	}
}
