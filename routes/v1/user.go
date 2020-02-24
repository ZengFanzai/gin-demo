package v1

import (
	"github.com/gin-gonic/gin"
	user "src/controllers/v1"
	"src/middleware"
)

func User(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	router.POST("/user/login", middleware.AlreadyLogin, user.UserLogin)
	router.POST("/user/register", user.UserRegister)
	router.GET("/user", middleware.NeedLogin, user.GetMySelf)
	router.GET("/user/:uid", user.GetUser)
	router.GET("/testredis", user.TestRedis)
}
