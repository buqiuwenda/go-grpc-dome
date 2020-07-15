package main

import (
	"context"
	pd "go-grpc-dome/both_stream/proto"
	"io"
	"log"
	"strconv"
)

const Address string = ":50501"

var streamClient pd.StreamClient

func main(){

}

func route(){
	req :=&pd.SimpleRequest{
		Data: "grpc stream",
	}
	// 调用我们的服务(Route方法)
	res,err :=streamClient.Route(context.Background(), req)
	if err !=nil{
		log.Fatalf("call route fial err:%v",err)
	}

	log.Println(res)
}
// conversations 调用服务端的Conversations方法
func conversations(){
	stream, err :=streamClient.Conversations(context.Background())
	if err !=nil{
		log.Fatalf("conversations fail err:%v", err)
	}

	for n:=0; n<5;n++{
		err :=stream.Send(&pd.StreamRequest{Question: "张三"+strconv.Itoa(n)})
		if err !=nil{
			log.Fatalf("stream client send err:%v", err)
		}

		res, err:=stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF{
			break
		}

		if err !=nil{
			log.Fatalf("conversations recv fail err:%v", err)
		}

		log.Println(res)
	}
	//  关闭流
	err =stream.CloseSend()
	if err !=nil{
		log.Fatalf("stream clien close fail err:%v", err)
	}

}
