package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"

	"URLShortener/service/cache"
	"URLShortener/service/models"
)

func ShortenUrl(req models.RequestBody, existingCache cache.StoreURLCache) (models.ResponseBody, error) {
	var (
		domain, path, newUrl       string
		generatedKey, generatedVal [32]byte
		cacheVal                   cache.CacheItem
	)

	parsedUrl, err := url.Parse(req.URL)
	if err != nil {
		return models.ResponseBody{}, errors.New("cannot parse URL")
	}

	domain = fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
	path = parsedUrl.Path
	generatedKey = sha256.Sum256([]byte(domain))
	generatedVal = sha256.Sum256([]byte(path))
	value := make(map[string]cache.EndPoint)
	value[hex.EncodeToString(generatedVal[:])] = cache.EndPoint{
		Domain: domain,
		Path:   path,
	}
	cacheVal = cache.CacheItem{
		ShortKey: hex.EncodeToString(generatedKey[:]),
		Value:    value,
	}
	existingCache.StoreUrl(cacheVal, hex.EncodeToString(generatedVal[:]))

	newUrl = fmt.Sprintf("http://shortenedURL/%d-%d", generatedKey, generatedVal)
	return models.ResponseBody{
		ShortUrl: newUrl,
	}, nil
}
