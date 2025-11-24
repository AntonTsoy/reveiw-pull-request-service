package server

import "github.com/AntonTsoy/review-pull-request-service/internal/transport/http/handlers"

type Server struct {
	handlers *handlers.Handlers
}

func New(handlers *handlers.Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}
