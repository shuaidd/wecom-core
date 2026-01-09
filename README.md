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

### 微信客服

企业微信客服服务，支持客服账号管理、接待人员管理、会话分配与消息收发，帮助企业快速搭建微信客服系统。

#### 客服账号管理

```go
// 添加客服账号
addResp, err := client.KF.AddAccount(ctx, &kf.AddAccountRequest{
    Name:    "客服小王",
    MediaID: "MEDIA_ID",  // 客服头像临时素材ID
})
if err != nil {
    log.Fatalf("添加客服账号失败: %v", err)
}
fmt.Printf("客服账号ID: %s\n", addResp.OpenKfID)

// 获取客服账号列表
listResp, err := client.KF.ListAccount(ctx, &kf.ListAccountRequest{
    Offset: 0,
    Limit:  100,
})
if err != nil {
    log.Fatalf("获取客服账号列表失败: %v", err)
}
for _, account := range listResp.AccountList {
    fmt.Printf("客服ID: %s, 名称: %s, 头像: %s\n",
        account.OpenKfID, account.Name, account.Avatar)
}

// 修改客服账号
err = client.KF.UpdateAccount(ctx, &kf.UpdateAccountRequest{
    OpenKfID: "OPEN_KFID",
    Name:     "资深客服小王",
    MediaID:  "NEW_MEDIA_ID",
})
if err != nil {
    log.Fatalf("修改客服账号失败: %v", err)
}

// 删除客服账号
err = client.KF.DeleteAccount(ctx, &kf.DeleteAccountRequest{
    OpenKfID: "OPEN_KFID",
})
if err != nil {
    log.Fatalf("删除客服账号失败: %v", err)
}
```

#### 获取客服链接

```go
// 获取客服账号链接（用于嵌入网页或生成二维码）
linkResp, err := client.KF.AddContactWay(ctx, &kf.AddContactWayRequest{
    OpenKfID: "OPEN_KFID",
    Scene:    "product_page",  // 可选：场景值，用于区分不同来源
})
if err != nil {
    log.Fatalf("获取客服链接失败: %v", err)
}
fmt.Printf("客服链接: %s\n", linkResp.URL)

// 可以在链接后拼接 scene_param 参数
// 例如: https://work.weixin.qq.com/kf/xxx?enc_scene=xxx&scene_param=a%3D1%26b%3D2
```

#### 接待人员管理

```go
// 添加接待人员
addServicerResp, err := client.KF.AddServicer(ctx, &kf.AddServicerRequest{
    OpenKfID:         "OPEN_KFID",
    UserIDList:       []string{"zhangsan", "lisi"},
    DepartmentIDList: []uint64{2, 4},
})
if err != nil {
    log.Fatalf("添加接待人员失败: %v", err)
}
// 查看操作结果
for _, result := range addServicerResp.ResultList {
    if result.UserID != "" {
        fmt.Printf("用户 %s: %s\n", result.UserID, result.ErrMsg)
    } else {
        fmt.Printf("部门 %d: %s\n", result.DepartmentID, result.ErrMsg)
    }
}

// 获取接待人员列表
servicerList, err := client.KF.ListServicer(ctx, &kf.ListServicerRequest{
    OpenKfID: "OPEN_KFID",
})
if err != nil {
    log.Fatalf("获取接待人员列表失败: %v", err)
}
for _, servicer := range servicerList.ServicerList {
    if servicer.UserID != "" {
        fmt.Printf("接待人员: %s, 状态: %d\n", servicer.UserID, servicer.Status)
    } else {
        fmt.Printf("接待部门: %d\n", servicer.DepartmentID)
    }
}

// 删除接待人员
deleteResp, err := client.KF.DeleteServicer(ctx, &kf.DeleteServicerRequest{
    OpenKfID:         "OPEN_KFID",
    UserIDList:       []string{"zhangsan"},
    DepartmentIDList: []uint64{2},
})
if err != nil {
    log.Fatalf("删除接待人员失败: %v", err)
}
```

#### 会话分配与消息收发

```go
// 获取会话状态
stateResp, err := client.KF.GetServiceState(ctx, &kf.GetServiceStateRequest{
    OpenKfID:       "OPEN_KFID",
    ExternalUserID: "EXTERNAL_USERID",
})
if err != nil {
    log.Fatalf("获取会话状态失败: %v", err)
}
fmt.Printf("会话状态: %d\n", stateResp.ServiceState)
// 状态说明：0-未处理, 1-由智能助手接待, 2-待接入池排队中, 3-由人工接待, 4-已结束/未开始

// 变更会话状态（例如：分配给人工接待）
transResp, err := client.KF.TransServiceState(ctx, &kf.TransServiceStateRequest{
    OpenKfID:       "OPEN_KFID",
    ExternalUserID: "EXTERNAL_USERID",
    ServiceState:   3,  // 变更为人工接待
    ServicerUserID: "zhangsan",  // 指定接待人员
})
if err != nil {
    log.Fatalf("变更会话状态失败: %v", err)
}
if transResp.MsgCode != "" {
    fmt.Printf("消息code: %s\n", transResp.MsgCode)
}

// 发送消息（文本消息示例）
sendResp, err := client.KF.SendMsg(ctx, &kf.SendMsgRequest{
    ToUser:   "EXTERNAL_USERID",
    OpenKfID: "OPEN_KFID",
    MsgType:  "text",
    Text: &kf.TextContent{
        Content: "您好，有什么可以帮您的吗？",
    },
})
if err != nil {
    log.Fatalf("发送消息失败: %v", err)
}
fmt.Printf("消息ID: %s\n", sendResp.MsgID)

// 发送图片消息
_, err = client.KF.SendMsg(ctx, &kf.SendMsgRequest{
    ToUser:   "EXTERNAL_USERID",
    OpenKfID: "OPEN_KFID",
    MsgType:  "image",
    Image: &kf.MediaContent{
        MediaID: "MEDIA_ID",
    },
})

// 发送菜单消息
_, err = client.KF.SendMsg(ctx, &kf.SendMsgRequest{
    ToUser:   "EXTERNAL_USERID",
    OpenKfID: "OPEN_KFID",
    MsgType:  "msgmenu",
    MsgMenu: &kf.MsgMenuContent{
        HeadContent: "您对本次服务是否满意？",
        List: []kf.MsgMenuItem{
            {
                Type: "click",
                Click: &kf.MsgMenuClickItem{
                    ID:      "101",
                    Content: "满意",
                },
            },
            {
                Type: "click",
                Click: &kf.MsgMenuClickItem{
                    ID:      "102",
                    Content: "不满意",
                },
            },
        },
        TailContent: "感谢您的反馈",
    },
})

// 发送欢迎语等事件响应消息
eventResp, err := client.KF.SendMsgOnEvent(ctx, &kf.SendMsgOnEventRequest{
    Code:    "WELCOME_CODE",  // 来自事件回调
    MsgType: "text",
    Text: &kf.TextContent{
        Content: "欢迎咨询，我们将竭诚为您服务！",
    },
})
if err != nil {
    log.Fatalf("发送欢迎语失败: %v", err)
}
fmt.Printf("欢迎语消息ID: %s\n", eventResp.MsgID)
```

#### 客户基础信息管理

```go
// 批量获取客户基础信息
customerResp, err := client.KF.BatchGetCustomer(ctx, &kf.BatchGetCustomerRequest{
    ExternalUserIDList:      []string{"wmxxxxxxxxxxxxxxxxxxxxxx", "wmyyyyyyyyyyyyyyyyyyyyyy"},
    NeedEnterSessionContext: 1,  // 返回客户48小时内最后一次进入会话的上下文信息
})
if err != nil {
    log.Fatalf("获取客户信息失败: %v", err)
}

// 遍历客户信息
for _, customer := range customerResp.CustomerList {
    fmt.Printf("客户: %s, 昵称: %s, 性别: %d\n",
        customer.ExternalUserID, customer.Nickname, customer.Gender)

    // 查看进入会话上下文
    if customer.EnterSessionContext != nil {
        fmt.Printf("  场景值: %s, 场景参数: %s\n",
            customer.EnterSessionContext.Scene,
            customer.EnterSessionContext.SceneParam)

        // 视频号信息
        if customer.EnterSessionContext.WechatChannels != nil {
            fmt.Printf("  视频号: %s, 场景: %d\n",
                customer.EnterSessionContext.WechatChannels.Nickname,
                customer.EnterSessionContext.WechatChannels.Scene)
        }
    }
}

// 查看无效的external_userid
for _, invalidID := range customerResp.InvalidExternalUserID {
    fmt.Printf("无效的客户ID: %s\n", invalidID)
}
```

#### 升级服务配置

```go
// 获取配置的专员与客户群
upgradeConfig, err := client.KF.GetUpgradeServiceConfig(ctx)
if err != nil {
    log.Fatalf("获取升级服务配置失败: %v", err)
}

// 查看专员服务配置
if upgradeConfig.MemberRange != nil {
    fmt.Printf("专员列表: %v\n", upgradeConfig.MemberRange.UserIDList)
    fmt.Printf("专员部门: %v\n", upgradeConfig.MemberRange.DepartmentIDList)
}

// 查看客户群配置
if upgradeConfig.GroupchatRange != nil {
    fmt.Printf("客户群列表: %v\n", upgradeConfig.GroupchatRange.ChatIDList)
}

// 为客户升级为专员服务
err = client.KF.UpgradeService(ctx, &kf.UpgradeServiceRequest{
    OpenKfID:       "kfxxxxxxxxxxxxxx",
    ExternalUserID: "wmxxxxxxxxxxxxxxxxxx",
    Type:           1,  // 1:专员服务, 2:客户群服务
    Member: &kf.UpgradeMember{
        UserID:  "zhangsan",
        Wording: "你好，我是你的专属服务专员张三",
    },
})
if err != nil {
    log.Fatalf("升级服务失败: %v", err)
}
fmt.Println("已为客户推荐专员服务")

// 为客户升级为客户群服务
err = client.KF.UpgradeService(ctx, &kf.UpgradeServiceRequest{
    OpenKfID:       "kfxxxxxxxxxxxxxx",
    ExternalUserID: "wmxxxxxxxxxxxxxxxxxx",
    Type:           2,  // 2:客户群服务
    Groupchat: &kf.UpgradeGroupchat{
        ChatID:  "wraaaaaaaaaaaaaaaa",
        Wording: "欢迎加入你的专属服务群",
    },
})
if err != nil {
    log.Fatalf("升级服务失败: %v", err)
}
fmt.Println("已为客户推荐客户群服务")

// 为客户取消推荐
err = client.KF.CancelUpgradeService(ctx, &kf.CancelUpgradeServiceRequest{
    OpenKfID:       "kfxxxxxxxxxxxxxx",
    ExternalUserID: "wmxxxxxxxxxxxxxxxxxx",
})
if err != nil {
    log.Fatalf("取消推荐失败: %v", err)
}
fmt.Println("已取消客户推荐")
```

支持的消息类型：
- **文本消息** (`text`)：纯文本内容
- **图片消息** (`image`)：通过media_id发送图片
- **语音消息** (`voice`)：通过media_id发送语音
- **视频消息** (`video`)：通过media_id发送视频
- **文件消息** (`file`)：通过media_id发送文件
- **图文链接消息** (`link`)：包含标题、描述、链接和缩略图
- **小程序消息** (`miniprogram`)：小程序卡片
- **菜单消息** (`msgmenu`)：交互式菜单
- **地理位置消息** (`location`)：位置信息
- **获客链接消息** (`ca_link`)：获客助手链接

功能说明：
- **客服账号管理**：创建、修改、删除和查询客服账号
- **客服链接获取**：获取带场景值的客服链接，可嵌入H5页面或生成二维码
- **接待人员管理**：添加、删除和查询接待人员（支持按用户和部门管理）
- **会话状态管理**：获取和变更会话状态，实现智能分配
- **消息收发**：支持10种消息类型的发送
- **事件响应消息**：发送欢迎语、提示语等场景化消息
- **场景追踪**：通过场景值(scene)和场景参数(scene_param)追踪用户咨询来源
- **企业限额**：一家企业最多可添加 **5000个** 客服账号，每个账号最多 **2000个** 接待人员

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

### 邮件服务

企业微信邮件服务，支持发送普通邮件、日程邮件、会议邮件，以及管理公共邮箱。

#### 发送普通邮件

```go
// 发送普通邮件
resp, err := client.Email.SendNormalEmail(ctx,
    &email.EmailRecipient{
        Emails:  []string{"user@example.com"},
        UserIDs: []string{"zhangsan"},
    },
    "邮件标题",
    "邮件正文内容",
    email.WithCC(&email.EmailRecipient{
        UserIDs: []string{"lisi"},
    }),
    email.WithAttachments([]*email.Attachment{
        {
            FileName: "document.pdf",
            Content:  "BASE64_ENCODED_CONTENT",
        },
    }),
)
if err != nil {
    log.Fatalf("发送邮件失败: %v", err)
}
fmt.Println("邮件发送成功")
```

#### 发送日程邮件

```go
// 发送日程邮件
resp, err := client.Email.SendScheduleEmail(ctx,
    &email.EmailRecipient{
        UserIDs: []string{"zhangsan", "lisi"},
    },
    "项目评审会议",
    "讨论Q1项目进展",
    &email.Schedule{
        Location:  "3楼会议室",
        StartTime: time.Now().Add(24 * time.Hour).Unix(),
        EndTime:   time.Now().Add(25 * time.Hour).Unix(),
        Reminders: &email.ScheduleReminders{
            IsRemind:              1,
            RemindBeforeEventMins: 15,
            Timezone:              8,
        },
    },
)
if err != nil {
    log.Fatalf("发送日程邮件失败: %v", err)
}
```

#### 发送会议邮件

```go
// 发送会议邮件
resp, err := client.Email.SendMeetingEmail(ctx,
    &email.EmailRecipient{
        UserIDs: []string{"zhangsan", "lisi", "wangwu"},
    },
    "季度总结会议",
    "回顾Q4工作成果",
    &email.Schedule{
        Location:  "线上会议",
        StartTime: time.Now().Add(48 * time.Hour).Unix(),
        EndTime:   time.Now().Add(50 * time.Hour).Unix(),
        Reminders: &email.ScheduleReminders{
            IsRemind:              1,
            RemindBeforeEventMins: 30,
            Timezone:              8,
        },
    },
    &email.Meeting{
        Hosts: &email.EmailRecipient{
            UserIDs: []string{"zhangsan"},
        },
        MeetingAdmins: &email.EmailRecipient{
            UserIDs: []string{"zhangsan"},
        },
        Option: &email.MeetingOption{
            Password:              "123456",
            AutoRecord:            2,  // 云录制
            EnableWaitingRoom:     true,
            AllowEnterBeforeHost:  true,
            EnableScreenWatermark: true,
        },
    },
)
if err != nil {
    log.Fatalf("发送会议邮件失败: %v", err)
}
```

#### 公共邮箱管理

```go
// 创建公共邮箱
createResp, err := client.Email.CreatePublicMail(ctx, &email.CreatePublicMailRequest{
    Email: "support@company.com",
    Name:  "客户支持",
    UserIDList: &email.StringList{
        List: []string{"zhangsan", "lisi"},
    },
    DepartmentList: &email.IDList{
        List: []uint32{1, 2},
    },
    CreateAuthCode: 1,  // 创建客户端专用密码
    AuthCodeInfo: &email.AuthCodeInfo{
        Remark: "办公电脑",
    },
})
if err != nil {
    log.Fatalf("创建公共邮箱失败: %v", err)
}
fmt.Printf("公共邮箱ID: %d\n", createResp.ID)
if createResp.AuthCode != "" {
    fmt.Printf("客户端专用密码: %s\n", createResp.AuthCode)
}

// 更新公共邮箱
updateResp, err := client.Email.UpdatePublicMail(ctx, &email.UpdatePublicMailRequest{
    ID:   createResp.ID,
    Name: "客户支持中心",
    AliasList: &email.StringList{
        List: []string{"service@company.com"},
    },
})
if err != nil {
    log.Fatalf("更新公共邮箱失败: %v", err)
}

// 获取公共邮箱详情
getResp, err := client.Email.GetPublicMail(ctx, []uint32{createResp.ID})
if err != nil {
    log.Fatalf("获取公共邮箱失败: %v", err)
}
for _, mailbox := range getResp.List {
    fmt.Printf("邮箱: %s, 名称: %s\n", mailbox.Email, mailbox.Name)
}

// 搜索公共邮箱
searchResp, err := client.Email.SearchPublicMail(ctx, 1, "support")
if err != nil {
    log.Fatalf("搜索公共邮箱失败: %v", err)
}
for _, mailbox := range searchResp.List {
    fmt.Printf("找到邮箱: %s\n", mailbox.Email)
}

// 获取客户端专用密码列表
authCodeList, err := client.Email.GetAuthCodeList(ctx, createResp.ID)
if err != nil {
    log.Fatalf("获取密码列表失败: %v", err)
}
for _, authCode := range authCodeList.AuthCodeList {
    fmt.Printf("密码ID: %d, 备注: %s\n", authCode.AuthCodeID, authCode.Remark)
}

// 删除客户端专用密码
err = client.Email.DeleteAuthCode(ctx, createResp.ID, authCodeID)
if err != nil {
    log.Fatalf("删除密码失败: %v", err)
}

// 删除公共邮箱
err = client.Email.DeletePublicMail(ctx, createResp.ID)
if err != nil {
    log.Fatalf("删除公共邮箱失败: %v", err)
}
```

#### 应用邮箱账号管理

```go
// 查询应用邮箱账号
aliasResp, err := client.Email.GetAppEmailAlias(ctx)
if err != nil {
    log.Fatalf("查询应用邮箱账号失败: %v", err)
}
fmt.Printf("主邮箱: %s\n", aliasResp.Email)
fmt.Printf("别名邮箱: %v\n", aliasResp.AliasList)

// 更新应用邮箱账号
// 原有的应用邮箱账号将会作为别名邮箱，具有收信能力
err = client.Email.UpdateAppEmailAlias(ctx, &email.UpdateAppEmailAliasRequest{
    NewEmail: "newemail@company.com",
})
if err != nil {
    log.Fatalf("更新应用邮箱账号失败: %v", err)
}
fmt.Println("应用邮箱账号更新成功")
```

#### 邮件群组管理

```go
// 创建邮件群组
err = client.Email.CreateGroup(ctx, &email.CreateGroupRequest{
    GroupID:   "sales@company.com",
    GroupName: "销售团队",
    EmailList: &email.StringList{
        List: []string{"zhangsan@company.com", "lisi@company.com"},
    },
    DepartmentList: &email.IDList{
        List: []uint32{1, 2},
    },
    AllowType: 3,  // 自定义成员
    AllowEmailList: &email.StringList{
        List: []string{"manager@company.com"},
    },
})
if err != nil {
    log.Fatalf("创建邮件群组失败: %v", err)
}
fmt.Println("邮件群组创建成功")

// 获取邮件群组详情
groupDetail, err := client.Email.GetGroup(ctx, "sales@company.com")
if err != nil {
    log.Fatalf("获取邮件群组失败: %v", err)
}
fmt.Printf("群组名称: %s\n", groupDetail.GroupName)
fmt.Printf("群组成员: %v\n", groupDetail.EmailList.List)
fmt.Printf("使用权限: %d\n", groupDetail.AllowType)

// 更新邮件群组
// 注意：Json数组类型传空值将会清空其内容，不传则保持不变
err = client.Email.UpdateGroup(ctx, &email.UpdateGroupRequest{
    GroupID:   "sales@company.com",
    GroupName: "销售一组",
    EmailList: &email.StringList{
        List: []string{"zhangsan@company.com", "wangwu@company.com"},
    },
})
if err != nil {
    log.Fatalf("更新邮件群组失败: %v", err)
}
fmt.Println("邮件群组更新成功")

// 模糊搜索邮件群组
searchResp, err := client.Email.SearchGroup(ctx, 1, "sales")
if err != nil {
    log.Fatalf("搜索邮件群组失败: %v", err)
}
fmt.Printf("找到 %d 个群组\n", searchResp.Count)
for _, group := range searchResp.Groups {
    fmt.Printf("群组ID: %s, 名称: %s\n", group.GroupID, group.GroupName)
}

// 获取全部邮件群组
allGroups, err := client.Email.SearchGroup(ctx, 0, "")
if err != nil {
    log.Fatalf("获取全部邮件群组失败: %v", err)
}
fmt.Printf("共有 %d 个邮件群组\n", allGroups.Count)

// 删除邮件群组
err = client.Email.DeleteGroup(ctx, "sales@company.com")
if err != nil {
    log.Fatalf("删除邮件群组失败: %v", err)
}
fmt.Println("邮件群组删除成功")
```

功能说明：
- **发送邮件**：支持发送普通邮件、日程邮件、会议邮件
- **附件支持**：支持添加附件（base64编码），总大小不超过50M，最多200个附件
- **收件人类型**：支持邮箱地址和企业内成员UserID两种方式
- **抄送密送**：支持设置抄送(CC)和密送(BCC)
- **日程管理**：支持创建、修改、取消日程，支持重复日程设置
- **会议功能**：支持设置会议密码、录制、等候室、水印等选项
- **公共邮箱**：支持创建、更新、删除、查询公共邮箱
- **权限管理**：支持按成员、部门、标签设置公共邮箱使用权限
- **客户端密码**：支持为公共邮箱创建和管理客户端专用密码
- **应用邮箱账号**：支持查询和更新应用邮箱账号及别名
- **邮件群组**：支持创建、获取、更新、搜索、删除邮件群组
- **群组权限**：支持设置群组使用权限（企业成员、任何人、组内成员、自定义成员）

### 微文档管理

企业微信微文档服务，支持文档、表格、智能表格的创建和管理，以及收集表的创建、编辑和数据收集。

#### 文档管理

```go
// 新建文档
createDocResp, err := client.Wedoc.CreateDoc(ctx, &wedoc.CreateDocRequest{
    DocType:    3, // 3:文档 4:表格 10:智能表格
    DocName:    "测试文档",
    AdminUsers: []string{"zhangsan"},
})
if err != nil {
    log.Fatalf("创建文档失败: %v", err)
}
fmt.Printf("文档创建成功: DocID=%s, URL=%s\n", createDocResp.DocID, createDocResp.URL)

// 获取文档基础信息
docInfo, err := client.Wedoc.GetDocBaseInfo(ctx, &wedoc.GetDocBaseInfoRequest{
    DocID: docID,
})
if err != nil {
    log.Fatalf("获取文档信息失败: %v", err)
}
fmt.Printf("文档信息: Name=%s, Type=%d\n", docInfo.DocBaseInfo.DocName, docInfo.DocBaseInfo.DocType)

// 重命名文档
err = client.Wedoc.RenameDoc(ctx, &wedoc.RenameDocRequest{
    DocID:   docID,
    NewName: "新文档名称",
})
if err != nil {
    log.Fatalf("重命名文档失败: %v", err)
}

// 分享文档
shareResp, err := client.Wedoc.ShareDoc(ctx, &wedoc.ShareDocRequest{
    DocID: docID,
})
if err != nil {
    log.Fatalf("分享文档失败: %v", err)
}
fmt.Printf("文档分享链接: %s\n", shareResp.ShareURL)

// 删除文档
err = client.Wedoc.DeleteDoc(ctx, &wedoc.DeleteDocRequest{
    DocID: docID,
})
if err != nil {
    log.Fatalf("删除文档失败: %v", err)
}
```

#### 文档内容管理

```go
// 获取文档数据
documentData, err := client.Wedoc.GetDocument(ctx, &wedoc.GetDocumentRequest{
    DocID: "DOCID",
})
if err != nil {
    log.Fatalf("获取文档数据失败: %v", err)
}
fmt.Printf("文档版本: %d\n", documentData.Version)

// 批量编辑文档内容
err = client.Wedoc.BatchUpdateDocument(ctx, &wedoc.BatchUpdateDocumentRequest{
    DocID:   "DOCID",
    Version: 10,
    Requests: []wedoc.UpdateRequest{
        {
            InsertText: &wedoc.InsertText{
                Text: "插入的文本内容",
                Location: wedoc.Location{
                    Index: 10,
                },
            },
        },
        {
            InsertTable: &wedoc.InsertTable{
                Rows: 3,
                Cols: 3,
                Location: wedoc.Location{
                    Index: 20,
                },
            },
        },
        {
            InsertImage: &wedoc.InsertImage{
                ImageID: "https://example.com/image.png",
                Location: wedoc.Location{
                    Index: 30,
                },
                Width:  300,
                Height: 200,
            },
        },
    },
})
if err != nil {
    log.Fatalf("批量编辑文档失败: %v", err)
}

// 上传文档图片
imageResp, err := client.Wedoc.ImageUpload(ctx, &wedoc.ImageUploadRequest{
    DocID:         "DOCID",
    Base64Content: "BASE64_ENCODED_IMAGE_CONTENT",
})
if err != nil {
    log.Fatalf("上传图片失败: %v", err)
}
fmt.Printf("图片URL: %s, 宽度: %d, 高度: %d\n", imageResp.URL, imageResp.Width, imageResp.Height)
```

#### 表格内容管理

```go
// 获取表格行列信息
properties, err := client.Wedoc.GetSheetProperties(ctx, &wedoc.GetSheetPropertiesRequest{
    DocID: "SPREADSHEET_DOCID",
})
if err != nil {
    log.Fatalf("获取表格行列信息失败: %v", err)
}
for _, prop := range properties.Properties {
    fmt.Printf("工作表: %s, 行数: %d, 列数: %d\n",
        prop.Title, prop.RowCount, prop.ColumnCount)
}

// 获取表格数据
rangeData, err := client.Wedoc.GetSheetRangeData(ctx, &wedoc.GetSheetRangeDataRequest{
    DocID:   "SPREADSHEET_DOCID",
    SheetID: "SHEET_ID",
    Range:   "A1:C10", // A1表示法
})
if err != nil {
    log.Fatalf("获取表格数据失败: %v", err)
}
for _, row := range rangeData.Data.Result.Rows {
    for _, cell := range row.Values {
        if cell.CellValue != nil {
            fmt.Printf("%s\t", cell.CellValue.Text)
        }
    }
    fmt.Println()
}

// 批量编辑表格内容
updateResp, err := client.Wedoc.BatchUpdateSpreadsheet(ctx, &wedoc.BatchUpdateSpreadsheetRequest{
    DocID: "SPREADSHEET_DOCID",
    Requests: []wedoc.SpreadsheetUpdateRequest{
        {
            // 新增工作表
            AddSheetRequest: &wedoc.SpreadsheetAddSheetRequest{
                Title:       "新工作表",
                RowCount:    100,
                ColumnCount: 26,
            },
        },
        {
            // 更新范围内单元格内容
            UpdateRangeRequest: &wedoc.UpdateRangeRequest{
                SheetID: "SHEET_ID",
                GridData: &wedoc.GridData{
                    StartRow:    0,
                    StartColumn: 0,
                    Rows: []*wedoc.RowData{
                        {
                            Values: []*wedoc.CellData{
                                {
                                    CellValue: &wedoc.CellValue{
                                        Text: "单元格内容",
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
        {
            // 删除连续的行
            DeleteDimensionRequest: &wedoc.DeleteDimensionRequest{
                SheetID:    "SHEET_ID",
                Dimension:  "ROW",
                StartIndex: 5,
                EndIndex:   10,
            },
        },
    },
})
if err != nil {
    log.Fatalf("批量编辑表格失败: %v", err)
}

// 查看操作结果
for _, resp := range updateResp.Data.Responses {
    if resp.AddSheetResponse != nil {
        fmt.Printf("新增工作表: %s\n", resp.AddSheetResponse.Properties.Title)
    }
    if resp.UpdateRangeResponse != nil {
        fmt.Printf("更新了 %d 个单元格\n", resp.UpdateRangeResponse.UpdatedCells)
    }
    if resp.DeleteDimensionResponse != nil {
        fmt.Printf("删除了 %d 行/列\n", resp.DeleteDimensionResponse.Deleted)
    }
}
```

#### 收集表管理

```go
// 创建收集表
createFormResp, err := client.Wedoc.CreateForm(ctx, &wedoc.CreateFormRequest{
    FormInfo: wedoc.FormInfo{
        FormTitle: "员工信息收集表",
        FormDesc:  "请填写您的个人信息",
        FormQuestion: wedoc.FormQuestion{
            Items: []wedoc.QuestionItem{
                {
                    QuestionID: 1,
                    Title:      "您的姓名",
                    Pos:        1,
                    Status:     1,
                    ReplyType:  1, // 1:文本
                    MustReply:  true,
                },
                {
                    QuestionID: 2,
                    Title:      "您的部门",
                    Pos:        2,
                    Status:     1,
                    ReplyType:  2, // 2:单选
                    MustReply:  true,
                    OptionItem: []wedoc.OptionItem{
                        {Key: 1, Value: "技术部", Status: 1},
                        {Key: 2, Value: "产品部", Status: 1},
                        {Key: 3, Value: "运营部", Status: 1},
                    },
                },
            },
        },
        FormSetting: wedoc.FormSetting{
            FillOutAuth:     0,     // 0:所有人可填写
            AllowMultiFill:  false, // 不允许多次填写
            CanNotifySubmit: true,  // 有回复时提醒
        },
    },
})
if err != nil {
    log.Fatalf("创建收集表失败: %v", err)
}
fmt.Printf("收集表创建成功: FormID=%s\n", createFormResp.FormID)

// 获取收集表信息
formInfo, err := client.Wedoc.GetFormInfo(ctx, &wedoc.GetFormInfoRequest{
    FormID: formID,
})
if err != nil {
    log.Fatalf("获取收集表信息失败: %v", err)
}
fmt.Printf("收集表: %s, 问题数=%d\n", formInfo.FormInfo.FormTitle, len(formInfo.FormInfo.FormQuestion.Items))

// 编辑收集表（修改设置）
err = client.Wedoc.ModifyForm(ctx, &wedoc.ModifyFormRequest{
    Oper:   2, // 2:全量修改设置
    FormID: formID,
    FormInfo: wedoc.FormInfo{
        FormSetting: wedoc.FormSetting{
            AllowMultiFill: true, // 修改为允许多次填写
        },
    },
})

// 获取收集表统计信息
statistic, err := client.Wedoc.GetFormStatistic(ctx, &wedoc.GetFormStatisticRequest{
    RepeatedID: repeatedID, // 从GetFormInfo获取
    ReqType:    1,          // 1:只获取统计结果
})
if err != nil {
    log.Fatalf("获取统计信息失败: %v", err)
}
fmt.Printf("已填写: %d次, 已填写人数: %d, 未填写人数: %d\n",
    statistic.FillCnt, statistic.FillUserCnt, statistic.UnfillUserCnt)

// 获取已提交列表
submitList, err := client.Wedoc.GetFormStatistic(ctx, &wedoc.GetFormStatisticRequest{
    RepeatedID: repeatedID,
    ReqType:    2,              // 2:获取已提交列表
    StartTime:  1700000000,     // 筛选开始时间
    EndTime:    1800000000,     // 筛选结束时间
    Limit:      20,
})

// 读取收集表答案
answers, err := client.Wedoc.GetFormAnswer(ctx, &wedoc.GetFormAnswerRequest{
    RepeatedID: repeatedID,
    AnswerIDs:  []uint64{1, 2, 3}, // 答案ID列表
})
if err != nil {
    log.Fatalf("读取答案失败: %v", err)
}
for _, answer := range answers.Answer.AnswerList {
    fmt.Printf("用户=%s, 回答数=%d\n", answer.UserName, len(answer.Reply.Items))
}
```

#### 智能表格内容管理

智能表格（Smartsheet）提供类似数据库的表格能力，支持对记录、字段、视图、子表和编组进行完整的CRUD操作。

##### 记录管理

```go
// 添加记录
addResp, err := client.Wedoc.AddRecords(ctx, &wedoc.AddRecordsRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
    KeyType: "CELL_VALUE_KEY_TYPE_FIELD_TITLE", // 使用字段标题作为key
    Records: []wedoc.AddRecord{
        {
            Values: map[string]interface{}{
                "姓名": []map[string]interface{}{
                    {"type": "text", "text": "张三"},
                },
                "年龄": 25,
                "是否在职": true,
            },
        },
    },
})

// 查询记录
records, err := client.Wedoc.GetRecords(ctx, &wedoc.GetRecordsRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
    Offset:  0,
    Limit:   100,
})
for _, record := range records.Records {
    fmt.Printf("记录ID=%s, 创建时间=%s\n", record.RecordID, record.CreateTime)
}

// 更新记录
updateResp, err := client.Wedoc.UpdateRecords(ctx, &wedoc.UpdateRecordsRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
    Records: []wedoc.UpdateRecord{
        {
            RecordID: "RECORD_ID",
            Values: map[string]interface{}{
                "年龄": 26,
            },
        },
    },
})

// 删除记录
err = client.Wedoc.DeleteRecords(ctx, &wedoc.DeleteRecordsRequest{
    DocID:     "DOCID",
    SheetID:   "SHEETID",
    RecordIDs: []string{"RECORD_ID_1", "RECORD_ID_2"},
})
```

##### 字段管理

```go
// 添加字段
addFieldsResp, err := client.Wedoc.AddFields(ctx, &wedoc.AddFieldsRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
    Fields: []wedoc.AddField{
        {
            FieldTitle: "电话号码",
            FieldType:  "FIELD_TYPE_PHONE_NUMBER",
        },
        {
            FieldTitle: "进度",
            FieldType:  "FIELD_TYPE_PROGRESS",
            PropertyProgress: &wedoc.ProgressFieldProperty{
                DecimalPlaces: 2, // 保留2位小数
            },
        },
    },
})

// 查询字段
fields, err := client.Wedoc.GetFields(ctx, &wedoc.GetFieldsRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
})
for _, field := range fields.Fields {
    fmt.Printf("字段: ID=%s, 标题=%s, 类型=%s\n",
        field.FieldID, field.FieldTitle, field.FieldType)
}

// 更新字段
updateFieldsResp, err := client.Wedoc.UpdateFields(ctx, &wedoc.UpdateFieldsRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
    Fields: []wedoc.UpdateField{
        {
            FieldID:    "FIELD_ID",
            FieldTitle: "新字段名称",
        },
    },
})

// 删除字段
err = client.Wedoc.DeleteFields(ctx, &wedoc.DeleteFieldsRequest{
    DocID:    "DOCID",
    SheetID:  "SHEETID",
    FieldIDs: []string{"FIELD_ID_1", "FIELD_ID_2"},
})
```

##### 视图管理

```go
// 添加视图
addViewResp, err := client.Wedoc.AddView(ctx, &wedoc.AddViewRequest{
    DocID:     "DOCID",
    SheetID:   "SHEETID",
    ViewTitle: "销售视图",
    ViewType:  "VIEW_TYPE_GRID", // 表格视图
    Property: &wedoc.ViewProperty{
        AutoSort: true,
        SortSpec: &wedoc.SortSpec{
            SortInfos: []wedoc.SortInfo{
                {FieldID: "FIELD_ID", Desc: true}, // 降序排序
            },
        },
        FilterSpec: &wedoc.FilterSpec{
            Conjunction: "CONJUNCTION_AND",
            Conditions: []wedoc.Condition{
                {
                    FieldID:  "STATUS_FIELD_ID",
                    Operator: "OPERATOR_IS",
                    StringValue: &wedoc.FilterStringValue{
                        Value: []string{"进行中"},
                    },
                },
            },
        },
    },
})

// 查询视图
views, err := client.Wedoc.GetViews(ctx, &wedoc.GetViewsRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
})

// 更新视图
updateViewResp, err := client.Wedoc.UpdateView(ctx, &wedoc.UpdateViewRequest{
    DocID:     "DOCID",
    SheetID:   "SHEETID",
    ViewID:    "VIEW_ID",
    ViewTitle: "更新后的视图名称",
})

// 删除视图
err = client.Wedoc.DeleteView(ctx, &wedoc.DeleteViewRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
    ViewID:  "VIEW_ID",
})
```

##### 子表管理

```go
// 添加子表
addSheetResp, err := client.Wedoc.AddSheet(ctx, &wedoc.AddSheetRequest{
    DocID: "DOCID",
    Properties: &wedoc.SheetProperty{
        Title: "新子表",
        Index: 0, // 插入位置
    },
})

// 查询子表
sheets, err := client.Wedoc.GetSheet(ctx, &wedoc.GetSheetRequest{
    DocID:            "DOCID",
    NeedAllTypeSheet: true, // 包含仪表盘和说明页
})
for _, sheet := range sheets.SheetList {
    fmt.Printf("子表: ID=%s, 标题=%s, 类型=%s\n",
        sheet.SheetID, sheet.Title, sheet.Type)
}

// 更新子表
updateSheetResp, err := client.Wedoc.UpdateSheet(ctx, &wedoc.UpdateSheetRequest{
    DocID:   "DOCID",
    SheetID: "SHEET_ID",
    Properties: &wedoc.SheetProperty{
        Title: "更新后的子表名称",
    },
})

// 删除子表
err = client.Wedoc.DeleteSheet(ctx, &wedoc.DeleteSheetRequest{
    DocID:   "DOCID",
    SheetID: "SHEET_ID",
})
```

##### 编组管理

```go
// 添加编组（字段分组）
addGroupResp, err := client.Wedoc.AddFieldGroup(ctx, &wedoc.AddFieldGroupRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
    Name:    "基础信息",
    Children: []wedoc.FieldGroupChildren{
        {FieldID: "FIELD_ID_1"},
        {FieldID: "FIELD_ID_2"},
    },
})

// 获取编组
groups, err := client.Wedoc.GetFieldGroup(ctx, &wedoc.GetFieldGroupRequest{
    DocID:   "DOCID",
    SheetID: "SHEETID",
})

// 更新编组
updateGroupResp, err := client.Wedoc.UpdateFieldGroup(ctx, &wedoc.UpdateFieldGroupRequest{
    DocID:        "DOCID",
    SheetID:      "SHEETID",
    FieldGroupID: "GROUP_ID",
    Name:         "更新后的编组名称",
})

// 删除编组
err = client.Wedoc.DeleteFieldGroup(ctx, &wedoc.DeleteFieldGroupRequest{
    DocID:        "DOCID",
    SheetID:      "SHEETID",
    FieldGroupID: "GROUP_ID",
})
```

#### 文档权限管理

```go
// 获取文档权限信息
authInfo, err := client.Wedoc.GetDocAuth(ctx, &wedoc.GetDocAuthRequest{
    DocID: "DOCID",
})
if err != nil {
    log.Fatalf("获取文档权限失败: %v", err)
}

// 查看文档查看规则
fmt.Printf("允许企业内成员浏览: %t\n", authInfo.AccessRule.EnableCorpInternal)
fmt.Printf("企业内成员获得权限: %d\n", authInfo.AccessRule.CorpInternalAuth)

// 查看文档安全设置
if authInfo.SecureSetting.Watermark.ShowText {
    fmt.Printf("水印文字: %s\n", authInfo.SecureSetting.Watermark.Text)
}

// 修改文档安全设置
enableReadonlyCopy := false
err = client.Wedoc.ModDocSaftySetting(ctx, &wedoc.ModDocSaftySettingRequest{
    DocID:              "DOCID",
    EnableReadonlyCopy: &enableReadonlyCopy,
    Watermark: &wedoc.Watermark{
        MarginType:      1,    // 1:稀疏 2:紧密
        ShowVisitorName: true,
        ShowText:        true,
        Text:            "企业机密",
    },
})
if err != nil {
    log.Fatalf("修改文档安全设置失败: %v", err)
}

// 修改文档查看规则
enableCorpInternal := true
corpInternalAuth := 1
err = client.Wedoc.ModDocJoinRule(ctx, &wedoc.ModDocJoinRuleRequest{
    DocID:              "DOCID",
    EnableCorpInternal: &enableCorpInternal,
    CorpInternalAuth:   &corpInternalAuth,  // 1:只读 2:读写
    CoAuthList: []wedoc.CoAuth{
        {
            Type:         2,  // 2:部门
            DepartmentID: 1,
            Auth:         1,  // 1:只读 2:读写
        },
    },
})
if err != nil {
    log.Fatalf("修改文档查看规则失败: %v", err)
}

// 修改文档通知范围及权限
err = client.Wedoc.ModDocMember(ctx, &wedoc.ModDocMemberRequest{
    DocID: "DOCID",
    UpdateFileMemberList: []wedoc.DocMember{
        {
            Type:   1,          // 1:用户
            UserID: "zhangsan",
            Auth:   7,          // 1:只读 2:读写 7:管理员
        },
    },
    DelFileMemberList: []wedoc.DocMember{
        {
            Type:   1,
            UserID: "lisi",
        },
    },
})
if err != nil {
    log.Fatalf("修改文档通知范围失败: %v", err)
}
```

#### 智能表格内容权限管理

```go
// 查询智能表格子表权限（全员权限）
privResp, err := client.Wedoc.GetSheetPriv(ctx, &wedoc.GetSheetPrivRequest{
    DocID: "DOCID",
    Type:  1,  // 1:全员权限 2:额外权限
})
if err != nil {
    log.Fatalf("查询子表权限失败: %v", err)
}

for _, rule := range privResp.RuleList {
    fmt.Printf("规则ID: %d, 类型: %d, 名称: %s\n", rule.RuleID, rule.Type, rule.Name)
    for _, priv := range rule.PrivList {
        fmt.Printf("  子表: %s, 权限: %d\n", priv.SheetID, priv.Priv)
    }
}

// 更新智能表格子表权限
err = client.Wedoc.UpdateSheetPriv(ctx, &wedoc.UpdateSheetPrivRequest{
    DocID: "DOCID",
    Type:  1,  // 1:全员权限
    PrivList: []wedoc.SheetPriv{
        {
            SheetID:                   "SHEET_ID",
            Priv:                      2,     // 2:可编辑
            CanInsertRecord:           true,
            CanDeleteRecord:           true,
            CanCreateModifyDeleteView: true,
            FieldPriv: &wedoc.FieldPriv{
                FieldRangeType: 1,  // 1:所有字段
            },
            RecordPriv: &wedoc.RecordPriv{
                RecordRangeType: 1,  // 1:全部记录
            },
        },
    },
})
if err != nil {
    log.Fatalf("更新子表权限失败: %v", err)
}

// 新增智能表格指定成员额外权限
createRuleResp, err := client.Wedoc.CreateRule(ctx, &wedoc.CreateRuleRequest{
    DocID: "DOCID",
    Name:  "销售团队权限",
})
if err != nil {
    log.Fatalf("创建权限规则失败: %v", err)
}
fmt.Printf("规则ID: %d\n", createRuleResp.RuleID)

// 更新智能表格指定成员额外权限
err = client.Wedoc.ModRuleMember(ctx, &wedoc.ModRuleMemberRequest{
    DocID:  "DOCID",
    RuleID: createRuleResp.RuleID,
    AddMemberRange: &wedoc.MemberRange{
        UserIDList: []string{"zhangsan", "lisi"},
    },
})
if err != nil {
    log.Fatalf("更新成员权限失败: %v", err)
}

// 为权限规则设置具体权限
err = client.Wedoc.UpdateSheetPriv(ctx, &wedoc.UpdateSheetPrivRequest{
    DocID:  "DOCID",
    Type:   2,  // 2:额外权限
    RuleID: createRuleResp.RuleID,
    Name:   "销售团队权限",
    PrivList: []wedoc.SheetPriv{
        {
            SheetID: "SHEET_ID",
            Priv:    2,  // 2:可编辑
            FieldPriv: &wedoc.FieldPriv{
                FieldRangeType: 2,  // 2:部分字段
                FieldRuleList: []wedoc.FieldRule{
                    {
                        FieldID:   "FIELD_ID_1",
                        FieldType: "FIELD_TYPE_TEXT",
                        CanEdit:   true,
                        CanInsert: true,
                        CanView:   true,
                    },
                },
            },
            RecordPriv: &wedoc.RecordPriv{
                RecordRangeType: 2,  // 2:满足任意条件的记录
                RecordRuleList: []wedoc.RecordRule{
                    {
                        FieldID:   "CREATED_USER",
                        OperType:  1,  // 1:包含自己
                    },
                },
                OtherPriv: 1,  // 1:不可编辑
            },
        },
    },
})

// 删除智能表格指定成员额外权限
err = client.Wedoc.DeleteRule(ctx, &wedoc.DeleteRuleRequest{
    DocID:      "DOCID",
    RuleIDList: []uint32{createRuleResp.RuleID},
})
if err != nil {
    log.Fatalf("删除权限规则失败: %v", err)
}
```

#### 高级功能账号管理

```go
// 分配高级功能账号
addResp, err := client.Wedoc.BatchAddVip(ctx, &wedoc.BatchAddVipRequest{
    UserIDList: []string{"zhangsan", "lisi", "wangwu"},
})
if err != nil {
    log.Fatalf("分配高级功能账号失败: %v", err)
}

fmt.Printf("分配成功: %v\n", addResp.SuccUserIDList)
fmt.Printf("分配失败: %v\n", addResp.FailUserIDList)

// 获取高级功能账号列表
listResp, err := client.Wedoc.ListVip(ctx, &wedoc.ListVipRequest{
    Limit: 100,
})
if err != nil {
    log.Fatalf("获取账号列表失败: %v", err)
}

fmt.Printf("高级功能账号列表: %v\n", listResp.UserIDList)
fmt.Printf("是否还有更多: %t\n", listResp.HasMore)

// 分页获取
if listResp.HasMore {
    nextPageResp, err := client.Wedoc.ListVip(ctx, &wedoc.ListVipRequest{
        Cursor: listResp.NextCursor,
        Limit:  100,
    })
    _ = nextPageResp
    _ = err
}

// 取消高级功能账号
delResp, err := client.Wedoc.BatchDelVip(ctx, &wedoc.BatchDelVipRequest{
    UserIDList: []string{"wangwu"},
})
if err != nil {
    log.Fatalf("取消高级功能账号失败: %v", err)
}

fmt.Printf("取消成功: %v\n", delResp.SuccUserIDList)
fmt.Printf("取消失败: %v\n", delResp.FailUserIDList)
```

支持的文档类型：
- **文档** (`DocType=3`)：在线文档编辑
- **表格** (`DocType=4`)：电子表格
- **智能表格** (`DocType=10`)：数据库式表格

支持的问题类型：
- **文本** (`ReplyType=1`)：文本输入，支持字符数、数字、邮箱、手机号等校验
- **单选** (`ReplyType=2`)：单选题，支持"其他"选项
- **多选** (`ReplyType=3`)：多选题，支持限制选项数量
- **位置** (`ReplyType=5`)：地理位置，支持自动定位和范围限制
- **图片** (`ReplyType=9`)：图片上传，支持数量和大小限制
- **文件** (`ReplyType=10`)：文件上传
- **日期** (`ReplyType=11`)：日期选择
- **时间** (`ReplyType=14`)：时间选择
- **下拉列表** (`ReplyType=15`)：下拉选择
- **体温** (`ReplyType=16`)：体温输入
- **签名** (`ReplyType=17`)：手写签名
- **部门** (`ReplyType=18`)：选择部门
- **成员** (`ReplyType=19`)：选择成员
- **时长** (`ReplyType=22`)：时长计算

功能说明：
- **文档管理**：创建、获取信息、重命名、分享、删除文档/表格/智能表格
- **收集表创建**：支持创建包含多种问题类型的收集表，最多200个问题
- **收集表编辑**：支持全量修改问题和设置
- **答案收集**：支持读取收集表答案，批量最多100个
- **统计查询**：支持查询统计信息、已提交列表、未提交列表
- **权限控制**：支持设置填写权限（所有人/指定人员/部门）
- **定时重复**：支持设置定时重复收集（每天/每周/每月）
- **问题校验**：支持为不同问题类型设置校验规则
- **文档权限管理**：获取文档权限信息、修改文档安全设置、修改文档查看规则、修改文档通知范围及权限
- **智能表格内容权限**：查询和更新子表权限、管理成员额外权限、设置字段和记录级别权限
- **高级功能账号管理**：分配和取消高级功能账号、获取账号列表

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

### 日历管理

企业微信日历服务，支持日历和日程的创建、查询、更新和删除，以及日程参与者管理。

#### 日历管理

```go
// 创建日历
createCalResp, err := client.Calendar.CreateCalendar(ctx, &calendar.CreateCalendarRequest{
    Calendar: calendar.Calendar{
        Summary:     "团队日历",
        Color:       "#FF3030",
        Description: "团队工作日历",
        Admins:      []string{"zhangsan", "lisi"},
        Shares: []calendar.Share{
            {UserID: "wangwu", Permission: 1}, // 1:可查看
        },
    },
})
if err != nil {
    log.Fatalf("创建日历失败: %v", err)
}
fmt.Printf("日历创建成功: CalID=%s\n", createCalResp.CalID)

// 获取日历详情
calResp, err := client.Calendar.GetCalendar(ctx, &calendar.GetCalendarRequest{
    CalIDList: []string{calID},
})
if err != nil {
    log.Fatalf("获取日历失败: %v", err)
}
for _, cal := range calResp.CalendarList {
    fmt.Printf("日历: %s, 颜色: %s\n", cal.Summary, cal.Color)
}

// 更新日历
updateResp, err := client.Calendar.UpdateCalendar(ctx, &calendar.UpdateCalendarRequest{
    Calendar: calendar.Calendar{
        CalID:       calID,
        Summary:     "更新后的日历名称",
        Color:       "#0000FF",
        Description: "更新后的描述",
    },
})
if err != nil {
    log.Fatalf("更新日历失败: %v", err)
}

// 删除日历
err = client.Calendar.DeleteCalendar(ctx, &calendar.DeleteCalendarRequest{
    CalID: calID,
})
if err != nil {
    log.Fatalf("删除日历失败: %v", err)
}
```

#### 日程管理

```go
// 创建日程
createSchResp, err := client.Calendar.CreateSchedule(ctx, &calendar.CreateScheduleRequest{
    Schedule: calendar.Schedule{
        Summary:     "需求评审会议",
        Description: "2.0版本需求初步评审",
        Location:    "10楼1005会议室",
        StartTime:   1571274600,
        EndTime:     1571320210,
        CalID:       calID,
        Attendees: []calendar.Attendee{
            {UserID: "zhangsan"},
            {UserID: "lisi"},
        },
        Reminders: &calendar.Reminders{
            IsRemind:              1,
            RemindBeforeEventSecs: 3600, // 提前1小时提醒
        },
    },
})
if err != nil {
    log.Fatalf("创建日程失败: %v", err)
}
fmt.Printf("日程创建成功: ScheduleID=%s\n", createSchResp.ScheduleID)

// 获取日程详情
schResp, err := client.Calendar.GetSchedule(ctx, &calendar.GetScheduleRequest{
    ScheduleIDList: []string{scheduleID},
})
if err != nil {
    log.Fatalf("获取日程失败: %v", err)
}
for _, sch := range schResp.ScheduleList {
    fmt.Printf("日程: %s, 地点: %s\n", sch.Summary, sch.Location)
}

// 更新日程
updateSchResp, err := client.Calendar.UpdateSchedule(ctx, &calendar.UpdateScheduleRequest{
    Schedule: calendar.Schedule{
        ScheduleID:  scheduleID,
        Summary:     "更新后的会议主题",
        Description: "更新后的会议描述",
        StartTime:   1571274600,
        EndTime:     1571320210,
    },
})
if err != nil {
    log.Fatalf("更新日程失败: %v", err)
}

// 新增日程参与者
err = client.Calendar.AddAttendees(ctx, &calendar.AddAttendeesRequest{
    ScheduleID: scheduleID,
    Attendees: []calendar.Attendee{
        {UserID: "wangwu"},
    },
})
if err != nil {
    log.Fatalf("新增参与者失败: %v", err)
}

// 删除日程参与者
err = client.Calendar.DeleteAttendees(ctx, &calendar.DeleteAttendeesRequest{
    ScheduleID: scheduleID,
    Attendees: []calendar.Attendee{
        {UserID: "wangwu"},
    },
})
if err != nil {
    log.Fatalf("删除参与者失败: %v", err)
}

// 获取日历下的日程列表
scheduleListResp, err := client.Calendar.GetScheduleByCalendar(ctx, &calendar.GetScheduleByCalendarRequest{
    CalID:  calID,
    Offset: 0,
    Limit:  100,
})
if err != nil {
    log.Fatalf("获取日程列表失败: %v", err)
}
for _, sch := range scheduleListResp.ScheduleList {
    fmt.Printf("日程: %s, 时间: %d - %d\n", sch.Summary, sch.StartTime, sch.EndTime)
}

// 取消日程
err = client.Calendar.DeleteSchedule(ctx, &calendar.DeleteScheduleRequest{
    ScheduleID: scheduleID,
})
if err != nil {
    log.Fatalf("取消日程失败: %v", err)
}
```

#### 重复日程管理

```go
// 创建重复日程（每周重复）
createSchResp, err := client.Calendar.CreateSchedule(ctx, &calendar.CreateScheduleRequest{
    Schedule: calendar.Schedule{
        Summary:   "周会",
        StartTime: 1571274600,
        EndTime:   1571320210,
        CalID:     calID,
        Reminders: &calendar.Reminders{
            IsRemind:   1,
            IsRepeat:   1,
            RepeatType: 1, // 1:每周
            Timezone:   8, // 东八区
        },
    },
})

// 更新重复日程 - 仅修改此日程
updateSchResp, err := client.Calendar.UpdateSchedule(ctx, &calendar.UpdateScheduleRequest{
    OpMode:      1,           // 1:仅修改此日程
    OpStartTime: 1663135200,  // 指定要修改的周期开始时间
    Schedule: calendar.Schedule{
        ScheduleID: scheduleID,
        StartTime:  1663142400, // 新的开始时间
        EndTime:    1663146000,
    },
})

// 更新重复日程 - 修改将来的所有日程
updateSchResp, err := client.Calendar.UpdateSchedule(ctx, &calendar.UpdateScheduleRequest{
    OpMode:      2,          // 2:修改将来的所有日程
    OpStartTime: 1663135200, // 从这个周期开始修改
    Schedule: calendar.Schedule{
        ScheduleID: scheduleID,
        Summary:    "更新后的周会",
        StartTime:  1663135200,
        EndTime:    1663138800,
    },
})

// 取消重复日程 - 仅删除此日程
err = client.Calendar.DeleteSchedule(ctx, &calendar.DeleteScheduleRequest{
    ScheduleID:  scheduleID,
    OpMode:      1,          // 1:仅删除此日程
    OpStartTime: 1663135200, // 指定要删除的周期开始时间
})

// 取消重复日程 - 删除本次及后续日程
err = client.Calendar.DeleteSchedule(ctx, &calendar.DeleteScheduleRequest{
    ScheduleID:  scheduleID,
    OpMode:      2,          // 2:删除本次及后续日程
    OpStartTime: 1663135200, // 从这个周期开始删除
})
```

### 会议管理

企业微信会议服务，支持会议的创建、修改、取消、查询详情等功能。

#### 创建预约会议

```go
// 创建预约会议
createResp, err := client.Meeting.Create(ctx, &meeting.CreateMeetingRequest{
    AdminUserID:      "zhangsan",
    Title:            "产品评审会议",
    MeetingStart:     1571274600,
    MeetingDuration:  3600,
    Description:      "2.0版本产品评审",
    Location:         "10楼1005会议室",
    Invitees: &meeting.Invitees{
        UserID: []string{"lisi", "wangwu"},
    },
    Settings: &meeting.Settings{
        RemindScope:           3, // 提醒所有成员
        Password:              "123456",
        EnableWaitingRoom:     true,
        AllowEnterBeforeHost:  true,
        EnableEnterMute:       1, // 入会时静音
        EnableScreenWatermark: false,
        Hosts: &meeting.Hosts{
            UserID: []string{"zhangsan"},
        },
        RingUsers: &meeting.RingUsers{
            UserID: []string{"lisi", "wangwu"},
        },
    },
    Reminders: &meeting.Reminders{
        IsRepeat:     0, // 非周期性会议
        RemindBefore: []int{900, 3600}, // 15分钟前和1小时前提醒
    },
})
if err != nil {
    log.Fatalf("创建会议失败: %v", err)
}
fmt.Printf("会议创建成功: MeetingID=%s\n", createResp.MeetingID)
```

#### 修改预约会议

```go
// 修改预约会议
err = client.Meeting.Update(ctx, &meeting.UpdateMeetingRequest{
    MeetingID:       meetingID,
    Title:           "更新后的会议标题",
    MeetingStart:    1571278800,
    MeetingDuration: 7200,
    Description:     "更新后的会议描述",
    Location:        "11楼会议室",
    Invitees: &meeting.Invitees{
        UserID: []string{"lisi", "wangwu", "zhaoliu"},
    },
})
if err != nil {
    log.Fatalf("修改会议失败: %v", err)
}
```

#### 取消预约会议

```go
// 取消预约会议
err = client.Meeting.Cancel(ctx, meetingID)
if err != nil {
    log.Fatalf("取消会议失败: %v", err)
}
fmt.Println("会议已取消")
```

#### 获取会议详情

```go
// 获取会议详情
info, err := client.Meeting.GetInfo(ctx, meetingID)
if err != nil {
    log.Fatalf("获取会议详情失败: %v", err)
}

fmt.Printf("会议标题: %s\n", info.Title)
fmt.Printf("会议状态: %d (1:待开始 2:会议中 3:已结束 4:已取消 5:已过期)\n", info.Status)
fmt.Printf("会议号: %s\n", info.MeetingCode)
fmt.Printf("入会链接: %s\n", info.MeetingLink)

// 遍历参会成员
for _, member := range info.Attendees.Member {
    fmt.Printf("成员: %s, 状态: %d, 入会次数: %d, 累计时长: %d秒\n",
        member.UserID, member.Status, member.TotalJoinCount, member.CumulativeTime)
}
```

#### 获取成员会议ID列表

```go
// 获取成员会议ID列表
resp, err := client.Meeting.GetUserMeetingIDs(ctx, &meeting.GetUserMeetingIDsRequest{
    UserID:    "zhangsan",
    Cursor:    "0", // 初次调用填"0"
    Limit:     100,
    BeginTime: 1570000000,
    EndTime:   1580000000,
})
if err != nil {
    log.Fatalf("获取成员会议列表失败: %v", err)
}

fmt.Printf("会议ID列表: %v\n", resp.MeetingIDList)

// 分页拉取
for resp.NextCursor != "" {
    resp, err = client.Meeting.GetUserMeetingIDs(ctx, &meeting.GetUserMeetingIDsRequest{
        UserID: "zhangsan",
        Cursor: resp.NextCursor,
        Limit:  100,
    })
    if err != nil {
        break
    }
    fmt.Printf("更多会议ID: %v\n", resp.MeetingIDList)
}
```

#### 创建周期性会议

```go
// 创建周期性会议（每周重复）
createResp, err := client.Meeting.Create(ctx, &meeting.CreateMeetingRequest{
    AdminUserID:      "zhangsan",
    Title:            "周会",
    MeetingStart:     1571274600,
    MeetingDuration:  3600,
    Description:      "每周例行会议",
    Location:         "10楼会议室",
    Invitees: &meeting.Invitees{
        UserID: []string{"lisi", "wangwu"},
    },
    Reminders: &meeting.Reminders{
        IsRepeat:      1,           // 周期性会议
        RepeatType:   1,           // 每周
        RepeatUntil:   1576876813,  // 重复结束时间
        RepeatInterval: 1,         // 重复间隔
        RemindBefore: []int{900},  // 15分钟前提醒
    },
})
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
  - ✅ 微信客服 (KF)
    - ✅ 客服账号管理（添加、删除、修改、获取列表）
    - ✅ 获取客服账号链接
    - ✅ 接待人员管理（添加、删除、获取列表）
    - ✅ 会话分配与消息收发（获取会话状态、变更会话状态）
    - ✅ 消息发送（文本、图片、语音、视频、文件、图文链接、小程序、菜单、地理位置、获客链接）
    - ✅ 事件响应消息（发送欢迎语等场景化消息）
    - ✅ 客户基础信息管理（批量获取客户基础信息）
    - ✅ 升级服务配置（获取配置的专员与客户群、为客户升级服务、取消推荐）
  - ✅ 邮件服务 (Email)
    - ✅ 发送邮件（普通邮件、日程邮件、会议邮件）
    - ✅ 公共邮箱管理（创建、更新、删除、获取、搜索）
    - ✅ 客户端专用密码管理（获取列表、删除）
    - ✅ 应用邮箱账号管理（查询、更新）
    - ✅ 邮件群组管理（创建、获取、更新、搜索、删除）
  - ✅ 微文档 (Wedoc)
    - ✅ 文档管理（新建文档/表格/智能表格、获取基础信息、重命名、分享、删除）
    - ✅ 文档内容管理（获取文档数据、批量编辑文档内容、上传文档图片）
    - ✅ 表格内容管理（获取表格行列信息、获取表格数据、批量编辑表格内容）
    - ✅ 收集表管理（创建收集表、获取信息、编辑收集表）
    - ✅ 收集表数据（读取答案、获取统计信息、已提交/未提交列表）
    - ✅ 智能表格内容管理（记录、字段、视图、子表、编组的完整CRUD操作）
    - ✅ 文档权限管理（获取权限信息、修改安全设置、修改查看规则、修改通知范围及权限）
    - ✅ 智能表格内容权限（查询和更新子表权限、管理成员额外权限、字段和记录级别权限）
    - ✅ 高级功能账号管理（分配、取消、获取高级功能账号列表）
  - ✅ 日历管理 (Calendar)
    - ✅ 日历管理（创建、获取、更新、删除日历）
    - ✅ 日程管理（创建、获取、更新、取消日程）
    - ✅ 日程参与者管理（新增、删除参与者）
    - ✅ 获取日历下的日程列表
    - ✅ 重复日程管理（支持不同操作模式）
  - ✅ 会议管理 (Meeting)
    - ✅ 创建预约会议
    - ✅ 修改预约会议
    - ✅ 取消预约会议
    - ✅ 获取会议详情
    - ✅ 获取成员会议ID列表
  - ⏳ OA 审批
  - 等 20+ 个模块

## 示例

查看 [examples](./examples) 目录获取更多示例：

- [基础示例](./examples/basic/main.go)
- [通讯录示例](./examples/contact/main.go)
- [应用管理示例](./examples/agent/main.go)
- [微文档示例](./examples/wedoc/main.go)

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
