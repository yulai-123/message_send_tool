package model

type ReceiverType string

var (
	OpenID ReceiverType = "open_id"
	ChatID ReceiverType = "chat_id"
)

type Message struct {
	Title         string       // 消息标题
	Content       string       // 消息内容
	ReceiveID     string       // 接收用户 open_id
	ReceiveIDType ReceiverType // 接收用户类型,默认 open_id，可选 chat_id
	ImageList     []FileInfo   // 图片列表
	URL           string       // 卡片消息的跳转链接
}

type FileInfo struct {
	FileName string
	FileData []byte
}
