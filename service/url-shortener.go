package service

import (
	"encoding/json"
	"fmt"
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
	StartMetrics(http.ResponseWriter)
}

type URL struct {
	httpClient *resty.Client
	storeCache cache.StoreURLCache
}

func (s *URL) StartShorteningUrl(request *http.Request, response http.ResponseWriter) {
	req, err := utils.GetRequestBodyShorten(request.Body)
	if err != nil {
		WriteResponse(response, "Request body improper", 400)
		return
	}

	fmt.Printf("Request - Full URL - %s\n", req.URL)
	res, err := utils.ShortenUrl(req, s.storeCache, request.Host)
	if err != nil {
		WriteResponse(response, err.Error(), 400)
	}

	fmt.Printf("Response - Shortened URL - %s\n", res.ShortUrl)
	WriteResponse(response, res, http.StatusOK)
}

func (s *URL) StartRedirectingUrl(request *http.Request, response http.ResponseWriter) {
	shortUrl := fmt.Sprintf("http://%s%s", request.Host, request.RequestURI)
	res, err := utils.RedirectUrl(shortUrl, s.storeCache)
	if err != nil {
		WriteResponse(response, err.Error(), 400)
	}

	fmt.Printf("Response - Full URL - %s\n", res.Url)
	WriteResponse(response, res, http.StatusOK)
}

func (s *URL) StartMetrics(response http.ResponseWriter) {
	res, err := utils.Metrics(s.storeCache)
	if err != nil {
		WriteResponse(response, err.Error(), 400)
	}
	WriteResponse(response, res, http.StatusOK)
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
