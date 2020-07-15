package main

import (
	"context"
	pd "go-grpc-dome/both_stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

const (
	// Address 监听地址
	Address string = ":50501"
	// Network 网络通信协议
	Network string = "tcp"
)

type StreamService struct{}

func main(){
	// 监听本地端口
	listener,err :=net.Listen(Network, Address)

	if err !=nil{
		log.Fatalf("listen port failed err:%v \n", err)
	}

	log.Println(Address+" address listener ")
	// 新建gRPC服务器实例
	grpcServer :=grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pd.RegisterStreamServer(grpcServer, &StreamService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err =grpcServer.Serve(listener)
	if err !=nil{
		log.Fatalf("grpc server failed err:%v", err)
	}
}


func(s *StreamService) Route(ctx context.Context, req *pd.SimpleRequest)(*pd.SimpleResponse, error){
	res :=&pd.SimpleResponse{
		Code:  200,
		Value: "hello" + req.Data,
	}

	return res,nil
}

// Conversations 实现Conversations方法
func (s *StreamService) Conversations(srv pd.Stream_ConversationsServer) error{
	n:=1

	for{
		req,err :=srv.Recv()
		if err == io.EOF{
			return nil
		}

		if err !=nil{
			return err
		}

		err = srv.Send(&pd.StreamResponse{
			Answer: strconv.Itoa(n)+". 你是谁？ request: "+req.Question,
		})

		if err !=nil{
			return err
		}
		n ++
		log.Printf("from stream client request:%v",req.Question)
	}
}
