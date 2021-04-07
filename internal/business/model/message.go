package model

import (
	"time"
)

type Message struct {
	// 主键
	Id int64 `db:"id" gorm:"id"`
	// 来源用户
	FromUser string `db:"from_user" gorm:"fromUser"`
	// 目标类型  1-用户 2-群 3-聊天室（保留）
	ToType int `db:"to_type" gorm:"toType"`
	// 目标id
	To string `db:"to" gorm:"to"`
	// 消息类型 1-文本 2-图片
	MessageType int `db:"message_type" gorm:"messageType"`
	// 消息内容
	MessageContent string `db:"message_content" gorm:"messageContent"`
	// 消息客户端id
	ClientId string `db:"client_id" gorm:"clientId"`
	// 是否撤回 0-否 1-是
	RecallFlag int `db:"recall_flag" gorm:"recallFlag"`
	// 撤回时间
	RecallTime time.Time `db:"recall_time" gorm:"recallTime"`
	// 是否送达 0-否 1-是
	IsSend int `db:"is_send" gorm:"isSend"`
	// 发送时间
	SendTime time.Time `db:"send_time" gorm:"sendTime"`
}
