package main

import (
	"context"
	"log"
	"net"
	"sync"

	pluginpb "Coconut-Peat-Supply-chain_core_system/plugins/grading/proto"
	pb "Coconut-Peat-Supply-chain_core_system/proto"

	"google.golang.org/grpc"
)

// PluginInfo stores information about a plugin.
type PluginInfo struct {
	Hostname string
	Port     string
}

// PluginRegistry stores all registered plugins.
type PluginRegistry struct {
	plugins map[string]PluginInfo // plugin_name -> PluginInfo
	mutex   sync.RWMutex
}

// Server is the core server that implements the PluginManager service.
type Server struct {
	pb.UnimplementedPluginManagerServer
	registry *PluginRegistry
}

// LoadPlugins loads hardcoded plugins into the registry.
func (r *PluginRegistry) LoadPlugins() {
	r.plugins = map[string]PluginInfo{
		"grading": {Hostname: "localhost", Port: "50052"},
		"PluginB": {Hostname: "localhost", Port: "50053"},
	}
}

// RegisterPlugin registers a new plugin.
func (s *Server) RegisterPlugin(ctx context.Context, req *pb.Pluginrequest) (*pb.Pluginresponse, error) {
	s.registry.mutex.Lock()
	defer s.registry.mutex.Unlock()

	if _, exists := s.registry.plugins[req.PluginName]; exists {
		return &pb.Pluginresponse{
			Success: false,
			Message: "Plugin already registered",
		}, nil
	}

	if pluginInfo, ok := s.registry.plugins[req.PluginName]; ok {
		s.registry.plugins[req.PluginName] = pluginInfo
		return &pb.Pluginresponse{
			Success: true,
			Message: "Plugin registered successfully",
		}, nil
	}

	return &pb.Pluginresponse{
		Success: true,
		Message: "Plugin registered successfully",
	}, nil
}

// ExecutePlugin forwards the request to the specific plugin.
func (s *Server) ExecutePlugin(ctx context.Context, req *pb.Pluginexecute) (*pb.Executionstatus, error) {
	s.registry.mutex.RLock()
	pluginInfo, exists := s.registry.plugins[req.PluginName]
	s.registry.mutex.RUnlock()

	if !exists {
		return &pb.Executionstatus{
			Success: false,
			Message: "Plugin not found",
		}, nil
	}

	address := pluginInfo.Hostname + ":" + pluginInfo.Port

	// Connect to the plugin using gRPC
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return &pb.Executionstatus{
			Success: false,
			Message: "Failed to connect to plugin: " + err.Error(),
		}, nil
	}
	defer conn.Close()

	client := pluginpb.NewGradingPluginClient(conn)
	pluginReq := &pluginpb.PluginExecute{
		PluginName: req.PluginName,
	}
	status, err := client.ExecutePlugin(ctx, pluginReq)
	if err != nil {
		return &pb.Executionstatus{
			Success: false,
			Message: "Error executing plugin: " + err.Error(),
		}, nil
	}

	return &pb.Executionstatus{
		Success: status.Success,
		Message: status.Message,
		Results: status.Results,
	}, nil
}

// UnregisterPlugin removes a plugin from the registry.
func (s *Server) UnregisterPlugin(ctx context.Context, req *pb.PluginUnRegister) (*pb.UnRegisterResponse, error) {
	s.registry.mutex.Lock()
	defer s.registry.mutex.Unlock()

	if _, exists := s.registry.plugins[req.PluginName]; !exists {
		return &pb.UnRegisterResponse{
			Success: false,
			Message: "Plugin not found",
		}, nil
	}

	delete(s.registry.plugins, req.PluginName)
	return &pb.UnRegisterResponse{
		Success: true,
		Message: "Plugin unregistered successfully",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pluginRegistry := &PluginRegistry{
		plugins: make(map[string]PluginInfo),
	}
	pluginRegistry.LoadPlugins()

	server := &Server{
		registry: pluginRegistry,
	}

	pb.RegisterPluginManagerServer(grpcServer, server)

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
