package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pd "go-grpc-dome/simple/proto"
)

const (
	// Address 监听地址
	Address string = ":50501"
	// Network 网络通信协议
	Network string = "tcp"
)

type SimpleService struct{}
// 简单模式
func main(){
	// 监听本地端口
	listener,err :=net.Listen(Network, Address)

	if err !=nil{
		log.Fatalf("listen port failed err:%v \n", err)
	}


	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	crets, err:=credentials.NewServerTLSFromFile("../pkg/tls/server.pem","../pkg/tls/server.key")
	if err !=nil{
		log.Fatalf("failed credentials err:%v", err)
	}

	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(resp interface{}, err error){
		//拦截普通方法请求，验证Token
		err = checkToken(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	// 新建gRPC服务器实例
	grpcServer :=grpc.NewServer(grpc.Creds(crets), grpc.UnaryInterceptor(interceptor))
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


func checkToken(ctx context.Context) error{
	//从上下文中获取元数据
	md, ok :=metadata.FromIncomingContext(ctx)
	if !ok{
		return errors.New("get appId, appSecret fail")
	}

	var (
		appId     string
		appSecret string
	)

	if value, ok := md["app_id"]; ok {
		appId = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appId !="grpc_token" || appSecret !="123456789"{
		return  status.Errorf(codes.Unauthenticated, "token 无效 appId=%s appSecret=%s", appId, appSecret)
	}

	return nil
}