package controller

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/18
 **/
import (
	"Gateway/gloabl"
	"common/Result"
	"common/call"
	"github.com/gin-gonic/gin"
)

func User(ctx *gin.Context) {
	request := Result.NewResult()
	userId := ctx.Query("user_id")
	token := ctx.Query("token")
	request.SetDataKey("token", token)
	request.SetDataKey("user_id", userId)
	response, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/user/", request)
	if httpError(err, ctx) {
		return
	}
	ctx.JSON(response.Code, gin.H{
		"status_code": response.Data["status_code"],
		"status_msg":  response.Message,
		"user":        response.Data["user"],
	})

}
func UserLogin(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	request := Result.NewResult()
	request.SetDataKey("username", username)
	request.SetDataKey("password", password)
	response, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/user/login", request)
	if httpError(err, ctx) {
		return
	}
	ctx.JSON(response.Code, gin.H{
		"status_code": response.Code,
		"status_msg":  response.Message,
		"user_id":     response.Data["userid"],
		"token":       response.Data["token"],
	})
	return
}
func UserRegister(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	request := Result.NewResult()
	request.OK().SetDataKey("username", username).SetDataKey("password", password)
	response, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/user/register", request)
	if httpError(err, ctx) {
		return
	}
	ctx.JSON(response.Code, gin.H{
		"status_code": response.Code,
		"status_msg":  response.Message,
		"user_id":     response.Data["userid"],
		"token":       response.Data["token"],
	})
	return
}
