package cache

import (
	"encoding/json"
)

//Author: zfz
//Date: 2020/2/24
//Description: 封装Get
//keyInfo：包含name和expire
//val：绑定的返回值
//f：缓存不存在，去数据库查询，并设置缓存
func Get(keyInfo KeyInfo, val interface{}, f func() interface{}) {
	value := cache.Client.Get(keyInfo.GetKey()).Val()
	if value == "" {
		// 此处会将无效查询结果也写入缓存，避免重复去数据库查询
		v, _ := json.Marshal(f())
		cache.Client.Set(keyInfo.GetKey(), v, keyInfo.GetExpire())
	} else {
		_ = json.Unmarshal([]byte(value), val)
	}
}
