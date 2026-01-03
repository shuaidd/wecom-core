package externalcontact

// ProductAttachment 商品图册附件
type ProductAttachment struct {
	Type  string               `json:"type"`
	Image *ImageAttachment     `json:"image,omitempty"`
}

// Product 商品信息
type Product struct {
	ProductID   string              `json:"product_id,omitempty"`
	Description string              `json:"description"`
	Price       int                 `json:"price"`
	ProductSN   string              `json:"product_sn,omitempty"`
	CreateTime  int64               `json:"create_time,omitempty"`
	Attachments []ProductAttachment `json:"attachments"`
}

// AddProductAlbumRequest 创建商品图册请求
type AddProductAlbumRequest struct {
	Description string              `json:"description"`
	Price       int                 `json:"price"`
	ProductSN   string              `json:"product_sn,omitempty"`
	Attachments []ProductAttachment `json:"attachments"`
}

// AddProductAlbumResponse 创建商品图册响应
type AddProductAlbumResponse struct {
	ProductID string `json:"product_id"`
}

// GetProductAlbumRequest 获取商品图册请求
type GetProductAlbumRequest struct {
	ProductID string `json:"product_id"`
}

// GetProductAlbumResponse 获取商品图册响应
type GetProductAlbumResponse struct {
	Product Product `json:"product"`
}

// GetProductAlbumListRequest 获取商品图册列表请求
type GetProductAlbumListRequest struct {
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

// GetProductAlbumListResponse 获取商品图册列表响应
type GetProductAlbumListResponse struct {
	NextCursor  string    `json:"next_cursor,omitempty"`
	ProductList []Product `json:"product_list"`
}

// UpdateProductAlbumRequest 编辑商品图册请求
type UpdateProductAlbumRequest struct {
	ProductID   string              `json:"product_id"`
	Description string              `json:"description,omitempty"`
	Price       int                 `json:"price,omitempty"`
	ProductSN   string              `json:"product_sn,omitempty"`
	Attachments []ProductAttachment `json:"attachments,omitempty"`
}

// DeleteProductAlbumRequest 删除商品图册请求
type DeleteProductAlbumRequest struct {
	ProductID string `json:"product_id"`
}
