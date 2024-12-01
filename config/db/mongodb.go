package mongo

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// connectMongoDB connects to the MongoDB server
func ConnectMongoDB() *mongo.Client {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatalf("DB_URL is empty")
	}
	clientOptions := options.Client().ApplyURI(DB_URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB!")
	MongoClient = client
	return client
}

// getCollection returns a MongoDB collection
func GetCollection(dbName, collectionName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(collectionName)
}
