package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"URLShortener/props"
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

func (s *Server) ConfigureAPI() {
	s.doOnce.Do(func() {
		configureApi(s.contextRoot, s.port)
	})
}

func configureApi(contextRoot string, port int) {
	var router = mux.NewRouter()
	router.HandleFunc(contextRoot+endpoint, func(rw http.ResponseWriter, r *http.Request) {

	})

	log.Printf("\nApplication is running in : %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatalf("Failure to start Go http server: %v", err)
	}
}
