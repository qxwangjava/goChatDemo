package web

import (
	"github.com/gin-gonic/gin"
	"goChatDemo/config"
	"goChatDemo/internal/business/api"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
	"net/http"
)

func ServiceWithAuth(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		r := gerror.ErrorMsg("需要认证")
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
		router.POST("/user/add", ErrorWrapper(api.AddUser))
		// 该拦截器以后内容需要拦截器过滤
		router.Use(ServiceWithAuth)
		err := router.Run(config.WebConfig.WebPort)
		if err != nil {
			logger.Logger.Error("初始化web失败==>", err)
			panic(err)
		}
	}()
	logger.Logger.Info("初始化web成功，监听端口:", config.WebConfig.WebPort)
}
