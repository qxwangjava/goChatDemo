package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ttjio/gorm-logrus"
	"goChatDemo/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var RedisClient *redis.Client
var Ctx context.Context

func InitDB(dbUrl string) {
	logger.Logger.Info("init DB begin")
	var err error
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: gorm_logrus.New(logger.Logger),
	})
	if err != nil {
		logger.Logger.Error("数据库连接失败")
		return
	}
	logger.Logger.Info("init DB success")
}

func InitRedisClient(addr, password string) {
	logger.Logger.Info("init redis begin")
	Ctx = context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	logger.Logger.Info("init redis success")
}
