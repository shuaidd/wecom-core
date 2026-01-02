package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/types/oauth"
)

func main() {
	// 从环境变量获取企业ID和Secret
	corpID := os.Getenv("WECOM_CORP_ID")
	corpSecret := os.Getenv("WECOM_CORP_SECRET")
	if corpID == "" || corpSecret == "" {
		log.Fatal("请设置环境变量: WECOM_CORP_ID 和 WECOM_CORP_SECRET")
	}

	// 创建企业微信客户端
	client, err := wecom.New(
		config.WithCorpID(corpID),
		config.WithCorpSecret(corpSecret),
		config.WithDebug(true),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	// 示例1: 构造网页授权链接（静默授权）
	fmt.Println("\n=== 示例1: 构造静默授权链接 ===")
	authURL, err := client.OAuth.BuildAuthorizeURL(oauth.BuildAuthorizeURLParams{
		CorpID:      corpID,
		RedirectURI: "http://example.com/callback",
		Scope:       oauth.ScopeBase,
		State:       "STATE123",
		AgentID:     "1000002",
	})
	if err != nil {
		log.Fatalf("构造授权链接失败: %v", err)
	}
	fmt.Printf("静默授权链接: %s\n", authURL)

	// 示例2: 构造网页授权链接（手动授权，可获取敏感信息）
	fmt.Println("\n=== 示例2: 构造手动授权链接（含敏感信息） ===")
	authURLPrivate, err := client.OAuth.BuildAuthorizeURL(oauth.BuildAuthorizeURLParams{
		CorpID:      corpID,
		RedirectURI: "http://example.com/callback",
		Scope:       oauth.ScopePrivateInfo,
		State:       "STATE456",
		AgentID:     "1000002", // snsapi_privateinfo 必须填写agentid
	})
	if err != nil {
		log.Fatalf("构造授权链接失败: %v", err)
	}
	fmt.Printf("手动授权链接: %s\n", authURLPrivate)

	// 示例3: 获取访问用户身份
	// 注意: 这里的code需要从回调URL中获取
	// 实际使用时，应该在Web回调处理函数中调用此方法
	code := os.Getenv("WECOM_AUTH_CODE")
	if code != "" {
		fmt.Println("\n=== 示例3: 获取访问用户身份 ===")
		userInfo, err := client.OAuth.GetUserInfo(ctx, code)
		if err != nil {
			log.Printf("获取用户身份失败: %v", err)
		} else {
			fmt.Printf("用户ID: %s\n", userInfo.UserID)
			if userInfo.UserTicket != "" {
				fmt.Printf("用户票据: %s\n", userInfo.UserTicket)

				// 示例4: 获取访问用户敏感信息
				// 只有在scope为snsapi_privateinfo且用户在应用可见范围内时才会返回user_ticket
				fmt.Println("\n=== 示例4: 获取访问用户敏感信息 ===")
				userDetail, err := client.OAuth.GetUserDetail(ctx, userInfo.UserTicket)
				if err != nil {
					log.Printf("获取用户敏感信息失败: %v", err)
				} else {
					fmt.Printf("用户ID: %s\n", userDetail.UserID)
					fmt.Printf("性别: %s\n", userDetail.Gender)
					fmt.Printf("头像: %s\n", userDetail.Avatar)
					fmt.Printf("二维码: %s\n", userDetail.QRCode)
					fmt.Printf("手机: %s\n", userDetail.Mobile)
					fmt.Printf("邮箱: %s\n", userDetail.Email)
					fmt.Printf("企业邮箱: %s\n", userDetail.BizMail)
					fmt.Printf("地址: %s\n", userDetail.Address)
				}
			}

			if userInfo.OpenID != "" {
				fmt.Printf("OpenID（非企业成员）: %s\n", userInfo.OpenID)
			}
			if userInfo.ExternalUserID != "" {
				fmt.Printf("外部联系人ID: %s\n", userInfo.ExternalUserID)
			}
		}
	}

	// 示例5: 获取用户二次验证信息
	// 注意: 这个code是用户触发二次验证时企业微信颁发的code
	tfaCode := os.Getenv("WECOM_TFA_CODE")
	if tfaCode != "" {
		fmt.Println("\n=== 示例5: 获取用户二次验证信息 ===")
		tfaInfo, err := client.OAuth.GetTFAInfo(ctx, tfaCode)
		if err != nil {
			log.Printf("获取二次验证信息失败: %v", err)
		} else {
			fmt.Printf("用户ID: %s\n", tfaInfo.UserID)
			fmt.Printf("二次验证授权码: %s\n", tfaInfo.TFACode)
			// tfa_code有效期五分钟，且只能使用一次
			// 验证用户身份信息无误后，可以调用通过二次验证接口，解锁企业微信终端
		}
	}

	fmt.Println("\n=== OAuth示例执行完成 ===")
	fmt.Println("\n提示:")
	fmt.Println("- 要测试获取用户身份，请设置环境变量 WECOM_AUTH_CODE（从OAuth回调URL获取）")
	fmt.Println("- 要测试二次验证，请设置环境变量 WECOM_TFA_CODE（从二次验证页面获取）")
}
