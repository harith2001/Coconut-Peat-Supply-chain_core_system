package sensor

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

// Global variables to hold sensor data
var (
	Qualified  int
	Acceptable int
	Rejected   int
)

// MQTT message handler
var LastUpdated time.Time

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	var q, a, r int
	_, err := fmt.Sscanf(string(msg.Payload()), "%d,%d,%d", &q, &a, &r)
	if err != nil {
		log.Println("Error parsing MQTT message:", err)
		return
	}

	Qualified, Acceptable, Rejected = q, a, r
	LastUpdated = time.Now()

	log.Printf("Updated values - Q: %d, A: %d, R: %d at %s\n", q, a, r, LastUpdated.Format(time.RFC3339))
}

// Connects to MQTT
func connectMQTT() mqtt.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using default settings")
	}

	mqttBroker := os.Getenv("MQTT_BROKER")
	clientID := os.Getenv("CLIENT_ID")
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	// Use TLS for secure connection
	opts.SetTLSConfig(&tls.Config{
		InsecureSkipVerify: false,
		ClientAuth:         tls.NoClientCert,
	})

	// Set automatic reconnect
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}

	log.Printf("Connected to MQTT broker")
	return client
}

// Subscribes to the topic
func subscribeToSensorData(client mqtt.Client) {
	topic := "grading/sensor_data"
	token := client.Subscribe(topic, 1, messageHandler)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Failed to subscribe to MQTT topic: %v", token.Error())
	}

	log.Printf("Subscribed to MQTT topic: %s\n", topic)
}

// starts listening for sensor data
func StartSensorSubscriber() {
	client := connectMQTT()
	subscribeToSensorData(client)

	// Wait for sensor data to be received
	select {}
}
