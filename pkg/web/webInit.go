package web

import (
	"github.com/gin-gonic/gin"
	"goChatDemo/internal/business/api"
)

func InitWeb() {
	router := gin.Default()
	router.POST("/user/add", api.AddUser)
	router.Run(":8080")
}
