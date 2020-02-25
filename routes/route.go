package routes

import (
	"github.com/gin-gonic/gin"
	v1 "src/routes/v1"
)

func Routers(router *gin.Engine) {
	version1 := router.Group("/v1")
	var _ = router.Group("/v2")
	v1.User(version1)
}
