package rpc

import (
	"goChatDemo/internal/business/service"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"google.golang.org/grpc"
	"net"
)

func InitRpc() {
	go func() {
		listen, err := net.Listen("tcp", RpcAddr)
		if err != nil {
			logger.Logger.Error("Failed to listen: %v", err)
		}

		// 实例化grpc Server
		s := grpc.NewServer()

		// 注册HelloService
		pb.RegisterHelloServer(s, service.HelloService)
		pb.RegisterUserServiceServer(s, service.UserService)
		logger.Logger.Info("Listen on " + RpcAddr)
		s.Serve(listen)
	}()

}
