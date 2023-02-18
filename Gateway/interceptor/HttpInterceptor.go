package interceptor

import (
	"Gateway/gloabl"
	"common/Result"
	"common/call"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/util/log"
)

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/18
 **/

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Query("userId")
		if userId == "" {
			userId = c.PostForm("userId")
		}
		if userId == "" {
			c.JSON(403, gin.H{
				"status_msg":  "请填写UserId",
				"status_code": 403,
			})
		}
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			c.JSON(403, gin.H{
				"status_msg":  "token",
				"status_code": 403,
			})
		}
		request := Result.NewResult()
		request.OK().SetDataKey("token", token)
		response, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/user/", request)
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{
				"status_msg":  err.Error(),
				"status_code": 500,
			})
		}
		if response.Code != 200 {
			c.JSON(response.Code, gin.H{
				"status_msg":  response.Message,
				"status_code": 403,
			})
		}
		userIdBd := response.Data["userId"].(string)
		if userIdBd != userId {
			c.JSON(403, gin.H{
				"status_msg":  "userId不对",
				"status_code": 403,
			})
		}
		c.Set("userId", userId)
		c.Next()
	}
}
