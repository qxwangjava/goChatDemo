package dao

import (
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
)

type groupUserDao struct {
}

var GroupUserDao = new(groupUserDao)

func (gud groupUserDao) GetGroupUserByUserId(userId string, groupId string) model.GroupUser {
	var groupUser model.GroupUser
	db.DB.Where("user_id = ? and group_id = ?", userId, groupId).Find(&groupUser)
	return groupUser
}

func (gud groupUserDao) GetGroupUsers(groupId string) []model.GroupUser {
	var groupUsers []model.GroupUser
	db.DB.Where("group_id = ?", groupId).Find(&groupUsers)
	return groupUsers
}
