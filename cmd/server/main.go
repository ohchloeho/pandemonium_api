package main

import (
    "log"
    "pandemonium_api/internal/database"
	"pandemonium_api/api"
)

func main() {
    db, err := database.NewDB()
    if err != nil {
        log.Fatalf("Could not connect to MongoDB: %v", err)
    }
    defer db.Close()
	// Set up the routes with the database connection
	router := api.SetupRouter(db.Database)

	// Start the server
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

    // Pass db.Database to other parts of your application, like your services
    log.Println("MongoDB setup complete, proceeding with application startup.")
}
