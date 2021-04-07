package config

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/pkg/errors"
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
				LogLevel:    "INFO",
				//LogDir: logger.InfoFileName,

			},
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		logger.Logger.Error("nacos初始化错误:", errors.Wrap(err, ""))
		return
	}

	content, err := nacosClient.GetConfig(vo.ConfigParam{DataId: nacosDataId, Group: nacosGroup})
	if err != nil {
		logger.Logger.Error("nacos读取配置错误:", err)
		return
	}

	err = defaultConfig.ReadConfig(strings.NewReader(content))
	if err != nil {
		logger.Logger.Error("Viper解析配置失败:", err)
		return
	}

	go func() {
		time.Sleep(time.Second * 5)
		err = nacosClient.ListenConfig(vo.ConfigParam{
			DataId: nacosDataId,
			Group:  nacosGroup,
			OnChange: func(namespace, group, dataId, data string) {
				logger.Logger.Info("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
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
