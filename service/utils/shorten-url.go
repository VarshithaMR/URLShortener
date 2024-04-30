package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"

	"URLShortener/service/cache"
	"URLShortener/service/models/shorten"
)

func ShortenUrl(req shorten.RequestBody, existingCache cache.StoreURLCache) (shorten.ResponseBody, error) {
	var (
		domain, path, newUrl       string
		generatedKey, generatedVal [32]byte
		cacheVal                   cache.CacheItem
	)

	parsedUrl, err := url.Parse(req.URL)
	if err != nil {
		return shorten.ResponseBody{}, errors.New("cannot parse URL")
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

	newUrl = fmt.Sprintf("http://shortURL/%s%s", hex.EncodeToString(generatedKey[:1]), hex.EncodeToString(generatedVal[:1]))
	existingCache.StoreUrl(cacheVal, hex.EncodeToString(generatedVal[:]), newUrl)

	return shorten.ResponseBody{
		ShortUrl: newUrl,
	}, nil
}
