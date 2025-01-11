package main

import (
	"grpc_study/model"
	"grpc_study/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()

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