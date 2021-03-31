package rpc_server

import (
	"context"
	"goChatDemo/internal/ws_conn"
	"goChatDemo/pkg/pb"
	"strconv"
)

type imService struct{}

var ImService = imService{}

func (is imService) CloseConn(ctx context.Context, req *pb.CloseConnReq) (*pb.CloseConnRsp, error) {
	ws_conn.CloseConn(int(req.DeviceType), req.UserId)
	return &pb.CloseConnRsp{}, nil
}

func (is imService) SendMsg(ctx context.Context, req *pb.SendMsgReq) (*pb.SendMsgRsp, error) {
	key := strconv.Itoa(int(req.DeviceType)) + "_" + req.DeviceId + "_" + req.UserId
	value, ok := ws_conn.LocalConnInfoManager.Load(key)
	if ok {
		//的判断消息类型
		if req.MessageType == 1 { //文本
			client := value.(ws_conn.Client)
			client.WriteChan <- []byte(req.MessageContent)
		}
		if req.MessageType == 2 { //图片
			client := value.(ws_conn.Client)
			client.WriteChan <- []byte(req.MessageContent)
		}

	}
	return &pb.SendMsgRsp{}, nil
}
