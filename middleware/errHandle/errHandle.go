package ErrHandle

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"src/global"
)

// 错误处理的结构体
type Error struct {
	StatusCode int         `json:"-"`
	Code       int         `json:"code"`
	Msg        interface{} `json:"msg"`
}

var (
	ServerError   = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	NotFound      = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	RegisterError = NewError(http.StatusOK, 2, "用户注册失败!")
	LoginError    = NewError(http.StatusOK, 3, "用户登录失败!")
	NotLogin      = NewError(http.StatusOK, 4, "用户未登录!")
	AlreadyLogin  = NewError(http.StatusOK, 5, "用户已登录!")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	switch e.Msg.(type) {
	case string:
		return e.Msg.(string)
	case []string:
		v, _ := json.Marshal(e.Msg)
		return string(v)
	}
	return "解析错误异常"
}

//new Error
func NewError(statusCode, Code int, msg interface{}) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode, err)
	c.Abort()
}

//author: zfz
//createAt: 2020/2/24
//description: 统一错误处理器，捕获request处理中的错误，并向用户统一返回相关信息
func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *Error
				if e, ok := err.(*Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = OtherError(e.Error())
				} else {
					Err = ServerError
					global.Logger.Error(string(debug.Stack()))
				}
				// 记录一个错误的日志
				c.JSON(Err.StatusCode, Err)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
