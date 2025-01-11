package main

import (
	"context"
	"fmt"
	"grpc_study/client/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("connect error: ", err)
	}
	defer conn.Close()
	productServiceClient := service.NewProductServiceClient(conn)

	resp, err := productServiceClient.GetProductStock(context.Background(),
		&service.ProductRequest{ProductId: 233})
	if err != nil {
		log.Fatal("call func error: ", err)
	}

	fmt.Println("stock = ", resp.ProductStock)
	
}
