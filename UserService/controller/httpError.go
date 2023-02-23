package controller

import "github.com/gin-gonic/gin"

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/20
 **/
func httpError(err error, ctx *gin.Context) bool {
	if err != nil {
		ctx.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  err.Error(),
		})
		return true
	}
	return false
}
