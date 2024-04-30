package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"URLShortener/props"
	"URLShortener/service"
)

const (
	endpoint = "/shorten"
)

type Server struct {
	host        string
	port        int
	contextRoot string
	doOnce      sync.Once
}

func NewServer(properties *props.Properties) *Server {
	server := new(Server)
	server.host = properties.Server.Host
	server.port = properties.Server.Port
	server.contextRoot = properties.Server.ContextRoot
	return server
}

func (s *Server) ConfigureAPI(shortener service.UrlShortenerApi) {
	s.doOnce.Do(func() {
		configureApi(s.contextRoot, s.port, shortener)
	})
}

func configureApi(contextRoot string, port int, shortener service.UrlShortenerApi) {
	var router = mux.NewRouter()
	router.HandleFunc(contextRoot+endpoint, func(rw http.ResponseWriter, r *http.Request) {
		HandleURLShortener(rw, r, shortener)
	})

	log.Printf("\nApplication is running in : %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatalf("Failure to start Go http server: %v", err)
	}
}
