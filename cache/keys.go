package cache

import (
	"time"
)

type KeyInfo interface {
	GetKey() string
	GetExpire() time.Duration
}

type UserinfoS struct {
	Key    string
	Expire time.Duration
}

func (user *UserinfoS) GetKey() string {
	return user.Key
}

func (user *UserinfoS) GetExpire() time.Duration {
	return user.Expire
}

var user = UserinfoS{Key: "", Expire: 1 * time.Minute}

func UserInfo(key string) *UserinfoS {
	user.Key = "userInfo:" + key
	return &user
}
