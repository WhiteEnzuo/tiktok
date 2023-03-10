package router

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"UserCenter/controller"
	"UserCenter/logic"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	//注册接口
	v1Group := r.Group("/douyin")
	v1Group.Handle("POST", "/user/register", controller.SHAMiddleWare(), logic.UserRegisterLogic)

}

func Login(r *gin.Engine) {
	//登录接口
	v1Group := r.Group("/douyin")
	v1Group.Handle("POST", "/user/login", controller.SHAMiddleWare(), logic.UserLoginLogic)

}

func UserInfo(r *gin.Engine) {
	//用户信息接口
	v1Group := r.Group("/douyin")
	v1Group.Handle("POST", "/user/", controller.JWTMiddleWare(), logic.UserInfoLogic)

}
func UserId(r *gin.Engine) {
	//用户信息接口
	v1Group := r.Group("/douyin")
	v1Group.Handle("POST", "/userId", controller.JWTMiddleWare(), logic.UserId)

}
func UserInfoByUserId(r *gin.Engine) {
	//用户信息接口
	v1Group := r.Group("/douyin")
	v1Group.Handle("POST", "/userInfoByIds", logic.UserInfoByUserIds)

}
