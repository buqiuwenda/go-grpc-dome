package auth

import "context"

type Token struct{
	AppId string
	AppSecret string
}

//获取当前请求认证所需的元数据
func(t *Token) GetRequestMetadata(ctx context.Context, uri ...string)(map[string]string, error){
	return map[string]string{"app_id": t.AppId, "app_secret": t.AppSecret}, nil
}

//  TLS 认证进行安全传输
func(t *Token) RequireTransportSecurity() bool{
	return true
}