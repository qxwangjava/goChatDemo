package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"goChatDemo/pkg/rpc"
	"net/http"
	"time"
)

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

var AddUser = func(c *gin.Context) {
	r := gin.Default()
	r.POST("/user/add")
	au := addUserVo{}
	err := c.ShouldBindJSON(&au)
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	birthday, err := time.Parse("2006-01-02 15:04:05", au.Birthday)
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	b, err := ptypes.TimestampProto(birthday)
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
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
	userId, err := rpc.UserServiceClient.AddUser(context.Background(), addUserDto)
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	c.JSON(http.StatusOK, rpc.Success(userId))
}
