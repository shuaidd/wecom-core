package interceptor

import (
	"context"
	"net/http"
)

// RequestInterceptor 请求拦截器，在发送请求前调用
// 可以修改请求的 Header 和 Body
// ctx: 请求上下文
// req: HTTP 请求对象，可以修改 Header
// body: 请求体对象（如果是 POST 请求）
type RequestInterceptor func(ctx context.Context, req *http.Request, body any) error

// ResponseInterceptor 响应拦截器，在解析响应前调用
// 可以访问原始的 HTTP 响应
// ctx: 请求上下文
// resp: HTTP 响应对象
type ResponseInterceptor func(ctx context.Context, resp *http.Response) error

// Response 响应对象
type Response struct {
	// ErrCode 错误码，0表示成功
	ErrCode int `json:"errcode"`
	// ErrMsg 错误信息
	ErrMsg string `json:"errmsg"`
	// Body 原始响应体
	Body []byte `json:"-"`
}

// AfterResponseInterceptor 响应后拦截器，在响应解析后调用
// 可以访问解析后的响应数据
// ctx: 请求上下文
// resp: 解析后的响应对象
type AfterResponseInterceptor func(ctx context.Context, resp *Response) error
