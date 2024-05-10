package server

import (
	"net/http"

	"URLShortener/service"
)

func HandleURLShortener(rw http.ResponseWriter, r *http.Request, shortener service.UrlShortenerApi) {
	switch r.Method {
	case http.MethodPost:
		shortener.StartShorteningUrl(r, rw)
	}
}

func HandleRedirector(rw http.ResponseWriter, r *http.Request, redirect service.UrlShortenerApi) {
	switch r.Method {
	case http.MethodGet:
		redirect.StartRedirectingUrl(r, rw)
	}
}

func HandleMetrics(rw http.ResponseWriter, r *http.Request, metrics service.UrlShortenerApi) {
	switch r.Method {
	case http.MethodGet:
		metrics.StartMetrics(rw)
	}
}
