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
	userId, ok := rawId.(int64)
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

func (p *ProxyUserInfo) UserInfoError(msg string) {
	re := Result.NewResult().Error()
	re.SetMessage(msg)
	p.c.JSON(http.StatusOK, re)
}

func (p *ProxyUserInfo) UserInfoOk(user *models.UserInfo) {
	re := Result.NewResult().OK()
	//m := make(map[string]interface{})
	//m["id"] = user.Id
	//m["name"] = user.Name
	//m["follow_count"] = user.FollowCount
	//m["follower_count"] = user.FollowerCount
	//m["is_follow"] = user.IsFollow
	re.SetMessage("Success")
	temp := make(map[string]interface{})
	temp["status_code"] = 0
	temp["status_msg"] = "Success"
	temp["user"] = user

	p.c.JSON(http.StatusOK, re)
}
