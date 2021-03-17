package main

import (
	"bufio"
	"github.com/gorilla/websocket"
	"goChatDemo/pkg/logger"
	"os"
	"strings"
)

func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:20000")
	//if err != nil {
	//	fmt.Println("err :", err)
	//	return
	//}
	//inputReader := bufio.NewReader(os.Stdin)
	//for {
	//	input, _ := inputReader.ReadString('\n') // 读取用户输入
	//	inputInfo := strings.Trim(input, "\r\n")
	//	if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
	//		return
	//	}
	//	_, err = conn.Write([]byte(inputInfo)) // 发送数据
	//	if err != nil {
	//		return
	//	}
	//	buf := [512]byte{}
	//	n, err := conn.Read(buf[:])
	//	if err != nil {
	//		fmt.Println("recv failed, err:", err)
	//		return
	//	}
	//	fmt.Println(string(buf[:n]))
	//}
	//=================================================

	//u := url.URL{Scheme: "ws", Host: "127.0.0.1:8081", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8081/ws", nil)
	if err != nil {
		logger.Logger.Error("连接失败:", err)
	}
	defer c.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		err = c.WriteMessage(websocket.TextMessage, []byte(inputInfo)) // 发送数据
		logger.Logger.Info("client send: ", inputInfo)
		_, message, err := c.ReadMessage()
		if err != nil {
			logger.Logger.Error("读取失败:", err)
			return
		}
		logger.Logger.Info("client recv: ", string(message))

	}

}
