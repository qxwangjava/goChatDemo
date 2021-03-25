package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"goChatDemo/internal/manager"
	"goChatDemo/pkg/gerror"
)

// 消息结构替
type ImAction struct {
	//事件
	Action string `json:"action"`
}

//消息大类
type Message struct {
	*ImAction
	//消息类型 1 文本 2图片
	MessageType int `json:"messageType"`
	//对方类型
	OtherSideType int `json:"otherSideType"`
	//对方id
	OtherSideId string `json:"otherSideId"`
}

type TextMessage struct {
	*Message
	//消息内容
	Text string `json:"text"`
}

//发送消息处理
func SendMessage(data []byte) []byte {
	var result = []byte{}
	var message = Message{}
	err := json.Unmarshal(data, &message)
	gerror.HandleError(err)
	messageType := message.MessageType
	switch messageType {
	case 1:
		result = sendTextMessage(data)
	case 2:

	default:
		result, _ = json.Marshal(gerror.ErrorMsg("找不到的消息类型"))
	}
	return result

}

//发送文本消息
func sendTextMessage(data []byte) []byte {
	textMessage := TextMessage{}
	err := json.Unmarshal(data, &textMessage)
	gerror.HandleError(err)
	if textMessage.OtherSideType == 1 { //对个人发消息
		connInfoList := manager.ConnManager[textMessage.OtherSideId]
		for i := range connInfoList {
			ConnMap := manager.ConnTypeMap[connInfoList[i].DeviceType]
			value, ok := ConnMap.Load(textMessage.OtherSideId)
			if ok {
				conn := value.(*websocket.Conn)
				_ = conn.WriteMessage(1, []byte(textMessage.Text))
				//TODO 不在线时消息存储
			}
		}
		result, _ := json.Marshal(gerror.SUCCESS)
		return result
	}
	if textMessage.OtherSideType == 2 { //群发消息
		result, _ := json.Marshal(gerror.SUCCESS)
		return result
	}
	result, _ := json.Marshal(gerror.ErrorMsg("找到对方类型"))
	return result
}
