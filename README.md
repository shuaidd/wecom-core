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

### 消息管理

```go
// 发送应用消息（敬请期待更多示例）
```

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
  - ✅ 外部联系人 (ExternalContact)
    - ✅ 客户管理（获取客户列表、获取客户详情、修改备注、批量获取）
    - ✅ 客户标签管理（企业标签、规则组标签、客户打标）
    - ✅ 客户联系规则组管理（规则组CRUD、管理范围）
    - ✅ 获客助手（获客链接、额度管理、使用统计）
    - ⏳ 客户群管理（规划中）
    - ⏳ 在职继承/离职继承（规划中）
    - ⏳ 消息推送（规划中）

- ⏳ **阶段三：更多业务模块**（规划中）
  - 应用管理
  - 素材管理
  - OA 审批
  - 会议管理
  - 日程管理
  - 等 20+ 个模块

## 示例

查看 [examples](./examples) 目录获取更多示例：

- [基础示例](./examples/basic/main.go)
- [通讯录示例](./examples/contact/main.go)

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
