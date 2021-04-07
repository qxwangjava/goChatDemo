package dao

import (
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
)

type friendDao struct {
}

var FriendDao = new(friendDao)

func (fd friendDao) GetFriend(userId, friendId string) model.Friend {
	var friend model.Friend
	db.DB.Where("user_id = ? and friend_id = ?", userId, friendId).Find(&friend)
	return friend
}
