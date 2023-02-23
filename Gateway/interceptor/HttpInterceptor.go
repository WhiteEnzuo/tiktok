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

func UserIdInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Query("user_id")
		if userId == "" {
			userId = c.PostForm("user_id")
		}
		if userId == "" {
			c.JSON(403, gin.H{
				"status_msg":  "请填写user_id",
				"status_code": 403,
			})
			c.Abort()
			return
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
			c.Abort()
			return
		}
		request := Result.NewResult()
		request.OK().SetDataKey("token", token)
		response, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/userId", request)
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{
				"status_msg":  err.Error(),
				"status_code": 500,
			})
			c.Abort()
			return
		}
		if response.Code != 200 {
			c.JSON(response.Code, gin.H{
				"status_msg":  response.Message,
				"status_code": 403,
			})
			c.Abort()
			return
		}
		userIdBd := response.Data["userId"]
		c.Set("userId2", userId)
		c.Set("userId", userIdBd)
		c.Next()
	}
}
func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			c.JSON(403, gin.H{
				"status_msg":  "token",
				"status_code": 403,
			})
			c.Abort()
			return
		}
		request := Result.NewResult()
		request.OK().SetDataKey("token", token)
		response, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/userId", request)
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{
				"status_msg":  err.Error(),
				"status_code": 500,
			})
			c.Abort()
			return
		}
		if response.Code != 200 {
			c.JSON(response.Code, gin.H{
				"status_msg":  response.Message,
				"status_code": 403,
			})
			c.Abort()
			return

		}
		userId := response.Data["userId"].(string)
		c.Set("userId", userId)
		c.Next()
	}
}
