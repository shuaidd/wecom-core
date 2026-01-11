package wedoc

// ==================== 获取文档数据 ====================

// GetDocumentRequest 获取文档数据请求
type GetDocumentRequest struct {
	DocID string `json:"docid"` // 文档的docid
}

// GetDocumentResponse 获取文档数据响应
type GetDocumentResponse struct {
	Version  uint32 `json:"version"`  // 文档版本
	Document *Node  `json:"document"` // 文档内容根节点
}

// Node 构成文档内容的节点
type Node struct {
	Begin    uint32    `json:"begin"`              // 起始位置
	End      uint32    `json:"end"`                // 结束位置
	Property *Property `json:"property,omitempty"` // 该节点的属性
	Type     string    `json:"type"`               // 节点类型
	Children []*Node   `json:"children,omitempty"` // 子节点
	Text     string    `json:"text,omitempty"`     // 文本内容，当节点类型为 Text 时有效
}

// Property 节点属性
type Property struct {
	SectionProperty   *SectionProperty   `json:"section_property,omitempty"`    // section 属性
	ParagraphProperty *ParagraphProperty `json:"paragraph_property,omitempty"`  // 段落属性
	RunProperty       *RunProperty       `json:"run_property,omitempty"`        // 文本属性
	TableProperty     *TableProperty     `json:"table_property,omitempty"`      // 表格属性
	TableRowProperty  *TableRowProperty  `json:"table_row_property,omitempty"`  // 表格行属性
	TableCellProperty *TableCellProperty `json:"table_cell_property,omitempty"` // 表格单元属性
	DrawingProperty   *DrawingProperty   `json:"drawing_property,omitempty"`    // drawing 属性
}

// SectionProperty Section 属性
type SectionProperty struct {
	PageSize    *PageSize    `json:"page_size,omitempty"`    // 页面尺寸
	PageMargins *PageMargins `json:"page_margins,omitempty"` // 页边距
}

// PageSize 页面尺寸
type PageSize struct {
	Width       float64 `json:"width"`                 // 页面宽度
	Height      float64 `json:"height"`                // 页面高度
	Orientation string  `json:"orientation,omitempty"` // 页面方向
}

// PageMargins 页边距
type PageMargins struct {
	Top    float64 `json:"top"`    // 上边距
	Right  float64 `json:"right"`  // 右边距
	Bottom float64 `json:"bottom"` // 下边距
	Left   float64 `json:"left"`   // 左边距
}

// ParagraphProperty 段落属性
type ParagraphProperty struct {
	NumberProperty *NumberProperty `json:"number_property,omitempty"` // 段落的编号属性
	Spacing        *Spacing        `json:"spacing,omitempty"`         // 段落间距
	Indent         *Indent         `json:"indent,omitempty"`          // 段落缩进
	AlignmentType  string          `json:"alignment_type,omitempty"`  // 文字水平方向的对齐类型
	TextDirection  string          `json:"text_direction,omitempty"`  // 文字方向
}

// NumberProperty 段落的编号属性
type NumberProperty struct {
	NestingLevel uint32 `json:"nesting_level"` // 编号缩进层级
	NumberID     string `json:"number_id"`     // 编号 ID
}

// Spacing 段落间距
type Spacing struct {
	Before   float64 `json:"before"`              // 段后间距，单位是像素（px）
	After    float64 `json:"after"`               // 段前间距，单位是像素（px）
	Line     float64 `json:"line"`                // 行间距数值，单位是像素（px）
	LineRule string  `json:"line_rule,omitempty"` // 行间距格式
}

// Indent 段落缩进
type Indent struct {
	Left           float64 `json:"left,omitempty"`             // 缩进左侧，单位是像素（px）
	LeftChars      uint32  `json:"left_chars,omitempty"`       // 缩进左侧字符数，单位 1/20 字符宽度
	Right          float64 `json:"right,omitempty"`            // 缩进右侧，单位是像素（px）
	RightChars     uint32  `json:"right_chars,omitempty"`      // 缩进右侧字符数，单位 1/20 字符宽度
	Hanging        float64 `json:"hanging,omitempty"`          // 垂直悬挂，单位是像素（px）
	HangingChars   uint32  `json:"hanging_chars,omitempty"`    // 垂直悬挂字符数，单位 1/20 字符宽度
	FirstLine      float64 `json:"first_line,omitempty"`       // 首行缩进，单位是像素（px）
	FirstLineChars uint32  `json:"first_line_chars,omitempty"` // 首行缩进字符数，单位 1/20 字符宽度
}

// RunProperty text 的属性
type RunProperty struct {
	Font          string   `json:"font,omitempty"`           // 字体
	Bold          bool     `json:"bold,omitempty"`           // 文字是否加粗
	Italics       bool     `json:"italics,omitempty"`        // 文字是否斜体表示
	Underline     bool     `json:"underline,omitempty"`      // 文字是否下划线
	Strike        bool     `json:"strike,omitempty"`         // 文字是否被删除线贯穿
	Color         string   `json:"color,omitempty"`          // 文字的颜色，颜色使用十六进制，RRGGBB格式
	Spacing       float64  `json:"spacing,omitempty"`        // 字符的间距
	Size          float64  `json:"size,omitempty"`           // 文字的大小，单位是半个点（half-points）
	Shading       *Shading `json:"shading,omitempty"`        // 文字阴影
	VerticalAlign string   `json:"vertical_align,omitempty"` // 垂直对齐类型
	IsPlaceholder bool     `json:"is_placeholder,omitempty"` // 本节点是否占位符
}

// Shading 阴影
type Shading struct {
	ForegroundColor string `json:"foreground_color,omitempty"` // 前景色，颜色使用十六进制RRGGBB 格式
	BackgroundColor string `json:"background_color,omitempty"` // 背景色，颜色使用十六进制 RRGGBB格式
}

// TableProperty 表格属性
type TableProperty struct {
	TableWidth              *TableWidth `json:"table_width,omitempty"`               // 表格宽度
	HorizontalAlignmentType string      `json:"horizontal_alignment_type,omitempty"` // 表格的水平对齐的方式
	TableLayout             string      `json:"table_layout,omitempty"`              // 表格布局
}

// TableWidth 表格宽度
type TableWidth struct {
	Width float64 `json:"width"` // 表格宽度，单位是像素（px）
	Type  string  `json:"type"`  // 表格宽度类型
}

// TableRowProperty 表格行属性
type TableRowProperty struct {
	IsHeader bool `json:"is_header"` // 本行是否是表头
}

// TableCellProperty 表格单元属性
type TableCellProperty struct {
	TableCellBorders  *Borders `json:"table_cell_borders,omitempty"` // 边界属性
	VerticalAlignment string   `json:"vertical_alignment,omitempty"` // 垂直方向对齐属性
}

// Borders 表格单元的边界属性
type Borders struct {
	Top    *BorderProperty `json:"top,omitempty"`    // 上边界
	Left   *BorderProperty `json:"left,omitempty"`   // 左边界
	Bottom *BorderProperty `json:"bottom,omitempty"` // 底部边界
	Right  *BorderProperty `json:"right,omitempty"`  // 右边界
}

// BorderProperty 边界属性
type BorderProperty struct {
	Color string `json:"color"` // 边界颜色，颜色使用十六进制RRGGBB格式
	Width uint32 `json:"width"` // 边界的宽度，单位是像素（px）
}

// DrawingProperty Drawing 属性
type DrawingProperty struct {
	InlineKeyword *Inline `json:"inline_keyword,omitempty"` // Drawing 类型中的实体，如一张图片
	Anchor        *Anchor `json:"anchor,omitempty"`         // Drawing 类型中的悬浮实体，如一张图片
	IsPlaceholder bool    `json:"is_placeholder,omitempty"` // 此处是否为占位符
}

// Inline Drawing 类型中的实体
type Inline struct {
	Picture *InlinePicture `json:"picture,omitempty"` // 图片内容
	Addon   *InlineAddon   `json:"addon,omitempty"`   // 插件信息
}

// InlinePicture 内联图片
type InlinePicture struct {
	URI          string           `json:"uri,omitempty"`           // 图片URI
	RelativeRect *RelativeRect    `json:"relative_rect,omitempty"` // 裁剪范围
	Shape        *ShapeProperties `json:"shape,omitempty"`         // 形状属性
}

// RelativeRect 相对矩形
type RelativeRect struct {
	Left   uint32 `json:"left"`   // 距左侧的距离
	Top    uint32 `json:"top"`    // 距顶部的距离
	Right  uint32 `json:"right"`  // 距右侧的距离
	Bottom uint32 `json:"bottom"` // 距底部的距离
}

// ShapeProperties 形状属性
type ShapeProperties struct {
	Transform *Transform2D `json:"transform,omitempty"` // 图片变换
}

// Transform2D 2D 变换
type Transform2D struct {
	Extent   *PositiveSize2D `json:"extent,omitempty"` // 边框
	Rotation int32           `json:"rotation"`         // 旋转角度
}

// PositiveSize2D 正数尺寸
type PositiveSize2D struct {
	CX int64 `json:"cx"` // 图片宽，单位是像素（px）
	CY int64 `json:"cy"` // 图片高，单位是像素（px）
}

// InlineAddon 内联插件
type InlineAddon struct {
	AddonID     string `json:"addon_id,omitempty"`     // 插件 ID
	AddonSource string `json:"addon_source,omitempty"` // 插件来源
}

// Anchor Drawing 类型中浮动的实体
type Anchor struct {
	Picture *AnchorPicture `json:"picture,omitempty"` // 图片内容
}

// AnchorPicture 锚点图片
type AnchorPicture struct {
	URI                string              `json:"uri,omitempty"`                 // 图片URI
	RelativeRect       *RelativeRect       `json:"relative_rect,omitempty"`       // 裁剪范围
	Shape              *ShapeProperties    `json:"shape,omitempty"`               // 形状属性
	PositionHorizontal *PositionHorizontal `json:"position_horizontal,omitempty"` // 水平位置
	PositionVertical   *PositionVertical   `json:"position_vertical,omitempty"`   // 竖直位置
	WrapNone           bool                `json:"wrap_none,omitempty"`           // 非文字包围
	WrapSquare         *WrapSquare         `json:"wrap_square,omitempty"`         // 四周型环绕
	WrapTopAndBottom   bool                `json:"wrap_top_and_bottom,omitempty"` // 上下型环绕
	BehindDoc          bool                `json:"behind_doc,omitempty"`          // 衬于文字下方
	AllowOverlap       bool                `json:"allow_overlap,omitempty"`       // 允许重叠
}

// PositionHorizontal 水平位置
type PositionHorizontal struct {
	PosOffset    int32  `json:"pos_offset"`    // 位置偏移
	RelativeFrom string `json:"relative_from"` // 相对位置类型
}

// PositionVertical 竖直位置
type PositionVertical struct {
	PosOffset    int32  `json:"pos_offset"`    // 位置偏移
	RelativeFrom string `json:"relative_from"` // 相对位置类型
}

// WrapSquare 四周型环绕
type WrapSquare struct {
	WrapText string `json:"wrap_text"` // 四周环绕文字类型
}

// ==================== 编辑文档内容 ====================

// BatchUpdateDocumentRequest 批量编辑文档内容请求
type BatchUpdateDocumentRequest struct {
	DocID    string          `json:"docid"`             // 文档的docid
	Version  uint32          `json:"version,omitempty"` // 操作的文档版本
	Requests []UpdateRequest `json:"requests"`          // 更新操作列表
}

// UpdateRequest 更新文档的操作
type UpdateRequest struct {
	ReplaceText        *ReplaceText        `json:"replace_text,omitempty"`         // 替换指定位置文本内容
	InsertText         *InsertText         `json:"insert_text,omitempty"`          // 在指定位置插入文本内容
	DeleteContent      *DeleteContent      `json:"delete_content,omitempty"`       // 删除指定位置内容
	InsertImage        *InsertImage        `json:"insert_image,omitempty"`         // 在指定位置插入图片
	InsertPageBreak    *InsertPageBreak    `json:"insert_page_break,omitempty"`    // 在指定位置插入分页符
	InsertTable        *InsertTable        `json:"insert_table,omitempty"`         // 在指定位置插入表格
	InsertParagraph    *InsertParagraph    `json:"insert_paragraph,omitempty"`     // 在指定位置插入段落
	UpdateTextProperty *UpdateTextProperty `json:"update_text_property,omitempty"` // 更新指定位置文本属性
}

// Range 表示从start_index开始的一段范围
type Range struct {
	StartIndex uint32 `json:"start_index"` // 起始位置，从0开始
	Length     uint32 `json:"length"`      // 长度
}

// Location 标准文档中的一个位置
type Location struct {
	Index uint32 `json:"index"` // 位置
}

// ReplaceText 替换文本
type ReplaceText struct {
	Text   string  `json:"text"`   // 要替换的文本
	Ranges []Range `json:"ranges"` // 要替换的文档范围
}

// InsertText 插入文本
type InsertText struct {
	Text     string   `json:"text"`     // 要插入的文本
	Location Location `json:"location"` // 插入的位置
}

// DeleteContent 删除指定位置内容
type DeleteContent struct {
	Range Range `json:"range"` // 要删除的范围
}

// InsertImage 插入图片
type InsertImage struct {
	ImageID  string   `json:"image_id"`         // 图片url，通过上传图片接口获得
	Location Location `json:"location"`         // 插入的位置
	Width    uint32   `json:"width,omitempty"`  // 图片的宽，单位是像素（px）
	Height   uint32   `json:"height,omitempty"` // 图片的高， 单位是像素（px）
}

// InsertPageBreak 插入分页符
type InsertPageBreak struct {
	Location Location `json:"location"` // 插入的位置
}

// InsertTable 插入表格
type InsertTable struct {
	Rows     uint32   `json:"rows"`     // 表格行数
	Cols     uint32   `json:"cols"`     // 表格列数
	Location Location `json:"location"` // 插入的位置
}

// InsertParagraph 插入段落
type InsertParagraph struct {
	Location Location `json:"location"` // 插入的位置
}

// TextProperty 文本属性
type TextProperty struct {
	Bold            bool   `json:"bold,omitempty"`             // 是否加粗
	Color           string `json:"color,omitempty"`            // 文字颜色，十六进制RRGGBB格式
	BackgroundColor string `json:"background_color,omitempty"` // 文字的背景颜色，十六进制RRGGBB 格式
}

// UpdateTextProperty 更新指定范围的文本属性
type UpdateTextProperty struct {
	TextProperty TextProperty `json:"text_property"` // 文本属性
	Ranges       []Range      `json:"ranges"`        // 更新文本属性的范围
}
