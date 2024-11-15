package services

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
	client mqtt.Client
}

// NewMQTTService initializes a new MQTT service and connects to the broker
func NewMQTTService(broker string, clientID string, topics []string, messageHandler mqtt.MessageHandler) *MQTTService {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)

	// Connect to the broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}

	// Subscribe to each topic with the provided message handler
	for _, topic := range topics {
		token := client.Subscribe(topic, 0, messageHandler)
		token.Wait()
		if token.Error() != nil {
			log.Fatalf("Failed to subscribe to topic %s: %v", topic, token.Error())
		}
	}

	return &MQTTService{client: client}
}

// Close disconnects the MQTT client
func (s *MQTTService) Close() {
	s.client.Disconnect(250)
}
