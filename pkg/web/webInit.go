package web

import (
	"github.com/gin-gonic/gin"
	"goChatDemo/internal/business/api"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
)

func InitWeb() {
	router := gin.Default()

	router.POST("/user/add", gerror.ErrorWrapper(api.AddUser))
	err := router.Run(":8080")
	if err != nil {
		logger.Logger.Error("初始化web失败==>", err)
		panic(err)
	}
}
