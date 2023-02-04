package consul

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	c "github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"io/ioutil"
	"net/http"
	"regexp"
)

func GetConsul() registry.Registry {
	return c.NewRegistry(registry.Addrs("8.130.28.213:8500"))

}

// GetConsulConfig 设置配置中心
func GetConsulConfig(path string, response *map[string]interface{}) error {
	get, err := http.Get("http://8.130.28.213:8500/v1/kv/" + path)
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(get.Body)

	compile := regexp.MustCompile(`"Value":"(.*?)"`)
	match := compile.FindStringSubmatch(string(all))
	if len(match) < 2 {
		return errors.New("解析器有问题")
	}
	value := match[1]
	decodeString, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}

	err = json.Unmarshal(decodeString, response)
	if err != nil {
		return err
	}
	return nil
}
