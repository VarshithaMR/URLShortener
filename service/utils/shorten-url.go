package utils

import (
	"URLShortener/service/cache"
	"fmt"
	"math/rand"
	"time"

	"URLShortener/service/models"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func ShortenUrl(req models.RequestBody, cache cache.StoreURLCache) models.ResponseBody {
	key := generateKeyForURL()
	newUrl := fmt.Sprintf("http://shortenedURL/%s", key)
	cache.StoreUrl(key, req.URL)
	return models.ResponseBody{
		ShortUrl: newUrl,
	}
}

func generateKeyForURL() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
