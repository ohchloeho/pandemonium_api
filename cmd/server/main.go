package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create a new Gin router
	router := gin.Default()

	// POST route definition
	router.POST("/test", func(c *gin.Context) {
		var json map[string]interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": json})
	})

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
	// // MongoDB client setup
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// client, err := mongo.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal("Failed to connect to MongoDB:", err)
	// }

	// // Ping the MongoDB server to ensure the connection is successful
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal("Failed to ping MongoDB:", err)
	// }

	// fmt.Println("Successfully connected to MongoDB")

	// // Connect to a specific database
	// db := client.Database("pandaDB")

	// // Initialize the Gin router
	// router := api.SetupRouter(db)

	// // Start the server
	// if err := router.Run(":8080"); err != nil {
	// 	log.Fatal("Unable to start server: ", err)
	// }
}
