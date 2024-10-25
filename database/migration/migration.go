package migration

import (
	"log"

	"github.com/canhbk/golang-gin-starter-kit/config"
	"github.com/canhbk/golang-gin-starter-kit/models"
)

// AutoMigrate will create or modify tables based on models
func AutoMigrate() {
	log.Println("Running database migrations...")

	err := config.DB.AutoMigrate(
		&models.User{},
		// Add more models here
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
}

// Rollback will drop all tables
func Rollback() {
	log.Println("Rolling back database migrations...")

	err := config.DB.Migrator().DropTable(
		&models.User{},
		// Add more models here
	)

	if err != nil {
		log.Fatalf("Failed to rollback database: %v", err)
	}

	log.Println("Database rollback completed successfully")
}
