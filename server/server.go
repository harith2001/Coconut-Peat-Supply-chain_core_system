package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "github.com/harith2001/Coconut-Peat-Supply-chain_core_system/proto"

	mongo "github.com/harith2001/Coconut-Peat-Supply-chain_core_system/config/db"
	sensor "github.com/harith2001/Coconut-Peat-Supply-chain_core_system/config/sensor"

	_ "net/http/pprof"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedMainServiceServer
}

func (s *Server) ClientFunction(ctx context.Context, req *pb.ClientRequest) (*pb.ClientResponse, error) {

	//get the plugin name and the plugin port number and the plugin name from the mongodb
	collection := mongo.MongoClient.Database("test").Collection("port")
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
	pluginName := req.PluginName

	//connecting the plugin
	//address := "0.0.0.0:" + pluginPort //local
	//address := pluginName + ":" + pluginPort //docker
	address := fmt.Sprintf("%s-plugin-service.default.svc.cluster.local:%s", pluginName, pluginPort) //kube
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
			WorkflowId:      req.WorkflowId,
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
			WorkflowId: req.WorkflowId,
		})
		if err != nil {
			return nil, err
		}

		// If execution is successful, send data to the blockchain
		if backendResp.Success == true {
			blockchainMain(backendResp.Results) // Pass the results to the blockchain function
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
			WorkflowId: req.WorkflowId,
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
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mongo.ConnectMongoDB()

	pb.RegisterMainServiceServer(grpcServer, &Server{})
	pb.RegisterNewPluginServiceServer(grpcServer, &NewPlugin{})
	reflection.Register(grpcServer)
	//sensor connection
	go sensor.SensorMain()

	//testing purpose
	// go func() {
	// 	log.Println("Starting pprof server on :6060")
	// 	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
	// 		log.Fatalf("pprof server failed: %v", err)
	// 	}
	// }()

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
