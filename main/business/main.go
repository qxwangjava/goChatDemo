package main

import (
	"goChatDemo/config"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/rpc"
)

func main() {
	db.InitDB(config.DbConfig.DbUrl)

	db.InitRedisClient(config.DbConfig.RedisUrl, "")

	rpc.InitRpc()
	select {}

}
