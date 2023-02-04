package call

import (
	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4/selector"
	"reflect"
)

import (
	"context"
)

func Call(s selector.Selector, context context.Context, serviceName string, path string, request interface{}, response interface{}) error {

	return nil
}

func Struct2map(obj any) (data map[string]any, err error) {
	// 通过反射将结构体转换成map
	data = make(map[string]any)
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		fileName, ok := objT.Field(i).Tag.Lookup("json")
		if ok {
			data[fileName] = objV.Field(i).Interface()
		} else {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
	}
	return data, nil
}
func Map2Struct(val map[interface{}]interface{}, target *interface{}) error {
	err := mapstructure.Decode(val, target)
	if err != nil {
		return err
	}
	return nil
}
