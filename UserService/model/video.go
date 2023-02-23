package model

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/20
 **/
type VideoIsLike struct {
	ID     int  `json:"ID"`
	IsLike bool `json:"isLike"`
}

type VideoInfo struct {
	ID            int `json:"ID"`
	FavoriteCount int `json:"favorite_count"`
	CommentCount  int `json:"comment_count"`
}
