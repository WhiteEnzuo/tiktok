package router

import (
	"Gateway/gloabl"
	"common/call"
	"github.com/gin-gonic/gin"
	"rpc/rpc/Test"
)

//路由注册
func Register(r *gin.Engine) {
	v1Group := r.Group("/service")
	v1Group.Handle("GET", "/Test", func(ctx *gin.Context) {
		req := Test.Request{
			Name: "123",
		}
		var resp Test.Response
		_ = call.Call(gloabl.GetConsul(), "service", "/v1/prod", req, &resp)
		ctx.JSON(200, resp)
	})

}
