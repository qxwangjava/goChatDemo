package ws_conn

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
	connInfo := manager.ConnInfo{
		UserId:     userId,
		DeviceId:   deviceId,
		DeviceType: deviceType,
	}
	userConnManager.Store(userId, conn)
	var userIdConn = manager.ConnManager[userId]
	if userIdConn == nil {
		userIdConn = []manager.ConnInfo{}
	}
	userIdConn = append(userIdConn, connInfo)
	manager.ConnManager[userId] = userIdConn
	logger.Logger.Info("建立连接成功")
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	var connCnt = len(manager.ConnManager)
	logger.Logger.Info("当前连接数量：", connCnt)
	go handleConn(&connInfo, conn)

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

func handleConn(connInfo *manager.ConnInfo, conn *websocket.Conn) {
	//for {
	//	mt, message, err := conn.ReadMessage()
	//	if err != nil {
	//		logger.Logger.Info("server read:", err)
	//		break
	//	}
	//	logger.Logger.Info("server recv: ", string(message))
	//	err = conn.WriteMessage(mt, message)
	//	logger.Logger.Info("server send: ", string(message))
	//	if err != nil {
	//		logger.Logger.Info("server write:", err)
	//		break
	//	}
	//}
	for {
		mt, data, err := conn.ReadMessage()
		handleResult := handlePackage(mt, data)
		if err != nil {
			logger.Logger.Error("服务端读取消息失败:", err)
			//连接失败，认为设备离线
			manager.ConnTypeMap[connInfo.DeviceType].Delete(connInfo.UserId)
			delete(manager.ConnManager, connInfo.UserId)
			gerror.HandleError(conn.Close())
			break
		}
		err = conn.WriteMessage(mt, handleResult)
		logger.Logger.Info("server send: ", string(handleResult))
		if err != nil {
			logger.Logger.Error("server write error:", err)
			break
		}
	}
}

/**
messageType 枚举
TextMessage = 1

BinaryMessage = 2

CloseMessage = 8

PingMessage = 9

PongMessage = 10

*/
func handlePackage(messageType int, data []byte) []byte {
	logger.Logger.Info("server recv: ", string(data))
	result := []byte{}
	switch messageType {
	case websocket.TextMessage: //处理文本类型的消息
		var imAction = service.ImAction{}
		err := json.Unmarshal(data, &imAction)
		gerror.HandleError(err)
		switch imAction.Action {
		case "sendMessage":
			result = service.SendMessage(data)
		default:
			result, _ = json.Marshal(gerror.ErrorMsg("找不到action"))
		}
	case websocket.PingMessage: //处理PING的消息
		logger.Logger.Info("收到ping消息")
	}
	return result
}
