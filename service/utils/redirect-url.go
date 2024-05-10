package utils

import (
	"URLShortener/service/cache"
	"URLShortener/service/models/redirect"
)

func RedirectUrl(shortUrl string, existingCache cache.StoreURLCache) (redirect.ResponseBody, error) {
	url, err := existingCache.GetFullUrl(shortUrl)
	if err != nil {
		return redirect.ResponseBody{}, err
	}

	return redirect.ResponseBody{
		Url: url,
	}, nil
}
