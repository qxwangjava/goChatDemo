package db

import (
	"github.com/go-redis/redis/v8"
	"goChatDemo/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitMysql(dbUrl string) {
	logger.Logger.Info("init DB begin")
	var err error
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		logger.Logger.Error("数据库连接失败")
		return
	}
	logger.Logger.Info("init DB success")
}
