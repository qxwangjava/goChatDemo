package rpc

import (
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"goChatDemo/pkg/rpc/server"
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
		pb.RegisterHelloServer(s, server.UserService)

		logger.Logger.Info("Listen on " + RpcAddr)
		s.Serve(listen)
	}()

}
