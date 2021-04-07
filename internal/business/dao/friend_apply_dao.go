package dao

import (
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
	"time"
)

type friendApplyDao struct {
}

var FriendApplyDao = new(friendApplyDao)

func (fd friendApplyDao) AddFriend(userId, friendId, friendName, remark, friendHeadImg string) {
	friendApply := model.FriendApply{
		CreateTime:    time.Now(),
		FriendId:      friendId,
		FriendName:    friendName,
		FriendHeadImg: friendHeadImg,
		Remark:        remark,
		UserId:        userId,
		Status:        1,
		UpdateTime:    time.Now(),
	}
	db.DB.Save(&friendApply)
}

func (fd friendApplyDao) GetFriendApply(userId string, pageSize int, pageNum int, friendName string, status int) []model.FriendApply {
	var applyList []model.FriendApply
	dbq := db.DB.Limit(pageSize).Offset(pageSize * (pageNum - 1))
	dbq.Where("userId = ? and status = ? and friendName like '%?%' ", userId, status, friendName).Find(&applyList)
	return applyList
}
