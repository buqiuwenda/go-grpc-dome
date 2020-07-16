package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pd "go-grpc-dome/simple/proto"
	"time"
)

const (
	// Address 监听地址
	Address string = ":50501"
)

var grpcClient pd.SimpleClient

func main(){
	// 连接服务器
	connect, err :=grpc.Dial(Address, grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("connect failed err:%v", err)
	}

	defer connect.Close()
	// 建立gRPC连接
	grpcClient =pd.NewSimpleClient(connect)

	// 调用我们的服务(Route方法)
	route(context.Background(), 3)
}


func route(ctx context.Context, timeout time.Duration){
	clientTimeout :=time.Now().Add(timeout * time.Second)
	ctx, cancel :=context.WithDeadline(ctx, clientTimeout)

	defer cancel()

	req :=&pd.SimpleRequest{
		Data: " hello grpc",
	}

	res, err :=grpcClient.Route(context.Background(), req)

	if err !=nil{
		statusCode, ok :=status.FromError(err)
		if ok{
			if statusCode.Code() == codes.DeadlineExceeded{
				log.Fatalf("client route timeout !")
			}
		}

		log.Fatalf("client route err:%v", err)
	}

	log.Println(res)
}