package logic

import (
	"common/Result"
	"github.com/gin-gonic/gin"
)

func UserRegisterLogic(c *gin.Context) {
	user, _ := c.Get("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	username, ok := user.(string)
	if !ok {
		re := Result.NewResult().OK()
		re.SetMessage("密码解析出错")
		c.JSON(200, re)
		return
	}
	registerResponse, err := PostUserLogin(username, password)
	if err != nil {
		re := Result.NewResult().OK()
		re.SetMessage(err.Error())
		c.JSON(200, re)
		return
	}
	re := Result.NewResult().OK()
	re.SetData(registerResponse)
	c.JSON(200, re)
}

func UserLoginLogic(c *gin.Context) {
	user, _ := c.Get("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	username, ok := user.(string)
	// 获取密码失败
	if !ok {
		re := Result.NewResult().OK()
		re.SetMessage("密码解析出错")
		c.JSON(200, re)
		return
	}
	// 查询数据库
	userLoginResponse, err := QueryUserLogin(username, password)

	//用户不存在返回对应的错误
	if err != nil {
		re := Result.NewResult().OK()
		re.SetMessage(err.Error())
		c.JSON(200, re)
		return
	}

	//用户存在，返回相应的id和token
	re := Result.NewResult().OK()
	c.JSON(re.Code, re.SetData(userLoginResponse))
}

func UserInfoLogic(c *gin.Context) {
	p := NewProxyUserInfo(c)
	//得到上层中间件根据token解析的userId
	rawId, ok := c.Get("user_id")
	if !ok {
		p.UserInfoError("解析userId出错")
		return
	}
	err := p.DoQueryUserInfoByUserId(rawId)
	if err != nil {
		p.UserInfoError(err.Error())
	}
	result := Result.NewResult()
	c.JSON(result.OK().Code, result.SetDataKey("userId", rawId))
	//p.UserInfoOk()
}
