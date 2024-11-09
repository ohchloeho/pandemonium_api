package api

import (
    "github.com/gin-gonic/gin"
    "pandemonium_api/api/handlers"
    "pandemonium_api/internal/services"
    "go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
    router := gin.Default()

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
