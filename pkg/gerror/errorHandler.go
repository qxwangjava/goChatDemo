package gerror

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goChatDemo/pkg/logger"
	"goChatDemo/pkg/rpc"
	"net/http"
)

type WrapperHandle func(c *gin.Context) error

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handle(c)
		if err != nil {
			//err = errors.Wrap(err, "")
			logger.Logger.Error("Error occurred", fmt.Sprintf("%+v", err))
			//logger.Logger.Error(err)
			r := rpc.ErrorMsg(err.Error())
			c.JSON(http.StatusOK, r)
		}
	}
}
