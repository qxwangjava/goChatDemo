package main

import (
	"context"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"goChatDemo/pkg/rpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(rpc.RpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)
	response, err := c.SayHello(context.Background(), &pb.Request{
		Name: "wqx",
	})
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	logger.Logger.Info(response)
}
