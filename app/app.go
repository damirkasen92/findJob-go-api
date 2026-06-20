package app

import (
	"net/http"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/config"
	"github.com/damir/jobfinder/internal/router"
	"github.com/damir/jobfinder/pkg/logger"
	"go.uber.org/zap"
)

type App struct {
	Router http.Handler
	Logger *zap.Logger
}

func New() (*App, error) {
	cfg := config.Load()
	db, err := NewDB(cfg)
	logger := logger.Init()
	defer logger.Sync()

	if err != nil {
		return nil, err
	}

	jwtManager := auth.NewJWTManager(cfg.JWTSecret)
	repositories := NewRepositories(db)
	services := NewServices(repositories, jwtManager)
	handlers := NewHandlers(services)

	r := router.NewRouter(
		router.Handlers{
			Auth:        handlers.Auth,
			Vacancy:     handlers.Vacancy,
			Resume:      handlers.Resume,
			Application: handlers.Application,
		},
		jwtManager,
		logger,
	)

	logger.Info("Server started")

	return &App{
		Router: r,
	}, nil
}
