package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nayanprasad/jobq-go/internal/service"
	"github.com/nayanprasad/jobq-go/internal/storage/postgres"
	"github.com/nayanprasad/jobq-go/internal/transport/http/handler"
)

func NewRouter(s *Server) http.Handler {
	r := chi.NewRouter()

	//middlewares
	r.Use(middleware.Logger)

	//routes
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	jobRepository := postgres.NewJobRepository(nil)
	jobService := service.NewService(jobRepository)
	jobHandler := handler.NewHandler(jobService)
	r.Post("/job/create", jobHandler.CreateJob)

	return r
}
