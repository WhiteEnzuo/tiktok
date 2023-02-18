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
		password := context.Query("password")
		if password == "" {
			password = context.PostForm("password")
		}
		// 将密码进行SHA加密
		o := sha1.New()
		o.Write([]byte(password))
		str := hex.EncodeToString(o.Sum(nil))
		context.Set("password", str)
		context.Next()
	}
}

// JWTMiddleWare 鉴权中间件，鉴权并设置user_id
func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
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
			re.SetMessage("token不正确")
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
		c.Set("user_id", tokenStruck.UserID)
		c.Next()
	}
}
