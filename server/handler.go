package server

import (
	"URLShortener/service"
	"net/http"
)

func HandleURLShortener(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		service.StartShorteningUrl(rw, r)
	}
}
