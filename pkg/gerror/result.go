package gerror

const (
	CodeSuccess        string = "0"
	CodeSuccessMessage string = "success"
	CodeFail           string = "-1"
	CodeFailMessage    string = "服务器异常"
)

//成功结果
var SUCCESS = Result{
	Success: true,
	Code:    CodeSuccess,
	Message: CodeSuccessMessage,
	Data:    nil,
}

//失败结果
var ERROR = Result{
	Success: true,
	Code:    CodeFail,
	Message: CodeFailMessage,
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
		Code:    CodeSuccess,
		Message: CodeSuccessMessage,
		Data:    data,
	}
}

func ErrorMsg(msg string) Result {
	return Result{
		Success: false,
		Code:    CodeFail,
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
