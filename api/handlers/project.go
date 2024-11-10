package handlers

import (
	"net/http"
	"pandemonium_api/internal/database/models"
	"pandemonium_api/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	c.JSON(http.StatusCreated, gin.H{"status": "project created", "project": createdProject})
}

// Get ALL projects
func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	// Call the service to retrieve all projects
	projects, err := h.projectService.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// Get Single Project
func (h *ProjectHandler) GetProject(c *gin.Context) {
	// Get the project ID from the URL parameter
	projectID := c.Param("id")

	// Convert the project ID to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Use the project service to retrieve the project record from the database
	project, err := h.projectService.GetProject(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve project"})
		return
	}

	c.JSON(http.StatusOK, project)
}
