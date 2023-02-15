package service

import (
	"common/Redis"
	"fmt"
)

var RdLikeUID *Redis.Redis //key:uid,value:vid
var RdLikeVID *Redis.Redis //key:vid,value:uid

func InitRD() {
	RdLikeUID = Redis.NewRedis("8.130.28.213", "6379", "2")
	RdLikeVID = Redis.NewRedis("8.130.28.213", "6379", "3")
	if RdLikeUID == nil || RdLikeVID == nil {
		fmt.Println("Redis 连接失败")
		return
	}
}
