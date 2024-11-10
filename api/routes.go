package api

import (
	"net/http"
	"pandemonium_api/api/handlers"
	"pandemonium_api/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// POST route that receives a JSON payload
	router.POST("/test", func(c *gin.Context) {
		var json map[string]interface{}

		// Bind JSON data to the `json` variable
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Respond with the received data
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "POST request received",
			"data":    json, // echo back the received JSON
		})
	})

	// Pass the MongoDB database instance to your services
	projectService := services.NewProjectService(db)
	projectHandler := handlers.NewProjectHandler(projectService)

	voiceNoteService := services.NewVoiceNoteService(db)
	voiceNoteHandler := handlers.NewVoiceNoteHandler(voiceNoteService)

	// Define your routes
	router.GET("/projects", projectHandler.GetAllProjects)
	router.GET("/projects/:id", projectHandler.GetProject)
	router.POST("/projects", projectHandler.CreateProject)
	router.POST("/projects/:id/voice-note", voiceNoteHandler.AddVoiceNoteToProject)

	return router
}
