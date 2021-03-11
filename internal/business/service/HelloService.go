package service

import (
	"context"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var name = in.Name
	logger.Logger.Info("收到请求：", name)
	var result = name + "，你好"
	response := pb.Response{
		Result: result,
	}
	return &response, nil
}
