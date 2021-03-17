package tcp_conn

import (
	"bufio"
	"fmt"
	"goChatDemo/pkg/logger"
	"net"
)

func InitTCPServer() {
	go func() {
		listen, err := net.Listen("tcp", "127.0.0.1:20000")
		if err != nil {
			logger.Logger.Error("listen failed, err:", err)
			return
		}
		logger.Logger.Info("TCP lisetn Success port 20000")
		for {
			conn, err := listen.Accept() // 建立连接
			logger.Logger.Info("建立连接成功。")
			if err != nil {
				logger.Logger.Error("accept failed, err:", err)
				continue
			}
			go process(conn) // 启动一个goroutine处理连接
		}
	}()

}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		_, _ = conn.Write([]byte(recvStr)) // 发送数据
	}
}
