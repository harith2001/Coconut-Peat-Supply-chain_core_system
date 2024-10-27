package main

import (
	"Coconut-Peat-Supply-chain_core_system/pkg/grpc"
	"Coconut-Peat-Supply-chain_core_system/pkg/mongo"
)

func main() {

	//connect to MongoDB
	mongo.ConnectMongoDB()

	//start gRPC server
	grpc.StartGrpcServer()
}
