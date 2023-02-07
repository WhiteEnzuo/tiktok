package models

import (
	"errors"
	"sync"
)

// UserLogin 用户登录表
type UserLogin struct {
	Id         int64 `gorm:"primary_key"`
	UserInfoId int64
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:200;notnull"`
}

// UserLoginDAO 用来创建一个单例对象
type UserLoginDAO struct {
}

var (
	userLoginDao  *UserLoginDAO
	userLoginOnce sync.Once
)

// NewUserLoginDao 创建一个单例对象
func NewUserLoginDao() *UserLoginDAO {
	userLoginOnce.Do(func() {
		userLoginDao = new(UserLoginDAO)
	})
	return userLoginDao
}

// QueryUserLogin 查询账号密码是否正确
func (u *UserLoginDAO) QueryUserLogin(username, password string, login *UserLogin) error {
	if login == nil {
		return errors.New("结构体指针为空")
	}
	DB.Where("username=? and password=?", username, password).First(login)
	if login.Id == 0 {
		return errors.New("用户不存在，账号或密码出错")
	}
	return nil
}

// IsUserExistByUsername 查询是否存在用户
func (u *UserLoginDAO) IsUserExistByUsername(username string) bool {
	var userLogin UserLogin
	DB.Where("username=?", username).First(&userLogin)
	if userLogin.Id == 0 {
		return false
	}
	return true
}
