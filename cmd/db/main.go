package main

import (
	"flag"
	"log"

	"github.com/canhbk/golang-gin-starter-kit/config"
	"github.com/canhbk/golang-gin-starter-kit/database/migration"
	"github.com/canhbk/golang-gin-starter-kit/database/seeder"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}
}

func main() {
	// Parse command line flags
	migrate := flag.Bool("migrate", false, "Run database migrations")
	rollback := flag.Bool("rollback", false, "Rollback database migrations")
	seed := flag.Bool("seed", false, "Seed database with initial data")
	refresh := flag.Bool("refresh", false, "Rollback, migrate, and seed database")
	flag.Parse()

	// Initialize database connection
	config.InitializeDB()

	// Execute commands based on flags
	if *refresh {
		migration.Rollback()
		migration.AutoMigrate()
		seeder.RunSeeders()
	} else {
		if *rollback {
			migration.Rollback()
		}
		if *migrate {
			migration.AutoMigrate()
		}
		if *seed {
			seeder.RunSeeders()
		}
	}
}
