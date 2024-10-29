package grpc

import (
	"Coconut-Peat-Supply-chain_core_system/pkg/plugin"
	"Coconut-Peat-Supply-chain_core_system/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type CoreServer struct {
	proto.UnimplementedCoreServiceServer
	registry *plugin.PluginRegistry
}

// NewCoreServer initializes CoreServer with a MongoDB-backed PluginRegistry.
func NewCoreServer() *CoreServer {
	return &CoreServer{
		registry: plugin.NewPluginRegistry(),
	}
}

func (s *CoreServer) PluginAction(ctx context.Context, req *proto.PluginActionRequest) (*proto.PluginActionResponse, error) {
	pluginName := req.PluginName
	action := req.Action
	log.Printf("Plugin Action: %s, Action: %s", pluginName, action)

	// Check if the action is "register"
	if action == "register" {
		// Register Plugin in plugin registry
		_, err := s.registry.RegisterPlugin(pluginName, "50052") // Replace with dynamic port if necessary
		if err != nil {
			return &proto.PluginActionResponse{Success: false, Message: err.Error()}, nil
		}
		log.Println("Plugin registered in registry successfully")

		// Connect to the PluginService to call Register gRPC method
		conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
		if err != nil {
			return &proto.PluginActionResponse{Success: false, Message: "Failed to connect to PluginService"}, nil
		}
		defer conn.Close()
		client := proto.NewPluginServiceClient(conn)

		// Call the Register method of PluginService
		registerResp, err := client.Register(ctx, &proto.RegisterRequest{PluginName: pluginName})
		if err != nil || !registerResp.Success {
			return &proto.PluginActionResponse{Success: false, Message: "Failed to register plugin in PluginService"}, nil
		}

		// After Register, call CreateChildPlugin with custom parameters if needed
		createChildReq := &proto.CreateChildPluginRequest{
			ParentPluginName:     pluginName,
			CustomizedParameters: req.Parameters, // Pass any custom parameters from the request
		}
		createChildResp, err := client.CreateChildPlugin(ctx, createChildReq)
		if err != nil || !createChildResp.Success {
			return &proto.PluginActionResponse{Success: false, Message: "Failed to create child plugin"}, nil
		}

		return &proto.PluginActionResponse{Success: true, Message: "Plugin registered and child plugin created successfully"}, nil
	}

	// Handle invalid actions
	return &proto.PluginActionResponse{Success: false, Message: "Invalid action"}, nil
}

// Start Grpc of the core
func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	CoreServer := NewCoreServer()

	proto.RegisterCoreServiceServer(grpcServer, CoreServer)

	log.Printf("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
