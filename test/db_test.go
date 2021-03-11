package test

import (
	"goChatDemo/config"
	"goChatDemo/internal/business/dao"
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"testing"
	"time"
)

func TestDB(T *testing.T) {
	db.InitDB(config.DbConfig.DbUrl)
	//新增
	birthDay, _ := time.Parse("2006-01-02 15:04:05", "1993-01-02 00:00:00")
	userId, _ := dao.UserDao.Add(&model.User{
		Birthday:  &birthDay,
		UserName:  "wangqx1",
		NickName:  "王奇轩",
		Mobile:    "17610565751",
		IsDeleted: 0,
	})
	logger.Logger.WithField("userId", userId).Info("保存成功")
	//根据主键删除
	userId = 4
	dao.UserDao.Del(4)
	logger.Logger.WithField("userId", userId).Info("删除成功")
	//根据条件删除
	user := model.User{
		IsDeleted: 1,
	}
	dao.UserDao.DelByCondition(&user)
	logger.Logger.WithField("user", user).Info("删除成功")
}
