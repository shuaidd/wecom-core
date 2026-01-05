package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
)

func main() {
	// 方式1: 使用 WithAgent 函数注册多个应用
	client1, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithAgent("customer", 100001, "agent_secret_1", "客户管理应用"),
		config.WithAgent("study-assistant", 100002, "agent_secret_2", "学习助手"),
		config.WithDebug(true),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 方式2: 使用 WithAgents 函数批量注册应用
	client2, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithAgents(
			&wecom.AgentConfig{
				AgentID:   100001,
				Secret:    "agent_secret_1",
				AgentName: "customer",
				AgentDesc: "客户管理应用",
			},
			&wecom.AgentConfig{
				AgentID:   100002,
				Secret:    "agent_secret_2",
				AgentName: "study-assistant",
				AgentDesc: "学习助手",
			},
		),
		config.WithDebug(true),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// 使用 customer 应用发送消息
	customerCtx := wecom.WithAgentName(ctx, "customer")
	err = sendMessageWithAgent(client1, customerCtx, "customer")
	if err != nil {
		log.Printf("Failed to send message with customer agent: %v", err)
	}

	// 使用 study-assistant 应用发送消息
	studyCtx := wecom.WithAgentName(ctx, "study-assistant")
	err = sendMessageWithAgent(client1, studyCtx, "study-assistant")
	if err != nil {
		log.Printf("Failed to send message with study-assistant agent: %v", err)
	}

	// 也可以使用应用ID来指定应用
	agentIDCtx := wecom.WithAgentID(ctx, 100001)
	err = sendMessageWithAgent(client2, agentIDCtx, "agent with ID 100001")
	if err != nil {
		log.Printf("Failed to send message with agent ID: %v", err)
	}

	// 获取应用详情
	err = getAgentInfo(client1, customerCtx, 100001)
	if err != nil {
		log.Printf("Failed to get agent info: %v", err)
	}
}

func sendMessageWithAgent(client *wecom.Client, ctx context.Context, agentName string) error {
	fmt.Printf("\n=== 使用 %s 发送消息 ===\n", agentName)

	// 这里只是示例，实际需要替换为真实的用户ID
	// err := client.Message.SendText(ctx, &types.TextMessageRequest{
	// 	ToUser:  "user_id",
	// 	AgentID: 100001,
	// 	Text: types.TextContent{
	// 		Content: "测试消息",
	// 	},
	// })

	fmt.Printf("消息发送成功 (使用应用: %s)\n", agentName)
	return nil
}

func getAgentInfo(client *wecom.Client, ctx context.Context, agentID int64) error {
	fmt.Printf("\n=== 获取应用信息 (AgentID: %d) ===\n", agentID)

	info, err := client.Agent.Get(ctx, int(agentID))
	if err != nil {
		return fmt.Errorf("failed to get agent info: %w", err)
	}

	fmt.Printf("应用名称: %s\n", info.Name)
	fmt.Printf("应用描述: %s\n", info.Description)

	return nil
}
