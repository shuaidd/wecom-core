package invoice

// ReimburseStatus 发票报销状态
type ReimburseStatus string

const (
	// ReimburseStatusInit 发票初始状态，未锁定
	ReimburseStatusInit ReimburseStatus = "INVOICE_REIMBURSE_INIT"
	// ReimburseStatusLock 发票已锁定，无法重复提交报销
	ReimburseStatusLock ReimburseStatus = "INVOICE_REIMBURSE_LOCK"
	// ReimburseStatusClosure 发票已核销，从用户卡包中移除
	ReimburseStatusClosure ReimburseStatus = "INVOICE_REIMBURSE_CLOSURE"
)

// InvoiceItem 发票项目信息
type InvoiceItem struct {
	// CardID 发票卡券的card_id
	CardID string `json:"card_id"`
	// EncryptCode 发票卡券的加密code
	EncryptCode string `json:"encrypt_code"`
}

// GetInvoiceInfoRequest 查询电子发票请求
type GetInvoiceInfoRequest struct {
	// CardID 发票id
	CardID string `json:"card_id"`
	// EncryptCode 加密code
	EncryptCode string `json:"encrypt_code"`
}

// InvoiceItemInfo 商品信息
type InvoiceItemInfo struct {
	// Name 项目（商品）名称
	Name string `json:"name"`
	// Num 项目数量
	Num int `json:"num,omitempty"`
	// Unit 项目单位
	Unit string `json:"unit,omitempty"`
	// Price 单价，以分为单位
	Price int `json:"price"`
	// Fee 金额，以分为单位
	Fee int `json:"fee,omitempty"`
}

// InvoiceUserInfo 发票用户信息
type InvoiceUserInfo struct {
	// Fee 发票加税合计金额，以分为单位
	Fee int `json:"fee"`
	// Title 发票的抬头
	Title string `json:"title"`
	// BillingTime 开票时间，为十位时间戳
	BillingTime int64 `json:"billing_time"`
	// BillingNo 发票代码
	BillingNo string `json:"billing_no"`
	// BillingCode 发票号码
	BillingCode string `json:"billing_code"`
	// Tax 税额,以分为单位
	Tax int `json:"tax"`
	// FeeWithoutTax 不含税金额，以分为单位
	FeeWithoutTax int `json:"fee_without_tax"`
	// Detail 发票详情
	Detail string `json:"detail"`
	// PdfURL 这张发票对应的PDF_URL
	PdfURL string `json:"pdf_url"`
	// TripPdfURL 其它消费凭证附件对应的URL
	TripPdfURL string `json:"trip_pdf_url,omitempty"`
	// CheckCode 校验码
	CheckCode string `json:"check_code"`
	// BuyerNumber 购买方纳税人识别号
	BuyerNumber string `json:"buyer_number,omitempty"`
	// BuyerAddressAndPhone 购买方地址、电话
	BuyerAddressAndPhone string `json:"buyer_address_and_phone,omitempty"`
	// BuyerBankAccount 购买方开户行及账号
	BuyerBankAccount string `json:"buyer_bank_account,omitempty"`
	// SellerNumber 销售方纳税人识别号
	SellerNumber string `json:"seller_number,omitempty"`
	// SellerAddressAndPhone 销售方地址、电话
	SellerAddressAndPhone string `json:"seller_address_and_phone,omitempty"`
	// SellerBankAccount 销售方开户行及账号
	SellerBankAccount string `json:"seller_bank_account,omitempty"`
	// Remarks 备注
	Remarks string `json:"remarks,omitempty"`
	// Cashier 收款人
	Cashier string `json:"cashier,omitempty"`
	// Maker 开票人
	Maker string `json:"maker,omitempty"`
	// ReimburseStatus 发报销状态
	ReimburseStatus ReimburseStatus `json:"reimburse_status"`
	// Info 商品信息结构
	Info []InvoiceItemInfo `json:"info,omitempty"`
	// OrderID 订单ID
	OrderID string `json:"order_id,omitempty"`
}

// GetInvoiceInfoResponse 查询电子发票响应
type GetInvoiceInfoResponse struct {
	// CardID 发票id
	CardID string `json:"card_id"`
	// BeginTime 发票的有效期起始时间
	BeginTime int64 `json:"begin_time"`
	// EndTime 发票的有效期截止时间
	EndTime int64 `json:"end_time"`
	// OpenID 用户标识
	OpenID string `json:"openid"`
	// Type 发票类型
	Type string `json:"type"`
	// Payee 发票的收款方
	Payee string `json:"payee"`
	// Detail 发票详情
	Detail string `json:"detail"`
	// UserInfo 发票的用户信息
	UserInfo InvoiceUserInfo `json:"user_info"`
}

// GetInvoiceInfoBatchRequest 批量查询电子发票请求
type GetInvoiceInfoBatchRequest struct {
	// ItemList 发票列表
	ItemList []InvoiceItem `json:"item_list"`
}

// InvoiceInfo 发票信息
type InvoiceInfo struct {
	// CardID 发票id
	CardID string `json:"card_id"`
	// BeginTime 发票的有效期起始时间
	BeginTime int64 `json:"begin_time,omitempty"`
	// EndTime 发票的有效期截止时间
	EndTime int64 `json:"end_time,omitempty"`
	// OpenID 用户标识
	OpenID string `json:"openid"`
	// Type 发票类型
	Type string `json:"type"`
	// Payee 发票的收款方
	Payee string `json:"payee"`
	// Detail 发票详情
	Detail string `json:"detail"`
	// UserInfo 发票的用户信息
	UserInfo InvoiceUserInfo `json:"user_info"`
}

// GetInvoiceInfoBatchResponse 批量查询电子发票响应
type GetInvoiceInfoBatchResponse struct {
	// ItemList 发票信息列表
	ItemList []InvoiceInfo `json:"item_list"`
}

// UpdateInvoiceStatusRequest 更新发票状态请求
type UpdateInvoiceStatusRequest struct {
	// CardID 发票id
	CardID string `json:"card_id"`
	// EncryptCode 加密code
	EncryptCode string `json:"encrypt_code"`
	// ReimburseStatus 发报销状态
	ReimburseStatus ReimburseStatus `json:"reimburse_status"`
}

// UpdateStatusBatchRequest 批量更新发票状态请求
type UpdateStatusBatchRequest struct {
	// OpenID 用户openid
	OpenID string `json:"openid"`
	// ReimburseStatus 发票报销状态
	ReimburseStatus ReimburseStatus `json:"reimburse_status"`
	// InvoiceList 发票列表，必须全部属于同一个openid
	InvoiceList []InvoiceItem `json:"invoice_list"`
}
