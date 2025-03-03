package sensor

import (
	"fmt"
	"log"
	"os"

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
func messageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received MQTT message on %s: %s\n", msg.Topic(), msg.Payload())

	//Parse the incoming message
	var q, a, r int
	_, err := fmt.Sscanf(string(msg.Payload()), "%d,%d,%d", &q, &a, &r)
	if err != nil {
		log.Println("Error parsing MQTT message:", err)
		return
	}

	//Update global variables
	Qualified, Acceptable, Rejected = q, a, r
	log.Printf("Updated values - Qualified: %d, Acceptable: %d, Rejected: %d\n", Qualified, Acceptable, Rejected)
}

// Connects to MQTT
func connectMQTT() mqtt.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using default settings")
	}

	mqttBroker := os.Getenv("MQTT_BROKER_URL")
	if mqttBroker == "" {
		mqttBroker = "tcp://localhost:1883"
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID("GradingPluginSubscriber")

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}

	log.Printf("Connected to MQTT broker at %s", mqttBroker)
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

	select {} // have to edit after the data is recived from the sensor to stop the subscriber
}
