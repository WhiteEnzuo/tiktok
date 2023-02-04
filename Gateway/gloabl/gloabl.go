package gloabl

import "go-micro.dev/v4/registry"

var consul registry.Registry

func GetConsul() registry.Registry {
	return consul
}
func SetConsul(c registry.Registry) {
	consul = c
}
