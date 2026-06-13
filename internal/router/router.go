package router

import (
	"net/http"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/handler"
	"github.com/damir/jobfinder/internal/middleware"
	"github.com/damir/jobfinder/internal/model"
	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Auth    *handler.AuthHandler
	Vacancy *handler.VacancyHandler
	Resume  *handler.ResumeHandler
}

func NewRouter(handlers Handlers, jwtManager *auth.JWTManager) *chi.Mux {
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
			handlers.Auth.Me,
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
		handlers.Auth.Register,
	)

	r.Post(
		"/auth/login",
		handlers.Auth.Login,
	)

	r.Post(
		"/auth/refresh",
		handlers.Auth.Refresh,
	)

	r.Get("/vacancies/{vacancyID}", handlers.Vacancy.GetByID)
	r.Get("/vacancies", handlers.Vacancy.GetList)

	r.Group(func(r chi.Router) {
		r.Use(
			authMiddleware.Handle,
		)

		r.Use(
			middleware.RequireRole(
				model.RoleAdmin,
				model.RoleCompany,
			),
		)

		r.Post(
			"/vacancies",
			handlers.Vacancy.Create,
		)

		r.Delete(
			"/vacancies/{vacancyID}",
			handlers.Vacancy.Delete,
		)
	})

	r.Get("/resumes", handlers.Resume.GetList)
	r.Get("/resumes/{resumeID}", handlers.Resume.GetByID)

	r.Group(func(r chi.Router) {
		r.Use(
			authMiddleware.Handle,
		)

		r.Get("/my/resumes", handlers.Resume.MyResumes)
	})

	r.Group(func(r chi.Router) {
		r.Use(
			authMiddleware.Handle,
		)

		r.Use(
			middleware.RequireRole(
				model.RoleUser,
			),
		)

		r.Post(
			"/resumes",
			handlers.Resume.Create,
		)

		r.Delete(
			"/resumes/{resumeID}",
			handlers.Resume.Delete,
		)

	})

	return r
}
