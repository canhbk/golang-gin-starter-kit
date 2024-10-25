package seeder

import (
	"log"
	"time"

	"github.com/canhbk/golang-gin-starter-kit/config"
	"github.com/canhbk/golang-gin-starter-kit/models"
	"golang.org/x/crypto/bcrypt"
)

// RunSeeders executes all seeders
func RunSeeders() {
	log.Println("Running database seeders...")

	// Run individual seeders
	seedUsers()
	// Add more seeder functions here

	log.Println("Database seeding completed successfully")
}

// seedUsers creates initial user records
func seedUsers() {
	log.Println("Seeding users...")

	// Hash password for users
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	users := []models.User{
		{
			Username:  "admin",
			Email:     "admin@example.com",
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Username:  "user",
			Email:     "user@example.com",
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		result := config.DB.FirstOrCreate(&user, models.User{Email: user.Email})
		if result.Error != nil {
			log.Printf("Error seeding user %s: %v", user.Email, result.Error)
		}
	}
}
