package seeders

import (
	"blog-app/config"
	"blog-app/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}

	users := []models.User{
		{Username: "admin", Email: "admin@example.com", Password: string(hashedPassword), IsAdmin: 1},
		{Username: "user1", Email: "user1@example.com", Password: string(hashedPassword), IsAdmin: 2},
		{Username: "user2", Email: "user2@example.com", Password: string(hashedPassword), IsAdmin: 2},
	}

	for _, user := range users {
		result := config.DB.Create(&user)
		if result.Error != nil {
			log.Fatal("Failed to seed users: ", result.Error)
		}
	}

	log.Println("Users seeded successfully")
}

func SeedAll() {
	SeedUsers()
	// Tambahkan fungsi seeding lainnya di sini
}
