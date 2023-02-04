package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"log"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(registry.Addrs("8.130.28.213:8500"))
}
func main() {
	router := gin.Default()
	v1Group := router.Group("/v1")
	{
		v1Group.Handle("GET", "/prod", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"a": "b",
			})
		})

	}

	server := web.NewService(
		web.Name("Test1"),
		web.Address(":8001"),
		web.Handler(router),
		web.Registry(consulReg),
	)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
