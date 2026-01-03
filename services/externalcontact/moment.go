package externalcontact

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// AddMomentTask 创建发表任务
// 企业和第三方应用可通过该接口创建客户朋友圈的发表任务
// 文档: https://developer.work.weixin.qq.com/document/path/95094
func (s *Service) AddMomentTask(ctx context.Context, req *externalcontact.AddMomentTaskRequest) (*externalcontact.AddMomentTaskResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddMomentTaskResponse](s.client, ctx, "/cgi-bin/externalcontact/add_moment_task", req)
}

// GetMomentTaskResult 获取任务创建结果
// 由于发表任务的创建是异步执行的，应用需要再调用该接口以获取创建的结果
// 文档: https://developer.work.weixin.qq.com/document/path/95094
func (s *Service) GetMomentTaskResult(ctx context.Context, jobID string) (*externalcontact.GetMomentTaskResultResponse, error) {
	params := url.Values{}
	params.Set("jobid", jobID)
	return client.GetAndUnmarshal[externalcontact.GetMomentTaskResultResponse](s.client, ctx, "/cgi-bin/externalcontact/get_moment_task_result", params)
}

// GetMomentList 获取企业全部的发表列表
// 企业和第三方应用可通过该接口获取企业全部的发表内容
// 文档: https://developer.work.weixin.qq.com/document/path/93333
func (s *Service) GetMomentList(ctx context.Context, req *externalcontact.GetMomentListRequest) (*externalcontact.GetMomentListResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetMomentListResponse](s.client, ctx, "/cgi-bin/externalcontact/get_moment_list", req)
}

// GetMomentTask 获取客户朋友圈企业发表的列表
// 企业和第三方应用可通过该接口获取企业发表的朋友圈成员执行情况
// 文档: https://developer.work.weixin.qq.com/document/path/93333
func (s *Service) GetMomentTask(ctx context.Context, req *externalcontact.GetMomentTaskRequest) (*externalcontact.GetMomentTaskResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetMomentTaskResponse](s.client, ctx, "/cgi-bin/externalcontact/get_moment_task", req)
}

// GetMomentCustomerList 获取客户朋友圈发表时选择的可见范围
// 企业和第三方应用可通过该接口获取客户朋友圈创建时，选择的客户可见范围
// 文档: https://developer.work.weixin.qq.com/document/path/93333
func (s *Service) GetMomentCustomerList(ctx context.Context, req *externalcontact.GetMomentCustomerListRequest) (*externalcontact.GetMomentCustomerListResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetMomentCustomerListResponse](s.client, ctx, "/cgi-bin/externalcontact/get_moment_customer_list", req)
}

// GetMomentSendResult 获取客户朋友圈发表后的可见客户列表
// 企业和第三方应用可通过该接口获取客户朋友圈发表后，可在微信朋友圈中查看的客户列表
// 文档: https://developer.work.weixin.qq.com/document/path/93333
func (s *Service) GetMomentSendResult(ctx context.Context, req *externalcontact.GetMomentSendResultRequest) (*externalcontact.GetMomentSendResultResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetMomentSendResultResponse](s.client, ctx, "/cgi-bin/externalcontact/get_moment_send_result", req)
}

// GetMomentComments 获取客户朋友圈的互动数据
// 企业和第三方应用可通过此接口获取客户朋友圈的互动数据
// 文档: https://developer.work.weixin.qq.com/document/path/93333
func (s *Service) GetMomentComments(ctx context.Context, req *externalcontact.GetMomentCommentsRequest) (*externalcontact.GetMomentCommentsResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetMomentCommentsResponse](s.client, ctx, "/cgi-bin/externalcontact/get_moment_comments", req)
}

// CancelMomentTask 停止发表企业朋友圈
// 企业和第三方应用可调用此接口，停止尚未发送的企业朋友圈发送任务
// 文档: https://developer.work.weixin.qq.com/document/path/97612
func (s *Service) CancelMomentTask(ctx context.Context, req *externalcontact.CancelMomentTaskRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/cancel_moment_task", req)
	return err
}

// ListMomentStrategy 获取规则组列表
// 企业可通过此接口获取企业配置的所有客户朋友圈规则组id列表
// 文档: https://developer.work.weixin.qq.com/document/path/94890
func (s *Service) ListMomentStrategy(ctx context.Context, req *externalcontact.ListMomentStrategyRequest) (*externalcontact.ListMomentStrategyResponse, error) {
	return client.PostAndUnmarshal[externalcontact.ListMomentStrategyResponse](s.client, ctx, "/cgi-bin/externalcontact/moment_strategy/list", req)
}

// GetMomentStrategy 获取规则组详情
// 企业可以通过此接口获取某个客户朋友圈规则组的详细信息
// 文档: https://developer.work.weixin.qq.com/document/path/94890
func (s *Service) GetMomentStrategy(ctx context.Context, req *externalcontact.GetMomentStrategyRequest) (*externalcontact.GetMomentStrategyResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetMomentStrategyResponse](s.client, ctx, "/cgi-bin/externalcontact/moment_strategy/get", req)
}

// GetMomentStrategyRange 获取规则组管理范围
// 企业可通过此接口获取某个朋友圈规则组管理的成员和部门列表
// 文档: https://developer.work.weixin.qq.com/document/path/94890
func (s *Service) GetMomentStrategyRange(ctx context.Context, req *externalcontact.GetMomentStrategyRangeRequest) (*externalcontact.GetMomentStrategyRangeResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetMomentStrategyRangeResponse](s.client, ctx, "/cgi-bin/externalcontact/moment_strategy/get_range", req)
}

// CreateMomentStrategy 创建新的规则组
// 企业可通过此接口创建一个新的客户朋友圈规则组
// 文档: https://developer.work.weixin.qq.com/document/path/94890
func (s *Service) CreateMomentStrategy(ctx context.Context, req *externalcontact.CreateMomentStrategyRequest) (*externalcontact.CreateMomentStrategyResponse, error) {
	return client.PostAndUnmarshal[externalcontact.CreateMomentStrategyResponse](s.client, ctx, "/cgi-bin/externalcontact/moment_strategy/create", req)
}

// EditMomentStrategy 编辑规则组及其管理范围
// 企业可通过此接口编辑规则组的基本信息和修改客户朋友圈规则组管理范围
// 文档: https://developer.work.weixin.qq.com/document/path/94890
func (s *Service) EditMomentStrategy(ctx context.Context, req *externalcontact.EditMomentStrategyRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/moment_strategy/edit", req)
	return err
}

// DeleteMomentStrategy 删除规则组
// 企业可通过此接口删除某个客户朋友圈规则组
// 文档: https://developer.work.weixin.qq.com/document/path/94890
func (s *Service) DeleteMomentStrategy(ctx context.Context, req *externalcontact.DeleteMomentStrategyRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/moment_strategy/del", req)
	return err
}
