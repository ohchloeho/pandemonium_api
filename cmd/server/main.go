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
	db := initDB()
	defer db.Close()

	// Initialize the Gin router
	router := api.SetupRouter(db.Database)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
