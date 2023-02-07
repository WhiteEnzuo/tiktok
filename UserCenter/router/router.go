package router

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"common/Result"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	//注册接口
	v1Group := r.Group("/v1")
	v1Group.Handle("POST", "/prod", func(context *gin.Context) {
		ok := Result.NewResult().OK()
		context.JSON(200, ok)
	})

}
