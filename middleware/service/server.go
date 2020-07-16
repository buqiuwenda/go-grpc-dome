package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	pd "go-grpc-dome/middleware/proto"
	"go-grpc-dome/middleware/service/middleware/auth"
	"go-grpc-dome/middleware/service/middleware/cred"
	"go-grpc-dome/middleware/service/middleware/recovery"
	"go-grpc-dome/middleware/service/middleware/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	// Address 监听地址
	Address string = ":50501"
	// Network 网络通信协议
	Network string = "tcp"
)

type SimpleService struct{}

func main(){
	// 监听本地端口
	listener,err :=net.Listen(Network, Address)

	if err !=nil{
		log.Fatalf("listen port failed err:%v \n", err)
	}


	// 新建gRPC服务器实例
	grpcServer :=grpc.NewServer(cred.TLSInterceptor(),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)

	// 在gRPC服务器注册我们的服务
	pd.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address+" address listener with tls token ...")
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err =grpcServer.Serve(listener)
	if err !=nil{
		log.Fatalf("grpc server failed err:%v", err)
	}
}


func(s *SimpleService) Route(ctx context.Context, req *pd.SimpleRequest)(*pd.SimpleResponse, error){
	res :=&pd.SimpleResponse{
		Code: 200,
		Value: "security "+req.Data,
	}

	return res, nil
}