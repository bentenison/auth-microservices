package main

import (
	"grpc-auth-app/auth-server/internal/handler"
	"grpc-auth-app/auth-server/pkg/config"
	"grpc-auth-app/auth-server/pkg/di"
)

func main() {
	//initialize configurations
	InitializeConfigurations()
	svc := di.BuildContainer()
	go handler.StartGrpcServer(svc)
	handler.StartHttpServer(svc)
}

func InitializeConfigurations() {
	config.InitConfig()
}
