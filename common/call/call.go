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
	"io"
	"io/ioutil"
	"net/http"
)

func Call(c registry.Registry, serviceName string, path string, request *Result.Result) (*Result.Result, error) {
	result := Result.NewResult()
	//注册中心获取服务，path方法，request参数，response
	/**
	服务名
	方法名
	参数
	参数类型
	*/
	service, _ := c.GetService(serviceName)
	next := selector.Random(service)
	node, _ := next()
	var req *http.Request
	var err error
	if request == nil {

		req, err = http.NewRequest("POST", "http://"+node.Address+path, nil)
		//if err != nil {
		//	return nil, err
		//}

	} else {
		marshal, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}
		buffer := bytes.NewBuffer(marshal)
		req, err = http.NewRequest("POST", "http://"+node.Address+path, buffer)
		//if err != nil {
		//	return nil, err
		//}
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	client := http.Client{}
	resp, err := client.Do(req.WithContext(context.TODO()))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
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

func CallForm(c registry.Registry, serviceName string, path string, contentType string, request *bytes.Buffer) (*Result.Result, error) {
	result := Result.NewResult()
	service, _ := c.GetService(serviceName)
	next := selector.Random(service)
	node, _ := next()
	resp, err := http.Post("http://"+node.Address+path, contentType, request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
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
