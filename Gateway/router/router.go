package router

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"Gateway/controller"
	"Gateway/interceptor"
	"github.com/gin-gonic/gin"
)

//路由注册
func Register(r *gin.Engine) {

	user := r.Group("/douyin/user")
	// /douyin/user/register/
	user.Handle("POST", "/register/", controller.UserRegister)
	// /douyin/user/login/
	user.Handle("POST", "/login/", controller.UserLogin)
	// /douyin/user/
	user.Handle("GET", "/", interceptor.UserIdInterceptor(), controller.User)

	feed := r.Group("/douyin/feed")
	// /douyin/feed/
	feed.Handle("GET", "/", controller.Feed)

	publish := r.Group("/douyin/publish")
	// /douyin/publish/action/
	publish.Handle("POST", "/action/", interceptor.Interceptor(), controller.PublishAction)
	// /douyin/publish/list/
	publish.Handle("GET", "/list/", interceptor.UserIdInterceptor(), controller.PublishList)

	favorite := r.Group("/douyin/favorite")
	favorite.Handle("POST", "/action/", interceptor.Interceptor(), controller.FavoriteAction)
	favorite.Handle("GET", "/list/", interceptor.Interceptor(), controller.FavoriteList)

}
