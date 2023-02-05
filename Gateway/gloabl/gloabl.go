package gloabl

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import "go-micro.dev/v4/registry"

var consul registry.Registry

func GetConsul() registry.Registry {
	return consul
}
func SetConsul(c registry.Registry) {
	consul = c
}
