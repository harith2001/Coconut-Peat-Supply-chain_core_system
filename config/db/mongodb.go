package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectMongoDB connects to the MongoDB server
// make the username and password hide
func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://harith:harith123@coconut-peat-supply-cha.qmatg.mongodb.net/?retryWrites=true&w=majority&appName=coconut-peat-supply-chain")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
	return client
}

// GetCollection returns a MongoDB collection
func GetCollection(dbName, collectionName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(collectionName)
}
