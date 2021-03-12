package main

import (
	"goChatDemo/config"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/rpc"
	"goChatDemo/pkg/web"
)

func main() {
	db.InitDB(config.DbConfig.DbUrl)

	db.InitRedisClient(config.DbConfig.RedisUrl, "")

	rpc.InitRpc()

	rpc.InitUserServiceClient()

	web.InitWeb()
	select {}

}
