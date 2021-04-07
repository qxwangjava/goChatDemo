package ws_action

type Action func(userId string, data []byte, ext ...interface{}) []byte

var WsActionMap = make(map[string]Action, 1)

func init() {
	WsActionMap[SendMessageAction] = SendMessage
	WsActionMap[CloseConnAction] = CloseConn
	WsActionMap[AddFriendAction] = AddFriend
	WsActionMap[GetFriendApplyAction] = GetFriendApply
}

// 消息结构替
type ImAction struct {
	//事件
	Action string `json:"action"`
}
