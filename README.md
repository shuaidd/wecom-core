# wecom-core

企业微信 Go SDK - 简洁、易用、功能完善的企业微信开发工具包

## 特性

- ✅ **统一日志监控**：支持自定义日志实现，完整的请求追踪
- ✅ **统一响应处理**：自动解析 JSON 响应，统一错误处理
- ✅ **统一重试逻辑**：智能重试机制，支持指数退避
- ✅ **Token 自动管理**：自动获取、缓存、刷新 access_token
- ✅ **并发安全**：所有操作都是并发安全的
- ✅ **接口化设计**：支持自定义 Logger、Cache 实现
- ✅ **易于扩展**：清晰的架构设计，易于添加新模块

## 安装

```bash
go get github.com/shuaidd/wecom-core
```

## 快速开始

```go
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
```

## 配置选项

```go
client, err := wecom.New(
    // 必填：企业ID和应用密钥
    config.WithCorpID("your_corp_id"),
    config.WithCorpSecret("your_corp_secret"),

    // 可选：自定义日志
    config.WithLogger(logger.NewStdLogger()),

    // 可选：设置超时时间（默认 30 秒）
    config.WithTimeout(60 * time.Second),

    // 可选：设置重试次数（默认 3 次）
    config.WithRetry(5),

    // 可选：设置退避时间
    config.WithBackoff(1*time.Second, 30*time.Second),

    // 可选：自定义缓存（默认使用内存缓存）
    config.WithCache(yourCustomCache),
)
```

## 功能模块

### 通讯录管理

#### 成员管理

```go
// 创建成员
createResp, err := client.Contact.CreateUser(ctx, &contact.CreateUserRequest{
    UserID:     "zhangsan",
    Name:       "张三",
    Mobile:     "13800000000",
    Department: []int{1},
    Email:      "zhangsan@example.com",
    Position:   "产品经理",
})

// 读取成员
user, err := client.Contact.GetUser(ctx, "zhangsan")

// 更新成员
err = client.Contact.UpdateUser(ctx, &contact.UpdateUserRequest{
    UserID:   "zhangsan",
    Position: "高级产品经理",
})

// 删除成员
err = client.Contact.DeleteUser(ctx, "zhangsan")

// 获取部门成员列表
users, err := client.Contact.ListUsers(ctx, 1, false)

// 获取部门成员详情
usersDetail, err := client.Contact.ListUsersDetail(ctx, 1, false)
```

#### 部门管理

```go
// 创建部门
deptID, err := client.Contact.CreateDepartment(ctx, &contact.CreateDepartmentRequest{
    Name:     "研发部",
    ParentID: 1,
    Order:    1,
})

// 获取部门详情
dept, err := client.Contact.GetDepartment(ctx, deptID)

// 更新部门
err = client.Contact.UpdateDepartment(ctx, &contact.UpdateDepartmentRequest{
    ID:   deptID,
    Name: "技术研发部",
})

// 删除部门
err = client.Contact.DeleteDepartment(ctx, deptID)

// 获取部门列表
departments, err := client.Contact.ListDepartments(ctx, 1)
```

### 外部联系人管理

#### 客户管理

```go
// 获取客户列表
contacts, err := client.ExternalContact.ListExternalContact(ctx, "zhangsan")

// 获取客户详情
detail, err := client.ExternalContact.GetExternalContact(ctx, "external_userid")

// 修改客户备注信息
err = client.ExternalContact.UpdateRemark(ctx, &externalcontact.UpdateRemarkRequest{
    UserID:         "zhangsan",
    ExternalUserID: "external_userid",
    Remark:         "重要客户",
    Description:    "产品负责人",
    RemarkCompany:  "某某科技公司",
})

// 批量获取客户详情
batchResp, err := client.ExternalContact.BatchGetByUser(ctx, &externalcontact.BatchGetByUserRequest{
    UserIDList: []string{"zhangsan", "lisi"},
    Limit:      100,
})
```

#### 客户标签管理

```go
// 获取企业标签库
tags, err := client.ExternalContact.GetCorpTagList(ctx, &externalcontact.GetCorpTagListRequest{})

// 添加企业客户标签
addResp, err := client.ExternalContact.AddCorpTag(ctx, &externalcontact.AddCorpTagRequest{
    GroupName: "客户类型",
    Tag: []externalcontact.AddCorpTagItem{
        {Name: "VIP客户", Order: 1},
        {Name: "潜在客户", Order: 2},
    },
})

// 为客户打标签
err = client.ExternalContact.MarkTag(ctx, &externalcontact.MarkTagRequest{
    UserID:         "zhangsan",
    ExternalUserID: "external_userid",
    AddTag:         []string{"tag_id_1", "tag_id_2"},
})

// 编辑企业客户标签
err = client.ExternalContact.EditCorpTag(ctx, &externalcontact.EditCorpTagRequest{
    ID:   "tag_id",
    Name: "核心客户",
})

// 删除企业客户标签
err = client.ExternalContact.DeleteCorpTag(ctx, &externalcontact.DeleteCorpTagRequest{
    TagID: []string{"tag_id_1", "tag_id_2"},
})
```

#### 客户群管理

```go
// 获取客户群列表
groups, err := client.ExternalContact.ListGroupChat(ctx, &externalcontact.ListGroupChatRequest{
    StatusFilter: 0,
    OwnerFilter: &externalcontact.OwnerFilter{
        UserIDList: []string{"zhangsan"},
    },
    Limit: 100,
})

// 获取客户群详情
groupDetail, err := client.ExternalContact.GetGroupChat(ctx, &externalcontact.GetGroupChatRequest{
    ChatID:   "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA",
    NeedName: 1,
})

// 客户群opengid转换
chatID, err := client.ExternalContact.OpenGIDToChatID(ctx, &externalcontact.OpenGIDToChatIDRequest{
    OpenGID: "oAAAAAAA",
})
```

#### 联系我与客户入群方式

```go
// 配置客户联系「联系我」方式
contactWayResp, err := client.ExternalContact.AddContactWay(ctx, &externalcontact.AddContactWayRequest{
    Type:       1,  // 1-单人, 2-多人
    Scene:      2,  // 1-在小程序中联系, 2-通过二维码联系
    SkipVerify: true,
    State:      "channel_001",
    User:       []string{"zhangsan"},
    Remark:     "市场推广活动",
})

// 获取企业已配置的「联系我」方式
contactWay, err := client.ExternalContact.GetContactWay(ctx, "config_id")

// 获取企业已配置的「联系我」列表
contactWayList, err := client.ExternalContact.ListContactWay(ctx, &externalcontact.ListContactWayRequest{
    Limit: 100,
})

// 更新企业已配置的「联系我」方式
err = client.ExternalContact.UpdateContactWay(ctx, &externalcontact.UpdateContactWayRequest{
    ConfigID: "config_id",
    Remark:   "更新后的备注",
})

// 删除企业已配置的「联系我」方式
err = client.ExternalContact.DeleteContactWay(ctx, "config_id")

// 结束临时会话
err = client.ExternalContact.CloseTempChat(ctx, "zhangsan", "external_userid")

// 配置客户群进群方式
joinWayResp, err := client.ExternalContact.AddJoinWay(ctx, &externalcontact.AddJoinWayRequest{
    Scene:          2,  // 1-群的小程序插件, 2-群的二维码插件
    Remark:         "产品交流群",
    AutoCreateRoom: 1,
    RoomBaseName:   "产品交流群",
    RoomBaseID:     1,
    ChatIDList:     []string{"wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA"},
    State:          "channel_group_001",
})

// 获取客户群进群方式配置
joinWay, err := client.ExternalContact.GetJoinWay(ctx, "config_id")

// 更新客户群进群方式配置
err = client.ExternalContact.UpdateJoinWay(ctx, &externalcontact.UpdateJoinWayRequest{
    ConfigID: "config_id",
    Scene:    2,
    Remark:   "更新后的备注",
    ChatIDList: []string{"wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA"},
})

// 删除客户群进群方式配置
err = client.ExternalContact.DeleteJoinWay(ctx, "config_id")
```

#### 企业服务人员管理

```go
// 获取配置了客户联系功能的成员列表
followUsers, err := client.ExternalContact.GetFollowUserList(ctx)
fmt.Printf("配置了客户联系功能的成员: %v\n", followUsers.FollowUser)
```

#### 统计管理

```go
// 获取群聊数据统计（按群主聚合）
groupChatStats, err := client.ExternalContact.GetGroupChatStatistic(ctx, &externalcontact.GroupChatStatisticRequest{
    DayBeginTime: 1600272000,
    DayEndTime:   1600444800,
    OwnerFilter: &externalcontact.OwnerFilter{
        UserIDList: []string{"zhangsan"},
    },
    Limit: 100,
})

// 获取群聊数据统计（按自然日聚合）
groupChatStatsByDay, err := client.ExternalContact.GetGroupChatStatisticGroupByDay(ctx, &externalcontact.GroupChatStatisticGroupByDayRequest{
    DayBeginTime: 1600272000,
    DayEndTime:   1600358400,
    OwnerFilter: &externalcontact.OwnerFilter{
        UserIDList: []string{"zhangsan"},
    },
})

// 获取联系客户统计数据
behaviorData, err := client.ExternalContact.GetUserBehaviorData(ctx, &externalcontact.GetUserBehaviorDataRequest{
    UserID:    []string{"zhangsan", "lisi"},
    StartTime: 1536508800,
    EndTime:   1536595200,
})
```

#### 客户朋友圈

```go
// 创建发表任务
taskResp, err := client.ExternalContact.AddMomentTask(ctx, &externalcontact.AddMomentTaskRequest{
    Text: &externalcontact.MomentText{
        Content: "分享产品动态",
    },
    Attachments: []externalcontact.MomentAttachment{
        {
            MsgType: "image",
            Image: &externalcontact.MomentImage{
                MediaID: "MEDIA_ID",
            },
        },
    },
    VisibleRange: &externalcontact.VisibleRange{
        SenderList: &externalcontact.SenderList{
            UserList: []string{"zhangsan", "lisi"},
        },
    },
})

// 获取任务创建结果
result, err := client.ExternalContact.GetMomentTaskResult(ctx, taskResp.JobID)

// 获取企业全部的发表列表
moments, err := client.ExternalContact.GetMomentList(ctx, &externalcontact.GetMomentListRequest{
    StartTime:  1605000000,
    EndTime:    1605172726,
    FilterType: 0,
    Limit:      20,
})

// 获取客户朋友圈企业发表的列表
tasks, err := client.ExternalContact.GetMomentTask(ctx, &externalcontact.GetMomentTaskRequest{
    MomentID: "moment_id",
    Limit:    100,
})

// 获取客户朋友圈发表时选择的可见范围
customers, err := client.ExternalContact.GetMomentCustomerList(ctx, &externalcontact.GetMomentCustomerListRequest{
    MomentID: "moment_id",
    UserID:   "zhangsan",
    Limit:    100,
})

// 获取客户朋友圈发表后的可见客户列表
sendResult, err := client.ExternalContact.GetMomentSendResult(ctx, &externalcontact.GetMomentSendResultRequest{
    MomentID: "moment_id",
    UserID:   "zhangsan",
    Limit:    100,
})

// 获取客户朋友圈的互动数据
comments, err := client.ExternalContact.GetMomentComments(ctx, &externalcontact.GetMomentCommentsRequest{
    MomentID: "moment_id",
    UserID:   "zhangsan",
})

// 停止发表企业朋友圈
err = client.ExternalContact.CancelMomentTask(ctx, &externalcontact.CancelMomentTaskRequest{
    MomentID: "moment_id",
})

// 获取朋友圈规则组列表
strategies, err := client.ExternalContact.ListMomentStrategy(ctx, &externalcontact.ListMomentStrategyRequest{
    Limit: 100,
})

// 创建朋友圈规则组
createResp, err := client.ExternalContact.CreateMomentStrategy(ctx, &externalcontact.CreateMomentStrategyRequest{
    StrategyName: "销售团队朋友圈",
    AdminList:    []string{"zhangsan", "lisi"},
    Range: []externalcontact.MomentStrategyRange{
        {Type: 1, UserID: "zhangsan"},
        {Type: 2, PartyID: 1},
    },
})

// 获取朋友圈规则组详情
strategyDetail, err := client.ExternalContact.GetMomentStrategy(ctx, &externalcontact.GetMomentStrategyRequest{
    StrategyID: strategyID,
})

// 编辑朋友圈规则组
err = client.ExternalContact.EditMomentStrategy(ctx, &externalcontact.EditMomentStrategyRequest{
    StrategyID:   strategyID,
    StrategyName: "销售一组朋友圈",
})

// 删除朋友圈规则组
err = client.ExternalContact.DeleteMomentStrategy(ctx, &externalcontact.DeleteMomentStrategyRequest{
    StrategyID: strategyID,
})
```

#### 客户联系规则组管理

```go
// 获取规则组列表
strategies, err := client.ExternalContact.ListStrategy(ctx, &externalcontact.ListStrategyRequest{
    Limit: 100,
})

// 创建规则组
createResp, err := client.ExternalContact.CreateStrategy(ctx, &externalcontact.CreateStrategyRequest{
    StrategyName: "销售团队",
    AdminList:    []string{"zhangsan", "lisi"},
    Range: []externalcontact.StrategyRange{
        {Type: 1, UserID: "zhangsan"},
        {Type: 2, PartyID: 1},
    },
})

// 获取规则组详情
strategy, err := client.ExternalContact.GetStrategy(ctx, strategyID)

// 编辑规则组
err = client.ExternalContact.EditStrategy(ctx, &externalcontact.EditStrategyRequest{
    StrategyID:   strategyID,
    StrategyName: "销售一组",
})

// 删除规则组
err = client.ExternalContact.DeleteStrategy(ctx, strategyID)
```

#### 消息推送

```go
// 创建企业群发
msgResp, err := client.ExternalContact.AddMsgTemplate(ctx, &externalcontact.AddMsgTemplateRequest{
    ChatType: "single",  // single-发送给客户, group-发送给客户群
    ExternalUserID: []string{"external_userid_1", "external_userid_2"},
    Sender: "zhangsan",
    Text: &externalcontact.TextContent{
        Content: "文本消息内容",
    },
    Attachments: []externalcontact.Attachment{
        {
            MsgType: "image",
            Image: &externalcontact.ImageAttachment{
                MediaID: "MEDIA_ID",
            },
        },
    },
})

// 获取群发记录列表
msgList, err := client.ExternalContact.GetGroupMsgListV2(ctx, &externalcontact.GetGroupMsgListV2Request{
    ChatType:  "single",
    StartTime: 1605171726,
    EndTime:   1605172726,
    Limit:     50,
})

// 获取群发成员发送任务列表
taskList, err := client.ExternalContact.GetGroupMsgTask(ctx, &externalcontact.GetGroupMsgTaskRequest{
    MsgID: msgResp.MsgID,
    Limit: 100,
})

// 获取企业群发成员执行结果
sendResult, err := client.ExternalContact.GetGroupMsgSendResult(ctx, &externalcontact.GetGroupMsgSendResultRequest{
    MsgID:  msgResp.MsgID,
    UserID: "zhangsan",
    Limit:  100,
})

// 发送新客户欢迎语
err = client.ExternalContact.SendWelcomeMsg(ctx, &externalcontact.SendWelcomeMsgRequest{
    WelcomeCode: "CALLBACK_CODE",  // 来自添加外部联系人事件
    Text: &externalcontact.TextContent{
        Content: "你好，欢迎添加我为好友！",
    },
    Attachments: []externalcontact.Attachment{
        {
            MsgType: "link",
            Link: &externalcontact.LinkAttachment{
                Title: "产品介绍",
                URL:   "https://example.com",
            },
        },
    },
})

// 停止企业群发
err = client.ExternalContact.CancelGroupMsgSend(ctx, &externalcontact.CancelGroupMsgSendRequest{
    MsgID: msgResp.MsgID,
})

// 提醒成员群发
err = client.ExternalContact.RemindGroupMsgSend(ctx, &externalcontact.RemindGroupMsgSendRequest{
    MsgID: msgResp.MsgID,
})

// 添加入群欢迎语素材
templateResp, err := client.ExternalContact.AddGroupWelcomeTemplate(ctx, &externalcontact.AddGroupWelcomeTemplateRequest{
    Text: &externalcontact.TextContent{
        Content: "亲爱的%NICKNAME%用户，你好",
    },
    Image: &externalcontact.ImageAttachment{
        MediaID: "MEDIA_ID",
    },
})

// 编辑入群欢迎语素材
err = client.ExternalContact.EditGroupWelcomeTemplate(ctx, &externalcontact.EditGroupWelcomeTemplateRequest{
    TemplateID: templateResp.TemplateID,
    Text: &externalcontact.TextContent{
        Content: "更新后的欢迎语",
    },
})

// 获取入群欢迎语素材
template, err := client.ExternalContact.GetGroupWelcomeTemplate(ctx, &externalcontact.GetGroupWelcomeTemplateRequest{
    TemplateID: templateResp.TemplateID,
})

// 删除入群欢迎语素材
err = client.ExternalContact.DelGroupWelcomeTemplate(ctx, &externalcontact.DelGroupWelcomeTemplateRequest{
    TemplateID: templateResp.TemplateID,
})
```

#### 在职继承

```go
// 分配在职成员的客户
transferResp, err := client.ExternalContact.OnJobTransferCustomer(ctx, &externalcontact.OnJobTransferCustomerRequest{
    HandoverUserID: "zhangsan",  // 原跟进成员
    TakeoverUserID: "lisi",      // 接替成员
    ExternalUserID: []string{"external_userid_1", "external_userid_2"},
    TransferSuccessMsg: "您好，您的服务已升级，后续将由我的同事李四接替我的工作，继续为您服务。",
})

// 分配在职成员的客户群
groupTransferResp, err := client.ExternalContact.OnJobTransferGroupChat(ctx, &externalcontact.OnJobTransferGroupChatRequest{
    ChatIDList: []string{"chat_id_1", "chat_id_2"},
    NewOwner:   "lisi",  // 新群主
})

// 查询客户接替状态
resultResp, err := client.ExternalContact.GetTransferResult(ctx, &externalcontact.TransferResultRequest{
    HandoverUserID: "zhangsan",
    TakeoverUserID: "lisi",
})

// 遍历接替结果
for _, customer := range resultResp.Customer {
    switch customer.Status {
    case 1:
        fmt.Printf("客户 %s 接替完毕\n", customer.ExternalUserID)
    case 2:
        fmt.Printf("客户 %s 等待接替\n", customer.ExternalUserID)
    case 3:
        fmt.Printf("客户 %s 拒绝接替\n", customer.ExternalUserID)
    case 4:
        fmt.Printf("客户 %s 接替成员客户达到上限\n", customer.ExternalUserID)
    }
}
```

#### 商品图册管理

```go
// 创建商品图册
productResp, err := client.ExternalContact.AddProductAlbum(ctx, &externalcontact.AddProductAlbumRequest{
    Description: "世界上最好的商品",
    Price:       30000,  // 单位为分
    ProductSN:   "SN123456",
    Attachments: []externalcontact.ProductAttachment{
        {
            Type: "image",
            Image: &externalcontact.ImageAttachment{
                MediaID: "MEDIA_ID",
            },
        },
    },
})

// 获取商品图册
product, err := client.ExternalContact.GetProductAlbum(ctx, &externalcontact.GetProductAlbumRequest{
    ProductID: productResp.ProductID,
})

// 获取商品图册列表
productList, err := client.ExternalContact.GetProductAlbumList(ctx, &externalcontact.GetProductAlbumListRequest{
    Limit: 50,
})

// 编辑商品图册
err = client.ExternalContact.UpdateProductAlbum(ctx, &externalcontact.UpdateProductAlbumRequest{
    ProductID:   productResp.ProductID,
    Description: "更新后的商品描述",
    Price:       35000,
})

// 删除商品图册
err = client.ExternalContact.DeleteProductAlbum(ctx, &externalcontact.DeleteProductAlbumRequest{
    ProductID: productResp.ProductID,
})
```

#### 聊天敏感词管理

```go
// 新建敏感词规则
ruleResp, err := client.ExternalContact.AddInterceptRule(ctx, &externalcontact.AddInterceptRuleRequest{
    RuleName: "敏感词规则1",
    WordList: []string{"敏感词1", "敏感词2"},
    SemanticsList: []int{1, 2, 3},  // 1：手机号、2：邮箱地址、3：红包
    InterceptType: 1,  // 1:警告并拦截发送；2:仅发警告
    ApplicableRange: &externalcontact.ApplicableRange{
        UserList:       []string{"zhangsan"},
        DepartmentList: []int{2, 3},
    },
})

// 获取敏感词规则列表
ruleList, err := client.ExternalContact.GetInterceptRuleList(ctx)

// 获取敏感词规则详情
ruleDetail, err := client.ExternalContact.GetInterceptRule(ctx, &externalcontact.GetInterceptRuleRequest{
    RuleID: ruleResp.RuleID,
})

// 修改敏感词规则
err = client.ExternalContact.UpdateInterceptRule(ctx, &externalcontact.UpdateInterceptRuleRequest{
    RuleID:   ruleResp.RuleID,
    RuleName: "更新后的规则名称",
    WordList: []string{"敏感词1", "敏感词2", "敏感词3"},
})

// 删除敏感词规则
err = client.ExternalContact.DelInterceptRule(ctx, &externalcontact.DelInterceptRuleRequest{
    RuleID: ruleResp.RuleID,
})
```

#### 获取已服务的外部联系人

```go
// 获取已服务的外部联系人
contactListResp, err := client.ExternalContact.GetContactList(ctx, &externalcontact.GetContactListRequest{
    Limit: 1000,
})

// 遍历结果
for _, info := range contactListResp.InfoList {
    if info.IsCustomer {
        fmt.Printf("客户: %s, 添加人: %s\n", info.ExternalUserID, info.FollowUserID)
    } else {
        fmt.Printf("其他外部联系人: %s, 添加人: %s\n", info.Name, info.FollowUserID)
    }
}

// 处理分页
if contactListResp.NextCursor != "" {
    // 获取下一页
    nextPageResp, err := client.ExternalContact.GetContactList(ctx, &externalcontact.GetContactListRequest{
        Cursor: contactListResp.NextCursor,
        Limit:  1000,
    })
    _ = nextPageResp
    _ = err
}
```

### 素材管理

企业微信素材管理服务，支持图片、语音、视频、文件等媒体资源的上传和下载。

#### 上传图片

```go
// 上传图片（本地文件）
imageResp, err := client.Media.UploadImage(ctx, "/path/to/image.jpg")
if err != nil {
    log.Fatalf("上传图片失败: %v", err)
}
fmt.Printf("图片URL（永久有效）: %s\n", imageResp.URL)

// 从 io.Reader 上传图片
file, _ := os.Open("/path/to/image.jpg")
defer file.Close()
imageResp, err = client.Media.UploadImageFromReader(ctx, file, "image.jpg")
```

#### 上传临时素材

```go
// 上传临时素材（本地文件）- 有效期3天
mediaResp, err := client.Media.UploadMedia(ctx, media.MediaTypeImage, "/path/to/image.jpg")
if err != nil {
    log.Fatalf("上传素材失败: %v", err)
}
fmt.Printf("MediaID: %s（3天内有效）\n", mediaResp.MediaID)

// 从 io.Reader 上传临时素材
file, _ := os.Open("/path/to/video.mp4")
defer file.Close()
mediaResp, err = client.Media.UploadMediaFromReader(ctx, media.MediaTypeVideo, file, "video.mp4")
```

支持的媒体类型：
- **图片（image）**: 10MB，支持JPG、PNG格式
- **语音（voice）**: 2MB，播放长度不超过60s，支持AMR格式
- **视频（video）**: 10MB，支持MP4格式
- **普通文件（file）**: 20MB

#### 获取临时素材

```go
// 获取临时素材
mediaData, err := client.Media.GetMedia(ctx, "MEDIA_ID")
if err != nil {
    log.Fatalf("获取素材失败: %v", err)
}
// 保存到文件
os.WriteFile("/path/to/download.jpg", mediaData, 0644)

// 使用Range分块下载（大文件）
mediaData, err = client.Media.GetMediaWithRange(ctx, "MEDIA_ID", "bytes=0-1048575")
```

#### 获取高清语音素材

```go
// 获取从JSSDK上传的高清语音素材（speex格式，16K采样率）
voiceData, err := client.Media.GetJSSDKMedia(ctx, "MEDIA_ID")
if err != nil {
    log.Fatalf("获取语音素材失败: %v", err)
}
os.WriteFile("/path/to/voice.speex", voiceData, 0644)
```

#### 异步上传大文件

```go
// 异步上传临时素材（支持最高200M）
uploadResp, err := client.Media.UploadByURL(ctx, &media.UploadByURLRequest{
    Scene:    1,  // 1-客户联系入群欢迎语素材
    Type:     "video",
    Filename: "large_video.mp4",
    URL:      "https://example.com/large_video.mp4",  // 必须支持Range分块下载
    MD5:      "file_md5_hash",
})
if err != nil {
    log.Fatalf("创建上传任务失败: %v", err)
}

// 查询异步上传任务结果
result, err := client.Media.GetUploadByURLResult(ctx, uploadResp.JobID)
if err != nil {
    log.Fatalf("查询任务失败: %v", err)
}

switch result.Status {
case media.UploadTaskStatusProcessing:
    fmt.Println("任务处理中...")
case media.UploadTaskStatusCompleted:
    fmt.Printf("上传成功，MediaID: %s\n", result.Detail.MediaID)
case media.UploadTaskStatusFailed:
    fmt.Printf("上传失败: %s\n", result.Detail.ErrMsg)
}
```

### 电子发票

企业微信电子发票管理服务，支持查询和更新电子发票的报销状态。

#### 查询电子发票

```go
// 查询单个电子发票
invoiceInfo, err := client.Invoice.GetInvoiceInfo(ctx, "CARD_ID", "ENCRYPT_CODE")
if err != nil {
    log.Fatalf("查询发票失败: %v", err)
}

fmt.Printf("发票类型: %s\n", invoiceInfo.Type)
fmt.Printf("发票抬头: %s\n", invoiceInfo.UserInfo.Title)
fmt.Printf("发票金额: %.2f元\n", float64(invoiceInfo.UserInfo.Fee)/100)
fmt.Printf("开票时间: %d\n", invoiceInfo.UserInfo.BillingTime)
fmt.Printf("发票号码: %s\n", invoiceInfo.UserInfo.BillingCode)
fmt.Printf("PDF链接: %s\n", invoiceInfo.UserInfo.PdfURL)
fmt.Printf("报销状态: %s\n", invoiceInfo.UserInfo.ReimburseStatus)

// 批量查询电子发票
batchResp, err := client.Invoice.GetInvoiceInfoBatch(ctx, []invoice.InvoiceItem{
    {CardID: "CARD_ID_1", EncryptCode: "ENCRYPT_CODE_1"},
    {CardID: "CARD_ID_2", EncryptCode: "ENCRYPT_CODE_2"},
})
if err != nil {
    log.Fatalf("批量查询发票失败: %v", err)
}

for _, inv := range batchResp.ItemList {
    fmt.Printf("发票: %s, 金额: %.2f元\n",
        inv.UserInfo.BillingCode,
        float64(inv.UserInfo.Fee)/100)
}
```

#### 更新发票状态

```go
// 更新单个发票状态 - 锁定发票
err = client.Invoice.UpdateInvoiceStatus(ctx,
    "CARD_ID",
    "ENCRYPT_CODE",
    invoice.ReimburseStatusLock,
)
if err != nil {
    log.Fatalf("锁定发票失败: %v", err)
}
fmt.Println("发票已锁定")

// 更新单个发票状态 - 核销发票（不可逆操作）
err = client.Invoice.UpdateInvoiceStatus(ctx,
    "CARD_ID",
    "ENCRYPT_CODE",
    invoice.ReimburseStatusClosure,
)
if err != nil {
    log.Fatalf("核销发票失败: %v", err)
}
fmt.Println("发票已核销")

// 批量更新发票状态
err = client.Invoice.UpdateStatusBatch(ctx,
    "USER_OPENID",
    invoice.ReimburseStatusLock,
    []invoice.InvoiceItem{
        {CardID: "CARD_ID_1", EncryptCode: "ENCRYPT_CODE_1"},
        {CardID: "CARD_ID_2", EncryptCode: "ENCRYPT_CODE_2"},
    },
)
if err != nil {
    log.Fatalf("批量更新发票状态失败: %v", err)
}
fmt.Println("批量操作成功")
```

发票状态说明：
- **INVOICE_REIMBURSE_INIT**: 发票初始状态，未锁定
- **INVOICE_REIMBURSE_LOCK**: 发票已锁定，无法重复提交报销
- **INVOICE_REIMBURSE_CLOSURE**: 发票已核销，从用户卡包中移除（不可逆）

注意事项：
1. 报销方须保证在报销、锁定、解锁后及时将状态同步至微信端
2. 批量更新为事务性操作，任一发票更新失败则所有操作回滚
3. 报销状态为不可逆状态，请谨慎调用

### 消息管理

```go
// 发送应用消息（敬请期待更多示例）
```

### 应用管理

企业微信应用管理服务，支持应用设置、菜单管理和工作台自定义展示。

#### 应用信息管理

```go
// 获取应用详情
agentInfo, err := client.Agent.Get(ctx, agentID)
fmt.Printf("应用ID: %d, 名称: %s\n", agentInfo.AgentID, agentInfo.Name)

// 获取应用列表
agentList, err := client.Agent.List(ctx)
for _, app := range agentList.AgentList {
    fmt.Printf("应用: [%d] %s\n", app.AgentID, app.Name)
}

// 设置应用信息
err = client.Agent.Set(ctx, &agent.SetAgentRequest{
    AgentID:     agentID,
    Name:        "应用名称",
    Description: "应用描述",
    HomeURL:     "https://example.com",
})
```

#### 菜单管理

```go
// 创建应用菜单
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

// 获取应用菜单
menu, err := client.Agent.GetMenu(ctx, agentID)

// 删除应用菜单
err = client.Agent.DeleteMenu(ctx, agentID)
```

#### 工作台自定义展示

```go
// 设置工作台模板（图片型）
err = client.Agent.SetWorkbenchTemplate(ctx, &agent.SetWorkbenchTemplateRequest{
    AgentID: agentID,
    Type:    agent.WorkbenchTypeImage,
    Image: &agent.ImageTemplate{
        URL:     "https://example.com/image.png",
        JumpURL: "https://example.com",
    },
})

// 设置工作台模板（关键数据型）
err = client.Agent.SetWorkbenchTemplate(ctx, &agent.SetWorkbenchTemplateRequest{
    AgentID: agentID,
    Type:    agent.WorkbenchTypeKeydata,
    Keydata: &agent.KeydataTemplate{
        Items: []agent.KeydataItem{
            {Key: "待审批", Data: "2", JumpURL: "https://example.com/approval"},
            {Key: "待办事项", Data: "5", JumpURL: "https://example.com/todo"},
        },
    },
    ReplaceUserData: false,
})

// 获取工作台模板
template, err := client.Agent.GetWorkbenchTemplate(ctx, agentID)

// 设置用户工作台数据
err = client.Agent.SetWorkbenchData(ctx, &agent.SetWorkbenchDataRequest{
    AgentID: agentID,
    UserID:  "zhangsan",
    Type:    agent.WorkbenchTypeKeydata,
    Keydata: &agent.KeydataTemplate{
        Items: []agent.KeydataItem{
            {Key: "待审批", Data: "2", JumpURL: "https://example.com/approval"},
        },
    },
})

// 批量设置用户工作台数据
err = client.Agent.BatchSetWorkbenchData(ctx, &agent.BatchSetWorkbenchDataRequest{
    AgentID:    agentID,
    UserIDList: []string{"zhangsan", "lisi"},
    Data: &agent.WorkbenchUserData{
        Type: agent.WorkbenchTypeKeydata,
        Keydata: &agent.KeydataTemplate{
            Items: []agent.KeydataItem{
                {Key: "待审批", Data: "0"},
            },
        },
    },
})

// 获取用户工作台数据
userData, err := client.Agent.GetWorkbenchData(ctx, agentID, "zhangsan")
```

支持的工作台模板类型：
- **关键数据型** (`WorkbenchTypeKeydata`)：展示关键业务数据
- **图片型** (`WorkbenchTypeImage`)：展示图片广告
- **列表型** (`WorkbenchTypeList`)：展示列表信息
- **Webview型** (`WorkbenchTypeWebview`)：嵌入网页内容

### 身份验证

```go
// OAuth 登录验证（敬请期待更多示例）
```

## 核心特性详解

### 自动 Token 管理

SDK 会自动处理 access_token 的获取、缓存和刷新：

- ✅ 首次调用时自动获取 token
- ✅ token 缓存在内存中（可自定义缓存实现）
- ✅ 提前 5 分钟自动刷新，避免过期
- ✅ 并发安全，防止重复获取
- ✅ token 失效时自动刷新并重试

### 智能重试机制

自动重试以下场景：

- ✅ Token 过期（errcode 40014, 42001）
- ✅ API 频率限制（errcode 45009）
- ✅ 系统繁忙（errcode 10001）
- ✅ 使用指数退避算法，避免频繁重试

### 统一日志记录

记录所有关键操作：

- 🔍 请求日志：URL、方法、耗时
- 🔍 响应日志：状态码、错误信息
- 🔍 Token 日志：获取、刷新、失效
- 🔍 重试日志：触发原因、次数

### 自定义 Logger

```go
// 实现 Logger 接口
type MyLogger struct{}

func (l *MyLogger) Debug(msg string, fields ...logger.Field) {}
func (l *MyLogger) Info(msg string, fields ...logger.Field)  {}
func (l *MyLogger) Warn(msg string, fields ...logger.Field)  {}
func (l *MyLogger) Error(msg string, fields ...logger.Field) {}

// 使用自定义 Logger
client, err := wecom.New(
    config.WithCorpID("your_corp_id"),
    config.WithCorpSecret("your_corp_secret"),
    config.WithLogger(&MyLogger{}),
)
```

### 自定义 Cache

```go
// 实现 Cache 接口
type MyCache struct{}

func (c *MyCache) Get(ctx context.Context, key string) (token string, expireAt time.Time, err error) {
    // 从 Redis 获取
}

func (c *MyCache) Set(ctx context.Context, key string, token string, expireAt time.Time) error {
    // 存储到 Redis
}

func (c *MyCache) Delete(ctx context.Context, key string) error {
    // 从 Redis 删除
}

// 使用自定义 Cache
client, err := wecom.New(
    config.WithCorpID("your_corp_id"),
    config.WithCorpSecret("your_corp_secret"),
    config.WithCache(&MyCache{}),
)
```

## 错误处理

```go
user, err := client.Contact.GetUser(ctx, "zhangsan")
if err != nil {
    // 判断是否为企业微信错误
    if errors.IsWecomError(err) {
        errCode := errors.GetErrorCode(err)
        // 根据错误码处理
        switch errCode {
        case errors.ErrCodeUserNotFound:
            // 成员不存在
        case errors.ErrCodeInvalidParameter:
            // 参数错误
        default:
            // 其他错误
        }
    }
    return err
}
```

## 项目结构

```
wecom-core/
├── wecom.go                    # 主入口
├── config/                     # 配置管理
├── internal/                   # 内部包（不对外暴露）
│   ├── client/                # HTTP 客户端
│   ├── auth/                  # Token 管理
│   ├── retry/                 # 重试逻辑
│   └── errors/                # 错误处理
├── pkg/                        # 公共包（可被外部引用）
│   ├── logger/                # 日志接口
│   └── cache/                 # 缓存接口
├── types/                      # 数据类型定义
│   ├── common/                # 通用类型
│   └── contact/               # 通讯录相关
└── services/                   # 业务服务
    └── contact/               # 通讯录服务
```

## 开发计划

详见 [开发计划.md](./开发计划.md)

- ✅ **阶段一：基础框架**（已完成）
  - 统一 HTTP 客户端
  - Token 自动管理
  - 智能重试机制
  - 错误处理

- ✅ **阶段二：核心业务模块**（已完成）
  - ✅ 通讯录管理 (Contact)
  - ✅ 身份验证 (OAuth)
  - ✅ 企业二维码 (QRCode)
  - ✅ IP 管理 (IP)
  - ✅ 上下游服务 (UpDown)
  - ✅ 企业互联 (CorpGroup)
  - ✅ 安全管理 (Security)
  - ✅ 消息管理 (Message)
  - ✅ 应用管理 (Agent)
    - ✅ 应用信息管理（获取应用详情、获取应用列表、设置应用）
    - ✅ 菜单管理（创建菜单、获取菜单、删除菜单）
    - ✅ 工作台自定义展示（设置/获取模板、设置/批量设置/获取用户数据）
  - ✅ 外部联系人 (ExternalContact)
    - ✅ 客户管理（获取客户列表、获取客户详情、修改备注、批量获取）
    - ✅ 客户标签管理（企业标签、规则组标签、客户打标）
    - ✅ 客户联系规则组管理（规则组CRUD、管理范围）
    - ✅ 获客助手（获客链接、额度管理、使用统计）
    - ✅ 客户群管理（获取群列表、获取群详情、opengid转换）
    - ✅ 客户朋友圈（发表任务、获取列表、互动数据、规则组管理）
    - ✅ 联系我与客户入群方式（「联系我」配置、客户群进群方式管理）
    - ✅ 企业服务人员管理（获取配置了客户联系功能的成员列表）
    - ✅ 统计管理（群聊数据统计、联系客户统计）
    - ✅ 消息推送（创建企业群发、获取群发记录、发送新客户欢迎语、入群欢迎语素材管理）
    - ✅ 在职继承（分配在职成员的客户、分配在职成员的客户群、查询客户接替状态）
    - ✅ 商品图册管理（创建、获取、列表、编辑、删除）
    - ✅ 聊天敏感词管理（新建、获取列表、获取详情、修改、删除）
    - ✅ 获取已服务的外部联系人
    - ⏳ 离职继承（部分完成）
    - ⏳ 上传附件资源（需要文件上传功能支持，待实现）

- ⏳ **阶段三：更多业务模块**（规划中）
  - ✅ 素材管理 (Media)
    - ✅ 上传图片（永久有效）
    - ✅ 上传临时素材（图片、语音、视频、文件）
    - ✅ 获取临时素材（支持Range分块下载）
    - ✅ 获取高清语音素材
    - ✅ 异步上传临时素材（支持200M大文件）
    - ✅ 查询异步上传任务结果
  - ✅ 电子发票 (Invoice)
    - ✅ 查询电子发票
    - ✅ 批量查询电子发票
    - ✅ 更新发票状态（锁定、解锁、核销）
    - ✅ 批量更新发票状态
  - ⏳ OA 审批
  - ⏳ 会议管理
  - ⏳ 日程管理
  - 等 20+ 个模块

## 示例

查看 [examples](./examples) 目录获取更多示例：

- [基础示例](./examples/basic/main.go)
- [通讯录示例](./examples/contact/main.go)
- [应用管理示例](./examples/agent/main.go)

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
