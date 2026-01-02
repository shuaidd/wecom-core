package qrcode

import (
	"context"
	"net/url"
	"strconv"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/qrcode"
)

// Service 企业二维码服务
type Service struct {
	client *client.Client
}

// NewService 创建企业二维码服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}

// GetJoinQRCode 获取加入企业二维码
// 文档: https://developer.work.weixin.qq.com/document/path/91714
func (s *Service) GetJoinQRCode(ctx context.Context, sizeType int) (string, error) {
	query := url.Values{}
	if sizeType > 0 {
		query.Set("size_type", strconv.Itoa(sizeType))
	}

	result, err := client.GetAndUnmarshal[qrcode.GetJoinQRCodeResponse](s.client, ctx, "/cgi-bin/corp/get_join_qrcode", query)
	if err != nil {
		return "", err
	}

	return result.JoinQRCode, nil
}

// BatchInvite 邀请成员
// 文档: https://developer.work.weixin.qq.com/document/path/90975
func (s *Service) BatchInvite(ctx context.Context, req *qrcode.BatchInviteRequest) (*qrcode.BatchInviteResponse, error) {
	return client.PostAndUnmarshal[qrcode.BatchInviteResponse](s.client, ctx, "/cgi-bin/batch/invite", req)
}
