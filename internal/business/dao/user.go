package dao

import (
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"strconv"
	"time"
)

type userDao struct {
}

// 用户Demo表
type User struct {
	CreatedAt time.Time `gorm:"created_at"` // 创建时间
	Id        int64     `gorm:"id"`         // 主键
	IsDeleted int32     `gorm:"is_deleted"` // 是否删除 1:是  -1:否
	Mobile    string    `gorm:"mobile"`     // 手机号
	NickName  string    `gorm:"nick_name"`  // 昵称
	UpdatedAt time.Time `gorm:"updated_at"` // 更新时间
	UserName  string    `gorm:"user_name"`  // 用户名

}

var UserDao = new(userDao)

func (ud userDao) Add(user *User) (int64, error) {
	err := db.DB.Create(user).Error
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	return user.Id, nil
}

func (ud userDao) Del(userId int64) {
	user := User{
		Id: userId,
	}
	err := db.DB.Delete(&user).Error
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
}

func (ud userDao) DelByCondition(user *User) {
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
		panic(err)
	}
}
