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
	collection := mongo.MongoClient.Database("test").Collection("plugins")
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
		"senosor_name":    "grading/sensor_data",
		"userRequirement": "",
		"workflow_id":     "null",
		"status":          true,
		"process":         "not",
		"steps":           steps,
		"created_at":      time.Now(),
		"updated_at":      time.Now(),
	}
	// Insert the new plugin details
	var err error
	_, err = collection.InsertOne(context.Background(), pluginDetails)
	if err != nil {
		return fmt.Errorf("error storing plugin details: %v", err)
	}

	log.Println("Plugin details stored successfully")
	return nil
}

// RegisterPlugin registers the grading plugin in MongoDB
func (s *GradingPluginServer) RegisterPlugin(ctx context.Context, req *proto.PluginRequest) (*proto.PluginResponse, error) {
	collection := mongo.MongoClient.Database("test").Collection("plugins")

	// Check if a plugin with workflow_id as null exists
	filter := bson.M{"plugin_name": req.PluginName, "workflow_id": "null"}
	var existingPlugin bson.M
	err := collection.FindOne(ctx, filter).Decode(&existingPlugin)

	if err == nil {
		// If a plugin with workflow_id as "null" exists, update it
		update := bson.M{
			"$set": bson.M{
				"userRequirement": req.UserRequirement,
				"workflow_id":     req.WorkflowId,
				"process":         "registered",
				"updated_at":      time.Now(),
			},
		}
		_, err = collection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Printf("Failed to update plugin: %v", err)
			return &proto.PluginResponse{Success: false, Message: "Failed to update existing plugin"}, err
		}
		return &proto.PluginResponse{Success: true, Message: "Existing plugin updated successfully"}, nil
	}

	// If no plugin with workflow_id as "null" exists, create a new one
	err = storePluginDetails()
	if err != nil {
		log.Printf("Failed to create new plugin: %v", err)
		return &proto.PluginResponse{Success: false, Message: "Failed to create new plugin"}, err
	}

	// Now, update the newly created plugin with the workflow details
	newFilter := bson.M{"plugin_name": req.PluginName, "workflow_id": "null"}
	update := bson.M{
		"$set": bson.M{
			"userRequirement": req.UserRequirement,
			"workflow_id":     req.WorkflowId,
			"process":         "registered",
			"updated_at":      time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, newFilter, update)
	if err != nil {
		log.Printf("Failed to update new plugin: %v", err)
		return &proto.PluginResponse{Success: false, Message: "Failed to update new plugin"}, err
	}

	return &proto.PluginResponse{Success: true, Message: "New plugin created and registered successfully"}, nil
}

// execute the grading plugin
func (s *GradingPluginServer) ExecutePlugin(ctx context.Context, req *proto.PluginExecute) (*proto.ExecutionStatus, error) {
	collection := mongo.MongoClient.Database("test").Collection("plugins")
	filter := bson.M{"plugin_name": req.PluginName, "workflow_id": req.WorkflowId}
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

// UnregisterPlugin deactivates
func (s *GradingPluginServer) UnregisterPlugin(ctx context.Context, req *proto.PluginUnregister) (*proto.UnregisterResponse, error) {
	collection := mongo.MongoClient.Database("test").Collection("plugins")
	filter := bson.M{"plugin_name": req.PluginName, "workflow_id": req.WorkflowId}
	update := bson.M{
		"$set": bson.M{
			"status":     false,
			"updated_at": time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return &proto.UnregisterResponse{Success: false, Message: "Failed to unregister plugin"}, err
	}

	return &proto.UnregisterResponse{Success: true, Message: "Plugin unregistered successfully"}, nil
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
