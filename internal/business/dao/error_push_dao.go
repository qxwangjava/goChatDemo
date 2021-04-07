package dao

import (
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
)

type errorPushDao struct {
}

var ErrorPushDao = new(errorPushDao)

func (epd *errorPushDao) AddErrorPush(push model.ErrorPush) (int64, error) {
	err := db.DB.Create(push).Error
	if err != nil {
		logger.Logger.Error(err)
		return 0, err
	}
	return push.Id, nil
}
