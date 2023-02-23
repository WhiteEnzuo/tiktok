package logic

import (
	"UserCenter/models"
	"common/token"
	"errors"
)

type QueryUserLoginFlow struct {
	username string
	password string

	data   map[string]interface{}
	userid int64
	token  string
}

// QueryUserLogin 查询用户是否存在，并返回token和id
func QueryUserLogin(username, password string) (map[string]interface{}, error) {
	return NewQueryUserLoginFlow(username, password).Do()
}

func NewQueryUserLoginFlow(username, password string) *QueryUserLoginFlow {
	return &QueryUserLoginFlow{username: username, password: password}
}

func (q *QueryUserLoginFlow) Do() (map[string]interface{}, error) {
	//对参数进行合法性验证
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	//准备好数据
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	//打包最终数据
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.data, nil
}

// 判断参数的合法性
func (q *QueryUserLoginFlow) checkNum() error {
	if q.username == "" {
		return errors.New("用户名为空")
	}
	if len(q.username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}
	if q.password == "" {
		return errors.New("密码为空")
	}
	return nil
}

// 准备数据
func (q *QueryUserLoginFlow) prepareData() error {
	userLoginDAO := models.NewUserLoginDao()
	var login models.UserLogin
	//准备好userid
	err := userLoginDAO.QueryUserLogin(q.username, q.password, &login)
	if err != nil {
		return err
	}
	q.userid = login.UserInfoId

	//准备颁发token
	q.token, err = token.GenToken(uint64(login.Id), login.Username)
	if err != nil {
		return err
	}
	return nil
}

// 打包数据
func (q *QueryUserLoginFlow) packData() error {
	m := make(map[string]interface{})
	m["userid"] = q.userid
	m["token"] = q.token
	q.data = m
	return nil
}
