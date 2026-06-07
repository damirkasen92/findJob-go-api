package app

import (
	"log"

	"github.com/damir/jobfinder/internal/config"
	"github.com/damir/jobfinder/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(
	cfg *config.Config,
) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(cfg.DBDSN),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Vacancy{},
	)

	if err != nil {
		return nil, err
	}

	log.Println("database connected")

	return db, nil
}
