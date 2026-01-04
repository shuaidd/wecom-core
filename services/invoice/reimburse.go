package invoice

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/invoice"
)

// GetInvoiceInfo 查询电子发票
// 报销方在获得用户选择的电子发票标识参数后，可以通过该接口查询电子发票的结构化信息，并获取发票PDF文件
// 文档: https://developer.work.weixin.qq.com/document/path/90284
func (s *Service) GetInvoiceInfo(ctx context.Context, cardID, encryptCode string) (*invoice.GetInvoiceInfoResponse, error) {
	req := &invoice.GetInvoiceInfoRequest{
		CardID:      cardID,
		EncryptCode: encryptCode,
	}
	return client.PostAndUnmarshal[invoice.GetInvoiceInfoResponse](s.client, ctx, "/cgi-bin/card/invoice/reimburse/getinvoiceinfo", req)
}

// GetInvoiceInfoBatch 批量查询电子发票
// 报销方在获得用户选择的电子发票标识参数后，可以通过该接口批量查询电子发票的结构化信息
// 文档: https://developer.work.weixin.qq.com/document/path/90285
func (s *Service) GetInvoiceInfoBatch(ctx context.Context, itemList []invoice.InvoiceItem) (*invoice.GetInvoiceInfoBatchResponse, error) {
	req := &invoice.GetInvoiceInfoBatchRequest{
		ItemList: itemList,
	}
	return client.PostAndUnmarshal[invoice.GetInvoiceInfoBatchResponse](s.client, ctx, "/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch", req)
}

// UpdateInvoiceStatus 更新发票状态
// 报销企业和报销服务商可以通过该接口对某一张发票进行锁定、解锁和报销操作
// 文档: https://developer.work.weixin.qq.com/document/path/90283
func (s *Service) UpdateInvoiceStatus(ctx context.Context, cardID, encryptCode string, status invoice.ReimburseStatus) error {
	req := &invoice.UpdateInvoiceStatusRequest{
		CardID:          cardID,
		EncryptCode:     encryptCode,
		ReimburseStatus: status,
	}
	_, err := client.PostAndUnmarshal[invoice.GetInvoiceInfoResponse](s.client, ctx, "/cgi-bin/card/invoice/reimburse/updateinvoicestatus", req)
	return err
}

// UpdateStatusBatch 批量更新发票状态
// 发票平台可以通过该接口对某个成员的一批发票进行锁定、解锁和报销操作
// 注意，报销状态为不可逆状态，请开发者慎重调用
// 文档: https://developer.work.weixin.qq.com/document/path/90286
func (s *Service) UpdateStatusBatch(ctx context.Context, openID string, status invoice.ReimburseStatus, invoiceList []invoice.InvoiceItem) error {
	req := &invoice.UpdateStatusBatchRequest{
		OpenID:          openID,
		ReimburseStatus: status,
		InvoiceList:     invoiceList,
	}
	_, err := client.PostAndUnmarshal[invoice.GetInvoiceInfoResponse](s.client, ctx, "/cgi-bin/card/invoice/reimburse/updatestatusbatch", req)
	return err
}
