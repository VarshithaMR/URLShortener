package utils

import (
	"URLShortener/service/models/redirect"
	"URLShortener/service/models/shorten"
	"encoding/json"
	"io"

	"github.com/opentracing/opentracing-go/log"
)

func GetRequestBodyShorten(body io.ReadCloser) (requestBody shorten.RequestBody, err error) {
	decoder := json.NewDecoder(body)
	if err = decoder.Decode(&requestBody); err != nil {
		log.Error(err)
		return
	}

	return
}

func GetRequestBodyRedirect(body io.ReadCloser) (requestBody redirect.RequestBody, err error) {
	decoder := json.NewDecoder(body)
	if err = decoder.Decode(&requestBody); err != nil {
		log.Error(err)
		return
	}

	return
}
