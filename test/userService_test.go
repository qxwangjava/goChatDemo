package test

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/pb"
	"goChatDemo/pkg/rpc"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestUserRpc(t *testing.T) {
	conn, err := grpc.Dial(rpc.RpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	birthday, err := time.ParseInLocation("2006-01-02 15:04:05", "1993-01-02 00:00:00", time.Local)
	if err != nil {
		logger.Logger.Error(err)
	}
	b, _ := ptypes.TimestampProto(birthday)
	response, err := c.AddUser(context.Background(), &pb.AddUserDto{
		UserName: "wqx",
		NickName: "王奇轩",
		Mobile:   "17610565751",
		Sex:      1,
		Birthday: b,
		HeadImg:  "123465",
		Email:    "3940@123",
	})
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	logger.Logger.Info(response)
}
