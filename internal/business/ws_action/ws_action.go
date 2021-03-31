package ws_action

import (
	"context"
	"goChatDemo/internal/business/rpc_client"
	"goChatDemo/internal/business/service"
	"goChatDemo/pkg/db"
	"strconv"
	"strings"
)

const (
	SendMessageAction = "sendMessage"
	CloseConnAction   = "closeConn"
)

type Action func(userId string, data []byte, ext ...interface{}) []byte

var WsActionMap = make(map[string]Action, 1)

func init() {
	WsActionMap[SendMessageAction] = SendMessage
	WsActionMap[CloseConnAction] = CloseConn
}

// 消息结构替
type ImAction struct {
	//事件
	Action string `json:"action"`
}

var SendMessage = func(userId string, data []byte, ext ...interface{}) []byte {
	var messageService = service.MessageService{}
	return messageService.SendMessage(userId, data)
}

var CloseConn = func(userId string, data []byte, ext ...interface{}) []byte {
	deviceType := ext[0].(int)
	deviceTypeOnLineCnt := db.RedisClient.LLen(context.Background(), userId).Val()
	oldOnlineIp := ""
	if deviceTypeOnLineCnt > 0 {
		deviceInfoList := db.RedisClient.LRange(context.Background(), userId, 0, -1).Val()
		for i := range deviceInfoList {
			oldDeviceInfo := deviceInfoList[i]
			oldDeviceType := strings.Split(oldDeviceInfo, "|")[0]
			oldDeviceId := strings.Split(oldDeviceInfo, "|")[1]
			if oldDeviceType == strconv.Itoa(deviceType) {
				oldConnKey := oldDeviceType + "_" + oldDeviceId + "_" + userId
				oldOnlineIp = db.RedisClient.Get(context.Background(), oldConnKey).Val()
				break
			}
		}
	}
	// grpc 下线 func CloseConn(key string)
	rpc_client.CloseConn(oldOnlineIp, userId, deviceType)
	return nil
}
