package utils

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type Logger struct {
}

var ginLogger *logrus.Logger
var gormLogger *logrus.Logger

//var file *os.File

func init() {
	initGinLogger()
	initGormLogger()
}

func initGinLogger() {
	ginLogger = logrus.New()
	ginLogger.Formatter.(*logrus.TextFormatter).TimestampFormat = "2006-01-02 15:04:05"
	var logPath = "./logs" // 日志打印到指定的目录
	// 目录不存在则创建
	if _, err := PathExists(logPath); err != nil {
		_ = os.MkdirAll(logPath, os.ModePerm)
	}
	ginLogger.SetLevel(logrus.DebugLevel)
	fileName := path.Join(logPath, "gin-api.log")
	//禁止logrus的输出
	//var err error
	//file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	// 设置日志输出的路径
	//logger.Out = file
	//apiLogPath := "gin-api.log"
	logWriter, _ := rotatelogs.New(
		fileName+".%Y-%m-%d",
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	ginLogger.AddHook(lfHook)
}

func initGormLogger() {
	gormLogger = logrus.New()
	gormLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    true,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
	var logPath = "./logs" // 日志打印到指定的目录
	// 目录不存在则创建
	if _, err := PathExists(logPath); err != nil {
		_ = os.MkdirAll(logPath, os.ModePerm)
	}
	gormLogger.SetLevel(logrus.DebugLevel)
	fileName := path.Join(logPath, "gorm.log")
	logWriter, _ := rotatelogs.New(
		fileName,
		//rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		//rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		//rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	gormLogger.AddHook(lfHook)
}

func (log Logger) SetGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		_path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// 这里是指定日志打印出来的格式。分别是状态码，执行时间,请求ip,请求方法,请求路由(等下我会截图)
		ginLogger.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, _path,
		)
	}
}

func (log Logger) GinLogger() *logrus.Logger {
	return ginLogger
}

func (log Logger) GormLogger() *logrus.Logger {
	return gormLogger
}

//func (log Logger) Close() {
//	_ = file.Close()
//}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
