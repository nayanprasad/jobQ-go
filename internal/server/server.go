package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func (*Server) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	return r
}

func (s *Server) Run(h http.Handler) error {
	svr := &http.Server{
		Addr:    s.config.Addr,
		Handler: h,
	}

	slog.Info("server started on", "port", s.config.Addr)

	return svr.ListenAndServe()
}
