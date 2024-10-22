package main

import (
	"context"
	"log"
	"time"

	"Coconut-Peat-Supply-chain_core_system/proto"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewPluginServiceClient(conn)

	// Simulate sending a child plugin configuration
	registerReq := &proto.RegisterPluginRequest{
		ParentPluginName: "GradingPlugin",
		CustomizedParameters: map[string]string{
			"userRequirementCount": "150",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	registerResp, err := client.RegisterChildPlugin(ctx, registerReq)
	if err != nil {
		log.Fatalf("could not register plugin: %v", err)
	}
	log.Printf("RegisterChildPlugin Response: %s", registerResp.Message)

	// Execute the grading plugin
	executeReq := &proto.ExecuteGradingRequest{
		ExecutionCount:  1,
		UserRequirement: 150,
	}

	executeResp, err := client.ExecuteGrading(ctx, executeReq)
	if err != nil {
		log.Fatalf("could not execute grading: %v", err)
	}
	log.Printf("ExecuteGrading Response: %s", executeResp.Message)
}
