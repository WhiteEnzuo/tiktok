package controller

import (
	"common/Result"
	"common/token"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"time"
)

// SHAMiddleWare gin.Context的Set方法设置password
func SHAMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := Result.NewResult()
		err := context.BindJSON(&request)
		if err != nil {
			context.JSON(500, gin.H{
				"status_msg":  err.Error(),
				"status_code": 500,
			})
		}
		password := request.Data["password"].(string)
		// 将密码进行SHA加密
		o := sha1.New()
		o.Write([]byte(password))
		str := hex.EncodeToString(o.Sum(nil))

		context.Set("password", str)
		context.Set("username", request.Data["username"].(string))
		context.Next()
	}
}

// JWTMiddleWare 鉴权中间件，鉴权并设置user_id
func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		request := Result.NewResult()
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(500, gin.H{
				"status_msg":  err.Error(),
				"status_code": 500,
			})
		}
		tokenStr := request.Data["token"].(string)

		//用户不存在
		if tokenStr == "" {
			re := Result.NewResult().Error()
			re.SetMessage("用户不存在")
			c.JSON(401, re)
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, err := token.ParseToken(tokenStr)
		if err != nil {
			re := Result.NewResult().Error()
			re.SetMessage(err.Error())
			c.JSON(403, re)
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			re := Result.NewResult().Error()
			re.SetMessage("token过期")
			c.JSON(402, re)
			c.Abort() //阻止执行
			return
		}
		//msg := Result.NewResult()
		c.Set("user_id", tokenStruck.UserID)
		c.Next()

		//log.Info("123")
	}
}
