package message_send_tool

import (
	"github.com/yulai-123/message_send_tool/impl/lark"
	"github.com/yulai-123/message_send_tool/model"
)

func NewLarkMessageHandler(appID, appSecret string) (model.MessageSender, error) {
	return lark.NewMessageHandler(appID, appSecret)
}
