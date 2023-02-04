package router

import (
	"github.com/gin-gonic/gin"
	"rpc/rpc/Test"
)

func Register(r *gin.Engine) {
	//注册接口
	v1Group := r.Group("/v1")
	v1Group.Handle("POST", "/prod", func(context *gin.Context) {
		var req Test.Request
		err := context.BindJSON(&req)
		var resp Test.Response
		resp.Test = "aaa"
		if err != nil {
			return
		}
		context.JSON(200, resp)
	})

}
