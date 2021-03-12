package rpc

const (
	CODE_SUCCESS         string = "0"
	CODE_SUCCESS_MESSAGE string = "success"
	CODE_FAIL            string = "-1"
	CODE_FAIL_MESSAGE    string = "服务器异常"
)

//成功结果
var SUCCESS = Result{
	Success: true,
	Code:    CODE_SUCCESS,
	Message: CODE_SUCCESS_MESSAGE,
	Data:    nil,
}

//失败结果
var ERROR = Result{
	Success: true,
	Code:    CODE_FAIL,
	Message: CODE_FAIL_MESSAGE,
	Data:    nil,
}

type Result struct {
	Success bool
	Code    string
	Message string
	Data    interface{}
}

func Success(data interface{}) Result {
	return Result{
		Success: true,
		Code:    CODE_SUCCESS,
		Message: CODE_SUCCESS_MESSAGE,
		Data:    data,
	}
}
