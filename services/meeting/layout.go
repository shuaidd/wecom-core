package meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/meeting"
)

// ==================== 布局管理相关接口 ====================

// ListLayout 获取会议布局列表
// 根据会议ID返回会议的基础和高级自定义布局信息列表
// 文档: https://developer.work.weixin.qq.com/document/path/93627
func (s *Service) ListLayout(ctx context.Context, meetingID string) (*meeting.ListLayoutResponse, error) {
	req := &meeting.ListLayoutRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[meeting.ListLayoutResponse](s.client, ctx, "/cgi-bin/meeting/advanced_layout/list", req)
}

// AddBasicLayout 添加会议基础布局
// 对API成功预定的会议添加会议基础布局，支持多个布局的添加，每个布局支持多页模板
// 一场会议最多添加10个布局
// 文档: https://developer.work.weixin.qq.com/document/path/93629
func (s *Service) AddBasicLayout(ctx context.Context, req *meeting.AddBasicLayoutRequest) (*meeting.AddBasicLayoutResponse, error) {
	return client.PostAndUnmarshal[meeting.AddBasicLayoutResponse](s.client, ctx, "/cgi-bin/meeting/layout/add", req)
}

// AddAdvancedLayout 添加会议高级布局
// 对当前会议添加高级布局，支持批量添加
// 单个会议最多允许添加20个高级布局
// 注意：高级布局目前仅支持 H.323/SIP 会议室终端
// 文档: https://developer.work.weixin.qq.com/document/path/93630
func (s *Service) AddAdvancedLayout(ctx context.Context, req *meeting.AddAdvancedLayoutRequest) (*meeting.AddAdvancedLayoutResponse, error) {
	return client.PostAndUnmarshal[meeting.AddAdvancedLayoutResponse](s.client, ctx, "/cgi-bin/meeting/advanced_layout/add", req)
}

// UpdateBasicLayout 修改会议基础布局
// 根据布局ID对设置好的会议基础布局进行修改
// 文档: https://developer.work.weixin.qq.com/document/path/93631
func (s *Service) UpdateBasicLayout(ctx context.Context, req *meeting.UpdateBasicLayoutRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/layout/update", req)
	return err
}

// UpdateAdvancedLayout 修改会议高级布局
// 对会议中的高级布局进行修改，注意修改的是布局定义
// 若修改的会议布局正被会议使用，新布局会自动应用到会议
// 接口仅支持全量更新，不支持部分字段单独更新
// 注意：高级布局目前仅支持 H.323/SIP 会议室终端
// 文档: https://developer.work.weixin.qq.com/document/path/93632
func (s *Service) UpdateAdvancedLayout(ctx context.Context, req *meeting.UpdateAdvancedLayoutRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/advanced_layout/update", req)
	return err
}

// BatchDeleteLayout 批量删除布局
// 根据布局ID批量删除布局，可以删除基础布局和高级布局
// 正在被应用的布局无法删除，请先设置成其他布局或恢复成默认原始布局后再行删除
// 接口不做布局是否存在的校验，删除不存在的布局不会有提示
// 注意：高级布局目前仅支持 H.323/SIP 会议室终端
// 文档: https://developer.work.weixin.qq.com/document/path/93633
func (s *Service) BatchDeleteLayout(ctx context.Context, req *meeting.BatchDeleteLayoutRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/advanced_layout/batch_delete", req)
	return err
}

// SetDefaultLayout 设置会议默认布局
// 对API成功预定的会议设置默认布局
// 文档: https://developer.work.weixin.qq.com/document/path/93634
func (s *Service) SetDefaultLayout(ctx context.Context, req *meeting.SetDefaultLayoutRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/layout/set_default", req)
	return err
}

// ApplyAdvancedLayout 设置高级布局
// 将会议中的高级自定义布局应用到指定成员或者整个会议
// 也可以恢复指定成员或整个会议的默认布局
// 注意：高级布局应用到指定成员目前仅支持 H.323/SIP 会议室终端
// 文档: https://developer.work.weixin.qq.com/document/path/93635
func (s *Service) ApplyAdvancedLayout(ctx context.Context, req *meeting.ApplyAdvancedLayoutRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/advanced_layout/apply", req)
	return err
}

// GetUserLayout 获取用户布局
// 根据会议ID和用户ID返回用户的布局设置信息
// 布局优先级：用户个性布局 > 会议自定义布局（高级布局、基础布局）> 会议默认布局
// 注意：高级布局目前仅支持 H.323/SIP 会议室终端
// 文档: https://developer.work.weixin.qq.com/document/path/93636
func (s *Service) GetUserLayout(ctx context.Context, req *meeting.GetUserLayoutRequest) (*meeting.GetUserLayoutResponse, error) {
	return client.PostAndUnmarshal[meeting.GetUserLayoutResponse](s.client, ctx, "/cgi-bin/meeting/advanced_layout/get_user_layout", req)
}

// ListLayoutTemplate 获取布局模板列表
// 获取企业下所有的布局模板列表
// 文档: https://developer.work.weixin.qq.com/document/path/93637
func (s *Service) ListLayoutTemplate(ctx context.Context) (*meeting.ListLayoutTemplateResponse, error) {
	return client.GetAndUnmarshal[meeting.ListLayoutTemplateResponse](s.client, ctx, "/cgi-bin/meeting/layout/list_template", nil)
}

// ==================== 背景管理相关接口 ====================

// ListBackground 获取会议背景列表
// 根据会议ID返回会议背景列表信息
// 文档: https://developer.work.weixin.qq.com/document/path/93638
func (s *Service) ListBackground(ctx context.Context, meetingID string) (*meeting.ListBackgroundResponse, error) {
	req := &meeting.ListBackgroundRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[meeting.ListBackgroundResponse](s.client, ctx, "/cgi-bin/meeting/layout/list_background", req)
}

// AddBackground 添加会议背景
// 对成功预定的会议添加会议背景，支持多个背景图片的添加
// 一场会议最多添加7个背景，且仅支持不超过10MB大小的PNG格式图片，分辨率最小为1920x1080
// 背景图片上传方式为异步上传，您可以通过订阅"素材上传结果"获取上传结果通知
// 文档: https://developer.work.weixin.qq.com/document/path/93639
func (s *Service) AddBackground(ctx context.Context, req *meeting.AddBackgroundRequest) (*meeting.AddBackgroundResponse, error) {
	return client.PostAndUnmarshal[meeting.AddBackgroundResponse](s.client, ctx, "/cgi-bin/meeting/layout/add_background", req)
}

// DeleteBackground 删除会议背景
// 根据背景ID删除单个会议背景
// 正在被会议应用的背景无法删除，请先设置成其他背景或恢复成会议的默认黑色背景后再行删除
// 文档: https://developer.work.weixin.qq.com/document/path/93640
func (s *Service) DeleteBackground(ctx context.Context, req *meeting.DeleteBackgroundRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/layout/delete_background", req)
	return err
}

// BatchDeleteBackground 批量删除会议背景
// 根据背景ID删除多个会议背景
// 正在被会议应用的背景无法删除，请先设置成其他背景或恢复成会议的默认黑色背景后再行删除
// 文档: https://developer.work.weixin.qq.com/document/path/93641
func (s *Service) BatchDeleteBackground(ctx context.Context, req *meeting.BatchDeleteBackgroundRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/layout/batch_delete_background", req)
	return err
}

// SetDefaultBackground 设置会议默认背景
// 对API成功预定的会议设置默认背景
// 文档: https://developer.work.weixin.qq.com/document/path/93642
func (s *Service) SetDefaultBackground(ctx context.Context, req *meeting.SetDefaultBackgroundRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/layout/set_default_background", req)
	return err
}
