package grpc

import (
	"Coconut-Peat-Supply-chain_core_system/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CoreServer struct {
	proto.UnimplementedCoreServiceServer
	registeredPlugins map[string]proto.PluginServiceClient // Registered plugins map
}

// Handles plugin actions: "register", "create_child", "execute"
func (s *CoreServer) PluginAction(ctx context.Context, req *proto.PluginActionRequest) (*proto.PluginActionResponse, error) {
	pluginName := req.PluginName
	action := req.Action

	// Register Plugin if not registered
	if action == "register" && s.registeredPlugins[pluginName] == nil {
		client, err := s.connectToPlugin(pluginName)
		if err != nil {
			return &proto.PluginActionResponse{Success: false, Message: "Failed to connect to plugin"}, nil
		}
		s.registeredPlugins[pluginName] = client
		return &proto.PluginActionResponse{Success: true, Message: "Plugin registered successfully"}, nil
	}

	// Create Child Plugin
	if action == "create_child" {
		client, ok := s.registeredPlugins[pluginName]
		if !ok {
			return &proto.PluginActionResponse{Success: false, Message: "Plugin not registered"}, nil
		}

		// Forward the create child request to the plugin
		response, err := client.Register(ctx, &proto.RegisterRequest{
			PluginName: pluginName,
		})
		if err != nil {
			return &proto.PluginActionResponse{Success: false, Message: "Failed to register child plugin"}, nil
		}

		return &proto.PluginActionResponse{Success: response.Success, Message: response.Message}, nil
	}

	// Execute Plugin Action
	if action == "execute" {
		client, ok := s.registeredPlugins[pluginName]
		if !ok {
			return &proto.PluginActionResponse{Success: false, Message: "Plugin not registered"}, nil
		}

		// Execute action with parameters
		execReq := &proto.ExecuteActionRequest{Parameters: req.Parameters}
		execRes, err := client.ExecuteAction(ctx, execReq)
		if err != nil {
			return &proto.PluginActionResponse{Success: false, Message: "Plugin execution failed"}, nil
		}

		return &proto.PluginActionResponse{Success: execRes.Success, Message: execRes.Message}, nil
	}

	return &proto.PluginActionResponse{Success: false, Message: "Invalid action"}, nil
}

// StartGrpcServer starts the gRPC server
func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// connectToPlugin initiates a gRPC connection to a plugin's independent server
func (s *CoreServer) connectToPlugin(pluginName string) (proto.PluginServiceClient, error) {
	// Assume plugins run on different ports, e.g., "localhost:50052" for GradingPlugin
	address := "localhost:50052"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return proto.NewPluginServiceClient(conn), nil
}
