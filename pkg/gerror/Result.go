package gerror

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
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) Result {
	return Result{
		Success: true,
		Code:    CODE_SUCCESS,
		Message: CODE_SUCCESS_MESSAGE,
		Data:    data,
	}
}

func ErrorMsg(msg string) Result {
	return Result{
		Success: false,
		Code:    CODE_FAIL,
		Message: msg,
		Data:    nil,
	}
}

func ErrorCode(code string, msg string) Result {
	return Result{
		Success: false,
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}
