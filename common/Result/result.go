package Result

import (
	"encoding/json"
)

type Result struct {
	Code      int                    `json:"code"`
	Data      map[string]interface{} `json:"data"`
	IsSuccess bool                   `json:"is_success"`
	Message   string                 `json:"message"`
}

func NewResult() *Result {
	r := new(Result)
	r.Code = 0
	r.Data = make(map[string]interface{})
	r.IsSuccess = false
	r.Message = ""
	return r
}
func (r *Result) SetCode(code int) *Result {
	r.Code = code
	return r
}
func (r *Result) SetData(data map[string]interface{}) *Result {
	r.Data = data
	return r
}
func (r *Result) SetSuccess(isSuccess bool) *Result {
	r.IsSuccess = isSuccess
	return r
}
func (r *Result) SetMessage(message string) *Result {
	r.Message = message
	return r
}
func (r *Result) OK() *Result {

	r.Code = 200
	r.IsSuccess = true
	r.Message = "Success"
	return r
}
func (r *Result) Error() *Result {
	r.Code = 500
	r.IsSuccess = false
	r.Message = "Error"
	return r
}
func (r *Result) SetDataKey(key string, val interface{}) *Result {
	data := r.Data
	data[key] = val
	return r
}
func (r *Result) ToJsonString() (string, error) {
	marshal, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

func (r *Result) ToJsonBytes() ([]byte, error) {
	marshal, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
