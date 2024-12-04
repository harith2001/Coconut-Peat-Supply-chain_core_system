package server

import (
	pb "Coconut-Peat-Supply-chain_core_system/proto"
	"context"
	"log"
	"net"
	"strconv"

	mongo "Coconut-Peat-Supply-chain_core_system/config/db"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	pb.UnimplementedMainServiceServer
}

func (s *Server) ClientFunction(ctx context.Context, req *pb.ClientRequest) (*pb.ClientResponse, error) {

	//get the plugin name and the plugin port number from the mongodb
	collection := mongo.MongoClient.Database("portDB").Collection("port")
	filter := bson.D{
		{Key: "plugin", Value: req.PluginName},
		{Key: "status", Value: true},
	}
	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatalf("Error while fetching the plugin details: %v", err)
	}
	pluginPort := strconv.Itoa(int(result["port"].(int32)))

	//connecting the plugin
	address := "0.0.0.0:" + pluginPort
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to backend service: %v", err)
	}
	defer conn.Close()

	// create a gRPC client for the backend service
	backendClient := pb.NewPluginClient(conn)

	//decide which action and call the backend service
	action := req.Action
	if action == "register" {
		backendResp, err := backendClient.RegisterPlugin(ctx, &pb.PluginRequest{
			PluginName:      req.PluginName,
			UserRequirement: req.UserRequirement,
		})
		if err != nil {
			return nil, err
		}
		return &pb.ClientResponse{
			Success: backendResp.Success,
			Message: backendResp.Message,
		}, nil

	} else if action == "execute" {
		backendResp, err := backendClient.ExecutePlugin(ctx, &pb.PluginExecute{
			PluginName: req.PluginName,
		})
		if err != nil {
			return nil, err
		}

		// Return the response to the client
		return &pb.ClientResponse{
			Success: backendResp.Success,
			Message: backendResp.Message,
			Results: backendResp.Results,
		}, nil
	} else if action == "unregister" {
		backendResp, err := backendClient.UnregisterPlugin(ctx, &pb.PluginUnregister{
			PluginName: req.PluginName,
		})
		if err != nil {
			return nil, err
		}
		// Return the response to the client
		return &pb.ClientResponse{
			Success: backendResp.Success,
			Message: backendResp.Message,
		}, nil
	} else {
		return nil, nil
	}
}

func StartServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mongo.ConnectMongoDB()

	pb.RegisterMainServiceServer(grpcServer, &Server{})
	pb.RegisterNewPluginServiceServer(grpcServer, &NewPlugin{})

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
