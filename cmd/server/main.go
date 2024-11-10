package main

import (
	"context"
	"fmt"
	"log"
	"pandemonium_api/api"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB client setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the MongoDB server to ensure the connection is successful
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Successfully connected to MongoDB")

	// Connect to a specific database
	db := client.Database("pandaDB")

	// Initialize the Gin router
	router := api.SetupRouter(db)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
