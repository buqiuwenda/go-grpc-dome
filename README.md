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

* Go gRPC教程-双向流式RPC（五）

### 拓展

* Go gRPC进阶-超时设置（六）

* Go gRPC进阶-TLS认证+自定义方法认证（七）
    *  生成私钥
    
    生成RSA私钥：openssl genrsa -out server.key 2048
    > 生成RSA私钥，命令的最后一个参数，将指定生成密钥的位数，如果没有指定，默认512 
                                                                                                                                                                                       
    * 生成公钥
    
    `openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650`
    > openssl req：生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密、-key指定私钥文件、-x509指输出证书、-days 3650为有效期
    
    此后则输入证书拥有者信息
    ```
    Country Name (2 letter code) [AU]:CN
    State or Province Name (full name) [Some-State]:YxYx
    Locality Name (eg, city) []:YxYx
    Organization Name (eg, company) [Internet Widgits Pty Ltd]:YX Co. Ltd
    Organizational Unit Name (eg, section) []:Develop
    Common Name (e.g. server FQDN or YOUR name) []:go-grpc-dome
    Email Address []:yxx@yxx.com
    ```
    
* Go gRPC进阶-go-grpc-middleware使用（八）

* Go gRPC进阶-proto数据验证（九）

* Go gRPC进阶-gRPC转换HTTP（十）



### 鸣谢
* [Go gRPC官方文档](https://grpc.io/docs/languages/go/quickstart/)

