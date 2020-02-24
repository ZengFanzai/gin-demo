package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"src/cache"
	"src/global"
	ErrHandle "src/middleware/errHandle"
	customValidator "src/middleware/validator"
	db "src/models"
	"src/routes"
	"src/utils"
	"time"
)

func main() {
	defer func() {
		fmt.Println("关闭相关连接")
		db.Close()
		cache.Close()
	}()
	//gin.DefaultWriter = utils.Logger{}.GetGinLogger().Writer()
	//TODO: XSS和CSRF防御
	server := gin.New()
	server.Use(gin.Recovery())
	// 自定义参数验证
	customValidator.RegisterCustomValidation()
	server.Use(utils.Logger{}.SetGinLogger())
	session := global.Config.GetStringMapString("session")
	store, _ := redis.NewStoreWithDB(1, session["network"], fmt.Sprintf("%s:%s", session["addr"], session["port"]), session["password"], session["db"], []byte("secret"))
	store.Options(sessions.Options{
		Path:     "",
		Domain:   "",
		MaxAge:   7 * 24 * 3600,
		Secure:   false,
		HttpOnly: true,
		SameSite: 0,
	})
	server.Use(sessions.Sessions("gin-session", store))
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(limits.RequestSizeLimiter(10 * 1024))
	//server.Use(func(context *gin.Context) {
	//	global.SetSession(context, "user", "root")
	//})
	server.Use(ErrHandle.ErrHandler())
	server.NoMethod(ErrHandle.HandleNotFound)
	server.NoRoute(ErrHandle.HandleNotFound)
	server.Static("/assets", "./assets")
	routes.Routers(server)

	// 优雅的关闭服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", global.Config.GetString("server.port")),
		Handler: server,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	//server.Run(":" + conf.GetString("server.port")) // 监听并在 0.0.0.0:8080 上启动服务
}
