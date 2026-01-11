package wedrive

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/wedrive"
)

// Service 微盘服务
type Service struct {
	client *client.Client
}

// New 创建微盘服务实例
func New(c *client.Client) *Service {
	return &Service{client: c}
}

// UploadFile 上传文件
func (s *Service) UploadFile(ctx context.Context, req *wedrive.FileUploadRequest) (*wedrive.FileUploadResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileUploadResponse](s.client, ctx, "/cgi-bin/wedrive/file_upload", req)
}

// DownloadFile 下载文件
func (s *Service) DownloadFile(ctx context.Context, req *wedrive.FileDownloadRequest) (*wedrive.FileDownloadResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileDownloadResponse](s.client, ctx, "/cgi-bin/wedrive/file_download", req)
}

// UploadInit 分块上传初始化
func (s *Service) UploadInit(ctx context.Context, req *wedrive.FileUploadInitRequest) (*wedrive.FileUploadInitResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileUploadInitResponse](s.client, ctx, "/cgi-bin/wedrive/file_upload_init", req)
}

// UploadPart 分块上传（单块）
func (s *Service) UploadPart(ctx context.Context, req *wedrive.FileUploadPartRequest) (*wedrive.FileUploadPartResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileUploadPartResponse](s.client, ctx, "/cgi-bin/wedrive/file_upload_part", req)
}

// UploadFinish 分块上传完成
func (s *Service) UploadFinish(ctx context.Context, req *wedrive.FileUploadFinishRequest) (*wedrive.FileUploadFinishResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileUploadFinishResponse](s.client, ctx, "/cgi-bin/wedrive/file_upload_finish", req)
}

// CreateFile 新建文件/文档/文件夹
func (s *Service) CreateFile(ctx context.Context, req *wedrive.FileCreateRequest) (*wedrive.FileCreateResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileCreateResponse](s.client, ctx, "/cgi-bin/wedrive/file_create", req)
}

// GetFileInfo 获取文件信息
func (s *Service) GetFileInfo(ctx context.Context, req *wedrive.FileInfoRequest) (*wedrive.FileInfoResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileInfoResponse](s.client, ctx, "/cgi-bin/wedrive/file_info", req)
}

// ListFiles 获取文件列表
func (s *Service) ListFiles(ctx context.Context, req *wedrive.FileListRequest) (*wedrive.FileListResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileListResponse](s.client, ctx, "/cgi-bin/wedrive/file_list", req)
}

// DeleteFiles 删除文件
func (s *Service) DeleteFiles(ctx context.Context, req *wedrive.FileDeleteRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/file_delete", req)
	return err
}

// MoveFiles 移动文件
func (s *Service) MoveFiles(ctx context.Context, req *wedrive.FileMoveRequest) (*wedrive.FileMoveResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileMoveResponse](s.client, ctx, "/cgi-bin/wedrive/file_move", req)
}

// RenameFile 重命名文件
func (s *Service) RenameFile(ctx context.Context, req *wedrive.FileRenameRequest) (*wedrive.FileRenameResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileRenameResponse](s.client, ctx, "/cgi-bin/wedrive/file_rename", req)
}

// ShareFile 获取文件分享链接
func (s *Service) ShareFile(ctx context.Context, req *wedrive.FileShareRequest) (*wedrive.FileShareResponse, error) {
	return client.PostAndUnmarshal[wedrive.FileShareResponse](s.client, ctx, "/cgi-bin/wedrive/file_share", req)
}

// SetFileSetting 修改文件分享设置
func (s *Service) SetFileSetting(ctx context.Context, req *wedrive.FileSettingRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/file_setting", req)
	return err
}

// SetFileSecureSetting 修改文件安全设置（水印等）
func (s *Service) SetFileSecureSetting(ctx context.Context, req *wedrive.FileSecureSettingRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/file_secure_setting", req)
	return err
}

// AddFileMembers 新增文件成员
func (s *Service) AddFileMembers(ctx context.Context, req *wedrive.FileACLAddRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/file_acl_add", req)
	return err
}

// RemoveFileMembers 删除文件成员
func (s *Service) RemoveFileMembers(ctx context.Context, req *wedrive.FileACLDelRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/file_acl_del", req)
	return err
}

// GetFilePermission 获取文件权限信息
func (s *Service) GetFilePermission(ctx context.Context, req *wedrive.GetFilePermissionRequest) (*wedrive.GetFilePermissionResponse, error) {
	return client.PostAndUnmarshal[wedrive.GetFilePermissionResponse](s.client, ctx, "/cgi-bin/wedrive/get_file_permission", req)
}

// SpaceCreate 新建空间
func (s *Service) SpaceCreate(ctx context.Context, req *wedrive.SpaceCreateRequest) (*wedrive.SpaceCreateResponse, error) {
	return client.PostAndUnmarshal[wedrive.SpaceCreateResponse](s.client, ctx, "/cgi-bin/wedrive/space_create", req)
}

// GetSpaceInfo 获取空间信息
func (s *Service) GetSpaceInfo(ctx context.Context, req *wedrive.SpaceInfoRequest) (*wedrive.SpaceInfoResponse, error) {
	return client.PostAndUnmarshal[wedrive.SpaceInfoResponse](s.client, ctx, "/cgi-bin/wedrive/space_info", req)
}

// ShareSpace 获取空间邀请链接
func (s *Service) ShareSpace(ctx context.Context, req *wedrive.SpaceShareRequest) (*wedrive.SpaceShareResponse, error) {
	return client.PostAndUnmarshal[wedrive.SpaceShareResponse](s.client, ctx, "/cgi-bin/wedrive/space_share", req)
}

// RenameSpace 重命名空间
func (s *Service) RenameSpace(ctx context.Context, req *wedrive.SpaceRenameRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/space_rename", req)
	return err
}

// DismissSpace 解散空间
func (s *Service) DismissSpace(ctx context.Context, req *wedrive.SpaceDismissRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/space_dismiss", req)
	return err
}

// SetSpaceSetting 修改空间安全设置
func (s *Service) SetSpaceSetting(ctx context.Context, req *wedrive.SpaceSettingRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/space_setting", req)
	return err
}

// AddSpaceMembers 在空间中新增成员/部门
func (s *Service) AddSpaceMembers(ctx context.Context, req *wedrive.SpaceCreateRequest) (*wedrive.SpaceCreateResponse, error) {
	// space_acl_add shares the same request structure as space create's auth_info
	return client.PostAndUnmarshal[wedrive.SpaceCreateResponse](s.client, ctx, "/cgi-bin/wedrive/space_acl_add", req)
}

// RemoveSpaceMembers 在空间中移除成员/部门
func (s *Service) RemoveSpaceMembers(ctx context.Context, req *wedrive.SpaceInfoRequest) error {
	// use space_acl_del endpoint payload; here using a generic request type from types if needed
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/wedrive/space_acl_del", req)
	return err
}
