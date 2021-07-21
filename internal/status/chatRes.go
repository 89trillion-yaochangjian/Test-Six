package status

type Err struct {
	Code int    `json:"code"` // 错误码
	Msg  string `json:"msg"`  // 错误描述
}

var CheckPra = Err{1001, "Please check the parameters"}
var RepeatCon = Err{1001, "Please check the parameters"}
var FisCon = Err{1001, "Please check the parameters"}
var SignIn = Err{1001, "Please check the parameters"}
