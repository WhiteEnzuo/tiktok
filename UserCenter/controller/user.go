package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
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
