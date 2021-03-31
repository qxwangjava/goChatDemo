package main

import (
	"bufio"
	"github.com/gorilla/websocket"
	"goChatDemo/pkg/logger"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//var serverIp = "192.168.122.129:8081"
	var serverIp = "192.168.249.27:8081"

	var requestHeader = http.Header{}
	token := []string{"1|1|3"}
	requestHeader["token"] = token
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+serverIp+"/ws", requestHeader)
	if err != nil {
		logger.Logger.Error("连接失败:", err)
	}
	defer conn.Close()
	var ch1 = make(chan int)
	defer close(ch1)
	//发送心跳
	go func() {
		i := 0
		for {
			<-time.After(time.Second)
			err := conn.WriteMessage(websocket.PingMessage, []byte{})
			//logger.Logger.Info("发送ping消息成功")
			if err != nil {
				i++
				if i >= 5 {
					logger.Logger.Error("服务器连接失败，客户端已掉线:", err)
					ch1 <- 1
					break
				}
			}
		}
	}()

	//客户端接收消息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				logger.Logger.Error("读取失败:", err)
				break
			}
			logger.Logger.Info("client recv: ", string(message))
		}
	}()

	//客户端发送消息
	go func() {
		inputReader := bufio.NewReader(os.Stdin)
		for {
			//data, _ := <-ch1
			//if data == 1 {
			//	break
			//}
			input, _ := inputReader.ReadString('\n') // 读取用户输入
			inputInfo := strings.Trim(input, "\r\n")
			if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, []byte(inputInfo)) // 发送数据
			logger.Logger.Info("client send: ", inputInfo)
		}
	}()
	logger.Logger.Info(token, ":客户端启动成功")
	select {
	case data := <-ch1:
		if data == 1 {
			conn.Close()
		}
	}

}
