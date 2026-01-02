package client

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest(MethodGet, "/api/test")

	assert.Equal(t, MethodGet, req.Method)
	assert.Equal(t, "/api/test", req.Path)
	assert.NotNil(t, req.Query)
	assert.Nil(t, req.Body)
}

func TestRequest_AddQuery(t *testing.T) {
	req := NewRequest(MethodGet, "/api/test")
	req.AddQuery("key1", "value1")
	req.AddQuery("key2", "value2")

	assert.Equal(t, "value1", req.Query.Get("key1"))
	assert.Equal(t, "value2", req.Query.Get("key2"))
}

func TestRequest_SetBody(t *testing.T) {
	req := NewRequest(MethodPost, "/api/test")
	body := map[string]string{"key": "value"}
	req.SetBody(body)

	assert.Equal(t, body, req.Body)
}

func TestRequest_BuildHTTPRequest_GET(t *testing.T) {
	ctx := context.Background()
	baseURL := "https://api.example.com"

	req := NewRequest(MethodGet, "/api/test")
	req.AddQuery("param1", "value1")
	req.AddQuery("param2", "value2")

	httpReq, err := req.BuildHTTPRequest(ctx, baseURL)
	require.NoError(t, err)

	assert.Equal(t, "GET", httpReq.Method)
	assert.Equal(t, "https://api.example.com/api/test?param1=value1&param2=value2", httpReq.URL.String())
	assert.Nil(t, httpReq.Body)
}

func TestRequest_BuildHTTPRequest_POST(t *testing.T) {
	ctx := context.Background()
	baseURL := "https://api.example.com"

	req := NewRequest(MethodPost, "/api/test")
	req.SetBody(map[string]string{"key": "value"})

	httpReq, err := req.BuildHTTPRequest(ctx, baseURL)
	require.NoError(t, err)

	assert.Equal(t, "POST", httpReq.Method)
	assert.Equal(t, "https://api.example.com/api/test", httpReq.URL.String())
	assert.NotNil(t, httpReq.Body)
	assert.Equal(t, "application/json; charset=utf-8", httpReq.Header.Get("Content-Type"))
}

func TestRequest_BuildHTTPRequest_InvalidBaseURL(t *testing.T) {
	ctx := context.Background()
	baseURL := "://invalid-url"

	req := NewRequest(MethodGet, "/api/test")
	_, err := req.BuildHTTPRequest(ctx, baseURL)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid base URL")
}

func TestRequest_BuildHTTPRequest_ChainedMethods(t *testing.T) {
	ctx := context.Background()
	baseURL := "https://api.example.com"

	req := NewRequest(MethodPost, "/api/test").
		AddQuery("param1", "value1").
		SetBody(map[string]string{"key": "value"})

	httpReq, err := req.BuildHTTPRequest(ctx, baseURL)
	require.NoError(t, err)

	assert.Equal(t, "POST", httpReq.Method)
	assert.Contains(t, httpReq.URL.String(), "param1=value1")
	assert.NotNil(t, httpReq.Body)
}

func TestRequest_AddQuery_URLEncoding(t *testing.T) {
	req := NewRequest(MethodGet, "/api/test")
	req.AddQuery("key", "value with spaces")

	expected := url.Values{}
	expected.Add("key", "value with spaces")
	assert.Equal(t, expected, req.Query)
}
