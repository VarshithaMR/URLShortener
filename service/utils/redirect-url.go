package utils

import (
	"URLShortener/service/cache"
	"URLShortener/service/models/redirect"
)

func RedirectUrl(req redirect.RequestBody, existingCache cache.StoreURLCache) (redirect.ResponseBody, error) {
	url, err := existingCache.GetFullUrl(req.ShortUrl)
	if err != nil {
		return redirect.ResponseBody{}, err
	}

	return redirect.ResponseBody{
		Url: url,
	}, nil
}
