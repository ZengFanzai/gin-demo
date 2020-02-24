# Gin-Demo

当前 go 环境为 1.13，gin 版本为 1.5.0

Demo 基于 Gin 框架进行二次开发，集成以下功能：

- 配置文件读取
- 日志记录
- ORM
- Redis 缓存
- 统一错误处理
- session 管理
- request 参数验证

TODO：web 安全处理，xss，csrf，SQL 注入(对原生 SQL)

**相关依赖**：

| 功能             | 依赖包                                                                  |
| ---------------- | ----------------------------------------------------------------------- |
| 配置文件读取     | [vpier](https://github.com/spf13/viper)                                 |
| 日志记录         | [logrus](https://github.com/sirupsen/logrus)                            |
| ORM              | [gorm](https://github.com/jinzhu/gorm")                                 |
| request 参数验证 | [validator.v9](https://github.com/go-playground/validator/tree/v9.31.0) |
| Redis 缓存       | [go-redis](https://github.com/go-redis/redis)                           |
| session 管理     | [gin-contrib\/sessions](https://github.com/gin-contrib/sessions)        |
