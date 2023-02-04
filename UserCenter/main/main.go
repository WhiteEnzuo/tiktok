package main

import "Service/admin"

func main() {
	server := admin.GetServer()
	err := server.Run()
	if err != nil {
		return
	}
}
