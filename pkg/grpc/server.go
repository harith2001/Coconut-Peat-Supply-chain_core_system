package grpc

import (
	"Coconut-Peat-Supply-chain_core_system/pkg/plugin"
	"Coconut-Peat-Supply-chain_core_system/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type PluginServer struct {
	proto.UnimplementedPluginServiceServer
}

// RegisterChildPlugin handles the gRPC request to register a new child plugin
func (s *PluginServer) RegisterChildPlugin(ctx context.Context, req *proto.RegisterPluginRequest) (*proto.PluginResponse, error) {
	// Pass request to the Plugin Manager to create a child plugin
	success, message := plugin.CreateChildPlugin(req.ParentPluginName, req.CustomizedParameters)

	return &proto.PluginResponse{
		Success: success,
		Message: message,
	}, nil
}

// ExecuteGrading handles the grading process
func (s *PluginServer) ExecuteGrading(ctx context.Context, req *proto.ExecuteGradingRequest) (*proto.ExecuteGradingResponse, error) {
	// Retrieve the grading plugin
	gradingPlugin, exists := plugin.ParentPlugins["GradingPlugin"]
	if !exists {
		return &proto.ExecuteGradingResponse{
			Success: false,
			Message: "GradingPlugin not found",
		}, nil
	}
	// Execute grading logic
	passed, message := plugin.ExecuteGradingPlugin(gradingPlugin, int(req.ExecutionCount), int(req.UserRequirement))

	return &proto.ExecuteGradingResponse{
		Success: passed,
		Message: message,
	}, nil
}

// StartGrpcServer starts the gRPC server
func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterPluginServiceServer(grpcServer, &PluginServer{})

	log.Printf("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
