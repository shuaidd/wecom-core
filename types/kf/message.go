package kf

import "github.com/shuaidd/wecom-core/types/common"

// SendMsgRequest 发送消息请求
// 支持发送消息类型:文本、图片、语音、视频、文件、图文、小程序、菜单消息、地理位置、获客链接
type SendMsgRequest struct {
	ToUser   string `json:"touser"`    // 指定接收消息的客户UserID
	OpenKfID string `json:"open_kfid"` // 指定发送消息的客服账号ID
	MsgID    string `json:"msgid,omitempty"` // 指定消息ID,不多于32字节,字符串取值范围:[0-9a-zA-Z_-]*
	MsgType  string `json:"msgtype"`   // 消息类型

	// 各种消息类型,根据msgtype选择填充对应的字段
	Text        *TextContent        `json:"text,omitempty"`        // 文本消息
	Image       *MediaContent       `json:"image,omitempty"`       // 图片消息
	Voice       *MediaContent       `json:"voice,omitempty"`       // 语音消息
	Video       *MediaContent       `json:"video,omitempty"`       // 视频消息
	File        *MediaContent       `json:"file,omitempty"`        // 文件消息
	Link        *LinkContent        `json:"link,omitempty"`        // 图文链接消息
	MiniProgram *MiniProgramContent `json:"miniprogram,omitempty"` // 小程序消息
	MsgMenu     *MsgMenuContent     `json:"msgmenu,omitempty"`     // 菜单消息
	Location    *LocationContent    `json:"location,omitempty"`    // 地理位置消息
	CALink      *CALinkContent      `json:"ca_link,omitempty"`     // 获客链接消息
}

// SendMsgResponse 发送消息响应
type SendMsgResponse struct {
	common.Response
	MsgID string `json:"msgid"` // 消息ID
}

// TextContent 文本消息内容
type TextContent struct {
	Content string `json:"content"` // 消息内容,最长不超过2048个字节
}

// MediaContent 媒体消息内容(图片、语音、视频、文件)
type MediaContent struct {
	MediaID string `json:"media_id"` // 媒体文件id,可以调用上传临时素材接口获取
}

// LinkContent 图文链接消息内容
type LinkContent struct {
	Title        string `json:"title"`                   // 标题,不超过128个字节,超过会自动截断
	Desc         string `json:"desc,omitempty"`          // 描述,不超过512个字节,超过会自动截断
	URL          string `json:"url"`                     // 点击后跳转的链接,最长2048字节,请确保包含了协议头(http/https)
	ThumbMediaID string `json:"thumb_media_id"`          // 缩略图的media_id,可以通过素材管理接口获得
}

// MiniProgramContent 小程序消息内容
type MiniProgramContent struct {
	AppID        string `json:"appid"`          // 小程序appid
	Title        string `json:"title,omitempty"` // 小程序消息标题,最多64个字节,超过会自动截断
	ThumbMediaID string `json:"thumb_media_id"` // 小程序消息封面的mediaid,封面图建议尺寸为520*416
	PagePath     string `json:"pagepath"`       // 点击消息卡片后进入的小程序页面路径,注意路径要以.html为后缀
}

// MsgMenuContent 菜单消息内容
type MsgMenuContent struct {
	HeadContent string          `json:"head_content,omitempty"` // 起始文本,不多于1024字节
	List        []MsgMenuItem   `json:"list,omitempty"`         // 菜单项配置,不超过50个,其中click/view/miniprogram的菜单类型加起来不超过10个
	TailContent string          `json:"tail_content,omitempty"` // 结束文本,不多于1024字节
}

// MsgMenuItem 菜单项
type MsgMenuItem struct {
	Type        string                  `json:"type"`                   // 菜单类型:click-回复菜单,view-超链接菜单,miniprogram-小程序菜单,text-文本
	Click       *MsgMenuClickItem       `json:"click,omitempty"`        // type为click的菜单项
	View        *MsgMenuViewItem        `json:"view,omitempty"`         // type为view的菜单项
	MiniProgram *MsgMenuMiniProgramItem `json:"miniprogram,omitempty"`  // type为miniprogram的菜单项
	Text        *MsgMenuTextItem        `json:"text,omitempty"`         // type为text的菜单项
}

// MsgMenuClickItem 回复菜单项
type MsgMenuClickItem struct {
	ID      string `json:"id,omitempty"` // 菜单ID,不少于1字节,不多于128字节
	Content string `json:"content"`      // 菜单显示内容,不少于1字节,不多于128字节
}

// MsgMenuViewItem 超链接菜单项
type MsgMenuViewItem struct {
	URL     string `json:"url"`     // 点击后跳转的链接,不少于1字节,不多于2048字节
	Content string `json:"content"` // 菜单显示内容,不少于1字节,不多于1024字节
}

// MsgMenuMiniProgramItem 小程序菜单项
type MsgMenuMiniProgramItem struct {
	AppID    string `json:"appid"`    // 小程序appid,不少于1字节,不多于32字节
	PagePath string `json:"pagepath"` // 点击后进入的小程序页面,不少于1字节,不多于1024字节
	Content  string `json:"content"`  // 菜单显示内容,不多于1024字节
}

// MsgMenuTextItem 文本菜单项
type MsgMenuTextItem struct {
	Content   string `json:"content"`             // 文本内容,支持\n换行,不少于1字节,不多于256字节
	NoNewline int    `json:"no_newline,omitempty"` // 内容后面是否不换行,0-换行,1-不换行,默认为0
}

// LocationContent 地理位置消息内容
type LocationContent struct {
	Name      string  `json:"name,omitempty"`    // 位置名
	Address   string  `json:"address,omitempty"` // 地址详情说明
	Latitude  float64 `json:"latitude"`          // 纬度,浮点数,范围为90 ~ -90
	Longitude float64 `json:"longitude"`         // 经度,浮点数,范围为180 ~ -180
}

// CALinkContent 获客链接消息内容
type CALinkContent struct {
	LinkURL string `json:"link_url"` // 通过获客助手创建的获客链接
}
