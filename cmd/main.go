package main

import (
	"Coconut-Peat-Supply-chain_core_system/pkg/grpc"
	"Coconut-Peat-Supply-chain_core_system/pkg/mongo"
	"Coconut-Peat-Supply-chain_core_system/pkg/plugin"
)

func main() {

	//connect to MongoDB
	mongo.ConnectMongoDB()

	//initialize grading parent plugin
	plugin.InitializeGradingParentPlugin()

	//start gRPC server
	grpc.StartGrpcServer()
}
