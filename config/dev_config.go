package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"goChatDemo/pkg/logger"
	"strings"
	"time"
)

var (
	defaultConfig *viper.Viper
	nacosIp       string
	nacosPort     uint64
	nacosDataId   string
	nacosGroup    string
)

func initConfig() {
	nacosIp = "192.168.249.27"
	nacosPort = 8848
	nacosDataId = "goChat.yaml"
	nacosGroup = "dev"

	defaultConfig = viper.New()
	defaultConfig.SetConfigType("yaml")
	//配置模型
	serverConfigs := []constant.ServerConfig{
		{IpAddr: nacosIp, Port: nacosPort},
	}

	//客户端
	nacosClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				TimeoutMs:   5000,
				NamespaceId: "go",
			},
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		logger.Logger.Error("nacos初始化错误:", err)
	}

	content, err := nacosClient.GetConfig(vo.ConfigParam{DataId: nacosDataId, Group: nacosGroup})
	if err != nil {
		logger.Logger.Error("nacos读取配置错误:" + content)
	}

	err = defaultConfig.ReadConfig(strings.NewReader(content))
	if err != nil {
		logger.Logger.Error("Viper解析配置失败:", err)
	}

	go func() {
		time.Sleep(time.Second * 10)
		err = nacosClient.ListenConfig(vo.ConfigParam{
			DataId: nacosDataId,
			Group:  nacosGroup,
			OnChange: func(namespace, group, dataId, data string) {
				fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
				err = defaultConfig.ReadConfig(strings.NewReader(data))
				if err != nil {
					logger.Logger.Error("Viper解析配置失败:", err)
				}
			},
		})
	}()

	DbConfig = DBConfig{
		DbUrl:    defaultConfig.GetString("db.DbUrl"),
		RedisUrl: defaultConfig.GetString("db.RedisUrl"),
	}
	WebConfig = WEBConfig{
		WebPort:       defaultConfig.GetString("web.WebPort"),
		WebSocketPort: defaultConfig.GetString("web.WebSocketPort"),
	}
	RpcConfig = RPCConfig{
		RpcPort: defaultConfig.GetString("rpc.port"),
	}
}
