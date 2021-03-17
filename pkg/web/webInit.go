package web

import (
	"github.com/gin-gonic/gin"
	"goChatDemo/internal/business/api"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/rpc"
	"net/http"
)

func ServiceWithAuth(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		r := rpc.ErrorMsg("需要认证")
		c.JSON(http.StatusOK, r)
		c.Abort()
	} else {
		c.Next()
	}

}

func ServiceWithoutAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "这是一个不用经过认证就能访问的接口"})
}

func InitWeb() {
	go func() {
		router := gin.Default()
		router.Use(ServiceWithAuth)
		router.POST("/user/add", gerror.ErrorWrapper(api.AddUser))
		err := router.Run(":8080")
		if err != nil {
			logger.Logger.Error("初始化web失败==>", err)
			panic(err)
		}
	}()
	logger.Logger.Info("初始化web 成功，监听端口 8080")
}
