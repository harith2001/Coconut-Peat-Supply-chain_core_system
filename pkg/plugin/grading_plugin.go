package plugin

import (
	"context"
	"log"
	"time"

	"Coconut-Peat-Supply-chain_core_system/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExecutionResult represents the result of a grading execution
type ExecutionResult struct {
	Qualified  int
	Acceptable int
	Rejected   int
}

// InitializeGradingParentPlugin adds the grading parent plugin to MongoDB
func InitializeGradingParentPlugin() {
	// Get the collection from MongoDB
	collection := mongo.GetCollection("pluginsDB", "plugins")

	// Define the grading parent plugin with default thresholds and customizable parameters
	gradingPlugin := bson.M{
		"plugin_name": "GradingPlugin",
		"type":        "parent",
		"sensor_name": "GradingSensor",
		"parameters": bson.M{
			"qualifiedThreshold":  0,   // Threshold for qualified husks
			"acceptableThreshold": 0,   // Threshold for acceptable husks
			"rejectedThreshold":   0,   // Threshold for rejected husks
			"userRequirement":     100, // Default user requirement
		},
		"customizable": []string{
			"userRequirement",
		},
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	// Insert the grading plugin if it doesn't exist
	filter := bson.M{"plugin_name": "GradingPlugin"}
	update := bson.M{
		"$setOnInsert": gradingPlugin,
	}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatalf("Failed to insert grading plugin: %v", err)
	}

	log.Println("Grading parent plugin inserted or already exists in MongoDB!")
}

func ExecuteGradingPlugin(plugin Plugin, executionCount int, userRequirement int) (bool, string) {
	totalQualified := 0
	totalAcceptable := 0

	for totalQualified+totalAcceptable < userRequirement {
		// Simulate grading a batch
		result := ExecutionResult{
			Qualified:  60, // Example count per batch
			Acceptable: 30,
			Rejected:   10,
		}

		// Update totals
		totalQualified += result.Qualified
		totalAcceptable += result.Acceptable

		// Store execution result in MongoDB
		collection := mongo.MongoClient.Database("pluginsDB").Collection("executions")
		execution := bson.M{
			"workflowId": "workflowId_example", // Replace with actual workflow ID
			"pluginName": plugin.PluginName,
			"status":     "completed",
			"result":     result,
			"timestamp":  time.Now(),
		}
		_, err := collection.InsertOne(context.Background(), execution)
		if err != nil {
			log.Printf("Failed to insert execution result: %v", err)
		}

		// Check if user wants to continue
		// This can be handled via a user input mechanism or a callback
		// For simplicity, let's assume we continue until the requirement is met
	}

	if (totalQualified + totalAcceptable) >= userRequirement {
		return true, "User requirement met. Moving to next plugin."
	}

	return false, "User requirement not met. Need to process another batch or decide to proceed."
}
