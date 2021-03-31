package ws_conn

import (
	"context"
	"goChatDemo/pkg/db"
	"strconv"
	"strings"
	"sync"
)

type UserInfo struct {
	Addr       string // 客户端地址
	UserId     string
	DeviceId   string
	DeviceType int   //1-android 2-ios 3-web
	LoginTime  int64 // 登录时间 登录以后才有
}

var ConnTypeMap = map[int]*sync.Map{}

// 安卓连接容器
var AndroidConn = sync.Map{}

// IOS连接容器
var IOSConn = sync.Map{}

// web连接容器
var WebConn = sync.Map{}

// 用户连接信息容器
var LocalConnInfoManager = sync.Map{}

// 解析token,获取用户id和设备id,反参 userId,deviceId,deviceType
func GetUserInfoFromToken(token string) (string, string, int) {
	userInfo := strings.Split(token, "|")
	deviceType, _ := strconv.Atoi(userInfo[2])
	return userInfo[0], userInfo[1], deviceType
}

//获取连接数量
func getConnCnt() int {
	result := 0
	LocalConnInfoManager.Range(func(k, v interface{}) bool {
		result++
		return true
	})
	return result
}

//初始化链接容器的容器
func InitConnMap() {
	ConnTypeMap[1] = &AndroidConn
	ConnTypeMap[2] = &IOSConn
	ConnTypeMap[3] = &WebConn
}

func GetOldOnLineIpAndKey(deviceType int, userId string) (string, string) {
	deviceTypeOnLineCnt := db.RedisClient.LLen(context.Background(), userId).Val()
	oldOnlineIp := ""
	oldConnKey := ""
	if deviceTypeOnLineCnt > 0 {
		deviceInfoList := db.RedisClient.LRange(context.Background(), userId, 0, -1).Val()
		for i := range deviceInfoList {
			oldDeviceInfo := deviceInfoList[i]
			oldDeviceType := strings.Split(oldDeviceInfo, "|")[0]
			oldDeviceId := strings.Split(oldDeviceInfo, "|")[1]
			if oldDeviceType == strconv.Itoa(deviceType) {
				oldConnKey = oldDeviceType + "_" + oldDeviceId + "_" + userId
				oldOnlineIp = db.RedisClient.Get(context.Background(), oldConnKey).Val()
			}
		}
	}
	return oldOnlineIp, oldConnKey
}

func CloseConn(deviceType int, userId string) {
	_, key := GetOldOnLineIpAndKey(deviceType, userId)
	value, ok := LocalConnInfoManager.Load(key)
	if ok {
		localClient := (value).(Client)
		LocalConnInfoManager.Delete(key)
		localClient.Conn.Close()
		close(localClient.WriteChan)
	}
}
