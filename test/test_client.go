package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"Coconut-Peat-Supply-chain_core_system/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the plugin server
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to plugin server: %v", err)
	}
	defer conn.Close()

	client := proto.NewPluginServiceClient(conn)

	// Make a test call, such as to the Register function
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &proto.RegisterRequest{PluginName: "GradingPlugin"}
	res, err := client.Register(ctx, req)
	if err != nil {
		log.Fatalf("Error calling Register: %v", err)
	}

	fmt.Printf("Response from Register: %v\n", res)
}
