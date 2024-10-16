package plugin

import (
	"context"
	"fmt"
	"log"

	"Coconut-Peat-Supply-chain_core_system/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type Plugin struct {
	PluginName         string            `bson:"plugin_name"`
	Type               string            `bson:"type"`
	Parameters         map[string]string `bson:"parameters"`
	CustomizableParams []string          `bson:"customizable"`
	ParentPlugin       string            `bson:"parentPlugin"`
	Version            int               `bson:"version"`
	CreatedAt          interface{}       `bson:"createdAt"`
	UpdatedAt          interface{}       `bson:"updatedAt"`
}

var ParentPlugins = map[string]Plugin{}
var ChildPlugins = []Plugin{}

// Initialize parent plugins from MongoDB
func InitPlugins() {
	log.Println("Initializing parent plugins from MongoDB")
	collection := mongo.MongoClient.Database("pluginsDB").Collection("plugins")
	cursor, err := collection.Find(context.Background(), bson.M{"type": "parent"})
	if err != nil {
		log.Fatalf("Failed to find parent plugins: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var plugin Plugin
		if err := cursor.Decode(&plugin); err != nil {
			log.Printf("Failed to decode plugin: %v", err)
			continue
		}
		ParentPlugins[plugin.PluginName] = plugin
	}

	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}

	log.Println("Initialized parent plugins from MongoDB")
}

func CreateChildPlugin(parentPluginName string, customizedParams map[string]string) (bool, string) {
	log.Printf("Attempting to create child plugin for parent: %s", parentPluginName)
	log.Printf("customized params: %s", customizedParams)
	parentPlugin, exists := ParentPlugins[parentPluginName]
	if !exists {
		return false, fmt.Sprintf("Parent plugin %s not found", parentPluginName)
	}

	// Validate customized parameters
	for key := range customizedParams {
		found := false
		for _, param := range parentPlugin.CustomizableParams {
			if param == key {
				found = true
				break
			}
		}
		if !found {
			return false, fmt.Sprintf("Parameter %s is not customizable", key)
		}
		log.Printf("Customizable parameters are valid for plugin: %s", parentPluginName)
	}

	// Create child plugin by inheriting from parent
	childPlugin := Plugin{
		PluginName:         parentPluginName + "_child",
		Type:               "child",
		Parameters:         make(map[string]string),
		CustomizableParams: parentPlugin.CustomizableParams,
		ParentPlugin:       parentPluginName,
		Version:            parentPlugin.Version,
		CreatedAt:          nil, // Handle timestamps appropriately
		UpdatedAt:          nil,
	}

	// Inherit parameters from parent
	for k, v := range parentPlugin.Parameters {
		childPlugin.Parameters[k] = v
	}

	// Override with customized parameters
	for k, v := range customizedParams {
		childPlugin.Parameters[k] = v
	}

	// Insert child plugin into MongoDB
	collection := mongo.MongoClient.Database("pluginsDB").Collection("plugins")
	_, err := collection.InsertOne(context.Background(), childPlugin)
	if err != nil {
		return false, fmt.Sprintf("Failed to insert child plugin: %v", err)
	}

	// Add to in-memory store
	ChildPlugins = append(ChildPlugins, childPlugin)

	return true, "Child plugin created successfully"
}
