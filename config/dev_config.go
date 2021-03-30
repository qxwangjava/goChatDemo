package config

import (
	"goChatDemo/pkg/util"
)

func initDevConf() {
	DbConfig = DBConfig{
		DbUrl:    "root:root@tcp(127.0.0.1:3306)/demo?parseTime=true",
		RedisUrl: "127.0.0.1:6379",
	}

	WebConfig = WEBConfig{
		WebPort:       ":8080",
		WebSocketAddr: util.GetServerIp() + ":8081",
	}

	RpcConfig = RPCConfig{
		RpcAddr: util.GetServerIp(),
		RpcPort: "50052",
	}

}
