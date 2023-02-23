package Vo

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/20
 **/

type Video struct {
	ID            int    `json:"id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
}
