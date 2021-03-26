package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"goChatDemo/config"
	"goChatDemo/internal/business/service"
	"goChatDemo/internal/manager"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
	"net/http"
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
		var imAction = service.ImAction{}
		err := json.Unmarshal(data, &imAction)
		gerror.HandleError(err)
		switch imAction.Action {
		case "sendMessage":
			//TODO 这里考虑分布式（1 grpc http2协议 2 mqtt）
			result = service.SendMessage(connInfo, data)
		default:
			result, _ = json.Marshal(gerror.ErrorMsg("找不到action"))
		}
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
	//websocket客户端一定是web端。
	var token = r.Header.Get("token")
	logger.Logger.Info("客户端信息：token:", token)
	//TODO 验证token的有效性获取userId和deviceId
	var userId, deviceId, deviceType = manager.GetUserInfoFromToken(token)
	//建立连接
	conn, err := upgrader.Upgrade(w, r, nil)
	//断开原来的web连接
	var userConnManager = manager.ConnTypeMap[deviceType]
	value, ok := userConnManager.Load(userId)
	if ok {
		oldConn := value.(*websocket.Conn)
		_ = oldConn.Close()
	}
	//TODO 读取多端登录配置,根据配置选择断开哪些连接

	//保存长连接
	userInfo := manager.UserInfo{
		UserId:     userId,
		DeviceId:   deviceId,
		DeviceType: deviceType,
	}
	userConnManager.Store(userId, conn)
	var userIdConn = manager.ConnManager[userId]
	if userIdConn == nil {
		userIdConn = []manager.UserInfo{}
	}
	userIdConn = append(userIdConn, userInfo)
	manager.ConnManager[userId] = userIdConn
	logger.Logger.Info("建立连接成功")
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	var connCnt = len(manager.ConnManager)
	logger.Logger.Info("当前连接数量：", connCnt)
	client := Client{
		Conn:      conn,
		UserInfo:  &userInfo,
		WriteChan: make(chan []byte, 1000),
	}
	go client.Read()
	go client.Write()
}
