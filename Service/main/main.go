package main

import (
	"common/consul"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/web"
	"log"
	"rpc/rpc/Test"
)

func main() {
	//创建gin
	router := gin.Default()
	//服务接口
	v1Group := router.Group("/v1")
	{
		v1Group.Handle("POST", "/prod", func(context *gin.Context) {
			var req Test.Request
			err := context.BindJSON(&req)
			if err != nil {
				return
			}
			fmt.Println(req)
			context.JSON(200, gin.H{
				"a": "b",
			})
		})

	}
	/**
		创建服务
	**/
	var server = web.NewService(
		web.Name("Test1"),                //服务名
		web.Address(":8001"),             //服务地址
		web.Handler(router),              //gin服务
		web.Registry(consul.GetConsul()), //注册中心
	)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

}
