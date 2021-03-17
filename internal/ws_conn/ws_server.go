package ws_conn

import (
	"github.com/gorilla/websocket"
	"goChatDemo/pkg/logger"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	logger.Logger.Info("建立连接成功")
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	handleConn(conn)

}

func InitWSServer(address string) {
	go func() {
		http.HandleFunc("/ws", wsHandler)
		err := http.ListenAndServe("127.0.0.1:8081", nil)
		if err != nil {
			logger.Logger.Error("webSocket 启动失败:", err)
		}

	}()
	logger.Logger.Info("webSocket启动成功,监听端口 8081")
}

func handleConn(conn *websocket.Conn) {
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			logger.Logger.Info("server read:", err)
			break
		}
		logger.Logger.Info("server recv: ", string(message))
		err = conn.WriteMessage(mt, message)
		logger.Logger.Info("server send: ", string(message))
		if err != nil {
			logger.Logger.Info("server write:", err)
			break
		}
	}
}
