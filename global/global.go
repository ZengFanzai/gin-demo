package global

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"src/cache"
	"src/config"
	db "src/models"
	"src/utils"
)

//定义一些全局变量

var Config = config.GetConfig()
var DBConfig = config.GetDBSettings()
var DB = db.GetDB()
var Cache = cache.GetCacheClient()
var Logger = utils.Logger{}.GetGinLogger()
var DBLogger = utils.Logger{}.GetGormLogger()

var SetSession = func(ctx *gin.Context, key, value interface{}) {
	session := sessions.Default(ctx)
	if v, err := json.Marshal(value); err != nil {
		Logger.Error("SetSession->MarshalError->", err)
	} else {
		session.Set(key, v)
		if err := session.Save(); err != nil {
			Logger.Error("SetSession->SaveError->", err)
		}
	}
}

var GetFromSession = func(ctx *gin.Context, key, value interface{}) {
	session := sessions.Default(ctx)
	switch val := session.Get(key).(type) {
	case []byte:
		if err := json.Unmarshal(val, value); err != nil {
			Logger.Error("GetSession->Unmarshal->", err)
		}
	case nil:
	default:
		Logger.Error("GetSession->UnknownType->", val)
	}
	if err := session.Save(); err != nil {
		Logger.Error(err)
	}
}
