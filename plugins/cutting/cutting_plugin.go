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
	collection := mongo.MongoClient.Database("test").Collection("plugins")
	steps := []string{
		"Insert the graded husk batch",
		"Start the cutting machine",
		"Check if all the graded husk are processed/cut",
	}
	pluginDetails := bson.M{
		"plugin_name":     "cutting",
		"senosor_name":    "null",
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

// Register registers the cutting plugin in MongoDB
func (s *CuttingPluginServer) RegisterPlugin(ctx context.Context, req *proto.PluginRequest) (*proto.PluginResponse, error) {
	collection := mongo.MongoClient.Database("test").Collection("plugins")

	// Check if a plugin with workflow_id as null exists
	filter := bson.M{"plugin_name": req.PluginName, "workflow_id": "null"}
	var existingPlugin bson.M
	err := collection.FindOne(ctx, filter).Decode(&existingPlugin)

	if err == nil {
		// If a plugin with workflow_id as "null" exists, update it
		userRequirement := req.UserRequirement
		if userRequirement == "" {
			userRequirement = "0"
		}

		update := bson.M{
			"$set": bson.M{
				"userRequirement": userRequirement,
				"workflow_id":     req.WorkflowId,
				"status":          true,
				"process":         "registered",
				"updated_at":      time.Now(),
			},
		}

		_, err = collection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Printf("Failed to update cutting plugin: %v", err)
			return &proto.PluginResponse{Success: false, Message: "Failed to update existing cutting plugin"}, err
		}
		return &proto.PluginResponse{Success: true, Message: "Existing cutting plugin updated successfully"}, nil
	}

	// If no plugin with workflow_id as "null" exists, create a new one
	err = storePluginDetails()
	if err != nil {
		log.Printf("Failed to create new cutting plugin: %v", err)
		return &proto.PluginResponse{Success: false, Message: "Failed to create new cutting plugin"}, err
	}

	// Now, update the newly created plugin with the workflow details
	newFilter := bson.M{"plugin_name": req.PluginName, "workflow_id": "null"}
	userRequirement := req.UserRequirement
	if userRequirement == "" {
		userRequirement = "0"
	}

	update := bson.M{
		"$set": bson.M{
			"userRequirement": userRequirement,
			"workflow_id":     req.WorkflowId,
			"status":          true,
			"process":         "registered",
			"updated_at":      time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, newFilter, update)
	if err != nil {
		log.Printf("Failed to update new cutting plugin: %v", err)
		return &proto.PluginResponse{Success: false, Message: "Failed to update new cutting plugin"}, err
	}

	return &proto.PluginResponse{Success: true, Message: "New cutting plugin created and registered successfully"}, nil
}

// execute grading
func (s *CuttingPluginServer) ExecutePlugin(ctx context.Context, req *proto.PluginExecute) (*proto.ExecutionStatus, error) {
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
		return &proto.ExecutionStatus{Success: false, Message: "Plugin cutting already completed"}, nil
	}

	if plugin["process"].(string) != "registered" {
		return &proto.ExecutionStatus{Success: false, Message: "Plugin is not registered"}, nil
	}

	filter = bson.M{"plugin_name": "grading", "workflow_id": req.WorkflowId}
	var gradingPlugin bson.M
	err = collection.FindOne(ctx, filter).Decode(&gradingPlugin)
	if err != nil {
		return &proto.ExecutionStatus{Success: false, Message: "Cutting plugin not found"}, err
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
	filter = bson.M{"plugin_name": req.PluginName, "workflow_id": req.WorkflowId}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return &proto.ExecutionStatus{Success: false, Message: "Failed to update plugin"}, err
	}

	return &proto.ExecutionStatus{Success: success, Message: message, Results: map[string]string{
		"total_count": strconv.Itoa(int(totalCount))}}, nil
}

// UnregisterPlugin deactivates
func (s *CuttingPluginServer) UnregisterPlugin(ctx context.Context, req *proto.PluginUnregister) (*proto.UnregisterResponse, error) {
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

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mongo.ConnectMongoDB()
	proto.RegisterPluginServer(grpcServer, &CuttingPluginServer{})
	storePluginDetails() //store mongoDB plugin details

	log.Println("gRPC server is running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
