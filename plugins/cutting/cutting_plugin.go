package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	mongo "cutting/config/db"
	"cutting/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type CuttingPluginServer struct {
	proto.UnimplementedPluginServer
}

// create the cutting plugin in MongoDB
func storePluginDetails() error {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")
	steps := []string{
		"Insert the graded husk batch",
		"Start the cutting machine",
		"Check if all the graded husk are processed/cut",
	}
	pluginDetails := bson.M{
		"plugin_name":     "cutting",
		"senosor_name":    "null",
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

// Register registers the cutting plugin in MongoDB
func (s *CuttingPluginServer) RegisterPlugin(ctx context.Context, req *proto.PluginRequest) (*proto.PluginResponse, error) {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")
	// Check if the cutting plugin is already registered
	filter := bson.M{"plugin_name": req.PluginName, "process": "registered"}
	var existingPlugin bson.M
	err := collection.FindOne(ctx, filter).Decode(&existingPlugin)
	if err == nil {
		return &proto.PluginResponse{Success: false, Message: "Plugin is already registered"}, nil
	}
	userRequirement := req.UserRequirement
	if userRequirement == "" {
		userRequirement = "0"
	}
	plugin := bson.M{
		"$set": bson.M{
			"userRequirement": userRequirement,
			"status":          true,
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

// execute grading
func (s *CuttingPluginServer) ExecutePlugin(ctx context.Context, req *proto.PluginExecute) (*proto.ExecutionStatus, error) {
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

	filter = bson.M{"plugin_name": "grading"}
	var gradingPlugin bson.M
	err = collection.FindOne(ctx, filter).Decode(&gradingPlugin)
	if err != nil {
		return &proto.ExecutionStatus{Success: false, Message: "Grading plugin not found"}, err
	}
	//get the total count of the grading plugin
	results := gradingPlugin["results"].(primitive.M)
	totalCount := results["total"].(int32)

	message := fmt.Sprintf("check if all the %d husks are cut", totalCount)
	success := true

	message = fmt.Sprintf("Cutting Completed Successfully")
	success = true

	// Update the cutting plugin status in MongoDB
	update := bson.M{
		"$set": bson.M{
			"process":    "completed",
			"results":    map[string]interface{}{"totalCount": totalCount},
			"updated_at": time.Now(),
		},
	}

	// Update the cutting plugin document
	filter = bson.M{"plugin_name": req.PluginName}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return &proto.ExecutionStatus{Success: false, Message: "Failed to update plugin"}, err
	}

	return &proto.ExecutionStatus{Success: success, Message: message, Results: map[string]string{
		"total_count": strconv.Itoa(int(totalCount))}}, nil
}

// UnregisterPlugin deactivates the grading plugin
func (s *CuttingPluginServer) UnregisterPlugin(ctx context.Context, req *proto.PluginUnregister) (*proto.UnregisterResponse, error) {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")
	filter := bson.M{"plugin_name": req.PluginName}
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

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mongo.ConnectMongoDB()
	proto.RegisterPluginServer(grpcServer, &CuttingPluginServer{})
	storePluginDetails()

	log.Println("gRPC server is running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
