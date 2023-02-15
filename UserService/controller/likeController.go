package controller

import (
	"UserService/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type likeResponse struct {
	StatusCode int8   `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// LikeAction 点赞或者取消赞操作 (还没调好)
func LikeAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		strUserId := c.GetString("userId") // token 鉴权中获取？？
		uid, _ := strconv.ParseInt(strUserId, 10, 64)
		strVideoId := c.Query("video_id")
		vid, _ := strconv.ParseInt(strVideoId, 10, 64)
		strActionType := c.Query("action_type")
		act, _ := strconv.ParseInt(strActionType, 10, 64)
		like := service.LikeServiceImpl{}
		err := like.LikeAction(int(uid), int(vid), int(act))
		if err == nil {
			log.Println("点赞成功")
			c.JSON(http.StatusOK, likeResponse{
				StatusCode: 0,
				StatusMsg:  "点赞成功",
			})
		} else {
			log.Println(err.Error())
			c.JSON(http.StatusOK, likeResponse{
				StatusCode: 1,
				StatusMsg:  "点赞失败",
			})
		}
	}
}
