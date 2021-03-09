package config

func initDevConf() {
	DbConfig = DBConfig{
		DbUrl:    "root:root@tcp(127.0.0.1:3306)/demo?parseTime=true",
		RedisUrl: "127.0.0.1:6379",
	}
}
