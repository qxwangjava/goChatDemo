package ws_conn

import (
	"context"
	"github.com/gorilla/websocket"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"strconv"
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
			key := strconv.Itoa(c.UserInfo.DeviceType) + "_" + c.UserInfo.DeviceId + "_" + c.UserInfo.UserId
			LocalConnInfoManager.Delete(key)
			db.RedisClient.Del(context.Background(), key)
			db.RedisClient.LRem(context.Background(),
				c.UserInfo.UserId,
				1,
				strconv.Itoa(c.UserInfo.DeviceType)+"|"+c.UserInfo.DeviceId)
			logger.Logger.Info(
				"用户id：",
				c.UserInfo.UserId,
				",设备类型：", c.UserInfo.DeviceType,
				",设备Id：", c.UserInfo.DeviceId,
				",客户端已掉线，当前连接数量：", getConnCnt())
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
			return

		}
	}
}
