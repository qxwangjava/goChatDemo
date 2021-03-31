package web_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"goChatDemo/internal/business/rpc_server"
	"goChatDemo/pkg/pb"
	"net/http"
	"time"
)

type addUserVo_back struct {
	// 邮箱
	Email string `json:"email,omitempty"`

	// 手机号
	Mobile string `json:"mobile,omitempty"`

	// 验证码
	VerificationCode string `json:"verificationCode,omitempty"`

	// 微信Id
	WechatId string `json:"wechatId,omitempty"`
}

type addUserVo struct {
	// 生日
	Birthday string `json:"birthday,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 扩展字段
	Ex string `json:"ex,omitempty"`
	// 头像地址
	HeadImg string `json:"headImg,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 昵称
	NickName string `json:"nickName,omitempty"`
	// 性别
	Sex int32 `json:"sex,omitempty"`
	// 用户名
	UserName string `json:"userName,omitempty"`
}

type updateUserVo struct {
	// 生日
	Birthday string `json:"birthday,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 扩展字段
	Ex string `json:"ex,omitempty"`
	// 头像地址
	HeadImg string `json:"headImg,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 昵称
	NickName string `json:"nickName,omitempty"`
	// 性别
	Sex int32 `json:"sex,omitempty"`
	// 用户名
	UserName string `json:"userName,omitempty"`
}

var AddUser = func(c *gin.Context) error {
	au := addUserVo{}
	err := c.ShouldBindJSON(&au)
	if err != nil {
		return errors.Wrap(err, "")
	}
	birthday, err := time.Parse("2006-01-02 15:04:05", au.Birthday)
	if err != nil {
		return errors.Wrap(err, "")
	}
	b, err := ptypes.TimestampProto(birthday)
	if err != nil {
		return errors.Wrap(err, "")
	}
	addUserDto := &pb.AddUserDto{
		Birthday: b,
		UserName: au.UserName,
		NickName: au.NickName,
		Email:    au.Email,
		Mobile:   au.Mobile,
		Sex:      au.Sex,
		Ex:       au.Ex,
		HeadImg:  au.HeadImg,
	}
	userId, err := rpc_server.UserServiceClient.AddUser(context.Background(), addUserDto)
	if err != nil {
		return errors.Wrap(err, "")
	}
	c.JSON(http.StatusOK, userId)
	return nil
}
