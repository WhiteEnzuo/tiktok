package main

import (
	"common/call"
	"common/config"
	"fmt"
	"rpc/rpc/Test"
)

type ServerConfig struct {
	Host string
	Port string
}

func main() {
	//读配置
	var s ServerConfig
	err1 := config.ReadConfig("server", &s)
	if err1 != nil {
		return
	}
	//fmt.Println(s.Port)
	//Response
	t := new(Test.Response)
	var t1 Test.Request
	t1.Name = "aaa"
	//服务调用
	err := call.Call("Test1", "/v1/prod", t1, t)
	fmt.Println(t.Test)
	if err != nil {
		fmt.Println(err)
	}
}
