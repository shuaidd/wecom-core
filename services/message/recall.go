package message

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/message"
)

// Recall 撤回应用消息
// 文档: https://developer.work.weixin.qq.com/document/path/94867
func (s *Service) Recall(ctx context.Context, req *message.RecallMessageRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/message/recall", req)
	return err
}
