package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	mongo "grading/config/db"
	sensor "grading/config/sensor"
	"grading/proto"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type GradingPluginServer struct {
	proto.UnimplementedPluginServer
}

// create the grading plugin in MongoDB
func storePluginDetails() error {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")
	steps := []string{
		"Unload all husk batch",
		"Starting the grading sensor",
		"Grading the husk based the color using the sensor",
		"Sorting the husk based on the color (qualified, acceptable, rejected)",
		"Counting the total usable husk (qualified + acceptable)",
		"Checking if the total usable husk is equal to the user requirement",
		"if the total usable husk is less than the user requirement, order another batch or decide to process",
	}
	pluginDetails := bson.M{
		"plugin_name":     "grading",
		"senosor_name":    "image_module",
		"userRequirement": "",
		"status":          true,
		"process":         "not",
		"steps":           steps,
		"created_at":      time.Now(),
		"updated_at":      time.Now(),
	}
	// Check if a plugin with the same name already exists
	filter := bson.M{"plugin_name": pluginDetails["plugin_name"]}
	var existingPlugin bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&existingPlugin)
	if err == nil {
		return fmt.Errorf("plugin with name '%s' already exists", pluginDetails["plugin_name"])
	}

	// Insert the new plugin details
	_, err = collection.InsertOne(context.Background(), pluginDetails)
	if err != nil {
		return fmt.Errorf("error storing plugin details: %v", err)
	}

	log.Println("Plugin details stored successfully")
	return nil
}

// Register registers the grading plugin in MongoDB
func (s *GradingPluginServer) RegisterPlugin(ctx context.Context, req *proto.PluginRequest) (*proto.PluginResponse, error) {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")
	// Check if the grading plugin is already registered
	filter := bson.M{"plugin_name": req.PluginName, "process": "registered"}
	var existingPlugin bson.M
	err := collection.FindOne(ctx, filter).Decode(&existingPlugin)
	if err == nil {
		return &proto.PluginResponse{Success: false, Message: "Plugin is already registered"}, nil
	}
	//if not registered, register the plugin
	userRequirement := req.UserRequirement
	plugin := bson.M{
		"$set": bson.M{
			"userRequirement": userRequirement,
			"process":         "registered",
			"updated_at":      time.Now(),
		},
	}
	update := bson.M{"plugin_name": req.PluginName}
	_, err = collection.UpdateOne(ctx, update, plugin)
	if err != nil {
		log.Printf("Failed to register plugin: %v", err)
		return &proto.PluginResponse{Success: false, Message: "Failed to register plugin"}, err
	}

	return &proto.PluginResponse{Success: true, Message: "Plugin registered successfully"}, nil
}

// execute the grading plugin
func (s *GradingPluginServer) ExecutePlugin(ctx context.Context, req *proto.PluginExecute) (*proto.ExecutionStatus, error) {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")
	filter := bson.M{"plugin_name": req.PluginName}
	var plugin bson.M
	err := collection.FindOne(ctx, filter).Decode(&plugin)
	if err != nil {
		return &proto.ExecutionStatus{Success: false, Message: "Plugin not found"}, err
	}

	if !plugin["status"].(bool) {
		return &proto.ExecutionStatus{Success: false, Message: "Plugin is deactivated"}, nil
	}

	if plugin["process"].(string) == "completed" {
		return &proto.ExecutionStatus{Success: false, Message: "Plugin grading already completed"}, nil
	}

	if plugin["process"].(string) != "registered" {
		return &proto.ExecutionStatus{Success: false, Message: "Plugin is not registered"}, nil
	}

	userRequirementStr, ok := plugin["userRequirement"].(string)
	if !ok {
		return &proto.ExecutionStatus{Success: false, Message: "Failed to get userRequirement"}, nil
	}

	userRequirement, err := strconv.Atoi(userRequirementStr)
	if err != nil {
		return &proto.ExecutionStatus{Success: false, Message: "userRequirement must be an integer"}, err
	}

	// **Use MQTT data instead of hardcoded values**
	totalCount := sensor.Qualified + sensor.Acceptable
	message := "Grading completed successfully"
	success := true

	// Check if the total count is less than the user requirement
	if totalCount < userRequirement {
		message = "Order another batch or decide to process"
		success = false
	}
	// Check if the total count is greater than the user requirement
	if totalCount > userRequirement {
		message = "User requirement exceeded"
		success = true
	}

	// Update the plugin status in MongoDB
	update := bson.M{
		"$set": bson.M{
			"process": "completed",
			"results": map[string]interface{}{
				"total":      totalCount,
				"qualified":  sensor.Qualified,
				"acceptable": sensor.Acceptable,
				"rejected":   sensor.Rejected,
			},
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return &proto.ExecutionStatus{Success: false, Message: "Failed to update plugin"}, err
	}

	return &proto.ExecutionStatus{Success: success, Message: message, Results: map[string]string{
		"qualified":       strconv.Itoa(sensor.Qualified),
		"acceptable":      strconv.Itoa(sensor.Acceptable),
		"rejected":        strconv.Itoa(sensor.Rejected),
		"total":           strconv.Itoa(totalCount),
		"userRequirement": strconv.Itoa(userRequirement)}}, nil
}

// start the grading plugin
func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mongo.ConnectMongoDB()
	proto.RegisterPluginServer(grpcServer, &GradingPluginServer{})
	storePluginDetails()
	go sensor.StartSensorSubscriber() //sensor connection

	log.Println("gRPC server is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
