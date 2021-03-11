package model

import (
	"time"
)

type User struct {
	// 生日
	Birthday *time.Time `db:"birthday" gorm:"birthday default:null"`
	// 创建时间
	CreatedAt time.Time `db:"created_at" gorm:"createdAt"`
	// 邮箱
	Email string `db:"email" gorm:"email"`
	// 扩展字段
	Ex string `db:"ex" gorm:"ex"`
	// 头像地址
	HeadImg string `db:"headImg" gorm:"headImg"`
	// 主键
	Id int64 `db:"id" gorm:"id"`
	// 是否删除 1:是  -1:否
	IsDeleted int `db:"is_deleted" gorm:"isDeleted"`
	// 手机号
	Mobile string `db:"mobile" gorm:"mobile"`
	// 昵称
	NickName string `db:"nick_name" gorm:"nickName"`
	// 性别
	Sex int `db:"sex" gorm:"sex"`
	// 更新时间
	UpdatedAt time.Time `db:"updated_at" gorm:"updatedAt"`
	// 用户名
	UserName string `db:"user_name" gorm:"userName"`
}
