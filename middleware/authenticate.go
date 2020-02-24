package middleware

import (
	"github.com/gin-gonic/gin"
	"src/global"
	ErrHandle "src/middleware/errHandle"
)

var AlreadyLogin = func(ctx *gin.Context) {
	loginUser := make(map[string]interface{})
	global.GetFromSession(ctx, "loginUser", &loginUser)
	if v, ok := loginUser["id"]; ok {
		if v != 0 {
			panic(ErrHandle.AlreadyLogin)
		}
	}
}

var NeedLogin = func(ctx *gin.Context) {
	loginUser := make(map[string]interface{})
	global.GetFromSession(ctx, "loginUser", &loginUser)
	if uint(loginUser["id"].(float64)) == 0 {
		panic(ErrHandle.NotLogin)
	}
}
