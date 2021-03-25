package rpc

import (
	"goChatDemo/config"
	"goChatDemo/internal/business/service"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"google.golang.org/grpc"
	"net"
)

var (
	UserServiceClient pb.UserServiceClient
)

func InitUserServiceClient() {
	conn, err := grpc.Dial(config.RpcConfig.RpcAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	UserServiceClient = pb.NewUserServiceClient(conn)
}

func InitRpc() {
	go func() {
		listen, err := net.Listen("tcp", config.RpcConfig.RpcAddr)
		if err != nil {
			logger.Logger.Error("grpc启动失败: ", err)
			panic(err)
		}

		// 实例化grpc Server
		s := grpc.NewServer()

		// 注册HelloService
		pb.RegisterHelloServer(s, service.HelloService)
		pb.RegisterUserServiceServer(s, service.UserService)
		logger.Logger.Info("grpc启动成功，监听端口：" + config.RpcConfig.RpcAddr)
		s.Serve(listen)
	}()

}
