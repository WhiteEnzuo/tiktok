package controller

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/18
 **/
import (
	"Gateway/gloabl"
	"Gateway/model"
	"Gateway/model/Vo"
	"common/Result"
	"common/call"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Feed(ctx *gin.Context) {
	token := ctx.Query("token")
	videoResponse, err := call.Call(gloabl.GetConsul(), "VideoService", "/tiktok/feed", nil)

	if httpError(err, ctx) || ResponseHandle(videoResponse, ctx) {
		return
	}
	ContributeTaskVoS := videoResponse.Data["ContributeTaskVoS"].([]interface{})
	Ids := make([]int, len(ContributeTaskVoS))
	UserId := make([]int, len(ContributeTaskVoS))
	videos := make([]Vo.Video, len(ContributeTaskVoS))
	for i := range ContributeTaskVoS {

		var c model.ContributeTaskVo
		err := mapstructure.Decode(ContributeTaskVoS[i], &c)
		if httpError(err, ctx) {
			return
		}
		Ids[i] = int(c.ID)
		UserId[i] = int(c.UserId)
		videos[i].ID = int(c.ID)
		videos[i].PlayUrl = c.VideoUrl
		videos[i].CoverUrl = c.PictureUrl
		videos[i].Title = c.VideoTitle
		videos[i].FavoriteCount = 0
		videos[i].CommentCount = 0
		videos[i].IsFavorite = false
	}
	request := Result.NewResult()
	request.SetDataKey("userIds", UserId)
	UserInfos, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/userInfoByIds", request)

	if httpError(err, ctx) || ResponseHandle(UserInfos, ctx) {
		return
	}
	userInfo := UserInfos.Data["userInfos"].([]map[string]interface{})
	for index := range ContributeTaskVoS {
		for j := range userInfo {
			if videos[index].ID == userInfo[j]["id"].(int) {
				videos[index].Author = Vo.Author{
					Id:            userInfo[j]["id"].(int),
					Name:          userInfo[j]["name"].(string),
					FollowCount:   userInfo[j]["follow_count"].(int),
					FollowerCount: userInfo[j]["follower_count"].(int),
					IsFollow:      false,
				}
				break
			}
		}
	}

	if token != "" {
		request = Result.NewResult()
		request.SetDataKey("token", token)
		response, err := call.Call(gloabl.GetConsul(), "UserCenter", "/douyin/userId", request)
		if ResponseHandle(response, ctx) {
			return
		}
		if httpError(err, ctx) {
			return
		}
		userId := response.Data["userId"].(string)
		request = Result.NewResult()
		request.SetDataKey("vids", Ids)
		request.SetDataKey("userId", userId)
		response, err = call.Call(gloabl.GetConsul(), "UserService", "/douyin/IsLikes/", request)
		if httpError(err, ctx) {
			return
		}
		if ResponseHandle(response, ctx) {
			return
		}
		videoIsLikes := response.Data["VideoIsLikes"].([]map[string]interface{})
		for i := range videoIsLikes {
			videos[i].IsFavorite = videoIsLikes[i]["isLike"].(bool)
		}

	}
	if httpError(err, ctx) {
		return
	}
	ctx.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "获取成功",
		"next_time":   0,
		"video_list":  videos,
	})
}
