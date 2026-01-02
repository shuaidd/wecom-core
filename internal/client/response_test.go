package client

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shuaidd/wecom-core/internal/errors"
)

func TestParseResponse_Success(t *testing.T) {
	body := `{"errcode":0,"errmsg":"ok","data":"test"}`
	httpResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}

	resp, err := ParseResponse(httpResp)
	require.NoError(t, err)

	assert.Equal(t, 0, resp.ErrCode)
	assert.Equal(t, "ok", resp.ErrMsg)
	assert.True(t, resp.IsSuccess())
}

func TestParseResponse_Error(t *testing.T) {
	body := `{"errcode":40014,"errmsg":"invalid access_token"}`
	httpResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}

	resp, err := ParseResponse(httpResp)
	require.Error(t, err)

	assert.Equal(t, 40014, resp.ErrCode)
	assert.Equal(t, "invalid access_token", resp.ErrMsg)
	assert.False(t, resp.IsSuccess())

	// Check error type
	assert.True(t, errors.IsWecomError(err))
	assert.Equal(t, 40014, errors.GetErrorCode(err))
}

func TestParseResponse_InvalidJSON(t *testing.T) {
	body := `{invalid json}`
	httpResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}

	_, err := ParseResponse(httpResp)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal response")
}

func TestResponse_Unmarshal(t *testing.T) {
	body := `{"errcode":0,"errmsg":"ok","userid":"zhangsan","name":"张三"}`
	httpResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}

	resp, err := ParseResponse(httpResp)
	require.NoError(t, err)

	type TestData struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		UserID  string `json:"userid"`
		Name    string `json:"name"`
	}

	var data TestData
	err = resp.Unmarshal(&data)
	require.NoError(t, err)

	assert.Equal(t, 0, data.ErrCode)
	assert.Equal(t, "ok", data.ErrMsg)
	assert.Equal(t, "zhangsan", data.UserID)
	assert.Equal(t, "张三", data.Name)
}

func TestResponse_Unmarshal_InvalidJSON(t *testing.T) {
	body := `{"errcode":0,"errmsg":"ok"}`
	httpResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}

	resp, err := ParseResponse(httpResp)
	require.NoError(t, err)

	type InvalidStruct struct {
		RequiredField int `json:"required_field"`
	}

	var data InvalidStruct
	// This should not error, just leave RequiredField as 0
	err = resp.Unmarshal(&data)
	require.NoError(t, err)
	assert.Equal(t, 0, data.RequiredField)
}

func TestResponse_IsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		errcode  int
		expected bool
	}{
		{
			name:     "success",
			errcode:  0,
			expected: true,
		},
		{
			name:     "error",
			errcode:  40014,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				CommonResponse: CommonResponse{
					ErrCode: tt.errcode,
				},
			}
			assert.Equal(t, tt.expected, resp.IsSuccess())
		})
	}
}
