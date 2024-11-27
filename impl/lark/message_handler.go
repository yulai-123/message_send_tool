package lark

import (
	"github.com/yulai-123/message_send_tool/model"
	"net/http"
	"time"
)

type MessageHandler struct {
	appID                  string
	appSecret              string
	tenantAccessToken      string
	tenantAccessExpireTime time.Time
	client                 *http.Client
}

func NewMessageHandler(appID, appSecret string) (*MessageHandler, error) {
	m := &MessageHandler{
		appID:                  appID,
		appSecret:              appSecret,
		tenantAccessExpireTime: time.Now().Add(-1 * time.Hour),
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	_, err := m.getTenantAccessToken()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *MessageHandler) SendMessage(message model.Message) error {
	return m.sendMessage(message)
}

func (m *MessageHandler) BatchSendMessage(messages []model.Message) error {
	for _, message := range messages {
		err := m.sendMessage(message)
		if err != nil {
			return err
		}
	}
	return nil
}
