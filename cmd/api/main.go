package main

import (
	"log"
	"net/http"

	"github.com/damir/jobfinder/internal/config"
	"github.com/damir/jobfinder/internal/model"
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

	// routing
	r := chi.NewRouter()

	r.Get("/health", func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", r)
}
