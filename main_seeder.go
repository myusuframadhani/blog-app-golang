package main

import (
	"blog-app/config"
	"blog-app/seeders"
)

func main_seeder() {
	config.ConnectDatabase()

	// Jalankan semua seeder
	seeders.SeedAll()
}
