package message

import "github.com/shuaidd/wecom-core/types/common"

// MessageType 消息类型
type MessageType string

const (
	// MessageTypeText 文本消息
	MessageTypeText MessageType = "text"
	// MessageTypeImage 图片消息
	MessageTypeImage MessageType = "image"
	// MessageTypeVoice 语音消息
	MessageTypeVoice MessageType = "voice"
	// MessageTypeVideo 视频消息
	MessageTypeVideo MessageType = "video"
	// MessageTypeFile 文件消息
	MessageTypeFile MessageType = "file"
	// MessageTypeTextCard 文本卡片消息
	MessageTypeTextCard MessageType = "textcard"
	// MessageTypeNews 图文消息
	MessageTypeNews MessageType = "news"
	// MessageTypeMPNews 图文消息（mpnews）
	MessageTypeMPNews MessageType = "mpnews"
	// MessageTypeMarkdown Markdown消息
	MessageTypeMarkdown MessageType = "markdown"
	// MessageTypeMiniProgramNotice 小程序通知消息
	MessageTypeMiniProgramNotice MessageType = "miniprogram_notice"
	// MessageTypeTemplateCard 模板卡片消息
	MessageTypeTemplateCard MessageType = "template_card"
)

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	// ToUser 成员ID列表（多个接收者用'|'分隔，最多支持1000个）
	ToUser string `json:"touser,omitempty"`
	// ToParty 部门ID列表，多个接收者用'|'分隔，最多支持100个
	ToParty string `json:"toparty,omitempty"`
	// ToTag 标签ID列表，多个接收者用'|'分隔，最多支持100个
	ToTag string `json:"totag,omitempty"`
	// MsgType 消息类型
	MsgType MessageType `json:"msgtype"`
	// AgentID 企业应用的id
	AgentID int `json:"agentid"`
	// Safe 表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0
	Safe *int `json:"safe,omitempty"`
	// EnableIDTrans 表示是否开启id转译，0表示否，1表示是，默认0
	EnableIDTrans *int `json:"enable_id_trans,omitempty"`
	// EnableDuplicateCheck 表示是否开启重复消息检查，0表示否，1表示是，默认0
	EnableDuplicateCheck *int `json:"enable_duplicate_check,omitempty"`
	// DuplicateCheckInterval 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	DuplicateCheckInterval *int `json:"duplicate_check_interval,omitempty"`

	// 各种消息类型的具体内容
	Text                *TextMessage                `json:"text,omitempty"`
	Image               *MediaMessage               `json:"image,omitempty"`
	Voice               *MediaMessage               `json:"voice,omitempty"`
	Video               *VideoMessage               `json:"video,omitempty"`
	File                *MediaMessage               `json:"file,omitempty"`
	TextCard            *TextCardMessage            `json:"textcard,omitempty"`
	News                *NewsMessage                `json:"news,omitempty"`
	MPNews              *MPNewsMessage              `json:"mpnews,omitempty"`
	Markdown            *MarkdownMessage            `json:"markdown,omitempty"`
	MiniProgramNotice   *MiniProgramNoticeMessage   `json:"miniprogram_notice,omitempty"`
	TemplateCard        *TemplateCardMessage        `json:"template_card,omitempty"`
}

// SendMessageResponse 发送消息响应
type SendMessageResponse struct {
	common.Response
	// InvalidUser 不合法的userid
	InvalidUser string `json:"invaliduser,omitempty"`
	// InvalidParty 不合法的partyid
	InvalidParty string `json:"invalidparty,omitempty"`
	// InvalidTag 不合法的标签id
	InvalidTag string `json:"invalidtag,omitempty"`
	// UnlicensedUser 没有基础接口许可(包含已过期)的userid
	UnlicensedUser string `json:"unlicenseduser,omitempty"`
	// MsgID 消息id，用于撤回应用消息
	MsgID string `json:"msgid,omitempty"`
	// ResponseCode 应用可使用response_code调用更新模版卡片消息接口
	ResponseCode string `json:"response_code,omitempty"`
}

// TextMessage 文本消息
type TextMessage struct {
	// Content 消息内容，最长不超过2048个字节
	Content string `json:"content"`
}

// MediaMessage 媒体消息（图片、语音、文件）
type MediaMessage struct {
	// MediaID 媒体文件id
	MediaID string `json:"media_id"`
}

// VideoMessage 视频消息
type VideoMessage struct {
	// MediaID 视频媒体文件id
	MediaID string `json:"media_id"`
	// Title 视频消息的标题
	Title string `json:"title,omitempty"`
	// Description 视频消息的描述
	Description string `json:"description,omitempty"`
}

// TextCardMessage 文本卡片消息
type TextCardMessage struct {
	// Title 标题，不超过128个字符
	Title string `json:"title"`
	// Description 描述，不超过512个字符
	Description string `json:"description"`
	// URL 点击后跳转的链接
	URL string `json:"url"`
	// BtnTxt 按钮文字，默认为"详情"
	BtnTxt string `json:"btntxt,omitempty"`
}

// NewsMessage 图文消息
type NewsMessage struct {
	// Articles 图文消息，一个图文消息支持1到8条图文
	Articles []NewsArticle `json:"articles"`
}

// NewsArticle 图文消息文章
type NewsArticle struct {
	// Title 标题
	Title string `json:"title"`
	// Description 描述
	Description string `json:"description,omitempty"`
	// URL 点击后跳转的链接
	URL string `json:"url,omitempty"`
	// PicURL 图文消息的图片链接
	PicURL string `json:"picurl,omitempty"`
	// AppID 小程序appid
	AppID string `json:"appid,omitempty"`
	// PagePath 点击消息卡片后的小程序页面
	PagePath string `json:"pagepath,omitempty"`
}

// MPNewsMessage 图文消息（mpnews）
type MPNewsMessage struct {
	// Articles 图文消息，一个图文消息支持1到8条图文
	Articles []MPNewsArticle `json:"articles"`
}

// MPNewsArticle 图文消息文章（mpnews）
type MPNewsArticle struct {
	// Title 标题
	Title string `json:"title"`
	// ThumbMediaID 图文消息缩略图的media_id
	ThumbMediaID string `json:"thumb_media_id"`
	// Author 图文消息的作者
	Author string `json:"author,omitempty"`
	// ContentSourceURL 图文消息点击"阅读原文"之后的页面链接
	ContentSourceURL string `json:"content_source_url,omitempty"`
	// Content 图文消息的内容，支持html标签
	Content string `json:"content"`
	// Digest 图文消息的描述
	Digest string `json:"digest,omitempty"`
}

// MarkdownMessage Markdown消息
type MarkdownMessage struct {
	// Content markdown内容
	Content string `json:"content"`
}

// MiniProgramNoticeMessage 小程序通知消息
type MiniProgramNoticeMessage struct {
	// AppID 小程序appid
	AppID string `json:"appid"`
	// Page 点击消息卡片后的小程序页面
	Page string `json:"page,omitempty"`
	// Title 消息标题
	Title string `json:"title"`
	// Description 消息描述
	Description string `json:"description,omitempty"`
	// EmphasisFirstItem 是否放大第一个content_item
	EmphasisFirstItem bool `json:"emphasis_first_item,omitempty"`
	// ContentItem 消息内容键值对
	ContentItem []ContentItem `json:"content_item,omitempty"`
}

// ContentItem 消息内容键值对
type ContentItem struct {
	// Key 长度10个汉字以内
	Key string `json:"key,omitempty"`
	// Value 长度30个汉字以内
	Value string `json:"value,omitempty"`
}

// RecallMessageRequest 撤回消息请求
type RecallMessageRequest struct {
	// MsgID 消息ID
	MsgID string `json:"msgid"`
}

// TemplateCardType 模板卡片类型
type TemplateCardType string

const (
	// TemplateCardTypeTextNotice 文本通知型
	TemplateCardTypeTextNotice TemplateCardType = "text_notice"
	// TemplateCardTypeNewsNotice 图文展示型
	TemplateCardTypeNewsNotice TemplateCardType = "news_notice"
	// TemplateCardTypeButtonInteraction 按钮交互型
	TemplateCardTypeButtonInteraction TemplateCardType = "button_interaction"
	// TemplateCardTypeVoteInteraction 投票选择型
	TemplateCardTypeVoteInteraction TemplateCardType = "vote_interaction"
	// TemplateCardTypeMultipleInteraction 多项选择型
	TemplateCardTypeMultipleInteraction TemplateCardType = "multiple_interaction"
)

// TemplateCardMessage 模板卡片消息
type TemplateCardMessage struct {
	// CardType 模板卡片类型
	CardType TemplateCardType `json:"card_type"`
	// Source 卡片来源样式信息
	Source *CardSource `json:"source,omitempty"`
	// ActionMenu 卡片右上角更多操作按钮
	ActionMenu *CardActionMenu `json:"action_menu,omitempty"`
	// TaskID 任务id
	TaskID string `json:"task_id,omitempty"`
	// MainTitle 一级标题
	MainTitle *CardMainTitle `json:"main_title,omitempty"`
	// QuoteArea 引用文献样式
	QuoteArea *CardQuoteArea `json:"quote_area,omitempty"`
	// EmphasisContent 关键数据样式（仅文本通知型支持）
	EmphasisContent *CardEmphasisContent `json:"emphasis_content,omitempty"`
	// SubTitleText 二级普通文本（仅文本通知型和按钮交互型支持）
	SubTitleText string `json:"sub_title_text,omitempty"`
	// HorizontalContentList 二级标题+文本列表
	HorizontalContentList []CardHorizontalContent `json:"horizontal_content_list,omitempty"`
	// JumpList 跳转指引样式的列表（仅文本通知型和图文展示型支持）
	JumpList []CardJump `json:"jump_list,omitempty"`
	// CardAction 整体卡片的点击跳转事件
	CardAction *CardAction `json:"card_action,omitempty"`
	// ImageTextArea 左图右文样式（仅图文展示型支持）
	ImageTextArea *CardImageTextArea `json:"image_text_area,omitempty"`
	// CardImage 图片样式（仅图文展示型支持）
	CardImage *CardImage `json:"card_image,omitempty"`
	// VerticalContentList 卡片二级垂直内容（仅图文展示型支持）
	VerticalContentList []CardVerticalContent `json:"vertical_content_list,omitempty"`
	// ButtonSelection 下拉式的选择器（仅按钮交互型支持）
	ButtonSelection *CardButtonSelection `json:"button_selection,omitempty"`
	// ButtonList 按钮列表（仅按钮交互型支持）
	ButtonList []CardButton `json:"button_list,omitempty"`
	// Checkbox 选择题样式（仅投票选择型支持）
	Checkbox *CardCheckbox `json:"checkbox,omitempty"`
	// SubmitButton 提交按钮样式（仅投票选择型和多项选择型支持）
	SubmitButton *CardSubmitButton `json:"submit_button,omitempty"`
	// SelectList 下拉式的选择器列表（仅多项选择型支持）
	SelectList []CardSelect `json:"select_list,omitempty"`
}

// CardSource 卡片来源样式信息
type CardSource struct {
	// IconURL 来源图片的url
	IconURL string `json:"icon_url,omitempty"`
	// Desc 来源图片的描述
	Desc string `json:"desc,omitempty"`
	// DescColor 来源文字的颜色，0(默认)灰色，1黑色，2红色，3绿色
	DescColor int `json:"desc_color,omitempty"`
}

// CardActionMenu 卡片右上角更多操作按钮
type CardActionMenu struct {
	// Desc 更多操作界面的描述
	Desc string `json:"desc,omitempty"`
	// ActionList 操作列表
	ActionList []CardActionMenuItem `json:"action_list"`
}

// CardActionMenuItem 操作列表项
type CardActionMenuItem struct {
	// Text 操作的描述文案
	Text string `json:"text"`
	// Key 操作key值
	Key string `json:"key"`
}

// CardMainTitle 一级标题
type CardMainTitle struct {
	// Title 一级标题
	Title string `json:"title,omitempty"`
	// Desc 标题辅助信息
	Desc string `json:"desc,omitempty"`
}

// CardQuoteArea 引用文献样式
type CardQuoteArea struct {
	// Type 引用文献样式区域点击事件，0或不填代表没有点击事件，1代表跳转url，2代表跳转小程序
	Type int `json:"type,omitempty"`
	// URL 点击跳转的url
	URL string `json:"url,omitempty"`
	// AppID 点击跳转的小程序的appid
	AppID string `json:"appid,omitempty"`
	// PagePath 点击跳转的小程序的pagepath
	PagePath string `json:"pagepath,omitempty"`
	// Title 引用文献样式的标题
	Title string `json:"title,omitempty"`
	// QuoteText 引用文献样式的引用文案
	QuoteText string `json:"quote_text,omitempty"`
}

// CardEmphasisContent 关键数据样式
type CardEmphasisContent struct {
	// Title 关键数据样式的数据内容
	Title string `json:"title,omitempty"`
	// Desc 关键数据样式的数据描述内容
	Desc string `json:"desc,omitempty"`
}

// CardHorizontalContent 二级标题+文本
type CardHorizontalContent struct {
	// Type 链接类型，0或不填代表不是链接，1代表跳转url，2代表下载附件，3代表点击跳转成员详情
	Type int `json:"type,omitempty"`
	// KeyName 二级标题
	KeyName string `json:"keyname"`
	// Value 二级文本
	Value string `json:"value,omitempty"`
	// URL 链接跳转的url
	URL string `json:"url,omitempty"`
	// MediaID 附件的media_id
	MediaID string `json:"media_id,omitempty"`
	// UserID 成员详情的userid
	UserID string `json:"userid,omitempty"`
}

// CardJump 跳转指引样式
type CardJump struct {
	// Type 跳转链接类型，0或不填代表不是链接，1代表跳转url，2代表跳转小程序
	Type int `json:"type,omitempty"`
	// Title 跳转链接样式的文案内容
	Title string `json:"title"`
	// URL 跳转链接的url
	URL string `json:"url,omitempty"`
	// AppID 跳转链接的小程序的appid
	AppID string `json:"appid,omitempty"`
	// PagePath 跳转链接的小程序的pagepath
	PagePath string `json:"pagepath,omitempty"`
}

// CardAction 整体卡片的点击跳转事件
type CardAction struct {
	// Type 跳转事件类型，0或不填代表不是链接，1代表跳转url，2代表打开小程序
	Type int `json:"type,omitempty"`
	// URL 跳转事件的url
	URL string `json:"url,omitempty"`
	// AppID 跳转事件的小程序的appid
	AppID string `json:"appid,omitempty"`
	// PagePath 跳转事件的小程序的pagepath
	PagePath string `json:"pagepath,omitempty"`
}

// CardImageTextArea 左图右文样式
type CardImageTextArea struct {
	// Type 左图右文样式区域点击事件，0或不填代表没有点击事件，1代表跳转url，2代表跳转小程序
	Type int `json:"type,omitempty"`
	// URL 点击跳转的url
	URL string `json:"url,omitempty"`
	// AppID 点击跳转的小程序的appid
	AppID string `json:"appid,omitempty"`
	// PagePath 点击跳转的小程序的pagepath
	PagePath string `json:"pagepath,omitempty"`
	// Title 左图右文样式的标题
	Title string `json:"title,omitempty"`
	// Desc 左图右文样式的描述
	Desc string `json:"desc,omitempty"`
	// ImageURL 左图右文样式的图片url
	ImageURL string `json:"image_url"`
}

// CardImage 图片样式
type CardImage struct {
	// URL 图片的url
	URL string `json:"url"`
	// AspectRatio 图片的宽高比
	AspectRatio float64 `json:"aspect_ratio,omitempty"`
}

// CardVerticalContent 卡片二级垂直内容
type CardVerticalContent struct {
	// Title 卡片二级标题
	Title string `json:"title"`
	// Desc 二级普通文本
	Desc string `json:"desc,omitempty"`
}

// CardButtonSelection 下拉式的选择器
type CardButtonSelection struct {
	// QuestionKey 下拉式的选择器的key
	QuestionKey string `json:"question_key"`
	// Title 下拉式的选择器左边的标题
	Title string `json:"title,omitempty"`
	// OptionList 选项列表
	OptionList []CardSelectOption `json:"option_list"`
	// SelectedID 默认选定的id
	SelectedID string `json:"selected_id,omitempty"`
}

// CardSelectOption 选择器选项
type CardSelectOption struct {
	// ID 选项的id
	ID string `json:"id"`
	// Text 选项的文案
	Text string `json:"text"`
}

// CardButton 按钮
type CardButton struct {
	// Type 按钮点击事件类型，0或不填代表回调点击事件，1代表跳转url
	Type int `json:"type,omitempty"`
	// Text 按钮文案
	Text string `json:"text"`
	// Style 按钮样式，目前可填1~4
	Style int `json:"style,omitempty"`
	// Key 按钮key值
	Key string `json:"key,omitempty"`
	// URL 跳转事件的url
	URL string `json:"url,omitempty"`
}

// CardCheckbox 选择题样式
type CardCheckbox struct {
	// QuestionKey 选择题key值
	QuestionKey string `json:"question_key"`
	// OptionList 选项list
	OptionList []CardCheckboxOption `json:"option_list"`
	// Mode 选择题模式，单选：0，多选：1
	Mode int `json:"mode,omitempty"`
}

// CardCheckboxOption 选择题选项
type CardCheckboxOption struct {
	// ID 选项id
	ID string `json:"id"`
	// Text 选项文案描述
	Text string `json:"text"`
	// IsChecked 该选项是否要默认选中
	IsChecked bool `json:"is_checked"`
}

// CardSubmitButton 提交按钮样式
type CardSubmitButton struct {
	// Text 按钮文案
	Text string `json:"text"`
	// Key 提交按钮的key
	Key string `json:"key"`
}

// CardSelect 下拉式的选择器
type CardSelect struct {
	// QuestionKey 下拉式的选择器题目的key
	QuestionKey string `json:"question_key"`
	// Title 下拉式的选择器上面的title
	Title string `json:"title,omitempty"`
	// OptionList 选项列表
	OptionList []CardSelectOption `json:"option_list"`
	// SelectedID 默认选定的id
	SelectedID string `json:"selected_id,omitempty"`
}

// SendSchoolNoticeRequest 发送学校通知请求
type SendSchoolNoticeRequest struct {
	// RecvScope 指定发送对象，0表示发送给家长，1表示发送给学生，2表示发送给家长和学生，默认为0
	RecvScope *int `json:"recv_scope,omitempty"`
	// ToParentUserID 家校通讯录家长列表（最多支持1000个）
	ToParentUserID []string `json:"to_parent_userid,omitempty"`
	// ToStudentUserID 家校通讯录学生列表（最多支持1000个）
	ToStudentUserID []string `json:"to_student_userid,omitempty"`
	// ToParty 家校通讯录部门列表（最多支持100个）
	ToParty []string `json:"to_party,omitempty"`
	// ToAll 1表示发送给所有人，0表示不发送给所有人，默认为0
	ToAll *int `json:"toall,omitempty"`
	// MsgType 消息类型
	MsgType MessageType `json:"msgtype"`
	// AgentID 企业应用的id
	AgentID int `json:"agentid"`
	// EnableIDTrans 表示是否开启id转译，0表示否，1表示是，默认0
	EnableIDTrans *int `json:"enable_id_trans,omitempty"`
	// EnableDuplicateCheck 表示是否开启重复消息检查，0表示否，1表示是，默认0
	EnableDuplicateCheck *int `json:"enable_duplicate_check,omitempty"`
	// DuplicateCheckInterval 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	DuplicateCheckInterval *int `json:"duplicate_check_interval,omitempty"`

	// 各种消息类型的具体内容
	Text     *TextMessage     `json:"text,omitempty"`
	Image    *MediaMessage    `json:"image,omitempty"`
	Voice    *MediaMessage    `json:"voice,omitempty"`
	Video    *VideoMessage    `json:"video,omitempty"`
	File     *MediaMessage    `json:"file,omitempty"`
	News     *NewsMessage     `json:"news,omitempty"`
	MPNews   *SchoolMPNews    `json:"mpnews,omitempty"`
	MiniProgram *SchoolMiniProgram `json:"miniprogram,omitempty"`
}

// SchoolMPNews 学校通知的图文消息（mpnews）
type SchoolMPNews struct {
	// Articles 图文消息，一个图文消息支持1到8条图文
	Articles []SchoolMPNewsArticle `json:"articles"`
}

// SchoolMPNewsArticle 学校通知的图文消息文章（mpnews）
type SchoolMPNewsArticle struct {
	// Title 标题
	Title string `json:"title"`
	// ThumbMediaID 图文消息缩略图的media_id
	ThumbMediaID string `json:"thumb_media_id"`
	// Author 图文消息的作者
	Author string `json:"author,omitempty"`
	// ContentSourceURL 图文消息点击"阅读原文"之后的页面链接
	ContentSourceURL string `json:"content_source_url,omitempty"`
	// Content 图文消息的内容，支持html标签，不超过666K个字节
	Content string `json:"content"`
	// Digest 图文消息的描述
	Digest string `json:"digest,omitempty"`
}

// SchoolMiniProgram 学校通知的小程序消息
type SchoolMiniProgram struct {
	// AppID 小程序appid
	AppID string `json:"appid"`
	// Title 小程序消息标题
	Title string `json:"title,omitempty"`
	// ThumbMediaID 小程序消息封面的mediaid
	ThumbMediaID string `json:"thumb_media_id"`
	// PagePath 点击消息卡片后进入的小程序页面路径
	PagePath string `json:"pagepath"`
}

// SendSchoolNoticeResponse 发送学校通知响应
type SendSchoolNoticeResponse struct {
	common.Response
	// InvalidParentUserID 不合法的家长userid
	InvalidParentUserID []string `json:"invalid_parent_userid,omitempty"`
	// InvalidStudentUserID 不合法的学生userid
	InvalidStudentUserID []string `json:"invalid_student_userid,omitempty"`
	// InvalidParty 不合法的部门id
	InvalidParty []string `json:"invalid_party,omitempty"`
}

// UpdateTemplateCardRequest 更新模板卡片消息请求
type UpdateTemplateCardRequest struct {
	// UserIDs 企业的成员ID列表（最多支持1000个）
	UserIDs []string `json:"userids,omitempty"`
	// PartyIDs 企业的部门ID列表（最多支持100个）
	PartyIDs []int `json:"partyids,omitempty"`
	// TagIDs 企业的标签ID列表（最多支持100个）
	TagIDs []int `json:"tagids,omitempty"`
	// AtAll 更新整个任务接收人员
	AtAll *int `json:"atall,omitempty"`
	// AgentID 应用的agentid
	AgentID int `json:"agentid"`
	// ResponseCode 更新卡片所需要消费的code
	ResponseCode string `json:"response_code"`
	// EnableIDTrans 表示是否开启id转译，0表示否，1表示是，默认0
	EnableIDTrans *int `json:"enable_id_trans,omitempty"`
	// Button 更新按钮为不可点击状态
	Button *UpdateButton `json:"button,omitempty"`
	// TemplateCard 更新为新的卡片
	TemplateCard *TemplateCardMessage `json:"template_card,omitempty"`
}

// UpdateButton 更新按钮
type UpdateButton struct {
	// ReplaceName 需要更新的按钮的文案
	ReplaceName string `json:"replace_name"`
}

// UpdateTemplateCardResponse 更新模板卡片消息响应
type UpdateTemplateCardResponse struct {
	common.Response
	// InvalidUser 不合法的userid
	InvalidUser []string `json:"invaliduser,omitempty"`
}

// AppChatSendRequest 应用推送消息请求
type AppChatSendRequest struct {
	// ChatID 群聊id
	ChatID string `json:"chatid"`
	// MsgType 消息类型
	MsgType MessageType `json:"msgtype"`
	// Safe 表示是否是保密消息，0表示否，1表示是，默认0
	Safe *int `json:"safe,omitempty"`

	// 各种消息类型的具体内容
	Text     *TextMessage     `json:"text,omitempty"`
	Image    *MediaMessage    `json:"image,omitempty"`
	Voice    *MediaMessage    `json:"voice,omitempty"`
	Video    *VideoMessage    `json:"video,omitempty"`
	File     *MediaMessage    `json:"file,omitempty"`
	TextCard *TextCardMessage `json:"textcard,omitempty"`
	News     *NewsMessage     `json:"news,omitempty"`
	MPNews   *MPNewsMessage   `json:"mpnews,omitempty"`
	Markdown *MarkdownMessage `json:"markdown,omitempty"`
}

// GetAppChatRequest 获取群聊会话请求
type GetAppChatRequest struct {
	// ChatID 群聊id
	ChatID string `json:"chatid"`
}

// GetAppChatResponse 获取群聊会话响应
type GetAppChatResponse struct {
	common.Response
	// ChatInfo 群聊信息
	ChatInfo *AppChatInfo `json:"chat_info,omitempty"`
}

// AppChatInfo 群聊信息
type AppChatInfo struct {
	// ChatID 群聊唯一标志
	ChatID string `json:"chatid"`
	// Name 群聊名
	Name string `json:"name"`
	// Owner 群主id
	Owner string `json:"owner"`
	// UserList 群成员id列表
	UserList []string `json:"userlist"`
	// ChatType 群聊类型，0表示普通群聊，1表示课程群聊
	ChatType int `json:"chat_type"`
}

// UpdateAppChatRequest 修改群聊会话请求
type UpdateAppChatRequest struct {
	// ChatID 群聊id
	ChatID string `json:"chatid"`
	// Name 新的群聊名
	Name string `json:"name,omitempty"`
	// Owner 新群主的id
	Owner string `json:"owner,omitempty"`
	// AddUserList 添加成员的id列表
	AddUserList []string `json:"add_user_list,omitempty"`
	// DelUserList 踢出成员的id列表
	DelUserList []string `json:"del_user_list,omitempty"`
}

// ListSmartsheetGroupChatRequest 获取群聊列表请求
type ListSmartsheetGroupChatRequest struct {
	// DocID 智能表格ID
	DocID string `json:"docid"`
	// Cursor 用于分页查询的游标
	Cursor string `json:"cursor,omitempty"`
	// Limit 每次请求返回的数据上限，默认100，最大200
	Limit int `json:"limit,omitempty"`
}

// ListSmartsheetGroupChatResponse 获取群聊列表响应
type ListSmartsheetGroupChatResponse struct {
	common.Response
	// HasMore 是否还有更多数据
	HasMore bool `json:"has_more"`
	// NextCursor 下一次请求的cursor值
	NextCursor string `json:"next_cursor,omitempty"`
	// ChatIDList 符合条件的群聊chatid列表
	ChatIDList []string `json:"chat_id_list"`
}
