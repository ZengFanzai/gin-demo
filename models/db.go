package models

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "gopkg.in/yaml.v2"
	"src/config"
	"src/utils"
	"text/template"
)

type DB = *gorm.DB

var db DB

func init() {
	initDB()
}

func initDB() {
	dbLogger := utils.Logger{}.GetGormLogger()
	url := "{{.UserName}}:{{.Password}}@tcp({{.Addr}})/{{.DBName}}?charset={{.Charset}}&parseTime=true"
	dbTemplate, _ := template.New("dbUrl").Parse(url)
	dbconf := bytes.Buffer{}
	err := dbTemplate.Execute(&dbconf, config.GetDBSettings())
	if err != nil {
		dbLogger.Panic("配置模板失败->", err)
	}
	db, err = gorm.Open("mysql", dbconf.String())
	if err != nil {
		dbLogger.Panic("数据库连接失败->", err)
	}
	db.DB().SetMaxOpenConns(200)
	db.DB().SetMaxIdleConns(40)
	db.SingularTable(true)
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tbl_" + defaultTableName
	}
	db.LogMode(true)
	db.SetLogger(dbLogger)
}

func GetDB() DB {
	return db
}

func Close() {
	_ = db.Close()
}
