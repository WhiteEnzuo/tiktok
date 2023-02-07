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
