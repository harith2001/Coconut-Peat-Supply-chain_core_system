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
	log.Printf("Plugin Action: ", pluginName, action)
	// Register Plugin
	if action == "register" {
		_, err := s.registry.RegisterPlugin(pluginName, "50052") // Use dynamic port if necessary
		if err != nil {
			return &proto.PluginActionResponse{Success: false, Message: err.Error()}, nil
		}
		return &proto.PluginActionResponse{Success: true, Message: "Plugin registered successfully"}, nil
	}
	//create child plugin with customized parameters

	//have to write the execute action function
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
