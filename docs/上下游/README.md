# 上下游接口对接完成

已完成 `docs/上下游` 目录下所有接口的对接，共实现 **18个接口**。

## 接口列表

### 企业管理相关 (7个接口)

1. **获取上下游列表** - `GetChainList`
   - 文档: https://developer.work.weixin.qq.com/document/path/93355
   - 路径: `/cgi-bin/corpgroup/corp/get_chain_list`

2. **获取上下游通讯录分组** - `GetChainGroup`
   - 文档: https://developer.work.weixin.qq.com/document/path/93355
   - 路径: `/cgi-bin/corpgroup/corp/get_chain_group`

3. **获取企业上下游通讯录分组下的企业详情列表** - `GetChainCorpInfoList`
   - 文档: https://developer.work.weixin.qq.com/document/path/93355
   - 路径: `/cgi-bin/corpgroup/corp/get_chain_corpinfo_list`

4. **获取企业上下游通讯录下的企业信息** - `GetChainCorpInfo`
   - 文档: https://developer.work.weixin.qq.com/document/path/93355
   - 路径: `/cgi-bin/corpgroup/corp/get_chain_corpinfo`

5. **移除企业** - `RemoveCorp`
   - 文档: https://developer.work.weixin.qq.com/document/path/95820
   - 路径: `/cgi-bin/corpgroup/corp/remove_corp`

6. **查询成员自定义id** - `GetChainUserCustomID`
   - 文档: https://developer.work.weixin.qq.com/document/path/95815
   - 路径: `/cgi-bin/corpgroup/corp/get_chain_user_custom_id`

7. **获取下级企业加入的上下游** - `GetCorpSharedChainList`
   - 文档: https://developer.work.weixin.qq.com/document/path/95816
   - 路径: `/cgi-bin/corpgroup/get_corp_shared_chain_list`

### 对接规则管理 (5个接口)

8. **新增对接规则** - `AddRule`
   - 文档: https://developer.work.weixin.qq.com/document/path/95792
   - 路径: `/cgi-bin/corpgroup/rule/add_rule`

9. **更新对接规则** - `ModifyRule`
   - 文档: https://developer.work.weixin.qq.com/document/path/95793
   - 路径: `/cgi-bin/corpgroup/rule/modify_rule`

10. **删除对接规则** - `DeleteRule`
    - 文档: https://developer.work.weixin.qq.com/document/path/95794
    - 路径: `/cgi-bin/corpgroup/rule/delete_rule`

11. **获取对接规则id列表** - `ListRuleIDs`
    - 文档: https://developer.work.weixin.qq.com/document/path/95795
    - 路径: `/cgi-bin/corpgroup/rule/list_ids`

12. **获取对接规则详情** - `GetRuleInfo`
    - 文档: https://developer.work.weixin.qq.com/document/path/95796
    - 路径: `/cgi-bin/corpgroup/rule/get_rule_info`

### 应用共享 (1个接口)

13. **获取应用共享信息** - `ListAppShareInfo`
    - 文档: https://developer.work.weixin.qq.com/document/path/93403
    - 路径: `/cgi-bin/corpgroup/corp/list_app_share_info`

### 客户关联 (3个接口)

14. **通过unionid和openid查询external_userid** - `UnionIDToExternalUserID`
    - 文档: https://developer.work.weixin.qq.com/document/path/95818
    - 路径: `/cgi-bin/corpgroup/unionid_to_external_userid`

15. **unionid查询pending_id** - `UnionIDToPendingID`
    - 文档: https://developer.work.weixin.qq.com/document/path/97357
    - 路径: `/cgi-bin/corpgroup/unionid_to_pending_id`

16. **external_userid查询pending_id** - `ExternalUserIDToPendingID`
    - 文档: https://developer.work.weixin.qq.com/document/path/97357
    - 路径: `/cgi-bin/corpgroup/batch/external_userid_to_pending_id`

### 联系人导入 (1个接口)

17. **批量导入上下游联系人** - `ImportChainContact`
    - 文档: https://developer.work.weixin.qq.com/document/path/95813
    - 路径: `/cgi-bin/corpgroup/import_chain_contact`

### 异步任务 (1个接口)

18. **获取异步任务结果** - `GetTaskResult`
    - 文档: https://developer.work.weixin.qq.com/document/path/95814
    - 路径: `/cgi-bin/corpgroup/getresult`

## 使用方法

```go
package main

import (
    "context"
    wecom "github.com/shuaidd/wecom-core"
    "github.com/shuaidd/wecom-core/config"
    "github.com/shuaidd/wecom-core/types/updown"
)

func main() {
    // 创建客户端
    client, _ := wecom.New(
        config.WithCorpID("your_corp_id"),
        config.WithCorpSecret("your_corp_secret"),
    )

    ctx := context.Background()

    // 使用上下游服务
    chains, err := client.UpDown.GetChainList(ctx)
    if err != nil {
        // 处理错误
    }

    // 更多示例请参考 examples/updown_example.go
}
```

## 代码结构

- **类型定义**: `types/updown/`
  - `corp.go` - 企业相关类型
  - `rule.go` - 规则相关类型
  - `app.go` - 应用共享相关类型
  - `customer.go` - 客户相关类型
  - `contact.go` - 联系人导入相关类型
  - `task.go` - 异步任务相关类型

- **服务实现**: `services/updown/`
  - `updown.go` - 服务主文件
  - `corp.go` - 企业管理接口实现
  - `rule.go` - 规则管理接口实现
  - `app.go` - 应用共享接口实现
  - `customer.go` - 客户关联接口实现
  - `contact.go` - 联系人导入接口实现
  - `task.go` - 异步任务接口实现

- **示例代码**: `examples/updown_example.go`

## 验证状态

- ✅ 所有代码已编译通过
- ✅ Go vet 检查通过
- ✅ 已集成到主客户端 `wecom.Client`
- ✅ 提供完整的示例代码
