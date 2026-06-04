package router

import (
	"net/http"

	"github.com/damir/jobfinder/internal/handler"
	"github.com/damir/jobfinder/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(authHandler *handler.AuthHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recover)

	r.Get("/health", func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.Write([]byte("ok"))
	})

	r.Post(
		"/auth/register",
		authHandler.Register,
	)

	return r
}
