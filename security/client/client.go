package main

import (
	"context"
	"go-grpc-dome/security/pkg/auth"
	pd "go-grpc-dome/simple/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	// Address 监听地址
	Address string = ":50501"
)

var grpcClient pd.SimpleClient

func main(){
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	crets, err:=credentials.NewClientTLSFromFile("../pkg/tls/server.pem","go-grpc-dome")
	if err !=nil{
		log.Fatalf("failed credentials err:%v", err)
	}

	// token
	token :=auth.Token{
		AppId: "grpc_token",
		AppSecret: "123456789",
	}

	// 连接服务器
	connect, err :=grpc.Dial(Address, grpc.WithTransportCredentials(crets), grpc.WithPerRPCCredentials(&token))
	if err !=nil{
		log.Fatalf("connect failed err:%v", err)
	}

	defer connect.Close()
	// 建立gRPC连接
	grpcClient =pd.NewSimpleClient(connect)

	// 调用我们的服务(Route方法)
	route()
}


func route(){
	req :=&pd.SimpleRequest{
		Data: " hello grpc",
	}

	res, err :=grpcClient.Route(context.Background(), req)

	if err !=nil{
		log.Fatalf("client route err:%v", err)
	}

	log.Println(res)
}