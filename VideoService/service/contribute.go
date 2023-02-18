package service

import (
	"VideoService/model"
	"common/Result"
	"github.com/gin-gonic/gin"
)

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/17
 **/

func GetContributeByUserId(c *gin.Context) {
	result := Result.NewResult()
	request := Result.NewResult()
	err := c.BindJSON(&request)
	if httpError(err, c) {
		return
	}
	data := request.Data
	i := data["userId"].(float64)
	userId := int64(i)

	//userId := int64()
	//atom, err := strconv.ParseInt(userId, 10, 64)
	if httpError(err, c) {
		return
	}
	vo := model.ContributeTaskVo{
		UserId: userId,
	}
	ContributeTaskVoS, err := vo.QueryByUserId()
	if httpError(err, c) {
		return
	}
	c.JSON(result.OK().Code, result.SetDataKey("ContributeTaskVoS", ContributeTaskVoS))
	return
}
func GetContributes(c *gin.Context) {
	result := Result.NewResult()
	vo := model.ContributeTaskVo{}
	ContributeTaskVoS, err := vo.QueryRandomId(30)
	if httpError(err, c) {
		return
	}
	c.JSON(result.OK().Code, result.SetDataKey("ContributeTaskVoS", ContributeTaskVoS))
	return
}
