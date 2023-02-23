package logic

import (
	"UserCenter/models"
	"common/Result"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProxyUserInfo struct {
	c *gin.Context
}

func NewProxyUserInfo(c *gin.Context) *ProxyUserInfo {
	return &ProxyUserInfo{c: c}
}

func (p *ProxyUserInfo) DoQueryUserInfoByUserId(rawId interface{}) error {
	userId, ok := rawId.(uint64)

	if !ok {
		return errors.New("解析userId失败")
	}
	//由于得到userinfo不需要组装model层的数据，所以直接调用model层的接口
	userinfoDAO := models.NewUserInfoDAO()

	var userInfo models.UserInfo
	err := userinfoDAO.QueryUserInfoById(userId, &userInfo)
	if err != nil {
		return err
	}
	p.UserInfoOk(&userInfo)
	return nil
}

func (p *ProxyUserInfo) QueryUserInfoByUserId(rawId interface{}) (error, *models.UserInfo) {
	userId, ok := rawId.(uint64)
	if !ok {
		return errors.New("解析userId失败"), nil
	}
	//由于得到userinfo不需要组装model层的数据，所以直接调用model层的接口
	userinfoDAO := models.NewUserInfoDAO()

	var userInfo *models.UserInfo
	err := userinfoDAO.QueryUserInfoById(userId, userInfo)
	if err != nil {
		return err, nil
	}
	return nil, userInfo
}

func (p *ProxyUserInfo) UserInfoError(msg string) {
	re := Result.NewResult().Error()
	re.SetMessage(msg)
	p.c.JSON(http.StatusOK, re)
}

func (p *ProxyUserInfo) UserInfoOk(user *models.UserInfo) {
	re := Result.NewResult().OK()
	re.SetMessage("Success")
	temp := make(map[string]interface{})
	temp["status_code"] = 0
	//temp.Message= "Success"
	temp["user"] = user
	re.SetData(temp)
	p.c.JSON(re.Code, re)
}
