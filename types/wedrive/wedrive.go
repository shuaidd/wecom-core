package wedrive

import "github.com/shuaidd/wecom-core/types/common"

// ==================== 文件上传 / 下载 ====================

type FileUploadRequest struct {
	SpaceID           string `json:"spaceid,omitempty"`
	FatherID          string `json:"fatherid,omitempty"`
	SelectedTicket    string `json:"selected_ticket,omitempty"`
	FileName          string `json:"file_name"`
	FileBase64Content string `json:"file_base64_content"`
}

type FileUploadResponse struct {
	common.Response
	FileID string `json:"fileid,omitempty"`
}

type FileDownloadRequest struct {
	FileID         string `json:"fileid,omitempty"`
	SelectedTicket string `json:"selected_ticket,omitempty"`
}

type FileDownloadResponse struct {
	common.Response
	DownloadURL string `json:"download_url,omitempty"`
	CookieName  string `json:"cookie_name,omitempty"`
	CookieValue string `json:"cookie_value,omitempty"`
}

// ==================== 分块上传 ====================

type FileUploadInitRequest struct {
	SpaceID        string   `json:"spaceid,omitempty"`
	FatherID       string   `json:"fatherid,omitempty"`
	SelectedTicket string   `json:"selected_ticket,omitempty"`
	FileName       string   `json:"file_name"`
	Size           uint64   `json:"size"`
	BlockSHA       []string `json:"block_sha"`
	SkipPushCard   bool     `json:"skip_push_card,omitempty"`
}

type FileUploadInitResponse struct {
	common.Response
	HitExist  bool   `json:"hit_exist,omitempty"`
	UploadKey string `json:"upload_key,omitempty"`
	FileID    string `json:"fileid,omitempty"`
}

type FileUploadPartRequest struct {
	UploadKey         string `json:"upload_key"`
	Index             int32  `json:"index"`
	FileBase64Content string `json:"file_base64_content"`
}

type FileUploadPartResponse struct {
	common.Response
}

type FileUploadFinishRequest struct {
	UploadKey string `json:"upload_key"`
}

type FileUploadFinishResponse struct {
	common.Response
	FileID string `json:"fileid,omitempty"`
}

// ==================== 文件/文档管理 ====================

type FileCreateRequest struct {
	SpaceID  string `json:"spaceid"`
	FatherID string `json:"fatherid"`
	FileType uint32 `json:"file_type"`
	FileName string `json:"file_name"`
}

type FileCreateResponse struct {
	common.Response
	FileID string `json:"fileid,omitempty"`
	URL    string `json:"url,omitempty"`
}

type FileInfo struct {
	FileID   string `json:"fileid,omitempty"`
	FileName string `json:"file_name,omitempty"`
	SpaceID  string `json:"spaceid,omitempty"`
	FatherID string `json:"fatherid,omitempty"`
	FileSize uint64 `json:"file_size,omitempty"`
	CTime    int64  `json:"ctime,omitempty"`
	MTime    int64  `json:"mtime,omitempty"`
	FileType uint32 `json:"file_type,omitempty"`
	FileStat uint32 `json:"file_status,omitempty"`
	Sha      string `json:"sha,omitempty"`
	Md5      string `json:"md5,omitempty"`
	URL      string `json:"url,omitempty"`
}

type FileInfoResponse struct {
	common.Response
	FileInfo *FileInfo `json:"file_info,omitempty"`
}

type FileInfoRequest struct {
	FileID string `json:"fileid"`
}

type SpaceACLAddRequest struct {
	SpaceID  string     `json:"spaceid"`
	AuthInfo []AuthInfo `json:"auth_info"`
}

type SpaceACLDelRequest struct {
	SpaceID  string     `json:"spaceid"`
	AuthInfo []AuthInfo `json:"auth_info"`
}

type FileListRequest struct {
	SpaceID  string `json:"spaceid"`
	FatherID string `json:"fatherid"`
	SortType uint32 `json:"sort_type,omitempty"`
	Start    uint32 `json:"start,omitempty"`
	Limit    uint32 `json:"limit,omitempty"`
}

type FileListResponse struct {
	common.Response
	HasMore   bool       `json:"has_more,omitempty"`
	NextStart uint32     `json:"next_start,omitempty"`
	FileList  []FileInfo `json:"file_list,omitempty"`
}

type FileDeleteRequest struct {
	FileID []string `json:"fileid"`
}

type FileDeleteResponse struct {
	common.Response
}

type FileMoveRequest struct {
	FatherID string   `json:"fatherid"`
	Replace  bool     `json:"replace,omitempty"`
	FileID   []string `json:"fileid"`
}

type FileMoveResponse struct {
	common.Response
	FileList []FileInfo `json:"file_list,omitempty"`
}

type FileRenameRequest struct {
	FileID  string `json:"fileid"`
	NewName string `json:"new_name"`
}

type FileRenameResponse struct {
	common.Response
	File *FileInfo `json:"file,omitempty"`
}

// ==================== 权限 / 分享 / 安全 ====================

type AuthInfo struct {
	Type         uint32 `json:"type,omitempty"`
	UserID       string `json:"userid,omitempty"`
	DepartmentID uint32 `json:"departmentid,omitempty"`
	Auth         uint32 `json:"auth,omitempty"`
}

type FileACLAddRequest struct {
	FileID   string     `json:"fileid"`
	AuthInfo []AuthInfo `json:"auth_info"`
}

type FileACLDelRequest struct {
	FileID   string     `json:"fileid"`
	AuthInfo []AuthInfo `json:"auth_info"`
}

type CommonOKResponse struct {
	common.Response
}

type FileShareRequest struct {
	FileID string `json:"fileid"`
}

type FileShareResponse struct {
	common.Response
	ShareURL string `json:"share_url,omitempty"`
}

type FileSettingRequest struct {
	FileID    string `json:"fileid"`
	AuthScope uint32 `json:"auth_scope"`
	Auth      uint32 `json:"auth,omitempty"`
}

type WatermarkSetting struct {
	Text              string `json:"text,omitempty"`
	MarginType        uint32 `json:"margin_type,omitempty"`
	ShowVisitorName   bool   `json:"show_visitor_name,omitempty"`
	ShowText          bool   `json:"show_text,omitempty"`
	ForceByAdmin      bool   `json:"force_by_admin,omitempty"`
	ForceBySpaceAdmin bool   `json:"force_by_space_admin,omitempty"`
}

type FileSecureSettingRequest struct {
	FileID    string           `json:"fileid"`
	Watermark WatermarkSetting `json:"watermark,omitempty"`
}

// 获取文件权限信息 response（简化版）
type GetFilePermissionRequest struct {
	FileID string `json:"fileid"`
}

type ShareRange struct {
	EnableCorpInternal bool   `json:"enable_corp_internal,omitempty"`
	CorpInternalAuth   uint32 `json:"corp_internal_auth,omitempty"`
	EnableCorpExternal bool   `json:"enable_corp_external,omitempty"`
	CorpExternalAuth   uint32 `json:"corp_external_auth,omitempty"`
}

type SecureSetting struct {
	EnableReadonlyCopy    bool `json:"enable_readonly_copy,omitempty"`
	ModifyOnlyByAdmin     bool `json:"modify_only_by_admin,omitempty"`
	EnableReadonlyComment bool `json:"enable_readonly_comment,omitempty"`
	BanShareExternal      bool `json:"ban_share_external,omitempty"`
}

type InheritFatherAuth struct {
	AuthList []AuthInfo `json:"auth_list,omitempty"`
	Inherit  bool       `json:"inherit,omitempty"`
}

type GetFilePermissionResponse struct {
	common.Response
	ShareRange     *ShareRange        `json:"share_range,omitempty"`
	SecureSetting  *SecureSetting     `json:"secure_setting,omitempty"`
	InheritFather  *InheritFatherAuth `json:"inherit_father_auth,omitempty"`
	FileMemberList []AuthInfo         `json:"file_member_list,omitempty"`
	Watermark      *WatermarkSetting  `json:"watermark,omitempty"`
}

// ==================== 空间管理 ====================

type SpaceCreateRequest struct {
	SpaceName    string     `json:"space_name"`
	AuthInfo     []AuthInfo `json:"auth_info,omitempty"`
	SpaceSubType uint32     `json:"space_sub_type,omitempty"`
}

type SpaceCreateResponse struct {
	common.Response
	SpaceID string `json:"spaceid,omitempty"`
}

type SpaceInfoRequest struct {
	SpaceID string `json:"spaceid"`
}

type SpaceInfo struct {
	SpaceID   string `json:"spaceid,omitempty"`
	SpaceName string `json:"space_name,omitempty"`
}

type SpaceInfoResponse struct {
	common.Response
	SpaceInfo *SpaceInfo `json:"space_info,omitempty"`
}

type SpaceShareRequest struct {
	SpaceID string `json:"spaceid"`
}

type SpaceShareResponse struct {
	common.Response
	SpaceShareURL string `json:"space_share_url,omitempty"`
}

type SpaceSettingRequest struct {
	SpaceID                      string  `json:"spaceid"`
	EnableWatermark              *bool   `json:"enable_watermark,omitempty"`
	ShareURLNoApprove            *bool   `json:"share_url_no_approve,omitempty"`
	ShareURLNoApproveDefaultAuth *uint32 `json:"share_url_no_approve_default_auth,omitempty"`
	EnableConfidentialMode       *bool   `json:"enable_confidential_mode,omitempty"`
	DefaultFileScope             *uint32 `json:"default_file_scope,omitempty"`
	BanShareExternal             *bool   `json:"ban_share_external,omitempty"`
}

type SpaceSettingResponse struct {
	common.Response
}

// 解散、重命名、移除成员等使用通用响应
type SpaceRenameRequest struct {
	SpaceID   string `json:"spaceid"`
	SpaceName string `json:"space_name"`
}

type SpaceDismissRequest struct {
	SpaceID string `json:"spaceid"`
}
