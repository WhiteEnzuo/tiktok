package service

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/7
 **/
import (
	"VideoService/model"
	"bytes"
	"common/Result"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go-micro.dev/v4/util/log"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

/*
	接受参数
	form:{
	userid
	title
	file
}
	返回
	{
	"msg":  "success",
	"video_id",
	"video_url":,
	"picture_url",
	"name",
}
*/
func UploadVideo(r *gin.Engine, c *gin.Context) {
	//获取上下文信息
	result := Result.NewResult()
	userId := c.PostForm("userId")
	atom, err := strconv.ParseInt(userId, 10, 64)
	if httpError(err, c) {
		return
	}
	title := c.PostForm("title")
	f, err := c.FormFile("file")
	//读取文件信息
	r.MaxMultipartMemory = 1 << 20
	if httpError(err, c) {
		return
	}
	fileMD5, err := fileHeaderMD5(f)
	if httpError(err, c) {
		return
	}
	//判断是否存在相同的文件
	file := model.File{Md5: fileMD5}
	err = file.QueryByMd5()
	if httpError(err, c) {
		return
	}

	if file.VideoUrl != "" {
		//如果相同就走这，不重复保存

		contributeTask := model.ContributeTask{
			UserId:     atom,
			VideoId:    fileMD5,
			VideoTitle: title,
		}
		err = contributeTask.Add()
		if httpError(err, c) {
			return
		}
		result.
			OK().
			SetDataKey("video_id", contributeTask.ID).
			SetDataKey("video_url", file.VideoUrl).
			SetDataKey("picture_url", file.ImageUrl).
			SetDataKey("name", file.FileName)
		c.JSON(result.Code, result)
		return
	}
	fileName := fileMD5
	dst := "/upload/video/" + fileName + ".mp4"
	err = c.SaveUploadedFile(f, "."+dst)
	if httpError(err, c) {
		return
	}
	//截取第一帧为图片
	snapshot, err := getSnapshot("."+dst, fileName, 0)
	if httpError(err, c) {
		return
	}
	file = model.File{VideoUrl: "/upload/video/" + fileName + ".mp4", FileName: fileName, Md5: fileMD5, ImageUrl: snapshot}
	err = file.Add()
	if httpError(err, c) {
		return
	}
	contributeTask := model.ContributeTask{
		UserId:     atom,
		VideoId:    fileMD5,
		VideoTitle: title,
	}
	err = contributeTask.Add()
	if httpError(err, c) {
		return
	}
	result.
		OK().
		SetDataKey("video_id", contributeTask.ID).
		SetDataKey("video_url", file.VideoUrl).
		SetDataKey("picture_url", file.ImageUrl).
		SetDataKey("name", file.FileName)
	c.JSON(result.Code, result)
	return

}

// 错误返回
func httpError(err error, c *gin.Context) bool {
	if err != nil {
		log.Error(err)
		result := Result.NewResult()
		c.JSON(result.Error().Code, result.SetMessage(err.Error()))
		return true
	}
	return false
}

//文件转md5确保唯一性
func fileHeaderMD5(video *multipart.FileHeader) (string, error) {
	f, err := video.Open()
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, err = io.Copy(hash, f)
	return hex.EncodeToString(hash.Sum(nil)), err
}

//获取第一帧图片
func getSnapshot(videoPath, imageName string, frameNum int) (ImagePath string, err error) {
	snapshotPath := "/upload/picture/" + imageName
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, "."+snapshotPath+".jpg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	imgPath := snapshotPath + ".jpg"

	return imgPath, nil
}
