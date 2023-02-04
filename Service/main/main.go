package main

import (
	//_ "common/call"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(registry.Addrs("8.130.28.213:8500"))

}
func main() {
	service, _ := consulReg.GetService("Test1")
	next := selector.Random(service)
	node, _ := next()
	fmt.Println(node)

}
