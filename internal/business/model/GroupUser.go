package model

import (
	"time"
)

type GroupUser struct {
	// 创建时间
	CreateTime time.Time `db:"create_time" gorm:"createTime"`
	// 附加属性
	Extra string `db:"extra" gorm:"extra"`
	// 组id
	GroupId int64 `db:"group_id" gorm:"groupId"`
	// 自增主键
	Id int64 `db:"id" gorm:"id"`
	// 成员类型，1：群主；2：管理员；3：普通成员
	MemberType int `db:"member_type" gorm:"memberType"`
	// 备注
	Remarks string `db:"remarks" gorm:"remarks"`
	// 禁言状态 0 否 1是
	Status int `db:"status" gorm:"status"`
	// 更新时间
	UpdateTime time.Time `db:"update_time" gorm:"updateTime"`
	// 用户id
	UserId int64 `db:"user_id" gorm:"userId"`
}
