package main

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import "UserCenter/admin"

func main() {
	server := admin.GetServer()
	err := server.Run()
	if err != nil {
		return
	}
}
