package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

func main() {
	// 创建企业微信客户端
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithCorpSecret("your_corp_secret"),
		config.WithLogger(logger.NewStdLogger()),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := context.Background()

	// 读取成员信息
	user, err := client.Contact.GetUser(ctx, "zhangsan")
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	fmt.Printf("成员信息: UserID=%s, Name=%s, Mobile=%s\n",
		user.UserID, user.Name, user.Mobile)
}
