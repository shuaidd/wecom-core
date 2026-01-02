# OAuth 身份验证服务

身份验证服务提供了企业微信OAuth2.0网页授权相关的功能,包括构造授权链接、获取用户身份信息、获取敏感信息以及二次验证等功能。

## 功能列表

### 1. 构造网页授权链接

用于构造OAuth2.0授权链接,引导用户进行身份授权。

```go
// 构造静默授权链接(获取基础信息)
authURL, err := client.OAuth.BuildAuthorizeURL(oauth.BuildAuthorizeURLParams{
    CorpID:      "your_corp_id",
    RedirectURI: "http://example.com/callback",
    Scope:       oauth.ScopeBase,  // 静默授权
    State:       "STATE123",
    AgentID:     "1000002",
})

// 构造手动授权链接(获取敏感信息)
authURL, err := client.OAuth.BuildAuthorizeURL(oauth.BuildAuthorizeURLParams{
    CorpID:      "your_corp_id",
    RedirectURI: "http://example.com/callback",
    Scope:       oauth.ScopePrivateInfo,  // 手动授权,可获取敏感信息
    State:       "STATE456",
    AgentID:     "1000002",  // snsapi_privateinfo时必填
})
```

**授权作用域说明:**
- `oauth.ScopeBase`: 静默授权,可获取成员的基础信息(UserId)
- `oauth.ScopePrivateInfo`: 手动授权,可获取成员的详细信息,包含头像、二维码等敏感信息

**注意事项:**
- `RedirectURI` 需要进行URL编码
- 当 `Scope` 为 `ScopePrivateInfo` 时,必须填写 `AgentID`
- 授权后用户会被重定向到 `RedirectURI`,并携带 `code` 参数

### 2. 获取访问用户身份

根据授权回调获得的code参数,获取用户的身份信息。

```go
// 从回调URL中获取code参数
code := r.URL.Query().Get("code")

// 获取用户身份信息
userInfo, err := client.OAuth.GetUserInfo(ctx, code)
if err != nil {
    log.Printf("获取用户身份失败: %v", err)
    return
}

// 企业成员
if userInfo.UserID != "" {
    fmt.Printf("用户ID: %s\n", userInfo.UserID)

    // scope为snsapi_privateinfo时会返回user_ticket
    if userInfo.UserTicket != "" {
        fmt.Printf("用户票据: %s\n", userInfo.UserTicket)
    }
}

// 非企业成员
if userInfo.OpenID != "" {
    fmt.Printf("OpenID: %s\n", userInfo.OpenID)
}

// 外部联系人
if userInfo.ExternalUserID != "" {
    fmt.Printf("外部联系人ID: %s\n", userInfo.ExternalUserID)
}
```

**返回字段说明:**
- `UserID`: 成员UserID(企业成员时返回)
- `UserTicket`: 成员票据,有效期1800s(scope为snsapi_privateinfo时返回)
- `OpenID`: 非企业成员的标识
- `ExternalUserID`: 外部联系人id

**注意事项:**
- `code` 只能使用一次,5分钟未使用自动过期
- 跳转域名须完全匹配应用的可信域名

### 3. 获取访问用户敏感信息

通过 `user_ticket` 获取成员授权的敏感字段信息。

```go
// 先获取用户身份
userInfo, err := client.OAuth.GetUserInfo(ctx, code)
if err != nil {
    return err
}

// 如果有user_ticket,获取敏感信息
if userInfo.UserTicket != "" {
    userDetail, err := client.OAuth.GetUserDetail(ctx, userInfo.UserTicket)
    if err != nil {
        log.Printf("获取用户敏感信息失败: %v", err)
        return
    }

    fmt.Printf("用户ID: %s\n", userDetail.UserID)
    fmt.Printf("性别: %s\n", userDetail.Gender)
    fmt.Printf("头像: %s\n", userDetail.Avatar)
    fmt.Printf("二维码: %s\n", userDetail.QRCode)
    fmt.Printf("手机: %s\n", userDetail.Mobile)
    fmt.Printf("邮箱: %s\n", userDetail.Email)
    fmt.Printf("企业邮箱: %s\n", userDetail.BizMail)
    fmt.Printf("地址: %s\n", userDetail.Address)
}
```

**返回字段说明:**
- `UserID`: 成员UserID
- `Gender`: 性别(0:未定义, 1:男性, 2:女性)
- `Avatar`: 头像URL
- `QRCode`: 员工个人二维码
- `Mobile`: 手机号(第三方应用不可获取)
- `Email`: 邮箱(第三方应用不可获取)
- `BizMail`: 企业邮箱(第三方应用不可获取)
- `Address`: 地址(第三方应用不可获取)

**注意事项:**
- 仅在用户同意 `snsapi_privateinfo` 授权时返回真实敏感信息
- 敏感字段需要管理员在应用详情里选择
- 成员oauth2授权时需要确认
- 第三方应用无法获取部分敏感字段

### 4. 获取用户二次验证信息

用于获取触发二次验证的成员信息。

```go
// 从二次验证回调中获取code
code := r.URL.Query().Get("code")

// 获取二次验证信息
tfaInfo, err := client.OAuth.GetTFAInfo(ctx, code)
if err != nil {
    log.Printf("获取二次验证信息失败: %v", err)
    return
}

fmt.Printf("用户ID: %s\n", tfaInfo.UserID)
fmt.Printf("二次验证授权码: %s\n", tfaInfo.TFACode)

// 验证用户身份信息无误后,可以调用通过二次验证接口
```

**返回字段说明:**
- `UserID`: 成员UserID
- `TFACode`: 二次验证授权码,有效期5分钟,只能使用一次

**权限说明:**
- 仅「通讯录同步」或者自建应用可调用
- 用户需要在二次验证范围和应用可见范围内
- 验证页面的链接必须填该自建应用的oauth2链接

## 使用场景

### Web应用OAuth授权流程

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/shuaidd/wecom-core"
    "github.com/shuaidd/wecom-core/config"
    "github.com/shuaidd/wecom-core/types/oauth"
)

var client *wecom.Client

func init() {
    var err error
    client, err = wecom.New(
        config.WithCorpID("your_corp_id"),
        config.WithCorpSecret("your_corp_secret"),
    )
    if err != nil {
        log.Fatal(err)
    }
}

// 步骤1: 引导用户授权
func handleLogin(w http.ResponseWriter, r *http.Request) {
    authURL, err := client.OAuth.BuildAuthorizeURL(oauth.BuildAuthorizeURLParams{
        CorpID:      "your_corp_id",
        RedirectURI: "http://example.com/callback",
        Scope:       oauth.ScopeBase,
        State:       "random_state",
        AgentID:     "1000002",
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 重定向到授权页面
    http.Redirect(w, r, authURL, http.StatusFound)
}

// 步骤2: 处理授权回调
func handleCallback(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    // 获取code参数
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "missing code", http.StatusBadRequest)
        return
    }

    // 获取用户身份
    userInfo, err := client.OAuth.GetUserInfo(ctx, code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 处理用户信息
    if userInfo.UserID != "" {
        // 企业成员登录成功
        log.Printf("用户登录: %s", userInfo.UserID)
        // TODO: 创建session, 设置cookie等
    }

    w.Write([]byte("登录成功"))
}

func main() {
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/callback", handleCallback)
    log.Fatal(http.Server{Addr: ":8080"}.ListenAndServe())
}
```

## API参考

### BuildAuthorizeURL

构造网页授权链接。

**参数:**
- `params.CorpID` (string, 必填): 企业CorpID
- `params.RedirectURI` (string, 必填): 授权后重定向的回调链接地址
- `params.Scope` (AuthScope, 必填): 应用授权作用域
- `params.State` (string, 可选): 重定向后会带上的state参数
- `params.AgentID` (string, 条件必填): 应用agentid(snsapi_privateinfo时必填)

**返回:**
- `string`: 授权链接
- `error`: 错误信息

### GetUserInfo

获取访问用户身份。

**参数:**
- `code` (string): 授权code

**返回:**
- `*oauth.GetUserInfoResponse`: 用户身份信息
- `error`: 错误信息

### GetUserDetail

获取访问用户敏感信息。

**参数:**
- `userTicket` (string): 用户票据

**返回:**
- `*oauth.GetUserDetailResponse`: 用户敏感信息
- `error`: 错误信息

### GetTFAInfo

获取用户二次验证信息。

**参数:**
- `code` (string): 二次验证code

**返回:**
- `*oauth.GetTFAInfoResponse`: 二次验证信息
- `error`: 错误信息

## 常见问题

### 1. 授权链接构造失败

**问题:** 调用 `BuildAuthorizeURL` 返回错误

**解决方案:**
- 检查必填参数是否都已提供
- 确认 `Scope` 为 `ScopePrivateInfo` 时是否填写了 `AgentID`
- 检查 `RedirectURI` 是否有效

### 2. 获取用户信息失败(错误码40029)

**问题:** 调用 `GetUserInfo` 返回 "invalid code"

**解决方案:**
- `code` 只能使用一次,不要重复使用
- `code` 有效期5分钟,请及时使用
- 确认 `code` 参数正确传递

### 3. 获取用户信息失败(错误码50001)

**问题:** 跳转域名不匹配

**解决方案:**
- 在企业微信管理后台配置可信域名
- 确保 `RedirectURI` 的域名与配置的可信域名完全匹配

### 4. 无法获取敏感信息

**问题:** `GetUserDetail` 返回的敏感字段为空

**解决方案:**
- 确认授权时使用的是 `ScopePrivateInfo`
- 在应用详情中勾选需要获取的敏感字段
- 确认用户在授权时同意了敏感信息授权
- 确认用户在应用可见范围内

## 相关文档

- [企业微信官方文档 - 网页授权登录](https://developer.work.weixin.qq.com/document/path/91335)
- [企业微信官方文档 - 身份验证](https://developer.work.weixin.qq.com/document/path/91023)
