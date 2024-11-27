package model

type MessageSender interface {
	SendMessage(message Message) error
	BatchSendMessage(messages []Message) error
}
