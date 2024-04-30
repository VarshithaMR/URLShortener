package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"URLShortener/service/cache"
	"URLShortener/service/models"
)

func ShortenUrl(req models.RequestBody, existingCache cache.StoreURLCache) (models.ResponseBody, error) {
	var (
		domain, path, newUrl       string
		generatedKey, generatedVal int
		cacheVal                   cache.CacheItem
	)

	urlPattern := `^(https?://[^/]+(?:/[^/]+)?)$`
	regex := regexp.MustCompile(urlPattern)
	fmt.Println(regex)
	match := regex.MatchString(req.URL)
	if !match {
		return models.ResponseBody{}, errors.New("improper url - regex doesnt match")
	}

	index := strings.Index(req.URL, "/")
	if index > 0 && req.URL[index-1:index+1] != "//" {
		domain = req.URL[:index]
		path = req.URL[index:]
		generatedKey = rand.Intn(len(domain))
		generatedVal = rand.Intn(len(path))
		value := make(map[int]cache.EndPoint)
		value[generatedVal] = cache.EndPoint{
			Domain: domain,
			Path:   path,
		}
		cacheVal = cache.CacheItem{
			ShortKey: strconv.Itoa(generatedKey),
			Value:    value,
		}
		existingCache.StoreUrl(cacheVal, generatedVal)
	}

	newUrl = fmt.Sprintf("http://shortenedURL/%d-%d", generatedKey, generatedVal)
	return models.ResponseBody{
		ShortUrl: newUrl,
	}, nil
}
