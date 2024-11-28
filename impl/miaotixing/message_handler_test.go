package miaotixing

import (
	"github.com/stretchr/testify/assert"
	"github.com/yulai-123/message_send_tool/model"
	"os"
	"testing"
)

func TestSendMessage(t *testing.T) {
	triggerID := os.Getenv("MIAOTIXING_TRIGGER_ID")
	m, err := NewMessageHandler(triggerID)
	assert.Nil(t, err)

	message := model.Message{
		Title:   "这是一个单測",
		Content: "这是单測内容",
	}

	err = m.SendMessage(message)
	assert.Nil(t, err)
}
