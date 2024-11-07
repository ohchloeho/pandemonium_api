package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response structure for the API
type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a response object
	response := Response{Message: "Hello, World!"}

	// Encode response to JSON and write it
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register the /hello endpoint
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
