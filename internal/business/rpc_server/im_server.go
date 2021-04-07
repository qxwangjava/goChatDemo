package rpc_server

import (
	"context"
	"encoding/json"
	"goChatDemo/internal/business/dao"
	"goChatDemo/internal/ws_conn"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"strconv"
)

type imService struct{}

type sendMessageStruct struct {
	FromId          string `json:"fromId"`
	FromType        int    `json:"fromType"`
	MessageContent  string `json:"messageContent"`
	MessageType     int    `json:"messageType"`
	MessageServerId int    `json:"messageServerId"`
	GroupId         int64  `json:"groupId"`
}

var ImService = imService{}

func (is imService) CloseConn(ctx context.Context, req *pb.CloseConnReq) (*pb.CloseConnRsp, error) {
	ws_conn.CloseConn(int(req.DeviceType), req.UserId)
	return &pb.CloseConnRsp{}, nil
}

func (is imService) SendMsg(ctx context.Context, req *pb.SendMsgReq) (*pb.SendMsgRsp, error) {
	key := strconv.Itoa(int(req.DeviceType)) + "_" + req.DeviceId + "_" + req.ToUserId
	value, ok := ws_conn.LocalConnInfoManager.Load(key)
	if ok {
		sms := sendMessageStruct{
			FromId:          req.FromId,
			FromType:        int(req.FromType),
			MessageType:     int(req.MessageType),
			MessageContent:  req.MessageContent,
			MessageServerId: int(req.MessageServerId),
			GroupId:         req.GroupId,
		}
		client := value.(ws_conn.Client)
		msg, err := json.Marshal(sms)
		if err != nil {
			logger.Logger.Error("json序列化失败", err)
		}
		client.WriteChan <- msg
		//判断消息类型
		//if req.MessageType == 1 { //文本
		//	client := value.(ws_conn.Client)
		//	msg, _ := json.Marshal(sms)
		//	client.WriteChan <- []byte(msg)
		//}
		//if req.MessageType == 2 { //图片
		//	client := value.(ws_conn.Client)
		//	client.WriteChan <- []byte(req.MessageContent)
		//}
		//更新消息发送结果为成功
		dao.MessageDao.UpdateMessageSendResult(req.MessageServerId)
	}
	return &pb.SendMsgRsp{}, nil
}
