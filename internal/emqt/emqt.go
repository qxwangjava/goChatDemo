package emqt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goChatDemo/internal/business/service"
	"goChatDemo/pkg/logger"
	"time"
)

const (
	SENDMESSAGE = "topic/sendMessage"
)

var Client mqtt.Client

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logger.Logger.Info("TOPIC:", msg.Topic())
	logger.Logger.Info("MSG: ", string(msg.Payload()))
}

var sendMessage = func(c mqtt.Client, message mqtt.Message) {
	service.SendMessage("2", message.Payload())
}

func subscribe() {
	//订阅主题
	if token := Client.Subscribe(SENDMESSAGE, 0, nil); token.Wait() && token.Error() != nil {
		logger.Logger.Error("订阅", SENDMESSAGE, "主题失败:", token.Error())
		return
	}
	Client.AddRoute(SENDMESSAGE, sendMessage)
	logger.Logger.Info("订阅", SENDMESSAGE, "主题成功！")
}

func InitEmqt() {
	go func() {
		opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:11883").SetClientID("first_client")
		opts.SetUsername("127.0.0.1")
		opts.SetKeepAlive(60 * time.Second)
		// 设置消息回调处理函数
		opts.SetDefaultPublishHandler(f)
		opts.SetPingTimeout(1 * time.Second)

		Client = mqtt.NewClient(opts)
		if token := Client.Connect(); token.Wait() && token.Error() != nil {
			logger.Logger.Error("emqtt初始化失败:", token.Error())
		}
		logger.Logger.Info("emqtt初始化成功！")
		subscribe()

		// 发布消息
		//token := c.Publish("testtopic/1", 0, false, "Hello World")
		//token.Wait()

		//time.Sleep(6 * time.Second)

		// 取消订阅
		//if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		//	fmt.Println(token.Error())
		//	os.Exit(1)
		//}

		// 断开连接
		//c.Disconnect(250)
		//time.Sleep(1 * time.Second)
	}()
}
