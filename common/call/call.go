package call

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"bytes"
	"common/Result"
	"context"
	"encoding/json"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"io/ioutil"
	"net/http"
)

func Call(c registry.Registry, serviceName string, path string, request *Result.Result) (*Result.Result, error) {
	result := Result.NewResult()
	service, _ := c.GetService(serviceName)
	next := selector.Random(service)
	node, _ := next()
	marshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(marshal)
	req, err := http.NewRequest("POST", "http://"+node.Address+path, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	client := http.Client{}
	resp, err := client.Do(req.WithContext(context.TODO()))
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(all, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
