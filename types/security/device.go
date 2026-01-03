package security

// ImportTrustDeviceRequest 导入可信企业设备请求
type ImportTrustDeviceRequest struct {
	DeviceList []DeviceInfo `json:"device_list"` // 设备列表，每次最多导入100条
}

// DeviceInfo 设备信息
type DeviceInfo struct {
	System          string   `json:"system"`                     // 设备的类型，Windows或Mac
	MacAddr         []string `json:"mac_addr,omitempty"`         // 设备MAC地址
	MotherboardUUID string   `json:"motherboard_uuid,omitempty"` // 主板UUID（Windows）
	HarddiskUUID    []string `json:"harddisk_uuid,omitempty"`    // 硬盘序列号（Windows）
	Domain          string   `json:"domain,omitempty"`           // Windows域名
	PCName          string   `json:"pc_name,omitempty"`          // Windows计算机名
	SeqNo           string   `json:"seq_no,omitempty"`           // Mac序列号
}

// ImportTrustDeviceResponse 导入可信企业设备响应
type ImportTrustDeviceResponse struct {
	ErrCode int                  `json:"errcode"`
	ErrMsg  string               `json:"errmsg"`
	Result  []ImportDeviceResult `json:"result"` // 导入结果列表
}

// ImportDeviceResult 导入设备结果
type ImportDeviceResult struct {
	DeviceIndex          int    `json:"device_index"`                     // 导入设备记录的标识，从1开始
	DeviceCode           string `json:"device_code,omitempty"`            // 设备的唯一标识
	DuplicatedDeviceCode string `json:"duplicated_device_code,omitempty"` // 重复导入时，冲突的设备的device_code
	Status               int    `json:"status"`                           // 导入结果，1-成功 2-重复导入 3-不支持的设备 4-数据格式错误
}

// ListTrustDeviceRequest 获取设备信息请求
type ListTrustDeviceRequest struct {
	Cursor string `json:"cursor,omitempty"` // 分页cursor
	Limit  int    `json:"limit,omitempty"`  // 查询返回的最大记录数，最高不超过100
	Type   int    `json:"type"`             // 查询设备类型，1-可信企业设备 2-未知设备 3-可信个人设备
}

// ListTrustDeviceResponse 获取设备信息响应
type ListTrustDeviceResponse struct {
	ErrCode    int           `json:"errcode"`
	ErrMsg     string        `json:"errmsg"`
	DeviceList []TrustDevice `json:"device_list"` // 设备列表
	NextCursor string        `json:"next_cursor"` // 分页游标
}

// TrustDevice 可信设备详细信息
type TrustDevice struct {
	DeviceCode       string   `json:"device_code"`                 // 设备编码
	System           string   `json:"system"`                      // 设备的类型，Windows或Mac
	MacAddr          []string `json:"mac_addr,omitempty"`          // 设备MAC地址
	MotherboardUUID  string   `json:"motherboard_uuid,omitempty"`  // 主板UUID
	HarddiskUUID     []string `json:"harddisk_uuid,omitempty"`     // 硬盘UUID
	Domain           string   `json:"domain,omitempty"`            // Windows域
	PCName           string   `json:"pc_name,omitempty"`           // 计算机名
	SeqNo            string   `json:"seq_no,omitempty"`            // Mac序列号
	LastLoginTime    int64    `json:"last_login_time,omitempty"`   // 设备最后登录时间戳
	LastLoginUserID  string   `json:"last_login_userid,omitempty"` // 设备最后登录成员userid
	ConfirmTimestamp int64    `json:"confirm_timestamp,omitempty"` // 设备归属/确认时间戳
	ConfirmUserID    string   `json:"confirm_userid,omitempty"`    // 设备归属/确认成员userid
	ApprovedUserID   string   `json:"approved_userid,omitempty"`   // 通过申报的管理员userid
	Source           int      `json:"source,omitempty"`            // 设备来源 0-未知 1-成员确认 2-管理员导入 3-成员自主申报
	Status           int      `json:"status,omitempty"`            // 设备状态 1-已导入未登录 2-待邀请 3-待管理员确认为企业设备 4-待管理员确认为个人设备 5-已确认为可信企业设备 6-已确认为可信个人设备
}

// GetDeviceByUserRequest 获取成员使用设备请求
type GetDeviceByUserRequest struct {
	LastLoginUserID string `json:"last_login_userid"` // 最后登录的成员userid
	Type            int    `json:"type"`              // 查询设备类型，1-可信企业设备 2-未知设备 3-可信个人设备
}

// GetDeviceByUserResponse 获取成员使用设备响应
type GetDeviceByUserResponse struct {
	ErrCode    int           `json:"errcode"`
	ErrMsg     string        `json:"errmsg"`
	DeviceList []TrustDevice `json:"device_list"` // 设备列表
}

// DeleteTrustDeviceRequest 删除设备信息请求
type DeleteTrustDeviceRequest struct {
	Type           int      `json:"type"`             // 删除设备类型，1-可信企业设备 2-未知设备 3-可信个人设备
	DeviceCodeList []string `json:"device_code_list"` // 设备编码列表，每次最多100个
}

// DeleteTrustDeviceResponse 删除设备信息响应
type DeleteTrustDeviceResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ApproveTrustDeviceRequest 确认为可信设备请求
type ApproveTrustDeviceRequest struct {
	DeviceCodeList []string `json:"device_code_list"` // 设备编码列表，每次最多100个
}

// ApproveTrustDeviceResponse 确认为可信设备响应
type ApproveTrustDeviceResponse struct {
	ErrCode     int      `json:"errcode"`
	ErrMsg      string   `json:"errmsg"`
	SuccessList []string `json:"success_list"` // 确认成功设备code列表
	FailList    []string `json:"fail_list"`    // 确认失败设备code列表
}

// RejectTrustDeviceRequest 驳回可信设备申请请求
type RejectTrustDeviceRequest struct {
	DeviceCodeList []string `json:"device_code_list"` // 设备编码列表，每次最多100个
}

// RejectTrustDeviceResponse 驳回可信设备申请响应
type RejectTrustDeviceResponse struct {
	ErrCode     int      `json:"errcode"`
	ErrMsg      string   `json:"errmsg"`
	SuccessList []string `json:"success_list"` // 驳回成功设备code列表
	FailList    []string `json:"fail_list"`    // 驳回失败设备code列表
}
