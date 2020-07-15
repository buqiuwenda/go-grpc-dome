package main

import (
	"context"
	pd "go-grpc-dome/simple/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type StreamService struct{}

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

	pd.RegisterStreamServerServer(grpcServer, &StreamService{})

	err =grpcServer.Serve(listener)
	if err !=nil{
		log.Fatalf("grpc server fail err:%v", err)
	}
}


func (s *StreamService) Route(ctx context.Context, req *pd.SimpleRequest)(*pd.SimpleResponse, error){
	res :=&pd.SimpleResponse{
		Code: 200,
		Value: "hello "+req.Data,
	}

	return res, nil
}

func (s *StreamService) ListValue(req *pd.SimpleRequest, srv pd.StreamServer_ListValueServer)error {
	for n:=0;n<5; n++{
		err :=srv.Send(&pd.StreamResponse{
			StreamValue: req.Data+strconv.Itoa(n),
		})

		if err !=nil{
			return err
		}

		time.Sleep(time.Second * 1)
		log.Println(n)
	}

	return nil
}