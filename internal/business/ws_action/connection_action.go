package ws_action

import (
	"context"
	"goChatDemo/internal/business/rpc_client"
	"goChatDemo/pkg/db"
	"strconv"
	"strings"
)

const (
	CloseConnAction = "closeConn"
)

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
