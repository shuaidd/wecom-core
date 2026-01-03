package externalcontact

// GetFollowUserListResponse 获取配置了客户联系功能的成员列表响应
type GetFollowUserListResponse struct {
	FollowUser []string `json:"follow_user"`
}
