package zap

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"log"
)

// ZapInterceptor 返回zap.logger实例(把日志输出到控制台)
func ZapInterceptor() *zap.Logger{
	logger, err :=zap.NewDevelopment()
	if err !=nil{
		log.Fatalf("init zap logger fail err:%v", err)
	}

	grpc_zap.ReplaceGrpcLogger(logger)

	return logger
}
