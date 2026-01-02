package main

import (
	"context"
	"fmt"
	"log"

	wecom "github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/types/updown"
)

func main() {
	// 创建企业微信客户端
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithCorpSecret("your_corp_secret"),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	// 示例1: 获取上下游列表
	chains, err := client.UpDown.GetChainList(ctx)
	if err != nil {
		log.Printf("获取上下游列表失败: %v", err)
	} else {
		fmt.Printf("上下游列表: %+v\n", chains.Chains)
	}

	// 示例2: 获取上下游通讯录分组
	if len(chains.Chains) > 0 {
		chainID := chains.Chains[0].ChainID
		groups, err := client.UpDown.GetChainGroup(ctx, &updown.GetChainGroupRequest{
			ChainID: chainID,
		})
		if err != nil {
			log.Printf("获取上下游通讯录分组失败: %v", err)
		} else {
			fmt.Printf("通讯录分组: %+v\n", groups.Groups)
		}

		// 示例3: 获取企业详情列表
		corpList, err := client.UpDown.GetChainCorpInfoList(ctx, &updown.GetChainCorpInfoListRequest{
			ChainID: chainID,
		})
		if err != nil {
			log.Printf("获取企业详情列表失败: %v", err)
		} else {
			fmt.Printf("企业详情列表: %+v\n", corpList.GroupCorps)
		}

		// 示例4: 新增对接规则
		ruleID, err := client.UpDown.AddRule(ctx, &updown.AddRuleRequest{
			ChainID: chainID,
			RuleInfo: &updown.RuleInfo{
				OwnerCorpRange: &updown.OwnerCorpRange{
					UserIDs: []string{"user1", "user2"},
				},
				MemberCorpRange: &updown.MemberCorpRange{
					CorpIDs: []string{"corp1", "corp2"},
				},
			},
		})
		if err != nil {
			log.Printf("新增对接规则失败: %v", err)
		} else {
			fmt.Printf("新增对接规则成功，规则ID: %d\n", ruleID)
		}

		// 示例5: 获取对接规则id列表
		ruleIDs, err := client.UpDown.ListRuleIDs(ctx, &updown.ListRuleIDsRequest{
			ChainID: chainID,
		})
		if err != nil {
			log.Printf("获取对接规则id列表失败: %v", err)
		} else {
			fmt.Printf("对接规则ID列表: %v\n", ruleIDs)
		}

		// 示例6: 查询成员自定义id
		if len(corpList.GroupCorps) > 0 && corpList.GroupCorps[0].CorpID != "" {
			userCustomID, err := client.UpDown.GetChainUserCustomID(ctx, &updown.GetChainUserCustomIDRequest{
				ChainID: chainID,
				CorpID:  corpList.GroupCorps[0].CorpID,
				UserID:  "some_user_id",
			})
			if err != nil {
				log.Printf("查询成员自定义id失败: %v", err)
			} else {
				fmt.Printf("成员自定义ID: %s\n", userCustomID)
			}
		}

		// 示例7: 获取应用共享信息
		appShareInfo, err := client.UpDown.ListAppShareInfo(ctx, &updown.ListAppShareInfoRequest{
			AgentID:      1000001,
			BusinessType: new(int), // 0: 企业互联/局校互联, 1: 上下游企业
		})
		if err != nil {
			log.Printf("获取应用共享信息失败: %v", err)
		} else {
			fmt.Printf("应用共享信息: %+v\n", appShareInfo.CorpList)
		}

		// 示例8: 批量导入上下游联系人
		jobID, err := client.UpDown.ImportChainContact(ctx, &updown.ImportChainContactRequest{
			ChainID: chainID,
			ContactList: []updown.ChainContact{
				{
					CorpName:  "测试企业",
					GroupPath: "华北区/北京市",
					CustomID:  "custom_001",
					ContactInfoList: []updown.ContactInfo{
						{
							Name:         "张三",
							IdentityType: 2, // 负责人
							Mobile:       "13800138000",
							UserCustomID: "100",
						},
					},
				},
			},
		})
		if err != nil {
			log.Printf("批量导入上下游联系人失败: %v", err)
		} else {
			fmt.Printf("批量导入任务已提交，任务ID: %s\n", jobID)

			// 示例9: 获取异步任务结果
			taskResult, err := client.UpDown.GetTaskResult(ctx, jobID)
			if err != nil {
				log.Printf("获取异步任务结果失败: %v", err)
			} else {
				fmt.Printf("任务状态: %d, 结果: %+v\n", taskResult.Status, taskResult.Result)
			}
		}

		// 示例10: unionid查询pending_id
		pendingID, err := client.UpDown.UnionIDToPendingID(ctx, &updown.UnionIDToPendingIDRequest{
			UnionID: "some_union_id",
			OpenID:  "some_open_id",
		})
		if err != nil {
			log.Printf("unionid查询pending_id失败: %v", err)
		} else {
			fmt.Printf("Pending ID: %s\n", pendingID)
		}
	}
}
