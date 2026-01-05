package client

import (
	"context"
	"net/http"

	"github.com/shuaidd/wecom-core/pkg/interceptor"
)

// RequestInterceptor 请求拦截器（内部使用）
type RequestInterceptor = interceptor.RequestInterceptor

// ResponseInterceptor 响应拦截器（内部使用）
type ResponseInterceptor = interceptor.ResponseInterceptor

// AfterResponseInterceptor 响应后拦截器（内部使用）
type AfterResponseInterceptor = interceptor.AfterResponseInterceptor

// Interceptors 拦截器集合
type Interceptors struct {
	// BeforeRequest 请求前拦截器列表
	BeforeRequest []RequestInterceptor
	// BeforeResponse 响应前拦截器列表（解析前）
	BeforeResponse []ResponseInterceptor
	// AfterResponse 响应后拦截器列表（解析后）
	AfterResponse []AfterResponseInterceptor
}

// NewInterceptors 创建拦截器集合
func NewInterceptors() *Interceptors {
	return &Interceptors{
		BeforeRequest:  make([]RequestInterceptor, 0),
		BeforeResponse: make([]ResponseInterceptor, 0),
		AfterResponse:  make([]AfterResponseInterceptor, 0),
	}
}

// AddRequestInterceptor 添加请求拦截器
func (i *Interceptors) AddRequestInterceptor(interceptor RequestInterceptor) {
	i.BeforeRequest = append(i.BeforeRequest, interceptor)
}

// AddResponseInterceptor 添加响应拦截器（解析前）
func (i *Interceptors) AddResponseInterceptor(interceptor ResponseInterceptor) {
	i.BeforeResponse = append(i.BeforeResponse, interceptor)
}

// AddAfterResponseInterceptor 添加响应后拦截器（解析后）
func (i *Interceptors) AddAfterResponseInterceptor(interceptor AfterResponseInterceptor) {
	i.AfterResponse = append(i.AfterResponse, interceptor)
}

// executeRequestInterceptors 执行请求拦截器
func (i *Interceptors) executeRequestInterceptors(ctx context.Context, req *http.Request, body any) error {
	for _, interceptor := range i.BeforeRequest {
		if err := interceptor(ctx, req, body); err != nil {
			return err
		}
	}
	return nil
}

// executeResponseInterceptors 执行响应拦截器（解析前）
func (i *Interceptors) executeResponseInterceptors(ctx context.Context, resp *http.Response) error {
	for _, interceptor := range i.BeforeResponse {
		if err := interceptor(ctx, resp); err != nil {
			return err
		}
	}
	return nil
}

// executeAfterResponseInterceptors 执行响应后拦截器（解析后）
func (i *Interceptors) executeAfterResponseInterceptors(ctx context.Context, resp *Response) error {
	// 转换 Response 为 interceptor.Response
	interceptorResp := &interceptor.Response{
		ErrCode: resp.ErrCode,
		ErrMsg:  resp.ErrMsg,
		Body:    resp.Body,
	}

	for _, interceptor := range i.AfterResponse {
		if err := interceptor(ctx, interceptorResp); err != nil {
			return err
		}
	}

	// 更新原始响应（如果拦截器修改了数据）
	resp.ErrCode = interceptorResp.ErrCode
	resp.ErrMsg = interceptorResp.ErrMsg
	resp.Body = interceptorResp.Body

	return nil
}
