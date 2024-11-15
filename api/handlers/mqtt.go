package handlers

import (
	"fmt"
	"log"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/studio-b12/gowebdav"
)

type MQTTHandler struct {
	nextcloudClient *gowebdav.Client
}

// NewMQTTHandler initializes an MQTT handler with optional dependencies
func NewMQTTHandler() *MQTTHandler {
	nextcloudUsername := os.Getenv("NEXTCLOUD_USERNAME")
	nextcloudPassword := os.Getenv("NEXTCLOUD_PASSWORD")
	nextcloudURL := "http://100.127.215.78/nextcloud"

	client := gowebdav.NewClient(nextcloudURL, nextcloudUsername, nextcloudPassword)
	return &MQTTHandler{
		nextcloudClient: client,
	}
}

// HandleMessage processes incoming MQTT messages
func (h *MQTTHandler) HandleMessage(client mqtt.Client, msg mqtt.Message) {
	// Parse message topic and payload
	topic := msg.Topic()
	payload := string(msg.Payload())

	// Log the received message
	fmt.Printf("Received message on topic %s: %s\n", topic, payload)

	// Example of message format: "create|filename|file_contents" for file creation
	// or "read|filename" for reading a file

	parts := strings.Split(payload, "|")
	if len(parts) < 2 {
		log.Printf("Invalid message format: %s", payload)
		return
	}

	command := parts[0]
	filename := parts[1]

	switch command {
	case "create":
		// Create a file in Nextcloud
		if len(parts) < 3 {
			log.Printf("Invalid message format for 'create' command")
			return
		}
		contents := parts[2]
		err := h.createFile(filename, contents)
		if err != nil {
			log.Printf("Failed to create file in Nextcloud: %v", err)
		} else {
			log.Printf("File '%s' created successfully in Nextcloud", filename)
		}
	case "read":
		// Read a file from Nextcloud
		contents, err := h.readFile(filename)
		if err != nil {
			log.Printf("Failed to read file from Nextcloud: %v", err)
		} else {
			log.Printf("File '%s' contents: %s", filename, contents)
		}
	default:
		log.Printf("Unknown command: %s", command)
	}
}

// createFile creates a file in Nextcloud
func (h *MQTTHandler) createFile(filename, contents string) error {
	// Use WebDAV to create a file in Nextcloud
	// WebDAV path: "/remote.php/webdav/"
	path := "/remote.php/webdav/files/" + filename

	// Write the contents to Nextcloud
	err := h.nextcloudClient.Write(path, []byte(contents), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file to Nextcloud: %w", err)
	}
	return nil
}

// readFile reads a file from Nextcloud
func (h *MQTTHandler) readFile(filename string) (string, error) {
	// Use WebDAV to read a file from Nextcloud
	// WebDAV path: "/remote.php/webdav/"
	path := "/remote.php/webdav/files/" + filename

	data, err := h.nextcloudClient.Read(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file from Nextcloud: %w", err)
	}
	return string(data), nil
}
