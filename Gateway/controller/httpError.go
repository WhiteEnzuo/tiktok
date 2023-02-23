package controller

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/20
 **/
import (
	"common/Result"
	"github.com/gin-gonic/gin"
)

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
func ResponseHandle(response *Result.Result, ctx *gin.Context) bool {
	if response.Code == 500 {
		ctx.JSON(500, response)
		return true
	}
	return false
}
