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
	douyin := r.Group("/douyin")
	douyin.Handle("POST", "/favorite/action/", controller.LikeAction())
	douyin.Handle("POST", "/favorite/list/", controller.GetLikeList)
	douyin.Handle("POST", "/favorite/IsLikes/", controller.IsLikes)
}
