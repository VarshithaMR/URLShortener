package service

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"

	"URLShortener/service/cache"
	"URLShortener/service/utils"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json; charset=utf-8"
)

type UrlShortenerApi interface {
	StartShorteningUrl(*http.Request, http.ResponseWriter)
	StartRedirectingUrl(*http.Request, http.ResponseWriter)
}

type URL struct {
	httpClient *resty.Client
	storeCache cache.StoreURLCache
}

func (s *URL) StartShorteningUrl(request *http.Request, response http.ResponseWriter) {
	req, err := utils.GetRequestBody(request.Body)
	if err != nil {
		WriteResponse(response, "Request body improper", 400)
		return
	}

	res := utils.ShortenUrl(req, s.storeCache)

	WriteResponse(response, res, http.StatusOK)
}

func (s *URL) StartRedirectingUrl(request *http.Request, response http.ResponseWriter) {
	req, err := utils.GetRequestBody(request.Body)
	if err != nil {
		WriteResponse(response, "Request body improper", 400)
		return
	}

	res := utils.RedirectUrl(req)
}

func WriteResponse(rw http.ResponseWriter, resp interface{}, responseCode int) {
	rw.WriteHeader(responseCode)
	rw.Header().Set(contentTypeKey, contentTypeValue)
	bytes, err := json.Marshal(resp)
	if err != nil {
		//TODO logging
	}
	rw.Write(bytes)
}

func NewURLShortener() UrlShortenerApi {
	newCache := cache.NewCache()
	return &URL{
		httpClient: resty.New(),
		storeCache: newCache,
	}
}
