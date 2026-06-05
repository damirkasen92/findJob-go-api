package app

import (
	"net/http"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/config"
	"github.com/damir/jobfinder/internal/router"
)

type App struct {
	Router http.Handler
}

func New() (*App, error) {
	cfg := config.Load()
	db, err := NewDB(cfg)

	if err != nil {
		return nil, err
	}

	jwtManager := auth.NewJWTManager(cfg.JWTSecret)
	repositories := NewRepositories(db)
	services := NewServices(repositories, jwtManager)
	handlers := NewHandlers(services)

	// TODO router accepts only Auth handlers
	r := router.NewRouter(
		handlers.Auth,
		jwtManager,
	)

	return &App{
		Router: r,
	}, nil
}
