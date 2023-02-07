package router

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"github.com/gin-gonic/gin"
)

//路由注册
func Register(r *gin.Engine) {
	//v1Group := r.Group("/service")
	//v1Group.Handle("GET", "/Test", func(ctx *gin.Context) {
	//	result := Result.NewResult()
	//	res, _ := call.Call(gloabl.GetConsul(), "service", "/v1/prod", result)
	//	ctx.JSON(200, res)
	//})
	//
	user := r.Group("/douyin/user")
	user.Handle("POST", "/register", func(ctx *gin.Context) {

	})
}
