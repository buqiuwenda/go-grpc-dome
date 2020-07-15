package main

import (
	"context"
	pd "go-grpc-dome/client_stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

const Address string = ":50501"

var grpcClient pd.StreamServerClient

func main(){
	connect, err:=grpc.Dial(Address, grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("connect fail err:%v", err)
	}
	defer connect.Close()

	grpcClient = pd.NewStreamServerClient(connect)
	route()
	routeList()
}


func route(){
	req :=&pd.SimpleRequest{
		Data: "grpc",
	}

	res,err :=grpcClient.Route(context.Background(), req)
	if err !=nil{
		log.Fatalf("call route fial err:%v",err)
	}

	log.Println(res)
}
// routeList 调用服务端RouteList方法
func routeList(){
	stream, err :=grpcClient.RouteList(context.Background())

	if err !=nil{
		log.Fatalf("Upload data err:%v",err)
	}

	for n:=0;n< 5; n++{
		//向流中发送消息
		err :=stream.Send(&pd.StreamRequest{
			StreamData: "stream client grpc "+strconv.Itoa(n),
		})

		if err == io.EOF{
			break
		}

		if err !=nil{
			log.Fatalf("stream client request err:%v", err)
		}
	}

	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err !=nil{
		log.Fatalf("close stream client fail err:%v", err)
	}

	log.Println(res)
}