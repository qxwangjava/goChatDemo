package controller

import (
	"goChatDemo/internal/business/service"
)

const (
	SEND_MESSGAE_ACTION = "sendMessage"
)

type MessageHandler func(userId string, data []byte) []byte

var MessageHandlerMap = make(map[string]MessageHandler, 1)

func init() {
	MessageHandlerMap[SEND_MESSGAE_ACTION] = SendMessage
}

// 消息结构替
type ImAction struct {
	//事件
	Action string `json:"action"`
}

var SendMessage = func(userId string, data []byte) []byte {
	var messageService = service.MessageService{}
	return messageService.SendMessage(userId, data)
}
