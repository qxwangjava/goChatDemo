package service

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"goChatDemo/internal/business/model"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
)

type userService struct{}

var UserService = userService{}

func (u userService) AddUser(ctx context.Context, addUserDto *pb.AddUserDto) (*pb.UserId, error) {
	birthday, _ := ptypes.Timestamp(addUserDto.Birthday)
	requestBody, _ := json.Marshal(addUserDto)
	logger.Logger.Info("收到请求：", string(requestBody))
	dbUser := model.User{
		Birthday: &birthday,
		Email:    addUserDto.Email,
		Ex:       addUserDto.Ex,
		HeadImg:  addUserDto.HeadImg,
		UserName: addUserDto.UserName,
		NickName: addUserDto.NickName,
		Mobile:   addUserDto.Mobile,
		Sex:      int(addUserDto.Sex),
	}

	db.DB.Save(&dbUser)
	result := pb.UserId{
		UserId: dbUser.Id,
	}
	return &result, nil
}

// 通过Id删除用户
func (u userService) DeleteUserById(context.Context, *pb.UserId) (*pb.NullMessage, error) {
	return &pb.NullMessage{}, nil
}

// 通过id获取用户信息
func (u userService) GetUserById(context.Context, *pb.UserId) (*pb.QueryUserDto, error) {
	return &pb.QueryUserDto{}, nil
}

// 通过手机号精确查询用户
func (u userService) GetUserByMobile(context.Context, *pb.UserMobile) (*pb.QueryUserDto, error) {
	return &pb.QueryUserDto{}, nil
}

// 更新用户信息
func (u userService) UpdateUser(context.Context, *pb.AddUserDto) (*pb.NullMessage, error) {
	return &pb.NullMessage{}, nil
}
