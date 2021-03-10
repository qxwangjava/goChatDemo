package server

import (
	"context"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
)

type userService struct{}

var UserService = userService{}

func (u userService) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var name = in.Name
	logger.Logger.Info("收到请求：", name)
	var result = name + "，你好"
	response := pb.Response{
		Result: result,
	}
	return &response, nil
}
