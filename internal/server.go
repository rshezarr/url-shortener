package internal

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"url-short/internal/config"
)

type Server struct {
	srv               *http.Server
	ServerErrorNotify chan error
}

func NewServer(cfg *config.Configuration, handler *chi.Mux) *Server {
	return &Server{
		srv: &http.Server{
			Handler:        handler,
			Addr:           ":" + cfg.HTTP.Port,
			ReadTimeout:    cfg.HTTP.Timeout,
			WriteTimeout:   cfg.HTTP.Timeout,
			IdleTimeout:    cfg.HTTP.IdleTimeout,
			MaxHeaderBytes: cfg.HTTP.HeaderBytes << 20,
		},
		ServerErrorNotify: make(chan error, 1),
	}
}

func (s *Server) Run() {
	s.ServerErrorNotify <- s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.ServerErrorNotify
}
