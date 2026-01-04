package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/types/agent"
)

func main() {
	// 从环境变量获取配置
	corpID := os.Getenv("WECOM_CORP_ID")
	corpSecret := os.Getenv("WECOM_CORP_SECRET")
	agentID := 1000002 // 替换为你的应用ID

	if corpID == "" || corpSecret == "" {
		log.Fatal("请设置环境变量 WECOM_CORP_ID 和 WECOM_CORP_SECRET")
	}

	// 创建客户端
	client, err := wecom.New(
		config.WithCorpID(corpID),
		config.WithCorpSecret(corpSecret),
		config.WithDebug(true),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	// 1. 获取应用详情
	fmt.Println("\n=== 获取应用详情 ===")
	agentInfo, err := client.Agent.Get(ctx, agentID)
	if err != nil {
		log.Fatalf("获取应用详情失败: %v", err)
	}
	fmt.Printf("应用ID: %d\n", agentInfo.AgentID)
	fmt.Printf("应用名称: %s\n", agentInfo.Name)
	fmt.Printf("应用描述: %s\n", agentInfo.Description)
	fmt.Printf("应用主页: %s\n", agentInfo.HomeURL)

	// 2. 获取应用列表
	fmt.Println("\n=== 获取应用列表 ===")
	agentList, err := client.Agent.List(ctx)
	if err != nil {
		log.Fatalf("获取应用列表失败: %v", err)
	}
	for _, app := range agentList.AgentList {
		fmt.Printf("- [%d] %s\n", app.AgentID, app.Name)
	}

	// 3. 设置应用（可选）
	fmt.Println("\n=== 设置应用 ===")
	err = client.Agent.Set(ctx, &agent.SetAgentRequest{
		AgentID:     agentID,
		Name:        "更新后的应用名称",
		Description: "这是一个测试应用",
	})
	if err != nil {
		log.Printf("设置应用失败: %v", err)
	} else {
		fmt.Println("应用设置成功")
	}

	// 4. 创建菜单
	fmt.Println("\n=== 创建菜单 ===")
	err = client.Agent.CreateMenu(ctx, &agent.CreateMenuRequest{
		AgentID: agentID,
		Button: []agent.MenuButton{
			{
				Type: "click",
				Name: "今日歌曲",
				Key:  "V1001_TODAY_MUSIC",
			},
			{
				Name: "菜单",
				SubButton: []agent.MenuButton{
					{
						Type: "view",
						Name: "搜索",
						URL:  "http://www.soso.com/",
					},
					{
						Type: "click",
						Name: "赞一下我们",
						Key:  "V1001_GOOD",
					},
				},
			},
		},
	})
	if err != nil {
		log.Printf("创建菜单失败: %v", err)
	} else {
		fmt.Println("菜单创建成功")
	}

	// 5. 获取菜单
	fmt.Println("\n=== 获取菜单 ===")
	menu, err := client.Agent.GetMenu(ctx, agentID)
	if err != nil {
		log.Printf("获取菜单失败: %v", err)
	} else {
		fmt.Printf("菜单按钮数量: %d\n", len(menu.Button))
	}

	// 6. 设置工作台模板（图片型）
	fmt.Println("\n=== 设置工作台模板 ===")
	err = client.Agent.SetWorkbenchTemplate(ctx, &agent.SetWorkbenchTemplateRequest{
		AgentID: agentID,
		Type:    agent.WorkbenchTypeImage,
		Image: &agent.ImageTemplate{
			URL:     "https://example.com/image.png",
			JumpURL: "https://example.com",
		},
		ReplaceUserData: false,
	})
	if err != nil {
		log.Printf("设置工作台模板失败: %v", err)
	} else {
		fmt.Println("工作台模板设置成功")
	}

	// 7. 获取工作台模板
	fmt.Println("\n=== 获取工作台模板 ===")
	template, err := client.Agent.GetWorkbenchTemplate(ctx, agentID)
	if err != nil {
		log.Printf("获取工作台模板失败: %v", err)
	} else {
		fmt.Printf("模板类型: %s\n", template.Type)
	}

	// 8. 设置用户工作台数据（关键数据型）
	fmt.Println("\n=== 设置用户工作台数据 ===")
	err = client.Agent.SetWorkbenchData(ctx, &agent.SetWorkbenchDataRequest{
		AgentID: agentID,
		UserID:  "zhangsan", // 替换为实际的用户ID
		Type:    agent.WorkbenchTypeKeydata,
		Keydata: &agent.KeydataTemplate{
			Items: []agent.KeydataItem{
				{
					Key:     "待审批",
					Data:    "2",
					JumpURL: "https://example.com/approval",
				},
				{
					Key:     "待办事项",
					Data:    "5",
					JumpURL: "https://example.com/todo",
				},
			},
		},
	})
	if err != nil {
		log.Printf("设置用户工作台数据失败: %v", err)
	} else {
		fmt.Println("用户工作台数据设置成功")
	}

	// 9. 删除菜单（谨慎使用）
	// fmt.Println("\n=== 删除菜单 ===")
	// err = client.Agent.DeleteMenu(ctx, agentID)
	// if err != nil {
	// 	log.Printf("删除菜单失败: %v", err)
	// } else {
	// 	fmt.Println("菜单删除成功")
	// }

	fmt.Println("\n=== 示例运行完成 ===")
}
