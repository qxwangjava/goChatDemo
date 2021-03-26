package service

import (
	"encoding/json"
	"goChatDemo/internal/business/dao"
	"goChatDemo/internal/manager"
	"goChatDemo/pkg/gerror"
	"strconv"
)

// 消息结构替
type ImAction struct {
	//事件
	Action string `json:"action"`
}

//消息大类
type Message struct {
	*ImAction
	//消息类型 1 文本 2图片
	MessageType int `json:"messageType"`
	//对方类型
	OtherSideType int `json:"otherSideType"`
	//对方id
	OtherSideId string `json:"otherSideId"`
}

type TextMessage struct {
	*Message
	//消息内容
	Text string `json:"text"`
}

//发送消息处理
func SendMessage(userId string, data []byte) []byte {
	var result = []byte{}
	var message = Message{}
	err := json.Unmarshal(data, &message)
	gerror.HandleError(err)
	messageType := message.MessageType
	switch messageType {
	case 1:
		result = sendTextMessage(userId, data)
	case 2:
		result = sendImageMessage(userId, data)
	default:
		result, _ = json.Marshal(gerror.ErrorMsg("不支持的消息类型"))
	}
	return result

}

func sendImageMessage(userId string, data []byte) []byte {
	result, _ := json.Marshal(gerror.ErrorMsg("暂不支持图片消息"))
	return result
}

//发送文本消息
func sendTextMessage(userId string, data []byte) []byte {
	textMessage := TextMessage{}
	err := json.Unmarshal(data, &textMessage)
	gerror.HandleError(err)
	if textMessage.OtherSideType == 1 { //对个人发消息
		//判断好友关系、黑名单等
		friendId := textMessage.OtherSideId
		friend := dao.FriendDao.GetFriend(userId, friendId)
		if friend.Id == 0 || friend.Status != 2 {
			result, _ := json.Marshal(gerror.ErrorMsg("发送失败，好友关系不正常"))
			return result
		}
		connList := manager.GetConnByUserId(friendId)
		//遍历多端，向长连接发送消息
		for i := range connList {
			conn := connList[i]
			_ = conn.WriteMessage(1, []byte(textMessage.Text))
			//TODO 不在线时消息存储
		}
		result, _ := json.Marshal(gerror.SUCCESS)
		return result
	}
	if textMessage.OtherSideType == 2 { //群发消息
		groupId := textMessage.OtherSideId
		//判断当前用户是否禁言
		groupUser := dao.GroupUserDao.GetGroupUserByUserId(userId, groupId)
		if groupUser.Id == 0 || groupUser.Status == 1 {
			result, _ := json.Marshal(gerror.ErrorMsg("发送失败，发送人被禁言，或已被提出群聊"))
			return result
		}
		//获取根据群id获取群用户
		groupUsers := dao.GroupUserDao.GetGroupUsers(groupId)
		userList := []string{}
		for i := range groupUsers {
			userList = append(userList, strconv.FormatInt(groupUsers[i].Id, 10))
		}

		//发送消息
		for i := range userList {
			connList := manager.GetConnByUserId(userList[i])
			//遍历多端，向长连接发送消息
			for i := range connList {
				conn := connList[i]
				_ = conn.WriteMessage(1, []byte(textMessage.Text))
				//TODO 不在线时消息存储
			}
		}

		result, _ := json.Marshal(gerror.SUCCESS)
		return result
	}
	result, _ := json.Marshal(gerror.ErrorMsg("找到对方类型"))
	return result
}
