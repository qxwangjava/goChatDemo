package main

import (
	"goChatDemo/config"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/rpc"
)

func main() {
	db.InitDB(config.DbConfig.DbUrl)

	db.InitRedisClient(config.DbConfig.RedisUrl, "")
	err := db.RedisClient.Set(db.Ctx, "key1", "value1", 0).Err()
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	val, getErr := db.RedisClient.Get(db.Ctx, "key1").Result()
	if getErr != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	logger.Logger.Info("val:", val)

	rpc.InitRpc()

	//新增
	//userId, _ := dao.UserDao.Add(&dao.User{
	//	UserName:  "wangqx1",
	//	NickName:  "王奇轩",
	//	Mobile:    "17610565751",
	//	IsDeleted: 0,
	//})
	//logger.Logger.WithField("userId", userId).Info("保存成功")
	//根据主键删除
	//userId = 4
	//dao.UserDao.Del(4)
	//logger.Logger.WithField("userId", userId).Info("删除成功")
	//根据条件删除
	//user := dao.User{
	//	IsDeleted: 1,
	//}
	//dao.UserDao.DelByCondition(&user)
	//logger.Logger.WithField("user", user).Info("删除成功")
}
