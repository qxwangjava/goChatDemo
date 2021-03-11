package test

import (
	"goChatDemo/config"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/logger"
	"testing"
)

func TestRedis(T *testing.T) {
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
}
