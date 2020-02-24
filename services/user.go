package services

import (
	"fmt"
	"src/cache"
	"src/models"
	"src/models/entity"
)

var db = models.GetDB()

//author: zfz
//createAt: 2020/2/24
//description: 获取用户信息
func GetUser(uid uint) entity.User {
	var user entity.User
	cache.Get(cache.UserInfo(fmt.Sprint(uid)), &user, func() interface{} {
		db.First(&user, uid)
		return user
	})
	return user
}

func CreateUser(user *entity.User) error {
	return db.Create(user).Error
}

//用户登录
func UserLogin(username, password string) (entity.User, error) {
	var user entity.User
	err := db.Where("username=? and password=?", username, password).First(&user).Error
	return user, err
}
