package web

const (
	CODE_SUCCESS         string = "0"
	CODE_SUCCESS_MESSAGE string = "success"
	CODE_FAIL            string = "-1"
	CODE_FAIL_MESSAGE    string = "服务器异常"
)

//成功结果
var SUCCESS = Result{
	success: true,
	code:    CODE_SUCCESS,
	message: CODE_SUCCESS_MESSAGE,
	data:    nil,
}

//失败结果
var ERROR = Result{
	success: true,
	code:    CODE_FAIL,
	message: CODE_FAIL_MESSAGE,
	data:    nil,
}

type Result struct {
	success bool
	code    string
	message string
	data    interface{}
}
