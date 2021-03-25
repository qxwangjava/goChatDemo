package gerror

import "goChatDemo/pkg/logger"

func HandleError(err error) {
	if err != nil {
		logger.Logger.Error(err)
	}
}
