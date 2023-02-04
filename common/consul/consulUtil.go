package consul

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

func GetConsul() registry.Registry {
	return consul.NewRegistry(registry.Addrs("8.130.28.213:8500"))
}
