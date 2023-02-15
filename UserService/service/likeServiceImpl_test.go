package service

import (
	"UserService/dao"
	"common/RabbitMq"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"testing"
)

func TestIsLike(t *testing.T) {
	dao.InitDB()
	InitRD()
	impl := LikeServiceImpl{}
	isExist, err := impl.IsLike(3, 4)
	fmt.Println(isExist)
	fmt.Println(err)
}

func TestVideoLikeCount(t *testing.T) {
	dao.InitDB()
	InitRD()
	impl := LikeServiceImpl{}
	count, err := impl.VideoLikeCount(4)
	fmt.Println(count)
	fmt.Println(err)
}

func TestLikeListCount(t *testing.T) {
	dao.InitDB()
	InitRD()
	impl := LikeServiceImpl{}
	count, err := impl.LikeListCount(3)
	fmt.Println(count)
	fmt.Println(err)
}

func TestLikeAction(t *testing.T) {
	dao.InitDB()
	InitRD()
	InitMQ()
	impl := LikeServiceImpl{}
	err := impl.LikeAction(4, 8, 1)
	fmt.Println(err)
}

func TestRedis(t *testing.T) {
	InitRD()
	_, err := RdLikeVID.Set("6", "2")
	if err != nil {
		log.Println(err)
		return
	}
	isExist, err := redis.Bool(RdLikeUID.Server.Do("SIsMember", "3", "9"))
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println(isExist)
	}
	RdLikeUID.Server.Do("SADD", 3, 2)
	RdLikeUID.Server.Do("SADD", 3, 4)
}

func TestMQ(t *testing.T) {
	RabbitMq.NewRabbitMq("admin", "admin", "8.130.28.213", "5672")
}
