package controller

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/18
 **/
import (
	"Gateway/gloabl"
	"Gateway/model/Vo"
	"common/Result"
	"common/call"
	"errors"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	query := ctx.Query("action_type")
	vid := ctx.Query("video_id")
	if query == "" {
		if httpError(errors.New("请填写action_type"), ctx) {
			return
		}
	}
	request := Result.NewResult()
	request.SetDataKey("userId", userId)
	request.SetDataKey("video_id", vid)
	response, err := call.Call(gloabl.GetConsul(), "UserService", "/favorite/action/", request)
	if httpError(err, ctx) || ResponseHandle(response, ctx) {
		return
	}
	ctx.JSON(response.Code, gin.H{
		"status_code": response.Data["status_code"],
		"status_msg":  response.Message,
	})
}
func FavoriteList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId2")
	request := Result.NewResult()
	request.SetDataKey("userId", userId)
	listResponse, err := call.Call(gloabl.GetConsul(), "UserService", "/favorite/list/", request)
	if httpError(err, ctx) {
		return
	}
	contributes, err := call.Call(gloabl.GetConsul(), "VideoService", "/tiktok/GetContributeByIds/", listResponse)
	ContributeTaskVoS := contributes.Data["ContributeTaskVoS"].([]map[string]interface{})
	Ids := make([]int, len(ContributeTaskVoS))
	UserId := make([]int, len(ContributeTaskVoS))
	videos := make([]Vo.Video, len(ContributeTaskVoS))
	for i := range ContributeTaskVoS {
		Ids[i] = ContributeTaskVoS[i]["ID"].(int)
		UserId[i] = ContributeTaskVoS[i]["userId"].(int)
		videos[i].ID = ContributeTaskVoS[i]["ID"].(int)
		videos[i].PlayUrl = ContributeTaskVoS[i]["videoUrl"].(string)
		videos[i].CoverUrl = ContributeTaskVoS[i]["pictureUrl"].(string)
		videos[i].Title = ContributeTaskVoS[i]["videoTitle"].(string)
		videos[i].FavoriteCount = 0
		videos[i].CommentCount = 0
		videos[i].IsFavorite = false
	}
	request = Result.NewResult()
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
	ctx.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "获取成功",
		"next_time":   0,
		"video_list":  videos,
	})
}
