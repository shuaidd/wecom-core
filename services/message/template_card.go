package message

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/message"
)

// UpdateTemplateCard 更新模板卡片消息
// 文档: https://developer.work.weixin.qq.com/document/path/94888
func (s *Service) UpdateTemplateCard(ctx context.Context, req *message.UpdateTemplateCardRequest) (*message.UpdateTemplateCardResponse, error) {
	return client.PostAndUnmarshal[message.UpdateTemplateCardResponse](s.client, ctx, "/cgi-bin/message/update_template_card", req)
}
