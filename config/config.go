package config

var (
	DbConfig DBConfig
)

type DBConfig struct {
	DbUrl    string
	RedisUrl string
}

func init() {
	initDevConf()
}
