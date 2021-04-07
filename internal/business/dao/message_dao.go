package dao

import (
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"time"
)

type messageDao struct {
}

var MessageDao = new(messageDao)

func (md *messageDao) AddMessage(msg *model.Message) (int64, error) {
	err := db.DB.Create(msg).Error
	if err != nil {
		logger.Logger.Error(err)
		return 0, err
	}
	return msg.Id, nil
}

func (md *messageDao) GetMsgByClientId(messageClientId string) *model.Message {
	var msg = &model.Message{}
	db.DB.Where("client_id = ?", messageClientId).Find(msg)
	return msg
}

func (md *messageDao) UpdateMessageSendResult(messageServerId int64) {
	db.DB.Where("id = ? ", messageServerId).Updates(model.Message{SendTime: time.Now(), IsSend: 1})
}
