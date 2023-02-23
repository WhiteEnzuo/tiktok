package router

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"VideoService/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	//注册接口
	video := r.Group("/tiktok")
	video.Handle("POST", "/upload", func(c *gin.Context) {
		service.UploadVideo(r, c)
	})
	video.Handle("POST", "/publish/list", service.GetContributeByUserId)
	video.Handle("POST", "/feed", service.GetContributes)
	video.Handle("POST", "/GetContributeByIds", service.GetContributeByIds)
	video.StaticFS("/video", http.Dir("./upload/video"))
	video.StaticFS("/picture", http.Dir("./upload/picture"))

}
