package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goChatDemo/pkg/gerror"
	"goChatDemo/pkg/logger"
	"net/http"
	"strings"
)

type WrapperHandle func(c *gin.Context) error

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handle(c)
		if err != nil {
			logger.Logger.Error("Error occurred:", fmt.Sprintf("%+v", err))
			errMsg := strings.Split(err.Error(), "desc = ")[1]
			r := gerror.ErrorMsg(errMsg)
			c.JSON(http.StatusOK, r)
		}
	}
}
