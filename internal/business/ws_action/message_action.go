package ws_action

import (
	"goChatDemo/internal/business/service"
)

const (
	SendMessageAction = "sendMessage"
)

var SendMessage = func(userId string, data []byte, ext ...interface{}) []byte {
	var messageService = service.MessageService{}
	return messageService.SendMessage(userId, data)
}
