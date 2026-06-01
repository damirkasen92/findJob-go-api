package main

import (
	"log"
	"net/http"

	"github.com/damir/jobfinder/internal/config"
	"github.com/damir/jobfinder/internal/handler"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
	"github.com/damir/jobfinder/internal/service"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO bloated main func
func main() {
	cfg := config.Load()

	db, err := gorm.Open(
		postgres.Open(cfg.DBDSN),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connected")

	err = db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(
		userRepo,
	)

	authHandler := handler.NewAuthHandler(
		userService,
	)

	// routing
	r := chi.NewRouter()

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

	http.ListenAndServe(":8080", r)
}
