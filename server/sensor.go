package server

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Map of plugin subscriptions
var pluginSubscriptions = map[string]string{
	"pluginA": "sensors/temperature",
	"pluginB": "sensors/motion",
}

var messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Core received message on %s: %s\n", msg.Topic(), msg.Payload())

	// Forward to relevant plugin
	for plugin, topic := range pluginSubscriptions {
		if msg.Topic() == topic {
			forwardToPlugin(plugin, msg.Payload())
		}
	}
}

func forwardToPlugin(plugin string, data []byte) {
	fmt.Printf("Forwarding data to %s: %s\n", plugin, data)
	// Send data via another MQTT topic if needed
}

func main() {
	broker := "tcp://broker.hivemq.com:1883" // use the env
	clientID := "CoreSystem"                 //use the env

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messageHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting:", token.Error())
		os.Exit(1)
	}

	fmt.Println("Core System Connected to MQTT Broker!")

	// Subscribe to all sensor topics
	for _, topic := range pluginSubscriptions {
		token := client.Subscribe(topic, 1, nil)
		token.Wait()
		fmt.Printf("Core Subscribed to: %s\n", topic)
	}

	// Keep connection alive
	for {
		time.Sleep(5 * time.Second)
	}
}
