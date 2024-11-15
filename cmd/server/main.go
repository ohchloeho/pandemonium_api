package main

import (
	"log"
	"pandemonium_api/api"
	"pandemonium_api/internal/database"
)

func initDB() *database.DB {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	return db
}

func main() {
	// Initialize MongoDB
	// db := initDB()
	// defer db.Close()

	router := api.SetupRouter()

	// Set trusted proxies to allow only specific IPs (e.g., your Apache server's IP
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
