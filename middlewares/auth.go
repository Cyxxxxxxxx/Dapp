package middlewares

import (
	"Dapp/controller"
	"Dapp/setting"
	"strings"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

// 判断管理后台是否有超管权限
func JudgePermissionWeb(c *gin.Context) {
	userid := c.Query("userid")

	if userid == "" {
		controller.ResponseError(c, controller.CodeErrPermission)
		c.Abort()
		return
	}

	//获得授权管理员列表
	authArray := strings.Split(setting.Conf.AuthID, ";")

	//循环查找是否有管理员权限
	for _, value := range authArray {
		if value == userid {
			c.Next()
			return
		}
	}

	controller.ResponseError(c, controller.CodeErrPermission)
	c.Abort()
	return
}
