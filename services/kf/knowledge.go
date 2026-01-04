package kf

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/kf"
)

// ===================== 知识库分组管理 =====================

// AddKnowledgeGroup 添加知识库分组
// 可通过此接口创建新的知识库分组
// 文档: https://developer.work.weixin.qq.com/document/path/95973
func (s *Service) AddKnowledgeGroup(ctx context.Context, req *kf.AddKnowledgeGroupRequest) (*kf.AddKnowledgeGroupResponse, error) {
	return client.PostAndUnmarshal[kf.AddKnowledgeGroupResponse](s.client, ctx, "/cgi-bin/kf/knowledge/add_group", req)
}

// DeleteKnowledgeGroup 删除知识库分组
// 可通过此接口删除已有的知识库分组，但不能删除系统创建的默认分组
// 文档: https://developer.work.weixin.qq.com/document/path/95973
func (s *Service) DeleteKnowledgeGroup(ctx context.Context, req *kf.DeleteKnowledgeGroupRequest) error {
	_, err := client.PostAndUnmarshal[kf.DeleteKnowledgeGroupResponse](s.client, ctx, "/cgi-bin/kf/knowledge/del_group", req)
	return err
}

// UpdateKnowledgeGroup 修改知识库分组
// 可通过此接口修改已有的知识库分组，但不能修改系统创建的默认分组
// 文档: https://developer.work.weixin.qq.com/document/path/95973
func (s *Service) UpdateKnowledgeGroup(ctx context.Context, req *kf.UpdateKnowledgeGroupRequest) error {
	_, err := client.PostAndUnmarshal[kf.UpdateKnowledgeGroupResponse](s.client, ctx, "/cgi-bin/kf/knowledge/mod_group", req)
	return err
}

// ListKnowledgeGroup 获取知识库分组列表
// 可通过此接口分页获取所有的知识库分组
// 文档: https://developer.work.weixin.qq.com/document/path/95973
func (s *Service) ListKnowledgeGroup(ctx context.Context, req *kf.ListKnowledgeGroupRequest) (*kf.ListKnowledgeGroupResponse, error) {
	return client.PostAndUnmarshal[kf.ListKnowledgeGroupResponse](s.client, ctx, "/cgi-bin/kf/knowledge/list_group", req)
}

// ===================== 知识库问答管理 =====================

// AddKnowledgeIntent 添加知识库问答
// 可通过此接口创建新的知识库问答
// 文档: https://developer.work.weixin.qq.com/document/path/95972
func (s *Service) AddKnowledgeIntent(ctx context.Context, req *kf.AddKnowledgeIntentRequest) (*kf.AddKnowledgeIntentResponse, error) {
	return client.PostAndUnmarshal[kf.AddKnowledgeIntentResponse](s.client, ctx, "/cgi-bin/kf/knowledge/add_intent", req)
}

// DeleteKnowledgeIntent 删除知识库问答
// 可通过此接口删除已有的知识库问答
// 文档: https://developer.work.weixin.qq.com/document/path/95972
func (s *Service) DeleteKnowledgeIntent(ctx context.Context, req *kf.DeleteKnowledgeIntentRequest) error {
	_, err := client.PostAndUnmarshal[kf.DeleteKnowledgeIntentResponse](s.client, ctx, "/cgi-bin/kf/knowledge/del_intent", req)
	return err
}

// UpdateKnowledgeIntent 修改知识库问答
// 可通过此接口修改已有的知识库问答
// question/similar_questions/answers这三部分可以按需更新，但更新的每一部分是覆盖写，需要传完整的字段
// 文档: https://developer.work.weixin.qq.com/document/path/95972
func (s *Service) UpdateKnowledgeIntent(ctx context.Context, req *kf.UpdateKnowledgeIntentRequest) error {
	_, err := client.PostAndUnmarshal[kf.UpdateKnowledgeIntentResponse](s.client, ctx, "/cgi-bin/kf/knowledge/mod_intent", req)
	return err
}

// ListKnowledgeIntent 获取知识库问答列表
// 可通过此接口分页获取的知识库问答详情列表
// 文档: https://developer.work.weixin.qq.com/document/path/95972
func (s *Service) ListKnowledgeIntent(ctx context.Context, req *kf.ListKnowledgeIntentRequest) (*kf.ListKnowledgeIntentResponse, error) {
	return client.PostAndUnmarshal[kf.ListKnowledgeIntentResponse](s.client, ctx, "/cgi-bin/kf/knowledge/list_intent", req)
}
