package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"Coconut-Peat-Supply-chain_core_system/pkg/mongo"
	"Coconut-Peat-Supply-chain_core_system/proto"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type GradingPluginServer struct {
	proto.UnimplementedPluginServiceServer
}

// Register registers the grading plugin in MongoDB
func (s *GradingPluginServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")
	// Check if the grading plugin is already registered
	filter := bson.M{"plugin_name": "GradingPlugin"}
	var existingPlugin bson.M
	err := collection.FindOne(ctx, filter).Decode(&existingPlugin)
	if err == nil {
		// Plugin already exists
		log.Println("GradingPlugin already registered in MongoDB")
		return &proto.RegisterResponse{Success: false, Message: "GradingPlugin is already registered"}, nil
	}
	gradingPlugin := bson.M{
		"plugin_name": "GradingPlugin",
		"type":        "parent",
		"sensor_name": "GradingSensor",
		"version":     "1.0.0",
		"parameters": bson.M{
			"qualifiedThreshold":  0,
			"acceptableThreshold": 0,
			"rejectedThreshold":   0,
			"userRequirement":     0,
		},
		"customizable": []string{"userRequirement"},
		"created_at":   time.Now(),
		"updated_at":   time.Now(),
	}

	_, err = collection.InsertOne(ctx, gradingPlugin)
	if err != nil {
		log.Printf("Failed to register grading plugin: %v", err)
		return &proto.RegisterResponse{Success: false, Message: "Failed to register grading plugin"}, nil
	}

	log.Println("GradingPlugin registered successfully in MongoDB")
	return &proto.RegisterResponse{Success: true, Message: "GradingPlugin registered successfully"}, nil
}

// CreateChildPlugin creates a child plugin with customized parameters
func (s *GradingPluginServer) CreateChildPlugin(ctx context.Context, req *proto.CreateChildPluginRequest) (*proto.CreateChildPluginResponse, error) {
	collection := mongo.MongoClient.Database("pluginDB").Collection("plugins")

	parentPluginName := req.ParentPluginName
	customizedParams := req.CustomizedParameters

	// Construct the child plugin by inheriting parameters from parent
	childPlugin := bson.M{
		"plugin_name":   parentPluginName + "_child",
		"type":          "child",
		"sensor_name":   "GradingSensor",
		"parameters":    bson.M{"userRequirement": customizedParams["userRequirement"]},
		"parent_plugin": parentPluginName,
		"created_at":    time.Now(),
		"updated_at":    time.Now(),
	}

	// Insert child plugin into MongoDB
	_, err := collection.InsertOne(ctx, childPlugin)
	if err != nil {
		log.Printf("Failed to create child plugin: %v", err)
		return &proto.CreateChildPluginResponse{Success: false, Message: "Failed to create child plugin"}, nil
	}

	log.Println("Child plugin created successfully in MongoDB")
	return &proto.CreateChildPluginResponse{Success: true, Message: "Child plugin created successfully"}, nil
}

// ExecuteGrading performs the grading action with simulated data
func (s *GradingPluginServer) ExecuteGrading(ctx context.Context, req *proto.ExecuteActionRequest) (*proto.ExecuteActionResponse, error) {
	userRequirementStr, ok := req.Parameters["userRequirement"]
	if !ok {
		return &proto.ExecuteActionResponse{
			Success: false,
			Message: "userRequirement parameter missing",
		}, nil
	}
	// Convert userRequirementStr to an int
	userRequirement, err := strconv.Atoi(userRequirementStr)
	if err != nil {
		return nil, errors.New("userRequirement must be an integer")
	}
	// Simulated grading results
	qualified := 60
	acceptable := 30
	rejected := 10

	totalCount := qualified + acceptable

	message := "Grading completed successfully"
	success := true
	if totalCount < userRequirement {
		message = "Order another batch or decide to process"
		success = false
	}

	return &proto.ExecuteActionResponse{
		Success: success,
		Message: message,
		Results: map[string]string{"qualified": strconv.Itoa(qualified), "acceptable": strconv.Itoa(acceptable), "rejected": strconv.Itoa(rejected)},
	}, nil
}

// start the grading plugin
func main() {
	port := flag.String("port", "50052", "Port for the gRPC server")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *port, err)
	}

	mongo.ConnectMongoDB()
	grpcServer := grpc.NewServer()
	proto.RegisterPluginServiceServer(grpcServer, &GradingPluginServer{})
	log.Printf("GradingPlugin gRPC server listening on port %s", *port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
