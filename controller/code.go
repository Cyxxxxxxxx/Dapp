package controller

//定义可能出现的错误码

type ResCode int64

const (
	CodeSuccess ResCode = 200 + iota
	CodeServerBusy
	CodeErrPermission
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:       "success",
	CodeServerBusy:    "服务繁忙",
	CodeErrPermission: "权限不足",
}

// Msg返回特定的错误提示信息
func (c ResCode) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		return CodeMsgMap[CodeServerBusy]
	}
	return msg
}
