package main

import (
	"blog-app/config"
	"blog-app/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Koneksi ke database
	config.ConnectDatabase()

	// Register routes
	routes.RegisterRoutes(router)

	// Jalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
