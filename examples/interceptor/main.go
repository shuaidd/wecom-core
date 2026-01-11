package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
)

func main() {
	// 从环境变量读取配置
	corpID := os.Getenv("WECOM_CORP_ID")
	corpSecret := os.Getenv("WECOM_CORP_SECRET")

	if corpID == "" || corpSecret == "" {
		log.Fatal("请设置环境变量：WECOM_CORP_ID 和 WECOM_CORP_SECRET")
	}

	// 创建客户端，配置拦截器
	client, err := wecom.New(
		config.WithCorpID(corpID),
		config.WithCorpSecret(corpSecret),
		config.WithDebug(true),

		// 请求前拦截器：添加自定义 Header
		config.WithRequestInterceptor(func(ctx context.Context, req *http.Request, body any) error {
			// 添加自定义 Header
			req.Header.Set("X-Custom-Header", "my-custom-value")
			req.Header.Set("X-Request-Time", time.Now().Format(time.RFC3339))

			fmt.Printf("[请求拦截器] 添加自定义 Header: %s\n", req.Header.Get("X-Custom-Header"))
			fmt.Printf("[请求拦截器] URL: %s\n", req.URL.String())

			// 可以根据请求内容修改 Header
			if body != nil {
				fmt.Printf("[请求拦截器] 请求体类型: %T\n", body)
			}

			return nil
		}),

		// 响应前拦截器：记录原始响应
		config.WithResponseInterceptor(func(ctx context.Context, resp *http.Response) error {
			fmt.Printf("[响应拦截器] HTTP Status: %d\n", resp.StatusCode)
			fmt.Printf("[响应拦截器] Content-Type: %s\n", resp.Header.Get("Content-Type"))
			return nil
		}),

		// 响应后拦截器：处理响应数据
		config.WithAfterResponseInterceptor(func(ctx context.Context, resp *wecom.InterceptorResponse) error {
			fmt.Printf("[响应后拦截器] ErrCode: %d\n", resp.ErrCode)
			fmt.Printf("[响应后拦截器] ErrMsg: %s\n", resp.ErrMsg)

			// 可以修改响应数据
			if resp.ErrCode != 0 {
				fmt.Printf("[响应后拦截器] 检测到错误，记录日志或发送告警\n")
			}

			return nil
		}),

		// 可以添加多个拦截器，它们会按顺序执行
		config.WithAfterResponseInterceptor(func(ctx context.Context, resp *wecom.InterceptorResponse) error {
			// 例如：记录性能指标
			fmt.Printf("[性能监控拦截器] 响应体大小: %d bytes\n", len(resp.Body))
			return nil
		}),
	)

	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	// 示例1：获取部门列表
	fmt.Println("\n=== 示例1：获取部门列表 ===")
	departments, err := client.Contact.ListDepartments(context.Background(), 1)
	if err != nil {
		log.Fatalf("获取部门列表失败: %v", err)
	}
	fmt.Printf("部门数量: %d\n", len(departments))
	for _, dept := range departments {
		fmt.Printf("  - ID: %d, 名称: %s\n", dept.ID, dept.Name)
	}

	// 示例2：使用 TraceID 追踪请求
	fmt.Println("\n=== 示例2：使用 TraceID 追踪请求 ===")
	ctx := wecom.WithTraceID(context.Background(), "trace-123456")
	user, err := client.Contact.GetUser(ctx, "user_id_example")
	if err != nil {
		// 这个请求可能失败，因为用户ID不存在，但能看到拦截器的输出
		fmt.Printf("获取用户失败（预期行为）: %v\n", err)
	} else {
		fmt.Printf("用户: %s\n", user.Name)
	}

	// 示例3：自定义拦截器 - 请求限流
	fmt.Println("\n=== 示例3：请求限流拦截器 ===")
	rateLimitedClient, err := wecom.New(
		config.WithCorpID(corpID),
		config.WithCorpSecret(corpSecret),
		config.WithRequestInterceptor(rateLimitInterceptor()),
	)
	if err != nil {
		log.Fatalf("创建限流客户端失败: %v", err)
	}

	// 发送多个请求测试限流
	for i := 0; i < 3; i++ {
		_, err := rateLimitedClient.Contact.ListDepartments(context.Background(), 1)
		if err != nil {
			fmt.Printf("请求 %d 失败: %v\n", i+1, err)
		} else {
			fmt.Printf("请求 %d 成功\n", i+1)
		}
	}
}

// rateLimitInterceptor 创建一个简单的限流拦截器
// 实际使用中应该使用更完善的限流算法（如令牌桶、漏桶等）
func rateLimitInterceptor() wecom.RequestInterceptor {
	lastRequestTime := time.Now().Add(-time.Second) // 初始化为1秒前

	return func(ctx context.Context, req *http.Request, body any) error {
		// 简单的时间窗口限流：每秒最多1个请求
		elapsed := time.Since(lastRequestTime)
		if elapsed < time.Second {
			waitTime := time.Second - elapsed
			fmt.Printf("[限流拦截器] 请求过快，等待 %v\n", waitTime)
			time.Sleep(waitTime)
		}
		lastRequestTime = time.Now()
		return nil
	}
}
