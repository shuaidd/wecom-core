package kf

import "github.com/shuaidd/wecom-core/types/common"

// GetCorpStatisticRequest 获取「客户数据统计」企业汇总数据请求
type GetCorpStatisticRequest struct {
	OpenKfID  string `json:"open_kfid"`  // 客服账号ID
	StartTime uint32 `json:"start_time"` // 起始日期的时间戳，填这一天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
	EndTime   uint32 `json:"end_time"`   // 结束日期的时间戳，填这一天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
}

// CorpStatistic 企业汇总统计数据
type CorpStatistic struct {
	SessionCnt                 uint64  `json:"session_cnt"`                    // 咨询会话数。客户发过消息并分配给接待人员或智能助手的客服会话数，转接不会产生新的会话
	CustomerCnt                uint64  `json:"customer_cnt"`                   // 咨询客户数。在会话中发送过消息的客户数量，若客户多次咨询只计算一个客户
	CustomerMsgCnt             uint64  `json:"customer_msg_cnt"`               // 咨询消息总数。客户在会话中发送的消息的数量
	UpgradeServiceCustomerCnt  uint64  `json:"upgrade_service_customer_cnt"`   // 升级服务客户数。通过「升级服务」功能成功添加专员或加入客户群的客户数
	AISessionReplyCnt          uint64  `json:"ai_session_reply_cnt"`           // 智能回复会话数。客户发过消息并分配给智能助手的咨询会话数
	AITransferRate             float64 `json:"ai_transfer_rate"`               // 转人工率。一个自然日内，客户给智能助手发消息的会话中，转人工的会话的占比
	AIKnowledgeHitRate         float64 `json:"ai_knowledge_hit_rate"`          // 知识命中率。一个自然日内，客户给智能助手发送的消息中，命中知识库的占比
	MsgRejectedCustomerCnt     uint64  `json:"msg_rejected_customer_cnt"`      // 被拒收消息的客户数。被接待人员设置了"不再接收消息"的客户数
}

// CorpStatisticItem 企业统计数据项
type CorpStatisticItem struct {
	StatTime  uint32         `json:"stat_time"`  // 数据统计日期，为当日0点的时间戳
	Statistic *CorpStatistic `json:"statistic"`  // 一天的统计数据。若当天未产生任何下列统计数据或统计数据还未计算完成则不会返回此项
}

// GetCorpStatisticResponse 获取「客户数据统计」企业汇总数据响应
type GetCorpStatisticResponse struct {
	common.Response
	StatisticList []CorpStatisticItem `json:"statistic_list"` // 统计数据列表
}

// GetServicerStatisticRequest 获取「客户数据统计」接待人员明细数据请求
type GetServicerStatisticRequest struct {
	OpenKfID       string `json:"open_kfid"`                  // 客服账号ID
	ServicerUserID string `json:"servicer_userid,omitempty"`  // 接待人员的userid。第三方应用为密文userid，即open_userid
	StartTime      uint32 `json:"start_time"`                 // 起始日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
	EndTime        uint32 `json:"end_time"`                   // 结束日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天
}

// ServicerStatistic 接待人员统计数据
type ServicerStatistic struct {
	SessionCnt                          uint64  `json:"session_cnt"`                             // 接入人工会话数。客户发过消息并分配给接待人员的咨询会话数
	CustomerCnt                         uint64  `json:"customer_cnt"`                            // 咨询客户数。在会话中发送过消息且接入了人工会话的客户数量，若客户多次咨询只计算一个客户
	CustomerMsgCnt                      uint64  `json:"customer_msg_cnt"`                        // 咨询消息总数。客户在会话中发送的消息的数量
	ReplyRate                           float64 `json:"reply_rate"`                              // 人工回复率。一个自然日内，客户给接待人员发消息的会话中，接待人员回复了的会话的占比
	FirstReplyAverageSec                float64 `json:"first_reply_average_sec"`                 // 平均首次响应时长，单位：秒
	SatisfactionInvestigateCnt          uint64  `json:"satisfaction_investgate_cnt"`             // 满意度评价发送数。当api托管了会话分配，满意度原生功能失效，满意度评价发送数为0
	SatisfactionParticipationRate       float64 `json:"satisfaction_participation_rate"`         // 满意度参评率。当api托管了会话分配，满意度原生功能失效
	SatisfiedRate                       float64 `json:"satisfied_rate"`                          // "满意"评价占比。在客户参评的满意度评价中，评价是"满意"的占比
	MiddlingRate                        float64 `json:"middling_rate"`                           // "一般"评价占比。在客户参评的满意度评价中，评价是"一般"的占比
	DissatisfiedRate                    float64 `json:"dissatisfied_rate"`                       // "不满意"评价占比。在客户参评的满意度评价中，评价是"不满意"的占比
	UpgradeServiceCustomerCnt           uint64  `json:"upgrade_service_customer_cnt"`            // 升级服务客户数。通过「升级服务」功能成功添加专员或加入客户群的客户数
	UpgradeServiceMemberInviteCnt       uint64  `json:"upgrade_service_member_invite_cnt"`       // 专员服务邀请数。接待人员通过「升级服务-专员服务」向客户发送服务专员名片的次数
	UpgradeServiceMemberCustomerCnt     uint64  `json:"upgrade_service_member_customer_cnt"`     // 添加专员的客户数。客户成功添加专员为好友的数量，若同一个客户添加多个专员，则计算多个客户数
	UpgradeServiceGroupchatInviteCnt    uint64  `json:"upgrade_service_groupchat_invite_cnt"`    // 客户群服务邀请数。接待人员通过「升级服务-客户群服务」向客户发送客户群二维码的次数
	UpgradeServiceGroupchatCustomerCnt  uint64  `json:"upgrade_service_groupchat_customer_cnt"`  // 加入客户群的客户数。客户成功加入客户群的数量，若同一个客户加多个客户群，则计算多个客户数
	MsgRejectedCustomerCnt              uint64  `json:"msg_rejected_customer_cnt"`               // 被拒收消息的客户数。被接待人员设置了"不再接收消息"的客户数
}

// ServicerStatisticItem 接待人员统计数据项
type ServicerStatisticItem struct {
	StatTime  uint32             `json:"stat_time"`  // 数据统计日期，为当日0点的时间戳
	Statistic *ServicerStatistic `json:"statistic"`  // 一天的统计数据。若当天未产生任何下列统计数据或统计数据还未计算完成则不会返回此项
}

// GetServicerStatisticResponse 获取「客户数据统计」接待人员明细数据响应
type GetServicerStatisticResponse struct {
	common.Response
	StatisticList []ServicerStatisticItem `json:"statistic_list"` // 统计数据列表
}
