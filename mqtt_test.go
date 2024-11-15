package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Helper function to create a new MQTT client for testing
func createMQTTClient(broker, clientID string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}
	return client
}

func TestMQTTConnection(t *testing.T) {
	broker := "tcp://localhost:1883" // Local broker
	topic := "test/topic"

	// Create and connect MQTT client
	client := createMQTTClient(broker, "test_mqtt_client")
	defer client.Disconnect(250)

	// Channel to handle received messages
	messageReceived := make(chan string, 1)

	// Message handler for when a message is received
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		message := string(msg.Payload())
		fmt.Printf("Received message: %s from topic: %s\n", message, msg.Topic())
		messageReceived <- message
	})

	// Publish a test message
	testMessage := "Hello from Go MQTT test"
	token := client.Publish(topic, 0, false, testMessage)
	token.Wait()

	select {
	case receivedMsg := <-messageReceived:
		if receivedMsg != testMessage {
			t.Errorf("Expected message %q but got %q", testMessage, receivedMsg)
		}
	case <-time.After(5 * time.Second):
		t.Error("Timeout waiting for message to be received")
	}
}

func TestMultipleMessages(t *testing.T) {
	broker := "tcp://localhost:1883" // Local broker
	topic := "test/topic"

	client := createMQTTClient(broker, "test_mqtt_client_multiple")
	defer client.Disconnect(250)

	// Channel to handle received messages
	messagesReceived := make(chan string, 3)

	// Message handler for when a message is received
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		message := string(msg.Payload())
		fmt.Printf("Received message: %s from topic: %s\n", message, msg.Topic())
		messagesReceived <- message
	})

	// Publish multiple test messages
	testMessages := []string{"Message 1", "Message 2", "Message 3"}
	for _, msg := range testMessages {
		token := client.Publish(topic, 0, false, msg)
		token.Wait()
	}

	// Wait for messages to be received and validate them
	for i := 0; i < len(testMessages); i++ {
		select {
		case receivedMsg := <-messagesReceived:
			if receivedMsg != testMessages[i] {
				t.Errorf("Expected message %q but got %q", testMessages[i], receivedMsg)
			}
		case <-time.After(5 * time.Second):
			t.Errorf("Timeout waiting for message %d to be received", i+1)
		}
	}
}

func TestMQTTSubscription(t *testing.T) {
	broker := "tcp://localhost:1883" // Local broker (e.g., Mosquitto)
	topic := "test/topic"

	client := createMQTTClient(broker, "test_mqtt_subscriber")
	defer client.Disconnect(250)

	// Channel to receive messages asynchronously
	messagesReceived := make(chan string)

	// Subscribe to the topic and log incoming messages
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		message := string(msg.Payload())
		fmt.Printf("Received message: %s from topic: %s\n", message, msg.Topic())
		messagesReceived <- message
	})

	fmt.Println("Subscribed to topic:", topic)
	fmt.Println("Listening for messages... (you can publish with mosquitto_pub)")

	// Run the test as a persistent subscriber
	timeout := time.After(30 * time.Second) // Timeout after 30 seconds
	for {
		select {
		case receivedMsg := <-messagesReceived:
			fmt.Printf("Message received: %s\n", receivedMsg)
		case <-timeout:
			fmt.Println("Test timeout reached, ending subscription.")
			return
		}
	}
}
