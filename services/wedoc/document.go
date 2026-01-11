package wedoc

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/wedoc"
)

const (
	getDocumentURL         = "/cgi-bin/wedoc/document/get"
	batchUpdateDocumentURL = "/cgi-bin/wedoc/document/batch_update"
)

// GetDocument 获取文档数据
// 该接口用于获取文档数据
func (s *Service) GetDocument(ctx context.Context, req *wedoc.GetDocumentRequest) (*wedoc.GetDocumentResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetDocumentResponse](s.client, ctx, getDocumentURL, req)
}

// BatchUpdateDocument 批量编辑文档内容
// 该接口可以对一个在线文档批量执行多个更新操作
func (s *Service) BatchUpdateDocument(ctx context.Context, req *wedoc.BatchUpdateDocumentRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, batchUpdateDocumentURL, req)
	return err
}
