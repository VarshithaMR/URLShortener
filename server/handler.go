package server

import (
	"URLShortener/service"
	"net/http"
)

func HandleURLShortener(rw http.ResponseWriter, r *http.Request, shortener service.UrlShortenerApi) {
	switch r.Method {
	case http.MethodPost:
		shortener.StartShorteningUrl(r, rw)
	}
}
