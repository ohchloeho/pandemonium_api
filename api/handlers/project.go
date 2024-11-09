package handlers

import (
	"net/http"
	"pandemonium_api/internal/database/models"
	"pandemonium_api/internal/services"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService *services.ProjectService
}

func NewProjectHandler(service *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: service}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project models.Project
	// Bind the JSON payload to the project struct
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the data to pass to the service layer
	projectData := services.ProjectData{
		Name:        project.Name,
		Description: project.Description,
		CreatedBy:   project.CreatedBy,
	}

	// Call the service to create the project
	createdProject, err := h.projectService.CreateProject(projectData)
	if err != nil {
		// Handle any errors that occur during project creation
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	// Respond with a success status and the created project information
	c.JSON(http.StatusCreated, gin.H{"status": "project created", "project": createdProject})
}
