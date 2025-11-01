package database

import (
	"log"

	"github.com/nepile/api-monitoring/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db

	err = DB.AutoMigrate(
		&models.User{},
		&models.Endpoint{},
		&models.CheckLog{},
	)
	if err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
}
