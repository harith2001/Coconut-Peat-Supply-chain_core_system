package main

import (
	"context"
	"testing"

	pb "github.com/harith2001/Coconut-Peat-Supply-chain_core_system/proto" // Adjust package path

	"google.golang.org/grpc"
)

func BenchmarkClientFunction(b *testing.B) {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMainServiceClient(conn)

	req := &pb.ClientRequest{
		PluginName:      "grading",
		UserRequirement: "100",
		Action:          "execute",
	}

	b.ResetTimer() // Reset timer to measure only function execution time

	for i := 0; i < b.N; i++ {
		_, err := client.ClientFunction(context.Background(), req)
		if err != nil {
			b.Errorf("gRPC request failed: %v", err)
		}
	}
}
