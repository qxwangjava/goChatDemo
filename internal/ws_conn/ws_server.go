package ws_conn

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"goChatDemo/config"
	"goChatDemo/internal/business/controller"
	"goChatDemo/internal/business/rpc_client"
	"goChatDemo/internal/manager"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/util"
	"net/http"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleMessage(connInfo *manager.UserInfo, messageType int, data []byte) []byte {
	logger.Logger.Info("server recv: ", string(data))
	result := []byte{}
	switch messageType {
	case websocket.TextMessage: //处理文本类型的消息
		var imAction = controller.ImAction{}
		err := json.Unmarshal(data, &imAction)
		gerror.HandleError(err)
		messageHandler := controller.MessageHandlerMap[imAction.Action]
		if messageHandler == nil {
			result, _ = json.Marshal(gerror.ErrorMsg("找不到action"))
			return result
		}
		//这里考虑分布式（1 grpc http2协议 ）
		return messageHandler(connInfo, data)

	case websocket.PingMessage: //处理PING的消息
		logger.Logger.Info("收到ping消息")
	}
	return result
}

func InitWSServer() {
	go func() {
		http.HandleFunc("/ws", wsHandler)
		err := http.ListenAndServe(config.WebConfig.WebSocketAddr, nil)
		if err != nil {
			logger.Logger.Error("webSocket 启动失败:", err)
			panic(err)
		}

	}()
	logger.Logger.Info("webSocket启动成功,监听端口 8081")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//获取参数 userId deviceId token
	var token = r.Header.Get("token")
	logger.Logger.Info("客户端信息：token:", token)
	//TODO 验证token的有效性获取userId和deviceId
	var userId, deviceId, deviceType = manager.GetUserInfoFromToken(token)
	//建立连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	key := strconv.Itoa(deviceType) + "_" + userId
	serverIp := util.GetServerIp()
	//断开原来的web连接
	_ = manager.CloseConn(deviceType, userId)
	//当前用户不在当前服务器上
	oldIp := db.RedisClient.Get(context.Background(), key)
	if oldIp != nil {
		// grpc 下线 func CloseConn(key string)
		rpc_client.CloseConn(oldIp.Val(), userId, deviceType)
	}
	//TODO 读取多端登录配置,根据配置选择断开哪些连接

	//redis 存储用户ip映射
	db.RedisClient.Set(context.Background(), key, serverIp, 0)

	//redis 存储用户在线状态
	deviceTypeOnLineCnt := db.RedisClient.LLen(context.Background(), userId).Val()
	if deviceTypeOnLineCnt != 0 {
		deviceTypeList := db.RedisClient.LRange(context.Background(), userId, 0, -1).Val()
		status := false
		for i := range deviceTypeList {
			if deviceTypeList[i] == strconv.Itoa(deviceType) {
				status = true
				break
			}
		}
		if !status {
			deviceTypeList = append(deviceTypeList, strconv.Itoa(deviceType))
			db.RedisClient.LPush(context.Background(), userId, deviceType)
		}
	} else {
		deviceTypeList := []int{}
		deviceTypeList = append(deviceTypeList, deviceType)
		db.RedisClient.LPush(context.Background(), userId, deviceType)
	}

	//保存长连接到本地
	userInfo := manager.UserInfo{
		Addr:       conn.RemoteAddr().String(),
		UserId:     userId,
		DeviceId:   deviceId,
		DeviceType: deviceType,
		LoginTime:  time.Now().Unix(),
	}
	userConnManager := manager.ConnTypeMap[deviceType]
	var userIdConn = manager.ConnManager[userId]
	if userIdConn == nil {
		userIdConn = []manager.UserInfo{}
	}
	userIdConn = append(userIdConn, userInfo)
	manager.ConnManager[userId] = userIdConn
	logger.Logger.Info("建立连接成功")
	var connCnt = len(manager.ConnManager)
	logger.Logger.Info("当前连接数量：", connCnt)
	client := Client{
		Conn:      conn,
		UserInfo:  &userInfo,
		WriteChan: make(chan []byte, 1000),
	}
	userConnManager.Store(userId, client)
	go client.Read()
	go client.Write()
}
