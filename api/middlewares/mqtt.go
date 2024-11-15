package middlewares

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client is the MQTT client structure
type Client struct {
	client mqtt.Client
	broker string
}

// NewClient initializes and returns an MQTT client
func NewClient(broker string, clientID string) *Client {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}

	return &Client{
		client: client,
		broker: broker,
	}
}

// Subscribe subscribes to a topic
func (c *Client) Subscribe(topic string, handler mqtt.MessageHandler) error {
	token := c.client.Subscribe(topic, 0, handler)
	token.Wait()
	return token.Error()
}

// Publish publishes a message to a topic
func (c *Client) Publish(topic string, message string) error {
	token := c.client.Publish(topic, 0, false, message)
	token.Wait()
	return token.Error()
}

// Disconnect disconnects the client from the broker
func (c *Client) Disconnect(gracePeriod uint) {
	c.client.Disconnect(gracePeriod)
}

// WaitForMessages waits for messages on the given topic for a specified timeout
func (c *Client) WaitForMessages(timeout time.Duration) error {
	timeoutChan := time.After(timeout)
	select {
	case <-timeoutChan:
		return fmt.Errorf("timeout waiting for message")
	}
}
