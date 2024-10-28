package plugin

import (
	"Coconut-Peat-Supply-chain_core_system/pkg/mongo"
	"Coconut-Peat-Supply-chain_core_system/proto"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

// PluginDetails represents the structure stored in MongoDB
type PluginDetails struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"plugin_name"`
	Port      string             `bson:"port"`
	CreatedAt time.Time          `bson:"created_at"`
}

// PluginRegistry manages the plugins in MongoDB.
type PluginRegistry struct {
	collection *mongoDriver.Collection
}

// NewPluginRegistry initializes the MongoDB plugin collection.
func NewPluginRegistry() *PluginRegistry {
	collection := mongo.MongoClient.Database("pluginDB").Collection("pluginRegsitry")
	return &PluginRegistry{collection: collection}
}

// RegisterPlugin adds a new plugin to MongoDB if not already registered.
func (pr *PluginRegistry) RegisterPlugin(pluginName, port string) (proto.PluginServiceClient, error) {
	log.Printf("Plugin Action 2: ", pluginName, port)
	filter := bson.M{"plugin_name": pluginName}
	var existingPlugin PluginDetails
	err := pr.collection.FindOne(context.TODO(), filter).Decode(&existingPlugin)
	if err == nil {
		get, err := pr.GetPlugin(pluginName)
		if err != nil {
			return nil, err
		}
		messg := fmt.Errorf("plugin %s already registered", pluginName)
		return get, messg
	}
	// Insert new plugin into MongoDB
	newPlugin := PluginDetails{
		Name:      pluginName,
		Port:      port,
		CreatedAt: time.Now(),
	}
	_, err = pr.collection.InsertOne(context.TODO(), newPlugin)
	if err != nil {
		return nil, fmt.Errorf("failed to register plugin %s", pluginName)
	}

	// Connect to the plugin
	client, err := pr.GetPlugin(pluginName)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetPlugin retrieves a plugin's details from MongoDB and establishes a gRPC connection.
func (pr *PluginRegistry) GetPlugin(pluginName string) (proto.PluginServiceClient, error) {
	log.Printf("Plugin Action 4: ", pluginName)
	filter := bson.M{"plugin_name": pluginName}
	var plugin PluginDetails
	err := pr.collection.FindOne(context.TODO(), filter).Decode(&plugin)
	if err != nil {
		if err == mongoDriver.ErrNoDocuments {
			return nil, fmt.Errorf("plugin %s not found", pluginName)
		}
		return nil, err
	}
	log.Printf("plugin founded", pluginName)
	// Connect to the plugin on the stored port
	return connectToPlugin(plugin.Port)
}

// ListPlugins retrieves all registered plugins from MongoDB.
func (pr *PluginRegistry) ListPlugins() ([]PluginDetails, error) {
	cursor, err := pr.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var plugins []PluginDetails
	if err := cursor.All(context.TODO(), &plugins); err != nil {
		return nil, err
	}
	return plugins, nil
}

// add : should be able to start any plugin server with name and port
// connectToPlugin establishes a gRPC connection to a plugin server.
func connectToPlugin(port string) (proto.PluginServiceClient, error) {

	// If the connection failed, attempt to start the plugin server
	log.Printf("Starting Grading Plugin server on port %s", port)
	cmd := exec.Command("go", "run", "./grading/grading_plugin.go", "-port", port)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	// Start the plugin process
	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start plugin server on port %s: %v", port, err)
	}
	log.Printf("Plugin server process started, waiting for it to initialize...")

	// Allow time for the server to start
	time.Sleep(5 * time.Second)

	// Attempt to connect to the plugin
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to plugin on port %s: %v", port, err)
	}

	client := proto.NewPluginServiceClient(conn)
	return client, nil
}
