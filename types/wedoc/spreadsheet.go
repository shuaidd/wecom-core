package wedoc

// ==================== 获取表格行列信息 ====================

// GetSheetPropertiesRequest 获取表格行列信息请求
type GetSheetPropertiesRequest struct {
	DocID string `json:"docid"` // 在线表格的docid
}

// GetSheetPropertiesResponse 获取表格行列信息响应
type GetSheetPropertiesResponse struct {
	Properties []Properties `json:"properties"` // 工作表属性
}

// Properties 工作表元数据相关的资源描述
type Properties struct {
	SheetID     string `json:"sheet_id"`     // 工作表ID，工作表的唯一标识
	Title       string `json:"title"`        // 工作表名称
	RowCount    uint32 `json:"row_count"`    // 表格的总行数
	ColumnCount uint32 `json:"column_count"` // 表格的总列数
}

// ==================== 获取表格数据 ====================

// GetSheetRangeDataRequest 获取表格数据请求
type GetSheetRangeDataRequest struct {
	DocID   string `json:"docid"`    // 在线表格唯一标识
	SheetID string `json:"sheet_id"` // 工作表ID，工作表的唯一标识
	Range   string `json:"range"`    // 查询的范围，格式遵循 A1表示法
}

// GetSheetRangeDataResponse 获取表格数据响应
type GetSheetRangeDataResponse struct {
	Data *GetSheetRangeDataResult `json:"data"` // 返回数据
}

// GetSheetRangeDataResult 获取表格数据结果
type GetSheetRangeDataResult struct {
	Result *GridData `json:"result"` // 表格数据
}

// GridData 表格的具体数据内容
type GridData struct {
	StartRow    uint32     `json:"start_row"`    // 起始行编号 （从0开始计算）
	StartColumn uint32     `json:"start_column"` // 起始列编号 （从0开始计算）
	Rows        []*RowData `json:"rows"`         // 各行的数据
}

// RowData 行数据的资源描述
type RowData struct {
	Values []*CellData `json:"values"` // 各个单元格的数据内容
}

// CellData 单元格的信息
type CellData struct {
	CellValue  *CellValue  `json:"cell_value,omitempty"`  // 单元格的数据内容
	CellFormat *CellFormat `json:"cell_format,omitempty"` // 单元格的样式信息
}

// CellValue 单元格的数据内容
type CellValue struct {
	Text string    `json:"text,omitempty"` // 文本内容
	Link *CellLink `json:"link,omitempty"` // 超链接内容
}

// CellLink 超链接的相关信息
type CellLink struct {
	URL  string `json:"url"`  // 链接url
	Text string `json:"text"` // 链接标题
}

// CellFormat 单元格的样式信息
type CellFormat struct {
	TextFormat *TextFormat `json:"text_format,omitempty"` // 文字样式
}

// TextFormat 文本样式信息
type TextFormat struct {
	Font          string `json:"font,omitempty"`          // 字体名称
	FontSize      uint32 `json:"font_size,omitempty"`     // 字体大小，最大72
	Bold          bool   `json:"bold,omitempty"`          // 字体加粗
	Italic        bool   `json:"italic,omitempty"`        // 斜体
	Strikethrough bool   `json:"strikethrough,omitempty"` // 字体删除线
	Underline     bool   `json:"underline,omitempty"`     // 下划线
	Color         *Color `json:"color,omitempty"`         // 字体颜色
}

// Color 颜色信息，采用 RGBA 表示法
type Color struct {
	Red   uint32 `json:"red"`   // 红色，取值范围：[0,255]
	Green uint32 `json:"green"` // 绿色，取值范围：[0,255]
	Blue  uint32 `json:"blue"`  // 蓝色，取值范围：[0,255]
	Alpha uint32 `json:"alpha"` // alpha通道，取值范围：[0,255]，默认值为255完全不透明
}

// ==================== 编辑表格内容 ====================

// BatchUpdateSpreadsheetRequest 批量编辑表格内容请求
type BatchUpdateSpreadsheetRequest struct {
	DocID    string                     `json:"docid"`    // 文档的docid
	Requests []SpreadsheetUpdateRequest `json:"requests"` // 更新操作列表
}

// BatchUpdateSpreadsheetResponse 批量编辑表格内容响应
type BatchUpdateSpreadsheetResponse struct {
	Data *BatchUpdateSpreadsheetData `json:"data"` // 返回数据
}

// BatchUpdateSpreadsheetData 批量编辑表格内容响应数据
type BatchUpdateSpreadsheetData struct {
	Responses []SpreadsheetUpdateResponse `json:"responses"` // 结果列表
}

// SpreadsheetUpdateRequest 更新请求
type SpreadsheetUpdateRequest struct {
	AddSheetRequest        *SpreadsheetAddSheetRequest    `json:"add_sheet_request,omitempty"`        // 新增工作表
	DeleteSheetRequest     *SpreadsheetDeleteSheetRequest `json:"delete_sheet_request,omitempty"`     // 删除工作表
	UpdateRangeRequest     *UpdateRangeRequest            `json:"update_range_request,omitempty"`     // 更新范围内单元格内容
	DeleteDimensionRequest *DeleteDimensionRequest        `json:"delete_dimension_request,omitempty"` // 删除表格连续的行或列
}

// SpreadsheetAddSheetRequest 新增工作表
type SpreadsheetAddSheetRequest struct {
	Title       string `json:"title"`        // 工作表名称
	RowCount    uint32 `json:"row_count"`    // 新增工作表的初始行数
	ColumnCount uint32 `json:"column_count"` // 新增工作表的初始列数
}

// SpreadsheetDeleteSheetRequest 删除工作表
type SpreadsheetDeleteSheetRequest struct {
	SheetID string `json:"sheet_id"` // 工作表唯一标识
}

// UpdateRangeRequest 更新范围内单元格内容请求
type UpdateRangeRequest struct {
	SheetID  string    `json:"sheet_id"`  // 工作表唯一标识
	GridData *GridData `json:"grid_data"` // 写入指定区域的数据
}

// DeleteDimensionRequest 删除表格连续的行（或列）的请求
type DeleteDimensionRequest struct {
	SheetID    string `json:"sheet_id"`    // 工作表唯一标识
	Dimension  string `json:"dimension"`   // 声明删除的维度为行或者列
	StartIndex uint32 `json:"start_index"` // 删除行列的起始序号（从1开始）
	EndIndex   uint32 `json:"end_index"`   // 删除行列的终止序号（从1开始）
}

// SpreadsheetUpdateResponse 更新操作对应的响应结构体类型
type SpreadsheetUpdateResponse struct {
	AddSheetResponse        *SpreadsheetAddSheetResponse    `json:"add_sheet_response,omitempty"`        // 新增工作表响应结构体
	DeleteSheetResponse     *SpreadsheetDeleteSheetResponse `json:"delete_sheet_response,omitempty"`     // 删除工作表响应结构体
	UpdateRangeResponse     *UpdateRangeResponse            `json:"update_range_response,omitempty"`     // 更新范围内单元格内容响应结构体
	DeleteDimensionResponse *DeleteDimensionResponse        `json:"delete_dimension_response,omitempty"` // 删除表格连续的行或列响应结构体
}

// SpreadsheetAddSheetResponse 新增子表操作的请求响应体结构
type SpreadsheetAddSheetResponse struct {
	Properties *Properties `json:"properties"` // 新增子表的属性
}

// SpreadsheetDeleteSheetResponse 删除工作表请求的相应结构体
type SpreadsheetDeleteSheetResponse struct {
	SheetID string `json:"sheet_id"` // 被删除的工作表的唯一标识
}

// UpdateRangeResponse 编辑区域内单元格内容请求响应体结构
type UpdateRangeResponse struct {
	UpdatedCells uint32 `json:"updated_cells"` // 数据更新的成功的单元格数量
}

// DeleteDimensionResponse 删除表格连续的行（或列），请求响应体结构
type DeleteDimensionResponse struct {
	Deleted int `json:"deleted"` // 被删除的行数（或列数）
}
