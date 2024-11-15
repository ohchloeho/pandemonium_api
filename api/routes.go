package api

import (
	"pandemonium_api/api/handlers"
	"pandemonium_api/api/middlewares"
	"pandemonium_api/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "this ping test is working for the REST API"})
	})

	// // Pass the MongoDB database instance to your services
	// projectService := services.NewProjectService(db)
	// projectHandler := handlers.NewProjectHandler(projectService)
	// voiceNoteService := services.NewVoiceNoteService(db)
	// voiceNoteHandler := handlers.NewVoiceNoteHandler(voiceNoteService)

	// MQTT
	mqttHandler := handlers.NewMQTTHandler()
	topics := []string{"test/topic", "projects/updates"}
	services.NewMQTTService("tcp://192.168.8.214:1883", "mqtt_api_client", topics, mqttHandler.HandleMessage)

	// Routes
	// router.GET("/projects", projectHandler.GetAllProjects)
	// router.GET("/projects/:id", projectHandler.GetProject)
	// router.POST("/projects", projectHandler.CreateProject)
	// router.POST("/projects/:id/voice-note", voiceNoteHandler.AddVoiceNoteToProject)

	router.POST("/ping-nc", middlewares.UploadToNextcloudMiddleware)

	return router
}
