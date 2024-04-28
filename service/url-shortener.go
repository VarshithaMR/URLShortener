package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"URLShortener/service/models"
	"URLShortener/service/utils"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json; charset=utf-8"
)

type UrlShortenerApi interface {
	StartShorteningUrl(*http.Request, http.ResponseWriter)
}

type Shortener struct {
	httpClient *resty.Client
}

func (s *Shortener) StartShorteningUrl(request *http.Request, response http.ResponseWriter) {
	req, err := utils.GetRequestBody(request.Body)
	if err != nil {
		WriteResponse(response, "Request body improper", 400)
		return
	}

	res := ShortenUrl(req)

	WriteResponse(response, res, http.StatusOK)
}

func ShortenUrl(req models.RequestBody) models.ResponseBody {
	encoded := EncodeURL(req.URL)
	newUrl := fmt.Sprintf("http://shortenedURL/%d", encoded)
	return models.ResponseBody{
		ShortUrl: newUrl,
	}
}

func EncodeURL(url string) int {
	readByte := []byte(url)
	destByte := make([]byte, hex.EncodedLen(len(readByte)))
	x := hex.Encode(destByte, readByte)
	return x
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
	return &Shortener{
		httpClient: resty.New(),
	}
}
