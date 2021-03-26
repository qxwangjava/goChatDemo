package main

import (
	"goChatDemo/config"
	"goChatDemo/internal/manager"
	"goChatDemo/internal/tcp_conn"
	"goChatDemo/internal/websocket"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/rpc"
	"goChatDemo/pkg/web"
)

func main() {
	db.InitDB(config.DbConfig.DbUrl)

	db.InitRedisClient(config.DbConfig.RedisUrl, "")

	rpc.InitRpc()

	rpc.InitUserServiceClient()

	tcp_conn.InitTCPServer()

	websocket.InitWSServer()

	web.InitWeb()

	manager.InitConnMap()

	select {}

}
