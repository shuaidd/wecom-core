package kf

import "github.com/shuaidd/wecom-core/types/common"

// ===================== 知识库分组管理 =====================

// AddKnowledgeGroupRequest 添加知识库分组请求
type AddKnowledgeGroupRequest struct {
	Name string `json:"name"` // 分组名。不超过12个字
}

// AddKnowledgeGroupResponse 添加知识库分组响应
type AddKnowledgeGroupResponse struct {
	common.Response
	GroupID string `json:"group_id"` // 分组ID
}

// DeleteKnowledgeGroupRequest 删除知识库分组请求
type DeleteKnowledgeGroupRequest struct {
	GroupID string `json:"group_id"` // 分组ID
}

// DeleteKnowledgeGroupResponse 删除知识库分组响应
type DeleteKnowledgeGroupResponse struct {
	common.Response
}

// UpdateKnowledgeGroupRequest 修改知识库分组请求
type UpdateKnowledgeGroupRequest struct {
	GroupID string `json:"group_id"` // 分组ID
	Name    string `json:"name"`     // 分组名。不超过12个字
}

// UpdateKnowledgeGroupResponse 修改知识库分组响应
type UpdateKnowledgeGroupResponse struct {
	common.Response
}

// ListKnowledgeGroupRequest 获取知识库分组列表请求
type ListKnowledgeGroupRequest struct {
	Cursor  string `json:"cursor,omitempty"`   // 上一次调用时返回的next_cursor，第一次拉取可以不填
	Limit   uint32 `json:"limit,omitempty"`    // 每次拉取的数据量，默认值500，最大值为1000
	GroupID string `json:"group_id,omitempty"` // 分组ID。可指定拉取特定的分组
}

// KnowledgeGroup 知识库分组
type KnowledgeGroup struct {
	GroupID   string `json:"group_id"`   // 分组ID
	Name      string `json:"name"`       // 分组名
	IsDefault uint32 `json:"is_default"` // 是否为默认分组。0-否 1-是。默认分组为系统自动创建，不可修改/删除
}

// ListKnowledgeGroupResponse 获取知识库分组列表响应
type ListKnowledgeGroupResponse struct {
	common.Response
	NextCursor string           `json:"next_cursor"` // 分页游标，再下次请求时填写以获取之后分页的记录
	HasMore    uint32           `json:"has_more"`    // 是否还有更多数据。0-没有 1-有
	GroupList  []KnowledgeGroup `json:"group_list"`  // 分组列表
}

// ===================== 知识库问答管理 =====================

// QuestionText 问题文本
type QuestionText struct {
	Content string `json:"content"` // 问题文本内容
}

// Question 问题
type Question struct {
	Text QuestionText `json:"text"` // 问题文本
}

// SimilarQuestionItem 相似问题项
type SimilarQuestionItem struct {
	Text QuestionText `json:"text"` // 相似问题文本
}

// SimilarQuestions 相似问题列表
type SimilarQuestions struct {
	Items []SimilarQuestionItem `json:"items,omitempty"` // 相似问题列表。最多支持100个
}

// AnswerText 回答文本
type AnswerText struct {
	Content string `json:"content"` // 回答文本内容
}

// AnswerAttachment 回答附件
type AnswerAttachment struct {
	MsgType     string                       `json:"msgtype"`               // 附件类型: image, video, link, miniprogram
	Image       *AnswerImageAttachment       `json:"image,omitempty"`       // 图片附件
	Video       *AnswerVideoAttachment       `json:"video,omitempty"`       // 视频附件
	Link        *AnswerLinkAttachment        `json:"link,omitempty"`        // 链接附件
	MiniProgram *AnswerMiniProgramAttachment `json:"miniprogram,omitempty"` // 小程序附件
}

// AnswerImageAttachment 图片附件
type AnswerImageAttachment struct {
	MediaID string `json:"media_id,omitempty"` // 图片的media_id（添加/修改时使用）
	Name    string `json:"name,omitempty"`     // 图片的文件名（获取时返回）
}

// AnswerVideoAttachment 视频附件
type AnswerVideoAttachment struct {
	MediaID string `json:"media_id,omitempty"` // 视频的media_id（添加/修改时使用）
	Name    string `json:"name,omitempty"`     // 视频的文件名（获取时返回）
}

// AnswerLinkAttachment 链接附件
type AnswerLinkAttachment struct {
	Title  string `json:"title"`             // 标题
	Desc   string `json:"desc,omitempty"`    // 描述
	URL    string `json:"url"`               // 点击后跳转的链接
	PicURL string `json:"pic_url,omitempty"` // 缩略图链接
}

// AnswerMiniProgramAttachment 小程序附件
type AnswerMiniProgramAttachment struct {
	Title        string `json:"title,omitempty"`          // 小程序消息标题。最多64个字节
	ThumbMediaID string `json:"thumb_media_id,omitempty"` // 小程序消息封面的mediaid（添加/修改时使用）
	AppID        string `json:"appid"`                    // 小程序appid
	PagePath     string `json:"pagepath"`                 // 点击消息卡片后进入的小程序页面路径
}

// Answer 回答
type Answer struct {
	Text        AnswerText         `json:"text"`                  // 回答文本
	Attachments []AnswerAttachment `json:"attachments,omitempty"` // 回答附件列表。最多支持4个
}

// AddKnowledgeIntentRequest 添加知识库问答请求
type AddKnowledgeIntentRequest struct {
	GroupID          string            `json:"group_id"`                    // 分组ID
	Question         Question          `json:"question"`                    // 主问题
	SimilarQuestions *SimilarQuestions `json:"similar_questions,omitempty"` // 相似问题
	Answers          []Answer          `json:"answers"`                     // 回答列表。目前仅支持1个
}

// AddKnowledgeIntentResponse 添加知识库问答响应
type AddKnowledgeIntentResponse struct {
	common.Response
	IntentID string `json:"intent_id"` // 问答ID
}

// DeleteKnowledgeIntentRequest 删除知识库问答请求
type DeleteKnowledgeIntentRequest struct {
	IntentID string `json:"intent_id"` // 问答ID
}

// DeleteKnowledgeIntentResponse 删除知识库问答响应
type DeleteKnowledgeIntentResponse struct {
	common.Response
}

// UpdateKnowledgeIntentRequest 修改知识库问答请求
type UpdateKnowledgeIntentRequest struct {
	IntentID         string            `json:"intent_id"`                   // 问答ID
	Question         *Question         `json:"question,omitempty"`          // 主问题
	SimilarQuestions *SimilarQuestions `json:"similar_questions,omitempty"` // 相似问题
	Answers          []Answer          `json:"answers,omitempty"`           // 回答列表。目前仅支持1个
}

// UpdateKnowledgeIntentResponse 修改知识库问答响应
type UpdateKnowledgeIntentResponse struct {
	common.Response
}

// ListKnowledgeIntentRequest 获取知识库问答列表请求
type ListKnowledgeIntentRequest struct {
	Cursor   string `json:"cursor,omitempty"`    // 上一次调用时返回的next_cursor，第一次拉取可以不填
	Limit    uint32 `json:"limit,omitempty"`     // 每次拉取的数据量，默认值500，最大值为1000
	GroupID  string `json:"group_id,omitempty"`  // 分组ID。可指定拉取特定分组下的问答
	IntentID string `json:"intent_id,omitempty"` // 问答ID。可指定拉取特定的问答
}

// KnowledgeIntent 知识库问答
type KnowledgeIntent struct {
	GroupID          string            `json:"group_id"`                    // 分组ID
	IntentID         string            `json:"intent_id"`                   // 问答ID
	Question         Question          `json:"question"`                    // 主问题
	SimilarQuestions *SimilarQuestions `json:"similar_questions,omitempty"` // 相似问题
	Answers          []Answer          `json:"answers"`                     // 回答列表
}

// ListKnowledgeIntentResponse 获取知识库问答列表响应
type ListKnowledgeIntentResponse struct {
	common.Response
	NextCursor string            `json:"next_cursor"` // 分页游标，再下次请求时填写以获取之后分页的记录
	HasMore    uint32            `json:"has_more"`    // 是否还有更多数据。0-没有 1-有
	IntentList []KnowledgeIntent `json:"intent_list"` // 问答列表
}
