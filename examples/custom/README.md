# 自定义 API 调用示例

本示例演示如何使用 wecom-core SDK 调用未封装的企业微信 API，同时完全复用 SDK 的 token 管理和重试逻辑。

## 功能特性

当 SDK 尚未封装某个企业微信 API 时，你可以使用以下方法自行调用：

- ✅ 自动注入 `access_token`
- ✅ 自动处理 token 过期并重试
- ✅ 自动处理速率限制（errcode 45009）并重试
- ✅ 自动处理系统繁忙（errcode 10001）并重试
- ✅ 支持 TraceID 传递
- ✅ 完整的请求/响应日志记录（Debug 模式）

## API 方法说明

### 1. CustomGet - GET 请求（手动解析）

```go
func (c *Client) CustomGet(ctx context.Context, path string, query map[string]string) (*Response, error)
```

发送 GET 请求，返回原始响应对象，需要手动调用 `resp.Unmarshal()` 解析。

**适用场景**：需要访问原始响应数据或自定义解析逻辑。

### 2. CustomPost - POST 请求（手动解析）

```go
func (c *Client) CustomPost(ctx context.Context, path string, body any) (*Response, error)
```

发送 POST 请求，返回原始响应对象，需要手动调用 `resp.Unmarshal()` 解析。

**适用场景**：需要访问原始响应数据或自定义解析逻辑。

### 3. CustomGetAndUnmarshal - GET 请求（自动解析）⭐推荐

```go
func CustomGetAndUnmarshal[T any](c *Client, ctx context.Context, path string, query map[string]string) (*T, error)
```

发送 GET 请求并自动解析响应到指定类型，一步到位。

**适用场景**：大多数情况下的首选方案，代码更简洁。

### 4. CustomPostAndUnmarshal - POST 请求（自动解析）⭐推荐

```go
func CustomPostAndUnmarshal[T any](c *Client, ctx context.Context, path string, body any) (*T, error)
```

发送 POST 请求并自动解析响应到指定类型，一步到位。

**适用场景**：大多数情况下的首选方案，代码更简洁。

## 使用示例

### 基础用法

```go
package main

import (
    "context"
    "log"

    "github.com/shuaidd/wecom-core"
    "github.com/shuaidd/wecom-core/config"
)

func main() {
    // 1. 创建客户端
    client, err := wecom.New(
        config.WithCorpID("your_corp_id"),
        config.WithCorpSecret("your_corp_secret"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // 2. 定义响应类型（必须嵌入 wecom.CommonResponse）
    type MyCustomResponse struct {
        wecom.CommonResponse          // 必须嵌入，包含 errcode 和 errmsg
        Data                string   `json:"data"`
    }

    // 3. 调用自定义 API
    result, err := wecom.CustomGetAndUnmarshal[MyCustomResponse](
        client,
        ctx,
        "/cgi-bin/your/custom/api",
        map[string]string{"param": "value"},
    )
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Result: %s", result.Data)
}
```

### GET 请求示例

```go
// 定义响应类型
type GetCallbackIPResponse struct {
    wecom.CommonResponse
    IPList []string `json:"ip_list"`
}

// 方式1：自动解析（推荐）
result, err := wecom.CustomGetAndUnmarshal[GetCallbackIPResponse](
    client,
    ctx,
    "/cgi-bin/getcallbackip",
    nil, // 无查询参数
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result.IPList)

// 方式2：手动解析
resp, err := client.CustomGet(ctx, "/cgi-bin/getcallbackip", nil)
if err != nil {
    log.Fatal(err)
}
var result GetCallbackIPResponse
if err := resp.Unmarshal(&result); err != nil {
    log.Fatal(err)
}
fmt.Println(result.IPList)
```

### POST 请求示例

```go
// 定义请求类型
type SendMessageRequest struct {
    ToUser  string `json:"touser"`
    MsgType string `json:"msgtype"`
    AgentID int    `json:"agentid"`
    Text    struct {
        Content string `json:"content"`
    } `json:"text"`
}

// 定义响应类型
type SendMessageResponse struct {
    wecom.CommonResponse
    MsgID string `json:"msgid"`
}

// 准备请求
req := SendMessageRequest{
    ToUser:  "UserID1",
    MsgType: "text",
    AgentID: 1000001,
}
req.Text.Content = "测试消息"

// 方式1：自动解析（推荐）
result, err := wecom.CustomPostAndUnmarshal[SendMessageResponse](
    client,
    ctx,
    "/cgi-bin/message/send",
    req,
)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("MsgID: %s\n", result.MsgID)

// 方式2：手动解析
resp, err := client.CustomPost(ctx, "/cgi-bin/message/send", req)
if err != nil {
    log.Fatal(err)
}
var result SendMessageResponse
if err := resp.Unmarshal(&result); err != nil {
    log.Fatal(err)
}
fmt.Printf("MsgID: %s\n", result.MsgID)
```

### 带查询参数的 GET 请求

```go
type MyResponse struct {
    wecom.CommonResponse
    Result string `json:"result"`
}

result, err := wecom.CustomGetAndUnmarshal[MyResponse](
    client,
    ctx,
    "/cgi-bin/some/api",
    map[string]string{
        "param1": "value1",
        "param2": "value2",
    },
)
```

## 重要说明

### 1. 响应类型定义

自定义响应类型**必须**嵌入 `wecom.CommonResponse`，这样 SDK 才能正确处理错误码和错误消息：

```go
type MyResponse struct {
    wecom.CommonResponse  // ✅ 必须嵌入
    MyField string       `json:"my_field"`
}
```

`wecom.CommonResponse` 包含：
- `ErrCode int` - 错误码，0 表示成功
- `ErrMsg string` - 错误消息

### 2. 错误处理

SDK 会自动处理以下错误并重试：
- Token 过期（errcode 40014, 42001）- 自动刷新 token 并重试
- 速率限制（errcode 45009）- 指数退避后重试
- 系统繁忙（errcode 10001）- 指数退避后重试

其他错误会直接返回，不会重试。

### 3. TraceID 传递

支持在 context 中传递 TraceID 用于请求追踪：

```go
ctx := wecom.WithTraceID(context.Background(), "your-trace-id")
result, err := wecom.CustomGetAndUnmarshal[MyResponse](client, ctx, "/path", nil)
```

### 4. Debug 模式

开启 Debug 模式可以查看详细的请求和响应日志：

```go
client, err := wecom.New(
    config.WithCorpID("your_corp_id"),
    config.WithCorpSecret("your_corp_secret"),
    config.WithDebug(true),  // 开启 Debug 模式
)
```

## 完整示例

运行本目录下的 `main.go` 查看完整示例：

```bash
# 修改 main.go 中的 corp_id 和 corp_secret
# 然后运行
go run examples/custom/main.go
```

## 最佳实践

1. **优先使用泛型方法**：`CustomGetAndUnmarshal` 和 `CustomPostAndUnmarshal` 代码更简洁
2. **正确定义响应类型**：必须嵌入 `wecom.CommonResponse`
3. **合理使用 Debug 模式**：开发时开启，生产环境关闭
4. **传递 TraceID**：便于问题追踪和日志关联
5. **错误处理**：始终检查返回的 error

## 参考文档

- [企业微信 API 文档](https://developer.work.weixin.qq.com/document/)
- [wecom-core 项目主页](https://github.com/shuaidd/wecom-core)
