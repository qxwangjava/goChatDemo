package manager

import (
	"context"
	"github.com/gorilla/websocket"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/gerror"
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
var ConnManager = map[string][]UserInfo{}

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

func CloseConn(deviceType int, userId string) gerror.Result {
	var userConnManager = ConnTypeMap[deviceType]
	value, ok := userConnManager.Load(userId)
	if ok { //当前用户的在当前服务器上
		key := strconv.Itoa(deviceType) + "_" + userId
		db.RedisClient.Del(context.Background(), key)
		oldConn := value.(*websocket.Conn)
	_:
		oldConn.Close()

	}
	return gerror.SUCCESS
}

func GetConnByUserId(userId string) []*websocket.Conn {
	//获取用户连接信息
	connInfoList := ConnManager[userId]
	//构建结果对象
	var result = make([]*websocket.Conn, 0)
	//遍历连接信息，根据设备类型从不通的容器中获取连接
	for i := range connInfoList {
		ConnMap := ConnTypeMap[connInfoList[i].DeviceType]
		value, ok := ConnMap.Load(userId)
		if ok {
			conn := value.(*websocket.Conn)
			result = append(result, conn)
		}
	}
	return result
}
