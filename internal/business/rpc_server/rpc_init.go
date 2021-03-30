package rpc_server

import (
	"goChatDemo/config"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"google.golang.org/grpc"
	"net"
)

var (
	UserServiceClient pb.UserServiceClient
)

//
//func InitUserServiceClient() {
//	conn, err := grpc.Dial(config.RpcConfig.RpcAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(Interceptor))
//	if err != nil {
//		logger.Logger.Error(err)
//		panic(err)
//	}
//	UserServiceClient = pb.NewUserServiceClient(conn)
//}

func InitRpc() {
	go func() {
		listen, err := net.Listen("tcp", config.RpcConfig.RpcPort)
		if err != nil {
			logger.Logger.Error("grpc启动失败: ", err)
			panic(err)
		}

		// 实例化grpc Server
		s := grpc.NewServer()

		// 注册HelloService
		pb.RegisterHelloServer(s, HelloService)
		pb.RegisterUserServiceServer(s, UserService)
		pb.RegisterImServerServer(s, ImService)
		logger.Logger.Info("grpc启动成功，监听端口" + config.RpcConfig.RpcPort)
		err = s.Serve(listen)
		if err != nil {
			logger.Logger.Error("grpc启动失败: ", err)
			panic(err)
		}

	}()

}
