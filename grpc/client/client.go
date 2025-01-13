package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc_study/client/service"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// creds, err := credentials.NewClientTLSFromFile("../cert/server.pem", "*.zhong")
	// if err != nil {
	// 	log.Fatal("证书错误", err)
	// }
	cert, _ := tls.LoadX509KeyPair("../cert/client.pem", "../cert/client.key")
	certPoll := x509.NewCertPool()
	ca, _ := os.ReadFile("../cert/ca.crt")
	certPoll.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName: "*.zhong",
		RootCAs: certPoll,
	})
	conn, err := grpc.NewClient("127.0.0.1:8002", grpc.WithTransportCredentials(creds))
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
