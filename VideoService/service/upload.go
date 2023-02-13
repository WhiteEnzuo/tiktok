package service

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/7
 **/
import (
	"VideoService/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-micro.dev/v4/util/log"
	"io"
	"mime/multipart"
	"net/http"
	"path"
)

func Upload(r *gin.Engine, c *gin.Context) {
	f, err := c.FormFile("file")
	r.MaxMultipartMemory = 1 << 20
	httpError(err, c)
	fileMD5, err := FileMD5(f)
	httpError(err, c)
	file := model.File{Md5: fileMD5}
	err = file.QueryByUrl()
	httpError(err, c)
	if file.Url != "" {
		c.JSON(200, gin.H{
			"msg":  "success",
			"size": f.Size,
			"url":  file.Url,
			"name": file.FileName,
		})
		return
	}
	u := uuid.New()
	fileName := u.String()
	dst := path.Join("./upload/", fileName, ".mp4")
	err = c.SaveUploadedFile(f, dst)
	httpError(err, c)
	file = model.File{Url: "/upload/" + fileName + ".mp4", Md5: fileMD5}
	err = file.Add()
	httpError(err, c)
	c.JSON(200, gin.H{
		"msg":  "success",
		"name": fileName,
		"size": f.Size,
		"url":  "/upload/" + fileName + ".mp4",
	})

}
func httpError(err error, c *gin.Context) {
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func FileMD5(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, f)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
