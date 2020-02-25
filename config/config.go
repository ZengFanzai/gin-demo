package config

import (
	"github.com/spf13/viper"
	"log"
)

var conf = viper.New()

func init() {
	initConfig()
}

func initConfig() {
	conf.SetConfigType("yaml")
	conf.SetConfigName("conf")
	conf.AddConfigPath("./config")
	if err := conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
	//conf.WatchConfig()
	//conf.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})
}

func Config() *viper.Viper {
	return conf
}

type Database struct {
	Addr     string
	Port     int
	UserName string
	Password string
	DBName   string
	Charset  string
}

var db *Database

func DBSettings() *Database {
	if err := conf.UnmarshalKey("database", &db); err != nil {
		log.Fatal("get DB settings error=>", err)
	}
	return db
}
