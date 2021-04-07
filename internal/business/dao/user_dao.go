package dao

import (
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"strconv"
)

type userDao struct {
}

var UserDao = new(userDao)

// 添加用户
func (ud userDao) Add(user *model.User) (int64, error) {
	err := db.DB.Create(user).Error
	if err != nil {
		logger.Logger.Error(err)
		return 0, err
	}
	return user.Id, nil
}

//根据Id删除
func (ud userDao) Del(userId int64) {
	user := model.User{
		Id: userId,
	}
	err := db.DB.Delete(&user).Error
	if err != nil {
		logger.Logger.Error(err)
	}
}

// 根据条件删除
func (ud userDao) DelByCondition(user *model.User) {
	whereCondition := " 1 = 1 "
	if user.IsDeleted != 0 {
		whereCondition += " and is_deleted = " + strconv.Itoa(int(user.IsDeleted))
	}
	if len(user.Mobile) != 0 {
		whereCondition += " and mobile = " + user.Mobile
	}
	if len(user.NickName) != 0 {
		whereCondition += " and nick_name = " + user.NickName
	}
	if len(user.UserName) != 0 {
		whereCondition += " and is_delete = " + user.UserName
	}
	err := db.DB.Where(whereCondition).Delete(&user).Error
	if err != nil {
		logger.Logger.Error(err)
	}
}

// 根据手机号查询
func (ud userDao) GetUserByMobile(mobile string) *[]model.User {
	var users []model.User
	db.DB.Where("mobile = ?", mobile).Find(&users)
	return &users
}
