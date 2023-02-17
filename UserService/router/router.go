package router

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"UserService/controller"
	"github.com/gin-gonic/gin"
)

// Like 点赞接口
func Like(r *gin.Engine) {
	v1Group := r.Group("/douyin")
	v1Group.Handle("POST", "/favorite/action/", controller.LikeAction())
}
