// package sensor

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sync"
// 	"time"

// 	mongo "Coconut-Peat-Supply-chain_core_system/config/db"

// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// // Set to store unique topics
// var (
// 	topicSet = make(map[string]bool)
// 	mu       sync.Mutex
// )

// func storeTopic(topic string) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Check if topic exists
// 	count, err := mongo.MongoClient.Database("topicDB").Collection("topics").CountDocuments(ctx, bson.M{"topic": topic})
// 	if err != nil {
// 		log.Printf("Error checking topic existence: %v", err)
// 		return
// 	}

// 	// Insert if not exists
// 	if count == 0 {
// 		_, err = mongo.MongoClient.Database("topicDB").Collection("topics").InsertOne(ctx, bson.M{"topic": topic, "timestamp": time.Now()})
// 		if err != nil {
// 			log.Printf("Error inserting topic: %v", err)
// 		} else {
// 			fmt.Printf("Stored topic: %s\n", topic)
// 		}
// 	}
// }

// func SensorMain() {

// 	if err := godotenv.Load(); err != nil {
// 		log.Println("Warning: No .env file found, using default settings")
// 	}

// 	mqttBroker := os.Getenv("MQTT_BROKER")
// 	if mqttBroker == "" {
// 		mqttBroker = "tcp://hivemq:1883"
// 	}
// 	clientID := os.Getenv("CLIENT_ID")
// 	if clientID == "" {
// 		clientID = "CoreClient"
// 	}
// 	opts := mqtt.NewClientOptions()
// 	opts.AddBroker(mqttBroker)
// 	opts.SetClientID(clientID)

// 	client := mqtt.NewClient(opts)
// 	if token := client.Connect(); token.Wait() && token.Error() != nil {
// 		log.Fatalf("Connection error: %v", token.Error())
// 	}

// 	fmt.Println("Connected to HiveMQ Broker!")

// 	// Subscribe to all topics
// 	topic := "#"
// 	token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
// 		mu.Lock()
// 		defer mu.Unlock()

// 		if _, exists := topicSet[msg.Topic()]; !exists {
// 			topicSet[msg.Topic()] = true
// 			fmt.Printf("New topic detected: %s\n", msg.Topic())
// 			storeTopic(msg.Topic())
// 		}
// 	})

// 	if token.Wait() && token.Error() != nil {
// 		log.Fatalf("Subscription error: %v", token.Error())
// 	}

// 	fmt.Println("Subscribed to all topics!")

// 	// Periodically print all unique topics
// 	go func() {
// 		for {
// 			time.Sleep(10 * time.Second) // Print every 10 seconds
// 			mu.Lock()
// 			fmt.Println("\nList of detected topics:")
// 			for topic := range topicSet {
// 				fmt.Println(topic)
// 			}
// 			mu.Unlock()
// 		}
// 	}()

// 	// Keep the connection alive
// 	select {} // Blocks forever
// }

package sensor

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	mongo "Coconut-Peat-Supply-chain_core_system/config/db"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	topicSet = make(map[string]bool)
	mu       sync.Mutex
)

func storeTopic(topic string) error {

	fmt.Printf("Topic: %s\n", topic)

	mu.Lock()
	defer mu.Unlock()

	fmt.Println("MongoDB client:", mongo.MongoClient)

	// Check if topic already exists
	filter := bson.M{"topic": topic}
	var exisitingTopic bson.M
	err := mongo.MongoClient.Database("topicDB").Collection("topics").FindOne(context.Background(), filter).Decode(&exisitingTopic)
	if err == nil {
		return fmt.Errorf("topic with name '%s' already exists", topic)
	}

	// Insert if not exists
	_, err = mongo.MongoClient.Database("topicDB").Collection("topics").InsertOne(context.Background(), bson.M{"topic": topic, "timestamp": time.Now()})
	if err != nil {
		return fmt.Errorf("error storing topic: %v", err)
	}

	fmt.Printf("Stored topic: %s\n", topic)
	return nil

}

func waitForShutdown(client mqtt.Client) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	<-sigs
	fmt.Println("\nShutting down...")

	client.Disconnect(250)
	if err := mongo.MongoClient.Disconnect(context.Background()); err != nil {
		log.Printf("Error closing MongoDB connection: %v", err)
	}
	os.Exit(0)
}

func SensorMain() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using default settings")
	}

	mqttBroker := os.Getenv("MQTT_BROKER")
	if mqttBroker == "" {
		mqttBroker = "tcp://hivemq:1883"
	}
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		clientID = "CoreClient"
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Connection error: %v", token.Error())
	}

	fmt.Println("Connected to HiveMQ Broker!")

	// Subscribe to all topics
	topic := "#"
	token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		mu.Lock()
		defer mu.Unlock()

		if _, exists := topicSet[msg.Topic()]; !exists {
			topicSet[msg.Topic()] = true
			fmt.Printf("New topic detected: %s\n", msg.Topic())
			storeTopic(msg.Topic())
		}
	})

	if token.Wait() && token.Error() != nil {
		log.Fatalf("Subscription error: %v", token.Error())
	}

	fmt.Println("Subscribed to all topics!")

	// Periodically print all unique topics
	go func() {
		for {
			time.Sleep(10 * time.Second)
			mu.Lock()
			fmt.Println("\nList of detected topics:")
			for topic := range topicSet {
				fmt.Println(topic)
			}
			mu.Unlock()
		}
	}()

	go waitForShutdown(client)

	select {} // Blocks forever
}
