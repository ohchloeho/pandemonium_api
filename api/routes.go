package api

import (
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

	// Pass the MongoDB database instance to your services
	projectService := services.NewProjectService(db)
	projectHandler := handlers.NewProjectHandler(projectService)

	voiceNoteService := services.NewVoiceNoteService(db)
	voiceNoteHandler := handlers.NewVoiceNoteHandler(voiceNoteService)

	// Define your routes
	router.POST("/projects", projectHandler.CreateProject)
	router.POST("/projects/:id/voice-note", voiceNoteHandler.AddVoiceNoteToProject)

	return router
}
