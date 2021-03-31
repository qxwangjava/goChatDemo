package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"goChatDemo/internal/business/dao"
	"goChatDemo/internal/business/rpc_client"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
	"strconv"
	"strings"
)

type MessageService struct {
}

//消息大类
type Message struct {
	//消息类型 1 文本 2图片
	MessageType int `json:"messageType"`
	//对方类型
	OtherSideType int `json:"otherSideType"`
	//对方id
	OtherSideId string `json:"otherSideId"`
	//消息内容
	MessageContent string `json:"messageContent"`
}

type TextMessage struct {
	*Message
	//消息内容
	Text string `json:"text"`
}

//发送消息处理
func (ms *MessageService) SendMessage(userId string, data []byte) []byte {
	var result = []byte{}
	var message = Message{}
	err := json.Unmarshal(data, &message)
	if err != nil {
		logger.Logger.Error("解析消息失败：", errors.Wrap(err, ""))
	}
	messageType := message.MessageType
	//找到需要发送的链接列表
	if message.OtherSideType == 1 { //对个人发消息
		//判断好友关系、黑名单等
		friendId := message.OtherSideId
		friend := dao.FriendDao.GetFriend(userId, friendId)
		if friend.Id == 0 || friend.Status != 2 {
			result, _ := json.Marshal(gerror.ErrorMsg("发送失败，好友关系不正常"))
			return result
		}
		//找到Ip,grpc调用
		deviceInfoList := db.RedisClient.LRange(context.Background(), friendId, 0, -1).Val()
		var connIp *redis.StringCmd
		for i := range deviceInfoList {
			deviceInfo := deviceInfoList[i]
			deviceType := strings.Split(deviceInfo, "|")[0]
			deviceId := strings.Split(deviceInfo, "|")[1]
			key := deviceType + "_" + deviceId + "_" + friendId
			connIp = db.RedisClient.Get(context.Background(), key)
			if connIp != nil {
				deviceTypeStr, _ := strconv.Atoi(deviceType)
				sendMsgResult := rpc_client.SendMsg(connIp.Val(), friendId, deviceTypeStr, deviceId, message.MessageType, message.MessageContent)
				if !sendMsgResult.Success {
					//TODO 消息发送失败的处理
				}
				result, _ := json.Marshal(gerror.SUCCESS)
				return result
			} else {
				result, _ := json.Marshal(gerror.ErrorMsg("没有找到连接"))
				return result
			}
		}
	}
	if message.OtherSideType == 2 { //群发消息
		groupId := message.OtherSideId
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

		result, _ := json.Marshal(gerror.SUCCESS)
		return result
	}

	switch messageType {
	//case 1:
	//	result = sendTextMessage(connInfo.UserId, data)
	//case 2:
	//	result = sendImageMessage(connInfo.UserId, data)
	default:
		result, _ = json.Marshal(gerror.ErrorMsg("不支持的消息类型"))
	}
	return result
}

//发送图片消息
func sendImageMessage(fromUserId string, data []byte) []byte {
	result, _ := json.Marshal(gerror.ErrorMsg("暂不支持图片消息"))
	return result
}
