package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Structure for the incoming JSON payload
type Recipe struct {
	Filename string `json:"filename"` // Name of the file to upload
	Content  string `json:"content"`  // Content of the file (e.g., markdown content)
}

// Middleware to handle file upload to Nextcloud with JSON data
func UploadToNextcloudMiddleware(c *gin.Context) {
	// Parse the JSON body from the PUT request
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Ensure that the filename and content are not empty
	if recipe.Filename == "" || recipe.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filename and content must be provided"})
		return
	}

	err := UploadToNextcloud(recipe.Filename, recipe.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload file '%s' to Nextcloud: %v", recipe.Filename, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully to Nextcloud"})
}

// Function to upload the file to Nextcloud
func UploadToNextcloud(filename string, content string) error {
	nextcloudUsername := os.Getenv("NEXTCLOUD_USERNAME")
	nextcloudPassword := os.Getenv("NEXTCLOUD_PASSWORD")
	nextcloudBaseURL := "http://127.0.0.1/nextcloud/remote.php/dav/files/"

	// Ensure that Nextcloud credentials are not empty
	if nextcloudUsername == "" || nextcloudPassword == "" {
		return fmt.Errorf("Nextcloud credentials are not set properly")
	}

	nextcloudURL := nextcloudBaseURL + nextcloudUsername + "/" + filename

	// Create a new HTTP PUT request
	req, err := http.NewRequest("PUT", nextcloudURL, bytes.NewBuffer([]byte(content)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set basic auth for Nextcloud
	req.SetBasicAuth(nextcloudUsername, nextcloudPassword)
	req.Header.Set("Content-Type", "text/plain") // Change content type as needed

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("upload request error: %w", err)
	}
	defer resp.Body.Close()

	// Check for success response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
