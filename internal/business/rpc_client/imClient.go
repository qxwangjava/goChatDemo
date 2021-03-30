package rpc_client

import (
	"context"
	"github.com/pkg/errors"
	"goChatDemo/config"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/intercepter"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"google.golang.org/grpc"
)

func getRpcAddr(ip string) string {
	return ip + ":" + config.RpcConfig.RpcPort
}

func SendMsg(ip string, toUserId string, deviceType int, messageType int, message string) gerror.Result {
	rpcConn, err := grpc.Dial(getRpcAddr(ip), grpc.WithInsecure(), grpc.WithUnaryInterceptor(intercepter.Interceptor))
	if err != nil {
		logger.Logger.Error("连接rpc失败，消息发送失败:", err)
		return gerror.ErrorMsg("连接rpc失败，消息发送失败")
	}
	imServerClient := pb.NewImServerClient(rpcConn)
	in := pb.SendMsgReq{
		UserId:         toUserId,
		DeviceType:     int32(deviceType),
		MessageType:    int32(messageType),
		MessageContent: message,
	}
	_, err = imServerClient.SendMsg(context.Background(), &in)
	if err != nil {
		logger.Logger.Error("发送rpc失败：", errors.Wrap(err, ""))
		return gerror.ErrorMsg("发送rpc失败：" + err.Error())
	}
	return gerror.SUCCESS
}

func CloseConn(ip string, userId string, deviceType int) gerror.Result {
	conn, err := grpc.Dial(getRpcAddr(ip), grpc.WithInsecure(), grpc.WithUnaryInterceptor(intercepter.Interceptor))
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	imRpcClient := pb.NewImServerClient(conn)
	in := &pb.CloseConnReq{
		DeviceType: int32(deviceType),
		UserId:     userId,
	}
	_, _ = imRpcClient.CloseConn(context.Background(), in)
	return gerror.SUCCESS
}
