package main

import (
	"goChatDemo/config"
	"goChatDemo/internal/emqt"
	"goChatDemo/internal/manager"
	"goChatDemo/internal/tcp_conn"
	"goChatDemo/internal/ws_conn"
	"goChatDemo/pkg/db"
	"goChatDemo/pkg/rpc"
	"goChatDemo/pkg/web"
)

func main() {

	db.InitDB(config.DbConfig.DbUrl)

	db.InitRedisClient(config.DbConfig.RedisUrl, "")

	emqt.InitEmqt()

	rpc.InitRpc()

	rpc.InitUserServiceClient()

	tcp_conn.InitTCPServer()

	web.InitWeb()

	manager.InitConnMap()

	ws_conn.InitWSServer()

	select {}

}
