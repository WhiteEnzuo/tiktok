package main

import (
	"Service/admin"
	"common/Result"
	"fmt"
)

func main() {
	//Result用来传输
	result := Result.NewResult()
	result.OK().SetCode(201).SetDataKey("test", 1)
	fmt.Println(result.ToJsonString())

	if false {
		server := admin.GetServer()
		err := server.Run()
		if err != nil {
			return
		}
	}

}
