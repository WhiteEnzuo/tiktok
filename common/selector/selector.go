package selector

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

func GetSelector(reg registry.Registry) selector.Selector {
	return selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)
}
