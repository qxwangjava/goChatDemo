package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goChatDemo/internal/emqt"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var sendMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	fmt.Printf("111: %s\n", msg.Payload())
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:11883").SetClientID("first_client")

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("emqtt初始化失败:", token.Error())
	}
	fmt.Println("emqtt初始化成功！")
	//发布消息
	//data := "{\"otherSideId\":\"2\",\"messageType\":1,\"otherSideType\":1,\"action\":\"sendMessage\",\"text\":\"大家好\"}"
	//if token := client.Publish(emqt.SENDMESSAGE, 2, false,data); token.Wait() && token.Error() != nil {
	//	fmt.Println("发布消息失败：", token.Error())
	//}
	//fmt.Println("发布消息成功！")
	//订阅消息
	if token := client.Subscribe(emqt.SENDMESSAGE, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println("订阅", emqt.SENDMESSAGE, "主题失败:", token.Error())
	}
	client.AddRoute(emqt.SENDMESSAGE, sendMessage)
	fmt.Println("订阅", emqt.SENDMESSAGE, "主题成功！")
	select {}
}
