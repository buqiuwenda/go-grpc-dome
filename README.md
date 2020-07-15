# go-grpc-dome

### 入门教程

* Go gRPC教程-环境安装（一） 
   
   * 安装proto (mac) 
      * 命令：`brew install protoc`
   * 安装 goalng 的proto编译支持
     * `go get -u github.com/golang/protobuf/protoc-gen-go`
   * 安装 gRPC包
     * `go get -u google.golang.org/grpc`   
   * 创建并编译proto文件
      * 新建proto文件夹，在里面新建simple.proto文件
      ```
        syntax = "proto3";// 协议为proto3
        
        package proto;
        
        // 定义发送请求信息
        message SimpleRequest{
            // 定义发送的参数
            // 参数类型 参数名 标识号(不可重复)
            string data = 1;
        }
        
        // 定义响应信息
        message SimpleResponse{
            // 定义接收的参数
            // 参数类型 参数名 标识号(不可重复)
            int32 code = 1;
            string value = 2;
        }
        
        // 定义我们的服务（可定义多个服务,每个服务可定义多个接口）
        service Simple{
            rpc Route (SimpleRequest) returns (SimpleResponse){};
        }
      ``` 

      * 编译proto文件 
      
      cmd进入simple.proto所在目录，运行以下指令进行编译
   
      protoc --go_out=plugins=grpc:./ ./simple.proto          
    
* Go gRPC教程-简单RPC（二）

* Go gRPC教程-服务端流式RPC（三）

* Go gRPC教程-客户端流式RPC（四）



### 感谢
* [Go gRPC官方文档](https://grpc.io/docs/languages/go/quickstart/)

