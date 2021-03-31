package main

import (
	"goChatDemo/config"
	"goChatDemo/internal/business/rpc_server"
	"goChatDemo/internal/tcp_conn"
	"goChatDemo/internal/web"
	"goChatDemo/internal/ws_conn"
	"goChatDemo/pkg/db"
)

func main() {
	db.InitDB(config.DbConfig.DbUrl)

	db.InitRedisClient(config.DbConfig.RedisUrl, "")

	rpc_server.InitRpc()

	tcp_conn.InitTCPServer()

	ws_conn.InitWSServer()

	web.InitWeb()

	ws_conn.InitConnMap()

	select {}

}
