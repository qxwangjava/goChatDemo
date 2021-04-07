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
	"time"
)

type MessageService struct {
}

//消息大类
type Message struct {
	//消息Id
	MessageClientId string `json:"messageClientId"`
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
	var result []byte
	var message = Message{}
	err := json.Unmarshal(data, &message)
	if err != nil {
		logger.Logger.Error("解析消息失败：", errors.Wrap(err, ""))
		result, _ := json.Marshal(gerror.ErrorMsg("解析消息失失败"))
		return result
	}
	//确认客户端消息Id,防止重复发送
	if message.MessageClientId == "" {
		result, _ := json.Marshal(gerror.ErrorMsg("客户端消息id不能为空"))
		return result
	}
	dbMessage := dao.MessageDao.GetMsgByClientId(message.MessageClientId)
	if dbMessage.Id != 0 {
		result, _ := json.Marshal(gerror.ErrorMsg("消息重复发送"))
		return result
	}
	//消息入库
	dbMessage.MessageType = message.MessageType
	dbMessage.MessageContent = message.MessageContent
	dbMessage.FromUser = userId
	dbMessage.To = message.OtherSideId
	dbMessage.ToType = message.OtherSideType
	dbMessage.ClientId = message.MessageClientId
	dbMessage.SendTime = time.Unix(0, 0)
	dbMessage.RecallTime = time.Unix(0, 0)
	messageServerId, err := dao.MessageDao.AddMessage(dbMessage)
	if err != nil {
		result, _ := json.Marshal(gerror.ErrorMsg("消息发送失败，请重试"))
		return result
	}

	messageType := message.MessageType
	//找到需要发送的链接列表
	var successKeyList []string
	var failKeyList []string
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
				sendMsgResult := rpc_client.SendMsg(
					connIp.Val(),
					friendId,
					deviceTypeStr,
					deviceId,
					message.MessageType, message.MessageContent, messageServerId, userId, 1, 0)
				if !sendMsgResult.Success {
					failKeyList = append(failKeyList, key)
				} else {
					successKeyList = append(successKeyList, key)
				}
			} else {
				failKeyList = append(failKeyList, key)
			}
		}
		result, _ := json.Marshal(gerror.Success(messageServerId))
		return result
	}
	if message.OtherSideType == 2 { //群发消息
		//TODO 群聊消息 对人入库
		groupId := message.OtherSideId
		//判断当前用户是否禁言
		groupUser := dao.GroupUserDao.GetGroupUserByUserId(userId, groupId)
		if groupUser.Id == 0 || groupUser.Status == 1 {
			result, _ := json.Marshal(gerror.ErrorMsg("发送失败，发送人被禁言，或已被提出群聊"))
			return result
		}
		//获取根据群id获取群用户
		groupUsers := dao.GroupUserDao.GetGroupUsers(groupId)
		var successKeyList []string
		var failKeyList []string
		for i := range groupUsers {
			groupUser := groupUsers[i]
			groupUserId := strconv.FormatInt(groupUser.UserId, 10)
			//找到Ip,grpc调用
			deviceInfoList := db.RedisClient.LRange(context.Background(), groupUserId, 0, -1).Val()
			var connIp *redis.StringCmd
			for i := range deviceInfoList {
				deviceInfo := deviceInfoList[i]
				deviceType := strings.Split(deviceInfo, "|")[0]
				deviceId := strings.Split(deviceInfo, "|")[1]
				key := deviceType + "_" + deviceId + "_" + groupUserId
				connIp = db.RedisClient.Get(context.Background(), key)
				if connIp != nil {
					deviceTypeStr, _ := strconv.Atoi(deviceType)
					iGroupId, _ := strconv.Atoi(groupId)
					sendMsgResult := rpc_client.SendMsg(
						connIp.Val(),
						groupUserId,
						deviceTypeStr,
						deviceId,
						message.MessageType, message.MessageContent, messageServerId, userId, 1, iGroupId)
					if !sendMsgResult.Success {
						failKeyList = append(failKeyList, key)
						//TODO 消息发送失败的处理
					} else {
						successKeyList = append(successKeyList, key)
					}
				} else {
					failKeyList = append(failKeyList, key)
				}
			}
		}

		result, _ := json.Marshal(gerror.Success(messageServerId))
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
