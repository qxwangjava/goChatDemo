package service

import (
	"encoding/json"
	"goChatDemo/internal/business/dao"
	"goChatDemo/pkg/gerror"
)

type FriendService struct{}

type addFriendData struct {
	friendId      string
	friendName    string
	friendHeadImg string
	remark        string
}

type getFriendApply struct {
	pageSize   int
	pageNum    int
	status     int
	friendName string
}

func (fr FriendService) AddFriend(userId string, data []byte) []byte {
	afd := addFriendData{}
	err := json.Unmarshal(data, &afd)
	if err != nil {
		result, _ := json.Marshal(gerror.ErrorMsg("数据解析失败"))
		return result
	}
	dao.FriendApplyDao.AddFriend(userId, afd.friendId, afd.friendName, afd.friendHeadImg, afd.remark)
	result, _ := json.Marshal(gerror.SUCCESS)
	return result
}

func (fr FriendService) GetFriendApply(userId string, data []byte) []byte {
	gfa := getFriendApply{}
	err := json.Unmarshal(data, &gfa)
	if err != nil {
		result, _ := json.Marshal(gerror.ErrorMsg("数据解析失败"))
		return result
	}
	applyList := dao.FriendApplyDao.GetFriendApply(userId, gfa.pageSize, gfa.pageNum, gfa.friendName, gfa.status)
	result, _ := json.Marshal(gerror.Success(applyList))
	return result
}
