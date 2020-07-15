package main

import (
	pd "go-grpc-dome/simple/proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
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
	listValue()
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

func listValue(){
	req :=pd.SimpleRequest{
		Data: "grpc stream server",
	}

	stream, err :=grpcClient.ListValue(context.Background(), &req)
	if err !=nil{
		log.Fatalf("call stream fial err:%v",err)
	}

	for{
		res, err := stream.Recv()
		if err ==io.EOF{
			break
		}

		if err !=nil{
			log.Fatalf("ListStr get stream err: %v", err)
		}

		log.Println(res.StreamValue)
		break;
	}
	stream.CloseSend()
}