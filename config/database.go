package config

import (
	"blog-app/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"path/filepath"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Mengambil variabel lingkungan (ENV) untuk koneksi database
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Mengatur DNS untuk PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Category{}, &models.Tag{})

	fmt.Println("Database connection successfully opened")
}

func init() {
	err := godotenv.Load(filepath.Join(".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
