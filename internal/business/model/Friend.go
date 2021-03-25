package model

import "time"

type Friend struct {
	// 创建时间
	CreateTime time.Time `db:"create_time" gorm:"createTime"`
	// 好友Id
	FriendId int64 `db:"friend_id" gorm:"friendId"`
	// 主键
	Id int64 `db:"id" gorm:"id"`
	// 备注
	Remark string `db:"remark" gorm:"remark"`
	// 状态 1-申请 2-同意 3-删除 4-拉黑
	Status int `db:"status" gorm:"status"`
	// 更新时间
	UpdateTime time.Time `db:"update_time" gorm:"updateTime"`
	// 用户Id
	UserId int64 `db:"user_id" gorm:"userId"`
}
