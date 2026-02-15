package api

import (
	"log/slog"
	"net/http"
)

type Config struct {
	Addr string
	DSN  string
}

type Server struct {
	config Config
}

func New(cnf Config) *Server {
	return &Server{
		config: cnf,
	}
}

func (s *Server) Mount() http.Handler {
	return NewRouter(s)
}

func (s *Server) Run(h http.Handler) error {
	svr := &http.Server{
		Addr:    s.config.Addr,
		Handler: h,
	}

	slog.Info("server started on", "port", s.config.Addr)

	return svr.ListenAndServe()
}
