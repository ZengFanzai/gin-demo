package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"src/config"
	"strconv"
)

type Cache struct {
	Client *redis.Client
}

var cache Cache

func init() {
	initRedis()
}

func initRedis() Cache {
	rdsConf := config.Config().GetStringMapString("redis")
	db, _ := strconv.Atoi(rdsConf["db"])
	cache.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdsConf["addr"], rdsConf["port"]),
		DB:       db,
		Password: rdsConf["password"],
	})
	return cache
}

func CacheClient() Cache {
	return cache
}

func Close() {
	cache.Client.Close()
}
