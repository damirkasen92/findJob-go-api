package router

import (
	"net/http"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/handler"
	"github.com/damir/jobfinder/internal/middleware"
	"github.com/damir/jobfinder/internal/model"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Handlers struct {
	Auth        *handler.AuthHandler
	Vacancy     *handler.VacancyHandler
	Resume      *handler.ResumeHandler
	Application *handler.ApplicationHandler
}

func NewRouter(handlers Handlers, jwtManager *auth.JWTManager, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()
	authMiddleware := middleware.NewAuthMiddleware(jwtManager)
	loggerMiddleware := middleware.Logger(logger)

	r.Use(middleware.RequestID)
	r.Use(loggerMiddleware)
	r.Use(middleware.Recover)
	r.Use(middleware.ErrorMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		// health
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

		// auth
		r.Post("/auth/register", handlers.Auth.Register)
		r.Post("/auth/login", handlers.Auth.Login)
		r.Post("/auth/refresh", handlers.Auth.Refresh)

		// protected
		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.Handle)
			r.Get("/me", handlers.Auth.Me)
		})

		// vacancies
		r.Get("/vacancies", handlers.Vacancy.GetList)
		r.Get("/vacancies/{vacancyID}", handlers.Vacancy.GetByID)

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.Handle)
			r.Use(middleware.RequireRole(model.RoleAdmin, model.RoleCompany))

			r.Post("/vacancies", handlers.Vacancy.Create)
			r.Delete("/vacancies/{vacancyID}", handlers.Vacancy.Delete)
			r.Get("/vacancies/{vacancyID}/applications", handlers.Application.ListByVacancy)
		})

		// resumes
		r.Get("/resumes", handlers.Resume.GetList)
		r.Get("/resumes/{resumeID}", handlers.Resume.GetByID)

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.Handle)
			r.Get("/my/resumes", handlers.Resume.MyResumes)
		})

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.Handle)
			r.Use(middleware.RequireRole(model.RoleUser))

			r.Post("/resumes", handlers.Resume.Create)
			r.Patch("/resumes", handlers.Resume.Update)
			r.Delete("/resumes/{resumeID}", handlers.Resume.Delete)
		})

		// applications
		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.Handle)
			r.Get("/my/applications", handlers.Application.GetByUser)
			r.Post("/applications", handlers.Application.Create)
			r.Patch("/applications/{id}/status", handlers.Application.UpdateStatus)
		})
	})

	return r
}
