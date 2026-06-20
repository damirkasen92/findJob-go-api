package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/damir/jobfinder/app"
	"go.uber.org/zap"
)

func main() {
	app, err := app.New()
	logger := app.Logger

	if err != nil {
		logger.Fatal(err.Error())
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: app.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen",
				zap.Any("err", err),
			)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Info("shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("server forced to shutdown",
			zap.Any("err", err),
		)
	}

	logger.Info("server exited")
}
