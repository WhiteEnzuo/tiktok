package main

import "UserCenter/admin"

func main() {
	server := admin.GetServer()
	err := server.Run()
	if err != nil {
		return
	}
}
