package controller

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/18
 **/
import (
	"Gateway/gloabl"
	"bytes"
	"common/Result"
	"common/call"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
)

func PublishList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId2")
	request := Result.NewResult()
	request.SetDataKey("userId", userId)
	response, err := call.Call(gloabl.GetConsul(), "VideoService", "/tiktok/publish/list", request)
	if httpError(err, ctx) {
		return
	}
	_ = response.Data["ContributeTaskVoS"].(map[string]interface{})
}
func PublishAction(ctx *gin.Context) {
	file, err := ctx.FormFile("data")
	if httpError(err, ctx) {
		return
	}
	token, b1 := ctx.GetPostForm("token")
	title, b2 := ctx.GetPostForm("title")
	if b1 == false || b2 == false {
		err := errors.New("获取token或者title错误")
		if httpError(err, ctx) {
			return
		}
	}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	err = bodyWriter.WriteField("title", title)
	if httpError(err, ctx) {
		return
	}
	err = bodyWriter.WriteField("token", token)
	if httpError(err, ctx) {
		return
	}
	formFile, err := bodyWriter.CreateFormFile("file", title)

	if httpError(err, ctx) {
		return
	}
	video, err := file.Open()
	if httpError(err, ctx) {
		return
	}
	_, err = io.Copy(formFile, video)
	if httpError(err, ctx) {
		return
	}

	response, err := call.CallForm(gloabl.GetConsul(), "", "", bodyWriter.FormDataContentType(), bodyBuf)
	if httpError(err, ctx) {
		return
	}
	ctx.JSON(response.Code, gin.H{
		"status_code": response.Code,
		"status_msg":  response.Message,
	})
}
