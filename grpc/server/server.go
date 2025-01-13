package main

import (
	"grpc_study/model"
	"grpc_study/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cred, err := credentials.NewServerTLSFromFile("../cert/server.pem", "../cert/server.key")
	if err != nil {
		log.Fatal("证书生成错误", err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(cred))

	service.RegisterProductServiceServer(rpcServer, model.ProductService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("start server error: ", err)
	}
}