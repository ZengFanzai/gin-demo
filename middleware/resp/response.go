package resp

import (
	"github.com/gin-gonic/gin"
	"src/global"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SetResponse(ctx *gin.Context, HttpCode int, resp Response) {
	if resp != (Response{}) {
		ctx.AsciiJSON(HttpCode, resp)
	} else {
		global.Logger.Panic("响应体设置无效")
	}
}
