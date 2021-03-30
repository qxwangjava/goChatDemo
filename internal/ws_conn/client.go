package ws_conn

import (
	"github.com/gorilla/websocket"
	"goChatDemo/pkg/logger"
)

type Client struct {
	UserInfo  UserInfo
	Conn      *websocket.Conn
	WriteChan chan []byte
}

func (c Client) Read() {
	conn := c.Conn
	for {
		mt, data, err := conn.ReadMessage()
		if err != nil {
			logger.Logger.Error("服务端读取消息失败:", err)
			//连接失败，认为设备离线
			ConnTypeMap[c.UserInfo.DeviceType].Delete(c.UserInfo.UserId)
			delete(ConnManager, c.UserInfo.UserId)
			conn.Close()
			close(c.WriteChan)
			var connCnt = len(ConnManager)
			logger.Logger.Info("当前连接数量：", connCnt)
			return
		}
		logger.Logger.Info("server recv: ", string(data))
		//处理数据
		handleResult := handleMessage(&c.UserInfo, mt, data)
		c.WriteChan <- handleResult
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

func (c Client) Write() {
	for {
		handleResult := <-c.WriteChan
		err := c.Conn.WriteMessage(websocket.TextMessage, handleResult)
		logger.Logger.Info("server send: ", string(handleResult))
		if err != nil {
			logger.Logger.Error("server write error:", err)
			break
		}
	}
}
