package main

import (
	"crypto/tls"
	"crypto/x509"
	"grpc_study/model"
	"grpc_study/service"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// cred, err := credentials.NewServerTLSFromFile("../cert/server.pem", "../cert/server.key")
	// if err != nil {
	// 	log.Fatal("证书生成错误", err)
	// }
	cert, err := tls.LoadX509KeyPair("../cert/server.pem", "../cert/server.key")
	if err != nil {
		log.Fatal("证书读取错误", err)
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("../cert/ca.crt")
	if err != nil {
		log.Fatal("ca证书读取错误")
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS((&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	}))

	rpcServer := grpc.NewServer(grpc.Creds(creds))

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