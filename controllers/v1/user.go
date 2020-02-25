package v1

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/cache"
	"src/global"
	ErrHandle "src/middleware/errHandle"
	"src/middleware/resp"
	"src/middleware/validator"
	"src/models/entity"
	"src/services"
)

//获取通过uid用户信息
var GetUser = func(ctx *gin.Context) {
	getUser := validator.GetUserUri{}
	validator.Validate(ctx, &getUser, nil, nil)
	user := services.GetUser(getUser.Uid)
	if user != (entity.User{}) {
		resp.SetResponse(ctx, http.StatusOK, resp.Response{Code: 0, Msg: "", Data: user})
	} else {
		resp.SetResponse(ctx, http.StatusOK, resp.Response{Code: 1, Msg: "当前用户信息不存在"})
	}
}

//获取当前用户的信息
var GetMySelf = func(ctx *gin.Context) {
	loginUser := make(map[string]interface{})
	global.GetFromSession(ctx, "loginUser", &loginUser)
	user := services.GetUser(uint(loginUser["id"].(float64)))
	if user != (entity.User{}) {
		resp.SetResponse(ctx, http.StatusOK, resp.Response{Code: 0, Msg: "", Data: user})
	} else {
		resp.SetResponse(ctx, http.StatusOK, resp.Response{Code: 1, Msg: "当前用户信息不存在"})
	}
}

//创建用户
var UserRegister = func(ctx *gin.Context) {
	createUser := validator.RegisterForm{}
	validator.Validate(ctx, nil, nil, &createUser)
	m := md5.New()
	m.Write([]byte(createUser.Password))
	user := entity.User{
		Username: createUser.Username,
		Password: hex.EncodeToString(m.Sum(nil)),
		Mobile:   createUser.Mobile,
		Email:    createUser.Email,
	}
	if err := services.CreateUser(&user); err != nil {
		// 向日志写入具体错误
		global.Logger.Error("CreateUserError", err)
		// 统一错误处理
		panic(ErrHandle.RegisterError)
	}
	resp.SetResponse(ctx, 200, resp.Response{
		Code: 0,
		Msg:  "创建成功",
		Data: user,
	})
}

//登录
var UserLogin = func(ctx *gin.Context) {
	loginUser := new(validator.LoginForm)
	validator.Validate(ctx, nil, nil, loginUser)
	m := md5.New()
	m.Write([]byte(loginUser.Password))
	if userInfo, err := services.UserLogin(loginUser.Username, hex.EncodeToString(m.Sum(nil))); err != nil {
		// 向日志写入具体错误
		global.Logger.Error("LoginError", err)
		// 统一错误处理
		panic(ErrHandle.LoginError)
	} else {
		user := map[string]interface{}{
			"id":       userInfo.ID,
			"username": userInfo.Username,
		}
		global.SetSession(ctx, "loginUser", user)
		resp.SetResponse(ctx, 200, resp.Response{
			Code: 0,
			Msg:  "登录成功",
			Data: userInfo,
		})
	}

}

//测试Redis
var TestRedis = func(ctx *gin.Context) {
	ctx.AsciiJSON(http.StatusOK, map[string]string{
		"key":   "test1",
		"value": cache.CacheClient().Client.Get("test1").Val(),
	})
}
