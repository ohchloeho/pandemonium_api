package handlers

import (
    "net/http"
	"pandemonium_api/internal/services"
    "pandemonium_api/internal/database/models"
    "github.com/gin-gonic/gin"
    "strconv"
)

type VoiceNoteHandler struct {
    voiceNoteService *services.VoiceNoteService
}

func NewVoiceNoteHandler(service *services.VoiceNoteService) *VoiceNoteHandler {
    return &VoiceNoteHandler{voiceNoteService: service}
}

func (h *VoiceNoteHandler) AddVoiceNoteToProject(c *gin.Context) {
    projectID := c.Param("id")

    var note models.VoiceNote
    projectIDInt, err := strconv.Atoi(projectID)
    if err != nil {
        // Handle the error, for example:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
        return
    }    
    
    note.ProjectID = projectIDInt // Assign to the correct project
    h.voiceNoteService.AddVoiceNoteToProject(projectID, note.Content)
    
    c.JSON(http.StatusCreated, note)
}
