package services

import (
	"context"
	"errors"
	"log"
	"time"

	"pandemonium_api/internal/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectData struct {
	Name        string
	Description string
	CreatedBy   string
}

type ProjectService struct {
	db *mongo.Database
}

func NewProjectService(db *mongo.Database) *ProjectService {
	return &ProjectService{db: db}
}

func (s *ProjectService) CreateProject(data ProjectData) (*models.Project, error) {
	// Validate input data
	if data.Name == "" {
		return nil, errors.New("project name cannot be empty")
	}

	// Create a new Project instance
	project := &models.Project{
		Name:        data.Name,
		Description: data.Description,
		CreatedBy:   data.CreatedBy,
		CreatedAt:   time.Now(),
	}

	// Insert the project into the "projects" collection in MongoDB
	collection := s.db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform the insertion
	result, err := collection.InsertOne(ctx, project)
	if err != nil {
		log.Println("Error inserting project into database:", err)
		return nil, err
	}

	// Set the inserted ID to the project (MongoDB generates an ID for us)
	project.ID = result.InsertedID.(primitive.ObjectID).Hex()

	log.Println("Project created successfully:", project)
	return project, nil
}

// Retrieve ALL projects
func (s *ProjectService) GetAllProjects() ([]*models.Project, error) {
	// Use the MongoDB driver to retrieve all projects from the database
	collection := s.db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var projects []*models.Project
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	if err := cur.All(ctx, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// Retrieve a project by ID
func (s *ProjectService) GetProject(id primitive.ObjectID) (*models.Project, error) {
	// Use the MongoDB driver to retrieve the project record from the database
	collection := s.db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var project models.Project
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}
