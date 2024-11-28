package miaotixing

import (
	"fmt"
	"github.com/yulai-123/message_send_tool/model"
	"net/http"
	"time"
)

/*
	喵提醒是一个允许给自己发送语音通知的工具。
	通过关注喵提醒公众号、绑定手机号、创建消息模版，会提供一个回调 url，调用回调 url，喵提醒会给你发送语音通知。
	需要注意调用频率，语音通知是有频率限制的
	喵提醒可以配置多种通知方式，不过这里只用它配置了语音通知，后续也可以用于配置短信、邮箱等通知
	https://www.showdoc.com.cn/miaotixing/9175237605891603
*/

type MessageHandler struct {
	triggerID string
	client    *http.Client
}

func NewMessageHandler(triggerID string) (*MessageHandler, error) {
	return &MessageHandler{
		triggerID: triggerID,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}, nil
}

func (m *MessageHandler) SendMessage(message model.Message) error {
	return m.sendMessage(message)
}

func (m *MessageHandler) BatchSendMessage(messages []model.Message) error {
	return fmt.Errorf("not implemented")
}
