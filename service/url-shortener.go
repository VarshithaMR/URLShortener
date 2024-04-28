package service

import "net/http"

type UrlShortenerApi interface {
	StartShorteningUrl(http.ResponseWriter, *http.Request)
}

func StartShorteningUrl(response http.ResponseWriter, request *http.Request) {

}
