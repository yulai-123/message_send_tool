package model

type Message struct {
	Title     string // 消息标题
	Content   string // 消息内容
	ReceiveID string // 接收用户 open_id
	ImageList []byte // 图片列表
	URL       string // 卡片消息的跳转链接
}
