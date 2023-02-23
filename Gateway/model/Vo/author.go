package Vo

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/20
 **/
type Author struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
