package server

import (
	"URLShortener/props"
	"sync"
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
