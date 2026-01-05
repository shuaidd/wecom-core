# 多应用配置示例

本示例演示如何配置和使用多个企业微信应用。

## 背景

在企业微信中，一个企业可能有多个自建应用，每个应用有独立的 AgentID 和 Secret。wecom-core SDK 支持在一个客户端中配置多个应用，并在调用 API 时通过 context 指定使用哪个应用的凭证。

## 配置方式

### 方式1: 使用 `WithAgent` 逐个添加应用

```go
client, err := wecom.New(
    config.WithCorpID("your_corp_id"),
    config.WithAgent("customer", 100001, "agent_secret_1", "客户管理应用"),
    config.WithAgent("study-assistant", 100002, "agent_secret_2", "学习助手"),
)
```

### 方式2: 使用 `WithAgents` 批量添加应用

```go
client, err := wecom.New(
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
)
```

### 方式3: 单应用模式（向后兼容）

如果只有一个应用，可以继续使用原来的配置方式：

```go
client, err := wecom.New(
    config.WithCorpID("your_corp_id"),
    config.WithCorpSecret("your_corp_secret"),
)
```

## 使用方式

### 通过应用名称指定

```go
ctx := context.Background()

// 使用 customer 应用
customerCtx := wecom.WithAgentName(ctx, "customer")
err := client.Message.SendText(customerCtx, req)

// 使用 study-assistant 应用
studyCtx := wecom.WithAgentName(ctx, "study-assistant")
err := client.Message.SendText(studyCtx, req)
```

### 通过应用ID指定

```go
ctx := context.Background()

// 使用应用ID为 100001 的应用
agentCtx := wecom.WithAgentID(ctx, 100001)
err := client.Message.SendText(agentCtx, req)
```

## 工作原理

1. **配置阶段**: 通过 `WithAgent` 或 `WithAgents` 注册多个应用到 TokenManager
2. **调用阶段**: 使用 `WithAgentName` 或 `WithAgentID` 将应用标识添加到 context
3. **Token获取**: SDK 从 context 中提取应用标识，使用对应应用的 secret 获取 access_token
4. **Token缓存**: 每个应用的 token 独立缓存，key 格式为 `wecom:token:{corpid}:{agent_key}`

## 注意事项

1. **应用标识唯一性**: 应用名称和应用ID都可以作为标识，SDK 内部会同时注册两个 key 指向同一个应用配置
2. **默认应用**: 如果调用 API 时没有指定应用，SDK 会尝试使用默认应用：
   - 如果配置了 `CorpSecret`（单应用模式），使用该凭证
   - 如果只配置了一个应用，使用该应用
   - 否则返回错误
3. **Token 隔离**: 不同应用的 token 完全隔离，互不影响
4. **向后兼容**: 原有的单应用配置方式继续有效，无需修改现有代码

## 参考配置文件

参考 `docs/企业微信配置示例.md` 中的 YAML 配置格式：

```yaml
qywx:
  baseUrl: https://qyapi.weixin.qq.com
  corpId: wxssfdsdsww
  agents:
    - agentName: customer
      agentId: 100001
      secret: SEosuierodsksls
      agentDesc: 客户管理应用
    - agentName: study-assitant
      agentId: 100002
      secret: YWISUkisksiusiu
      agentDesc: 学习助手
```

## 运行示例

```bash
# 修改示例代码中的企业ID和应用配置
# 然后运行
go run examples/multi-agent/main.go
```
