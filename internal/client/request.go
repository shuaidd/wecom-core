package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// HTTPMethod HTTP方法
type HTTPMethod string

const (
	// MethodGet GET方法
	MethodGet HTTPMethod = http.MethodGet
	// MethodPost POST方法
	MethodPost HTTPMethod = http.MethodPost
)

// Request 表示一个HTTP请求
type Request struct {
	// Method HTTP方法
	Method HTTPMethod
	// Path API路径
	Path string
	// Query 查询参数
	Query url.Values
	// Body 请求体
	Body any
	// RawBody 原始请求体（用于multipart）
	RawBody []byte
	// ContentType 内容类型（用于multipart）
	ContentType string
}

// NewRequest 创建新请求
func NewRequest(method HTTPMethod, path string) *Request {
	return &Request{
		Method: method,
		Path:   path,
		Query:  url.Values{},
	}
}

// AddQuery 添加查询参数
func (r *Request) AddQuery(key, value string) *Request {
	r.Query.Add(key, value)
	return r
}

// SetBody 设置请求体
func (r *Request) SetBody(body any) *Request {
	r.Body = body
	return r
}

// NewMultipartRequest 创建multipart/form-data请求
func NewMultipartRequest(path string, body []byte, contentType string) *Request {
	return &Request{
		Method:      MethodPost,
		Path:        path,
		Query:       url.Values{},
		RawBody:     body,
		ContentType: contentType,
	}
}

// BuildHTTPRequest 构建http.Request
func (r *Request) BuildHTTPRequest(ctx context.Context, baseURL string) (*http.Request, error) {
	// 构建完整URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}
	u.Path = r.Path
	u.RawQuery = r.Query.Encode()

	// 构建请求体
	var body io.Reader
	if r.Method == MethodPost {
		if r.RawBody != nil {
			// 使用原始body（用于multipart）
			body = bytes.NewReader(r.RawBody)
		} else if r.Body != nil {
			// 使用JSON body
			jsonData, err := json.Marshal(r.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal request body: %w", err)
			}
			body = bytes.NewReader(jsonData)
		}
	}

	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, string(r.Method), u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	// 设置请求头
	if r.Method == MethodPost {
		if r.ContentType != "" {
			// 使用自定义ContentType（用于multipart）
			req.Header.Set("Content-Type", r.ContentType)
		} else if r.Body != nil {
			// 使用JSON ContentType
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
	}

	return req, nil
}
