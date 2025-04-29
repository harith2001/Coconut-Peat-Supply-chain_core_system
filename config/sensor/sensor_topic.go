package sensor

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	mongo "github.com/harith2001/Coconut-Peat-Supply-chain_core_system/config/db"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	topicSet = make(map[string]bool)
	mu       sync.RWMutex // Use RWMutex for read-heavy operations
)

// Graceful shutdown handling
func waitForShutdown(client mqtt.Client, cancel context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	<-sigs
	fmt.Println("\nShutting down...")

	cancel() // Cancel context
	client.Disconnect(250)
	if err := mongo.MongoClient.Disconnect(context.Background()); err != nil {
		log.Printf("Error closing MongoDB connection: %v", err)
	}
	os.Exit(0)
}

// Preload topics from database to minimize queries
func preloadTopics(ctx context.Context) {
	mu.Lock()
	defer mu.Unlock()

	collection := mongo.MongoClient.Database("test").Collection("topics")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error preloading topics: %v", err)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var topicData struct {
			Topic string `bson:"topic"`
		}
		if err := cursor.Decode(&topicData); err == nil {
			topicSet[topicData.Topic] = true
		}
	}
}

// Optimized message handler
func messageHandler(client mqtt.Client, msg mqtt.Message) {
	mu.RLock()
	_, exists := topicSet[msg.Topic()]
	mu.RUnlock()

	if exists {
		return // Skip already stored topics
	}

	mu.Lock()
	topicSet[msg.Topic()] = true
	mu.Unlock()

	fmt.Printf("New topic detected: %s\n", msg.Topic())

	// Upsert topic to MongoDB (Insert if not exists)
	filter := bson.M{"topic": msg.Topic()}
	update := bson.M{"$set": bson.M{"timestamp": time.Now()}}
	opts := options.Update().SetUpsert(true)

	_, err := mongo.MongoClient.Database("test").Collection("topics").UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		log.Printf("Error storing topic: %v", err)
	} else {
		fmt.Printf("Stored topic: %s\n", msg.Topic())
	}
}

// Main function
func SensorMain() error {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using default settings")
	}

	// Set up MQTT Broker details
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
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Connection error: %v", token.Error())
		return token.Error()
	}

	fmt.Println("Connected to HiveMQ Cloud!")

	// Preload topics from MongoDB
	ctx, cancel := context.WithCancel(context.Background())
	preloadTopics(ctx)

	// Subscribe to all topics
	topic := "#"
	token := client.Subscribe(topic, 1, messageHandler)
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Subscription error: %v", token.Error())
	}

	fmt.Println("Subscribed to all topics!")

	// Periodic topic list print
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(10 * time.Second):
				mu.RLock()
				fmt.Println("\n List of detected topics:")
				for topic := range topicSet {
					fmt.Println(topic)
				}
				mu.RUnlock()
			}
		}
	}()

	// Handle shutdown
	go waitForShutdown(client, cancel)

	<-ctx.Done() // Blocks until shutdown signal is received
	return nil
}
