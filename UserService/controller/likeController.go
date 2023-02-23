package controller

import (
	"UserService/model"
	"UserService/service"
	"common/Result"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// LikeAction 点赞或者取消赞操作
func LikeAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := Result.NewResult()
		err := c.BindJSON(&request)
		if httpError(err, c) {
			return
		}
		strUserId := request.Data["userId"].(string) // token 鉴权中获取？？
		uid, _ := strconv.ParseInt(strUserId, 10, 64)
		strVideoId := request.Data["video_id"].(string)
		vid, _ := strconv.ParseInt(strVideoId, 10, 64)
		strActionType := request.Data["action_type"].(string)
		act, _ := strconv.ParseInt(strActionType, 10, 64)
		like := service.LikeServiceImpl{}
		err = like.LikeAction(int(uid), int(vid), int(act))
		response := Result.NewResult()
		if err == nil {
			log.Println("点赞成功")
			c.JSON(response.OK().Code, response.SetMessage("点赞成功").SetDataKey("StatusCode", 0))
		} else {
			log.Println(err.Error())
			c.JSON(response.Error().Code, response.SetMessage("点赞失败").SetDataKey("StatusCode", 1))

		}
	}
}

func GetLikeList(c *gin.Context) {
	request := Result.NewResult()
	err := c.BindJSON(&request)
	if httpError(err, c) {
		return
	}
	strUserId := request.Data["userId"].(string) // token 鉴权中获取？？
	uid, _ := strconv.Atoi(strUserId)
	like := service.LikeServiceImpl{}
	list, err := like.GetLikeList(uid)
	if httpError(err, c) {
		return
	}
	response := Result.NewResult()
	c.JSON(response.OK().Code, response.SetMessage("成功获取").SetDataKey("list", list))
}

func IsLikes(c *gin.Context) {
	request := Result.NewResult()
	err := c.BindJSON(&request)
	if httpError(err, c) {
		return
	}
	strUserId := request.Data["userId"].(string) // token 鉴权中获取？？
	vids := request.Data["vids"].([]int)         // token 鉴权中获取？？
	uid, _ := strconv.Atoi(strUserId)
	like := service.LikeServiceImpl{}
	response := Result.NewResult()

	VideoIsLikes := make([]model.VideoIsLike, len(vids))

	for index, vid := range vids {
		isLike, err := like.IsLike(uid, vid)
		if httpError(err, c) {
			return
		}
		VideoIsLikes[index] = model.VideoIsLike{
			ID:     vid,
			IsLike: isLike,
		}
	}
	//httpError(err,c)
	c.JSON(response.OK().Code, response.SetMessage("成功获取").SetDataKey("VideoIsLikes", VideoIsLikes))
}
func GetLikeCounts(c *gin.Context) {
	request := Result.NewResult()
	err := c.BindJSON(&request)
	if httpError(err, c) {
		return
	}
	vids := request.Data["vids"].([]int) // token 鉴权中获取？？
	like := service.LikeServiceImpl{}
	response := Result.NewResult()
	VideoInfos := make([]model.VideoInfo, len(vids))

	for index, vid := range vids {
		count, err := like.VideoLikeCount(vid)
		if httpError(err, c) {
			return
		}
		VideoInfos[index] = model.VideoInfo{
			ID:            vid,
			FavoriteCount: count,
			CommentCount:  0,
		}

	}
	c.JSON(response.OK().Code, response.SetMessage("成功获取").SetDataKey("VideoInfos", VideoInfos))

}
