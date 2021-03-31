package ws_conn

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"goChatDemo/config"
	"goChatDemo/internal/business/ws_action"
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

func handleMessage(connInfo *UserInfo, messageType int, data []byte) []byte {
	logger.Logger.Info("server recv: ", string(data))
	result := []byte{}
	switch messageType {
	case websocket.TextMessage: //处理文本类型的消息
		var imAction = ws_action.ImAction{}
		err := json.Unmarshal(data, &imAction)
		if err != nil {
			logger.Logger.Error("解析请求内容失败：", errors.Wrap(err, ""))
			result, _ = json.Marshal(gerror.ErrorMsg("解析请求内容失败，请检查请求内容"))
			return result
		}
		messageHandler := ws_action.WsActionMap[imAction.Action]
		if messageHandler == nil {
			result, _ = json.Marshal(gerror.ErrorMsg("找不到action"))
			return result
		}
		//这里考虑分布式（1 grpc http2协议 ）
		return messageHandler(connInfo.UserId, data)

	case websocket.PingMessage: //处理PING的消息
		logger.Logger.Info("收到ping消息")
	}
	return result
}

func InitWSServer() {
	go func() {
		http.HandleFunc("/ws", wsHandler)
		err := http.ListenAndServe(config.WebConfig.WebSocketPort, nil)
		if err != nil {
			logger.Logger.Error("webSocket 启动失败:", err)
			panic(err)
		}

	}()
	logger.Logger.Info("webSocket启动成功,监听端口", config.WebConfig.WebSocketPort)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//获取参数 userId deviceId token
	var token = r.Header.Get("token")
	logger.Logger.Info("客户端信息：token:", token)
	//TODO 验证token的有效性获取userId和deviceId
	var userId, deviceId, deviceType = GetUserInfoFromToken(token)
	//建立连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	//断开原来同类型设备的连接
	ws_action.WsActionMap[ws_action.CloseConnAction](userId, []byte{}, deviceType)

	//TODO 读取多端登录配置,根据配置选择断开哪些连接

	//redis存储Ip映射
	key := strconv.Itoa(deviceType) + "_" + deviceId + "_" + userId
	serverIp := util.GetServerIp()
	db.RedisClient.Set(context.Background(), key, serverIp, 0)
	//redis 存储用户在线状态
	deviceTypeOnLineCnt := db.RedisClient.LLen(context.Background(), userId).Val()
	deviceInfo := strconv.Itoa(deviceType) + "|" + deviceId
	if deviceTypeOnLineCnt != 0 {
		deviceInfoList := db.RedisClient.LRange(context.Background(), userId, 0, -1).Val()
		status := false
		for i := range deviceInfoList {
			if deviceInfoList[i] == deviceInfo {
				status = true
				break
			}
		}
		if !status {
			deviceInfoList = append(deviceInfoList, deviceInfo)
			db.RedisClient.LPush(context.Background(), userId, deviceInfo)
		}
	} else {
		var deviceInfoList []string
		deviceInfoList = append(deviceInfoList, deviceInfo)
		db.RedisClient.LPush(context.Background(), userId, deviceInfo)
	}

	//保存用户信息到服务器本地
	newUserInfo := UserInfo{
		Addr:       conn.RemoteAddr().String(),
		UserId:     userId,
		DeviceId:   deviceId,
		DeviceType: deviceType,
		LoginTime:  time.Now().Unix(),
	}
	client := Client{
		Conn:      conn,
		UserInfo:  newUserInfo,
		WriteChan: make(chan []byte, 1000),
	}
	LocalConnInfoManager.Store(key, client)
	logger.Logger.Info("建立连接成功")
	var connCnt = getConnCnt()
	logger.Logger.Info("当前连接数量：", connCnt)

	go client.Read()
	go client.Write()
}
