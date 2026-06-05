package router

import (
	"net/http"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/handler"
	"github.com/damir/jobfinder/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(authHandler *handler.AuthHandler, jwtManager *auth.JWTManager) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recover)

	authMiddleware :=
		middleware.NewAuthMiddleware(
			jwtManager,
		)

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(
			authMiddleware.Handle,
		)

		r.Get(
			"/me",
			authHandler.Me,
		)
	})

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

	r.Post(
		"/auth/login",
		authHandler.Login,
	)

	return r
}
