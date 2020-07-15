package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pd "go-grpc-dome/simple/proto"
)

const (
	// Address 监听地址
	Address string = ":50501"
)


func main(){
	// 连接服务器
	connect, err :=grpc.Dial(Address, grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("connect failed err:%v", err)
	}

	defer connect.Close()
	// 建立gRPC连接
	grpcClient :=pd.NewSimpleClient(connect)

	req :=&pd.SimpleRequest{
		Data: " hello grpc",
	}

	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err :=grpcClient.Route(context.Background(), req)

	if err !=nil{
		log.Fatalf("Call Route err: %v", err)
	}

	log.Println(res)
}