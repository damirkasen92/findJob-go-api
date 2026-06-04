package main

import (
	"log"
	"net/http"

	"github.com/damir/jobfinder/internal/config"
	"github.com/damir/jobfinder/internal/handler"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
	"github.com/damir/jobfinder/internal/router"
	"github.com/damir/jobfinder/internal/service"
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

	r := router.NewRouter(
		authHandler,
	)

	err = http.ListenAndServe(":9000", r)

	if err != nil {
		log.Fatal(err)
	}
}
