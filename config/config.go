package config

var (
	DbConfig  DBConfig
	WebConfig WEBConfig
	RpcConfig RPCConfig
)

type DBConfig struct {
	DbUrl    string
	RedisUrl string
}

type WEBConfig struct {
	WebPort       string
	WebSocketPort string
}

type RPCConfig struct {
	RpcPort string
}

func init() {
	initConfig()
}
