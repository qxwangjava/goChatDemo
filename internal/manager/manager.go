package manager

import (
	"strconv"
	"strings"
	"sync"
)

type ConnInfo struct {
	UserId     string
	DeviceId   string
	DeviceType int //1-android 2-ios 3-web
}

var ConnTypeMap = map[int]*sync.Map{}

// 安卓连接容器
var AndroidConn = sync.Map{}

// IOS连接容器
var IOSConn = sync.Map{}

// web连接容器
var WebConn = sync.Map{}

// 用户连接信息容器
var ConnManager = map[string][]ConnInfo{}

// 解析token,获取用户id和设备id,反参 userId,deviceId,deviceType
func GetUserInfoFromToken(token string) (string, string, int) {
	userInfo := strings.Split(token, "|")
	deviceType, _ := strconv.Atoi(userInfo[2])
	return userInfo[0], userInfo[1], deviceType
}

//初始化链接容器的容器
func InitConnMap() {
	ConnTypeMap[1] = &AndroidConn
	ConnTypeMap[2] = &IOSConn
	ConnTypeMap[3] = &WebConn
}
