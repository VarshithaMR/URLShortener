package cache

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/patrickmn/go-cache"

	"URLShortener/service/models/metrics"
)

type StoreURLCache interface {
	StoreUrl(CacheItem, string, string)
	GetFullUrl(string) (string, error)
	GetMetrics() ([]metrics.ResponseBody, error)
}
type storeCache struct {
	shortUrlsCache *cache.Cache
	redirectCache  *cache.Cache
}

type CacheItem struct {
	ShortKey string
	Value    map[string]EndPoint
}

type EndPoint struct {
	Domain string
	Path   string
}

type Pair struct {
	shortKey string
	shortVal string
}

func (s *storeCache) StoreUrl(cacheItem CacheItem, shortVal string, newUrl string) {
	value, found := s.shortUrlsCache.Get(cacheItem.ShortKey)
	if !found {
		s.shortUrlsCache.Set(cacheItem.ShortKey, cacheItem.Value, 10*time.Minute)
		fmt.Printf("Key-Value %s-%s stored in cache\n", cacheItem.ShortKey, shortVal)

		pair := Pair{
			shortKey: cacheItem.ShortKey,
			shortVal: shortVal,
		}
		s.redirectCache.Set(newUrl, pair, 10*time.Minute)
		fmt.Printf("NewUrl-shortKey-shortVal, %s-%s-%s stored in cache\n", newUrl, cacheItem.ShortKey, shortVal)
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

	s.shortUrlsCache.Set(cacheItem.ShortKey, cacheData, 10*time.Minute)
	fmt.Printf("Key-Value %s-%s stored in cache", cacheItem.ShortKey, shortVal)

	pair := Pair{
		shortKey: cacheItem.ShortKey,
		shortVal: shortVal,
	}

	s.redirectCache.Set(newUrl, pair, 10*time.Minute)
	fmt.Printf("NewUrl-shortKey-shortVal, %s-%s-%s stored in cache\n", newUrl, cacheItem.ShortKey, shortVal)
}

func (s *storeCache) GetFullUrl(url string) (string, error) {
	value, found := s.redirectCache.Get(url)
	if !found {
		fmt.Printf("This url doesnt exist in cache. Please request existing ones")
		return "", errors.New("short url doesnt exist")
	}

	pair := value.(Pair)
	value, found = s.redirectCache.Get(pair.shortKey)
	if !found {
		return "", errors.New("short key doesnt exist")
	}

	cacheData := value.(map[string]EndPoint)
	urlPart := cacheData[pair.shortVal]
	return urlPart.Domain + urlPart.Path, nil

}

func (s *storeCache) GetMetrics() (metric []metrics.ResponseBody, err error) {
	cacheItems := s.shortUrlsCache.Items()
	for _, val := range cacheItems {
		cacheItemVal := val.Object.(map[string]EndPoint)
		var keys []string
		for key := range cacheItemVal {
			keys = append(keys, key)
		}

		metric = append(metric, metrics.ResponseBody{
			Name: cacheItemVal[keys[0]].Domain,
			Hits: len(cacheItemVal),
		})
	}
	sort.Slice(metric, func(i, j int) bool {
		return metric[i].Hits > metric[j].Hits
	})
	if len(metric) < 1 {
		return []metrics.ResponseBody{}, errors.New("no urls have been hit to show the metrics")
	}
	return metric[:3], nil
}

func NewCache() StoreURLCache {
	return &storeCache{
		shortUrlsCache: cache.New(0, 360*time.Minute),
		redirectCache:  cache.New(0, 360*time.Minute),
	}
}
