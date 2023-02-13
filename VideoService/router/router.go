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
	video := r.Group("/video")
	video.Handle("POST", "/upload", func(c *gin.Context) {
		//result := Result.NewResult()
		service.Upload(r, c)
		//c.JSON(200, result.OK().SetMessage("123").SetDataKey("url", "Test"))
	})
	video.StaticFS("/file", http.Dir("./upload"))

}
