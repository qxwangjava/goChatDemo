package model

import (
	"time"
)

type ErrorPush struct {
	// 主键
	Id int64 `db:"id" gorm:"id"`
	//
	CreateTime time.Time `db:"create_time" gorm:"createTime"`
	// 推送内容
	PushContent string `db:"push_content" gorm:"pushContent"`
	// 用户Id
	UserId string `db:"user_id" gorm:"userId"`
	// 是否处理
	IsHandle int `db:"is_handle" gorm:"isHandle"`
	// 处理时间
	HandleTime time.Time `db:"handle_time" gorm:"handleTime"`
}
