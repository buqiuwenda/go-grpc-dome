package cred

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// TLSInterceptor TLS证书认证

func TLSInterceptor() grpc.ServerOption{
	creds, err :=credentials.NewServerTLSFromFile("../tls/server.pem", "../tls/server.key")
	if err!=nil{
		log.Fatalf("fail TLSInterceptor err:%v", err)
	}

	return grpc.Creds(creds)
}
