package gerror

import (
	"github.com/gin-gonic/gin"
	"goChatDemo/pkg/rpc"
	"net/http"
)

type WrapperHandle func(c *gin.Context) error

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handle(c)
		if err != nil {
			r := rpc.ErrorMsg(err.Error())
			c.JSON(http.StatusOK, r)
			panic(err)
		}
	}
}
