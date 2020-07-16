package auth

import (
	"context"
	"errors"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type TokenInfo struct{
	ID string
	Role []string
}

// AuthInterceptor 认证拦截器，对以authorization为头部，形式为`bearer token`的Token进行验证
func AuthInterceptor(ctx context.Context) (context.Context, error){
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err !=nil {
		return nil, err
	}

	tokenInfo, err :=parseToken(token)
	if err != nil{
		return nil, grpc.Errorf(codes.Unauthenticated, "token Unauthenticated error:%v", err)
	}

	newCtx :=context.WithValue(ctx, tokenInfo.ID, tokenInfo.Role)

	return newCtx, nil
}

//解析token，并进行验证
func parseToken(token string)(TokenInfo, error){
	var tokenInfo TokenInfo
	if token == "grpc.auth.token"{
		tokenInfo.ID="1"
		tokenInfo.Role=[]string{"admin"}
		return tokenInfo,nil
	}

	return tokenInfo, errors.New("无效token bearer")

}

//从token中获取用户唯一标识
func userClaimFromToken(tokenInfo TokenInfo) string {
	return tokenInfo.ID
}