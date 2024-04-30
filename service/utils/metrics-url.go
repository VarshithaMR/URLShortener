package utils

import (
	"URLShortener/service/cache"
	"URLShortener/service/models/metrics"
)

func Metrics(existingCache cache.StoreURLCache) ([]metrics.ResponseBody, error) {
	metric, err := existingCache.GetMetrics()
	if err != nil {
		return []metrics.ResponseBody{}, err
	}
	return metric, nil
}
