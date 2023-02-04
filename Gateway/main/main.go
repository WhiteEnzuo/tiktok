package main

import (
	"common/call"
	"fmt"
	"rpc/rpc/Test"
)

func main() {
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
