package main

import (
	pd "../proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type SimpleService struct{}

const (
	// Address 监听地址
	Address string = ":50501"
	// Network 网络通信协议
	Network string = "tcp"
)


func main(){
	// 监听本地端口
	listener,err :=net.Listen(Network, Address)

	if err !=nil{
		log.Fatalf("listen port failed err:%v \n", err)
	}

	log.Println(Address+" address listener ")

	grpcServer :=grpc.NewServer()

	pd.RegisterStreamServerServer(grpcServer, &SimpleService{})

	err =grpcServer.Serve(listener)
	if err !=nil{
		log.Fatalf("grpc server fail err:%v", err)
	}
}


func (s *SimpleService) Route(ctx context.Context, req *pd.SimpleRequest)(*pd.SimpleResponse, error){
	res :=&pd.SimpleResponse{
		Code: 200,
		Value: "hello "+req.Data,
	}

	return res, nil
}

func (s *SimpleService) RouteList(srv pd.StreamServer_RouteListServer) error{
	for{
		res, err :=srv.Recv()
		if err == io.EOF{
			return srv.SendAndClose(&pd.SimpleResponse{Value: "ok"})
		}

		if err !=nil{
			return err
		}

		log.Println(res.StreamData)
	}
}