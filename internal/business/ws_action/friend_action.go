package ws_action

import "goChatDemo/internal/business/service"

const (
	AddFriendAction      = "addFriendAction"
	GetFriendApplyAction = "getFriendApply"
)

//好友申请
var AddFriend = func(userId string, data []byte, ext ...interface{}) []byte {
	var friendService = service.FriendService{}
	return friendService.AddFriend(userId, data)
}

//获取好友申请
var GetFriendApply = func(userId string, data []byte, ext ...interface{}) []byte {
	var friendService = service.FriendService{}
	return friendService.GetFriendApply(userId, data)
}

// TODO 同意好友申请

// TODO 拒绝好友申请
