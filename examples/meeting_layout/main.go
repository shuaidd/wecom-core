package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/pkg/logger"
	"github.com/shuaidd/wecom-core/types/meeting"
)

func main() {
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithCorpSecret("your_corp_secret"),
		config.WithLogger(logger.NewStdLogger()),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := context.Background()
	meetingID := "your_meeting_id"

	// ==================== 布局管理示例 ====================

	// 1. 获取会议布局列表
	layoutList, err := client.Meeting.ListLayout(ctx, meetingID)
	if err != nil {
		log.Fatalf("获取会议布局列表失败: %v", err)
	}
	fmt.Printf("当前应用的布局ID: %s\n", layoutList.SelectedLayoutID)
	fmt.Printf("布局数量: %d\n", len(layoutList.LayoutList))

	// 2. 获取布局模板列表
	templates, err := client.Meeting.ListLayoutTemplate(ctx)
	if err != nil {
		log.Fatalf("获取布局模板列表失败: %v", err)
	}
	fmt.Printf("布局模板数量: %d\n", len(templates.LayoutTemplateList))
	for _, tpl := range templates.LayoutTemplateList {
		fmt.Printf("  模板ID: %s, 缩略图: %s\n", tpl.LayoutTemplateID, tpl.ThumbnailURL)
	}

	// 3. 添加会议基础布局
	addBasicResp, err := client.Meeting.AddBasicLayout(ctx, &meeting.AddBasicLayoutRequest{
		MeetingID: meetingID,
		LayoutList: []meeting.BasicLayoutInput{
			{
				PageList: []meeting.BasicLayoutPageInput{
					{
						LayoutTemplateID: "1",
						UserSeatList: []meeting.UserSeat{
							{
								GridID:   "1",
								GridType: 1,
								UserID:   "zhangsan",
								NickName: "张三",
							},
						},
					},
				},
			},
		},
		DefaultLayoutOrder: 1,
	})
	if err != nil {
		log.Fatalf("添加会议基础布局失败: %v", err)
	}
	fmt.Printf("添加布局成功，应用的布局ID: %s\n", addBasicResp.SelectedLayoutID)

	// 4. 添加会议高级布局
	addAdvancedResp, err := client.Meeting.AddAdvancedLayout(ctx, &meeting.AddAdvancedLayoutRequest{
		MeetingID: meetingID,
		LayoutList: []meeting.AdvancedLayoutInput{
			{
				LayoutName: "自定义布局",
				PageList: []meeting.LayoutPage{
					{
						LayoutTemplateID: "3",
						EnablePolling:    true,
						PollingSetting: &meeting.PollingSetting{
							PollingIntervalUnit: 1,
							PollingInterval:     10,
							IgnoreUserNoVideo:   false,
							IgnoreUserAbsence:   false,
						},
						UserSeatList: []meeting.UserSeat{
							{
								GridID:    "1",
								GridType:  1,
								VideoType: 3,
								UserList: []meeting.GridUser{
									{
										UserID:   "zhangsan",
										NickName: "dGVzdA==", // base64编码的"test"
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("添加会议高级布局失败: %v", err)
	}
	fmt.Printf("添加高级布局成功，布局数量: %d\n", len(addAdvancedResp.LayoutList))

	// 5. 修改会议基础布局
	if len(addBasicResp.LayoutList) > 0 {
		layoutID := addBasicResp.LayoutList[0].LayoutID
		err = client.Meeting.UpdateBasicLayout(ctx, &meeting.UpdateBasicLayoutRequest{
			MeetingID: meetingID,
			LayoutID:  layoutID,
			PageList: []meeting.BasicLayoutPageInput{
				{
					LayoutTemplateID: "2",
					UserSeatList: []meeting.UserSeat{
						{
							GridID:   "1",
							GridType: 1,
							UserID:   "lisi",
							NickName: "李四",
						},
					},
				},
			},
			EnableSetDefault: true,
		})
		if err != nil {
			log.Fatalf("修改会议基础布局失败: %v", err)
		}
		fmt.Println("修改基础布局成功")
	}

	// 6. 设置会议默认布局
	err = client.Meeting.SetDefaultLayout(ctx, &meeting.SetDefaultLayoutRequest{
		MeetingID:        meetingID,
		SelectedLayoutID: addBasicResp.SelectedLayoutID,
	})
	if err != nil {
		log.Fatalf("设置会议默认布局失败: %v", err)
	}
	fmt.Println("设置默认布局成功")

	// 7. 设置高级布局（应用到会议）
	err = client.Meeting.ApplyAdvancedLayout(ctx, &meeting.ApplyAdvancedLayoutRequest{
		MeetingID: meetingID,
		LayoutID:  addAdvancedResp.LayoutList[0].LayoutID,
	})
	if err != nil {
		log.Fatalf("应用高级布局失败: %v", err)
	}
	fmt.Println("应用高级布局成功")

	// 8. 获取用户布局
	userLayout, err := client.Meeting.GetUserLayout(ctx, &meeting.GetUserLayoutRequest{
		MeetingID:  meetingID,
		TmpOpenID:  "tmp_openid_123",
		InstanceID: 1,
	})
	if err != nil {
		log.Fatalf("获取用户布局失败: %v", err)
	}
	fmt.Printf("用户布局ID: %s, 布局类型: %d\n", userLayout.SelectedLayoutID, userLayout.LayoutType)

	// 9. 批量删除布局
	err = client.Meeting.BatchDeleteLayout(ctx, &meeting.BatchDeleteLayoutRequest{
		MeetingID:    meetingID,
		LayoutIDList: []string{"layout_id_1", "layout_id_2"},
	})
	if err != nil {
		log.Fatalf("批量删除布局失败: %v", err)
	}
	fmt.Println("批量删除布局成功")

	// ==================== 背景管理示例 ====================

	// 1. 获取会议背景列表
	bgList, err := client.Meeting.ListBackground(ctx, meetingID)
	if err != nil {
		log.Fatalf("获取会议背景列表失败: %v", err)
	}
	fmt.Printf("当前应用的背景ID: %s\n", bgList.SelectedBackgroundID)
	fmt.Printf("背景数量: %d\n", len(bgList.BackgroundList))

	// 2. 添加会议背景
	addBgResp, err := client.Meeting.AddBackground(ctx, &meeting.AddBackgroundRequest{
		MeetingID: meetingID,
		ImageList: []meeting.BackgroundImage{
			{
				ImageMD5: "abc123def456",
				ImageURL: "https://example.com/background1.png",
			},
			{
				ImageMD5: "xyz789uvw012",
				ImageURL: "https://example.com/background2.png",
			},
		},
		DefaultImageOrder: 1,
	})
	if err != nil {
		log.Fatalf("添加会议背景失败: %v", err)
	}
	fmt.Printf("添加背景成功，应用的背景ID: %s\n", addBgResp.SelectedBackgroundID)

	// 3. 设置会议默认背景
	if len(addBgResp.BackgroundList) > 0 {
		err = client.Meeting.SetDefaultBackground(ctx, &meeting.SetDefaultBackgroundRequest{
			MeetingID:            meetingID,
			SelectedBackgroundID: addBgResp.BackgroundList[0].BackgroundID,
		})
		if err != nil {
			log.Fatalf("设置会议默认背景失败: %v", err)
		}
		fmt.Println("设置默认背景成功")
	}

	// 4. 删除会议背景
	if len(addBgResp.BackgroundList) > 0 {
		err = client.Meeting.DeleteBackground(ctx, &meeting.DeleteBackgroundRequest{
			MeetingID:    meetingID,
			BackgroundID: addBgResp.BackgroundList[0].BackgroundID,
		})
		if err != nil {
			log.Fatalf("删除会议背景失败: %v", err)
		}
		fmt.Println("删除背景成功")
	}

	// 5. 批量删除会议背景
	err = client.Meeting.BatchDeleteBackground(ctx, &meeting.BatchDeleteBackgroundRequest{
		MeetingID:        meetingID,
		BackgroundIDList: []string{"bg_id_1", "bg_id_2"},
	})
	if err != nil {
		log.Fatalf("批量删除会议背景失败: %v", err)
	}
	fmt.Println("批量删除背景成功")
}
