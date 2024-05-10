package utils

import (
	"encoding/json"
	"io"

	"github.com/opentracing/opentracing-go/log"

	"URLShortener/service/models/shorten"
)

func GetRequestBodyShorten(body io.ReadCloser) (requestBody shorten.RequestBody, err error) {
	decoder := json.NewDecoder(body)
	if err = decoder.Decode(&requestBody); err != nil {
		log.Error(err)
		return
	}

	return
}
