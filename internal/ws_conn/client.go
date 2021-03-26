package ws_conn

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"goChatDemo/internal/business/service"
	"goChatDemo/internal/emqt"
	"goChatDemo/internal/manager"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
)

type Client struct {
	UserInfo  *manager.UserInfo
	Conn      *websocket.Conn
	WriteChan chan []byte
}

//func subscribe(){
//	//订阅主题
//	if token := emqt.Client.Subscribe(emqt.SENDMESSAGE, 1, nil); token.Wait() && token.Error() != nil {
//		logger.Logger.Error("订阅",emqt.SENDMESSAGE,"主题失败:",token.Error())
//		return
//	}
//	//emqt.Client.AddRoute("sendMessage", sendMessage)
//	logger.Logger.Info("订阅",emqt.SENDMESSAGE,"主题成功！")
//}

func (c Client) Read() {
	conn := c.Conn
	for {
		mt, data, err := conn.ReadMessage()
		if err != nil {
			logger.Logger.Error("服务端读取消息失败:", err)
			//连接失败，认为设备离线
			manager.ConnTypeMap[c.UserInfo.DeviceType].Delete(c.UserInfo.UserId)
			delete(manager.ConnManager, c.UserInfo.UserId)
			conn.Close()
			close(c.WriteChan)
			var connCnt = len(manager.ConnManager)
			logger.Logger.Info("当前连接数量：", connCnt)
			return
		}
		logger.Logger.Info("server recv: ", string(data))
		//处理数据
		handleResult := handleMessage(c.UserInfo, mt, data)
		c.WriteChan <- handleResult
	}
}

func publicMqtt(data []byte) {
	//发布消息
	if token := emqt.Client.Publish(emqt.SENDMESSAGE, 2, false, data); token.Wait() && token.Error() != nil {
		logger.Logger.Error("发布消息失败：", token.Error())
	}
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
			publicMqtt(data)
			result, _ = json.Marshal(gerror.SUCCESS)
		default:
			result, _ = json.Marshal(gerror.ErrorMsg("找不到action"))
		}
	case websocket.PingMessage: //处理PING的消息
		logger.Logger.Info("收到ping消息")
	}
	return result
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
