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
	WebSocketAddr string
}

type RPCConfig struct {
	RpcAddr string
	RpcPort string
}

func init() {
	initDevConf()
}
