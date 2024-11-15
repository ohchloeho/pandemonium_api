package services

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
	client mqtt.Client
}

// NewMQTTService creates and initializes a new MQTT client, subscribes to topics, and starts the message handler.
func NewMQTTService(brokerURL, clientID string, topics []string, messageHandler mqtt.MessageHandler) *MQTTService {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messageHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to broker: %v", token.Error())
	}
	log.Println("Connected to MQTT broker")

	for _, topic := range topics {
		if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
			log.Printf("Failed to subscribe to topic %s: %v", topic, token.Error())
		} else {
			log.Printf("Subscribed to topic %s", topic)
		}
	}

	return &MQTTService{
		client: client,
	}
}

// Start keeps the client alive and blocks indefinitely to process incoming messages.
func (s *MQTTService) Start() {
	log.Println("MQTT Service started, waiting for messages...")
	s.client.Connect().Wait()
}

// Close gracefully disconnects from the MQTT broker.
func (s *MQTTService) Close() {
	if s.client.IsConnected() {
		s.client.Disconnect(250)
		log.Println("Disconnected from MQTT broker")
	} else {
		log.Println("MQTT client is already disconnected")
	}
}
