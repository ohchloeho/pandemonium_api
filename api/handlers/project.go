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
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	projectData := services.ProjectData{
		Name:        project.Name,
		Description: project.Description,
		CreatedBy:   project.CreatedBy,
	}
	h.projectService.CreateProject(projectData)
	c.JSON(200, gin.H{"status": "project created"})
}
