package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": //程序中的错误码
	"msg": //提示信息
	“data" : //数据
}
*/

type ResponseCode struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// ResponeseError 返回一个code类型的错误
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseCode{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
