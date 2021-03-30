package controller

import (
	"goChatDemo/internal/business/service"
	"goChatDemo/internal/manager"
)

const (
	SEND_MESSGAE_ACTION = "sendMessage"
)

type MessageHandler func(connInfo *manager.UserInfo, data []byte) []byte

var MessageHandlerMap = make(map[string]MessageHandler, 1)

func init() {
	MessageHandlerMap[SEND_MESSGAE_ACTION] = SendMessage
}

// 消息结构替
type ImAction struct {
	//事件
	Action string `json:"action"`
}

var SendMessage = func(connInfo *manager.UserInfo, data []byte) []byte {
	userId := connInfo.UserId
	var messageService = service.MessageService{}
	return messageService.SendMessage(userId, data)
}
