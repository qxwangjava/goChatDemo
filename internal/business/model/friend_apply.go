package model

import (
	"time"
)

type FriendApply struct {
	// 主键
	Id int64 `db:"id" gorm:"id"`
	// 用户Id
	UserId string `db:"user_id" gorm:"userId"`
	// 好友Id
	FriendId string `db:"friend_id" gorm:"friendId"`
	// 好友头像
	FriendHeadImg string `db:"friend_head_img" gorm:"friendHeadImg"`
	// 好友昵称
	FriendName string `db:"friend_name" gorm:"friendName"`
	// 创建时间
	CreateTime time.Time `db:"create_time" gorm:"createTime"`
	// 更新时间
	UpdateTime time.Time `db:"update_time" gorm:"updateTime"`
	// 备注
	Remark string `db:"remark" gorm:"remark"`
	// 状态 1-申请 2-同意 3-拒绝
	Status int `db:"status" gorm:"status"`
}
