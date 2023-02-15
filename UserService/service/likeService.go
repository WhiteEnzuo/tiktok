package service

type LikeService interface {
	// IsLike 当前用户 uid 是否点赞了当前视频 vid
	IsLike(uid int, vid int) (bool, error)
	// VideoLikeCount 当前视频 vid 被点赞数
	VideoLikeCount(vid int) (int, error)
	// LikeListCount 当前用户 uid 点赞视频总数
	LikeListCount(uid int) (int, error)
	// LikeAction 当前用户 uid 对视频 vid 的点赞操作，1-点赞，2-取消点赞
	LikeAction(uid int, vid int, act int) error
	// GetLikeList 当前用户 uid 所有点赞视频
	GetLikeList(uid int, curId int) ([]int, error)
	//// UserLikeCount 当前用户 uid 被点赞总数
	//UserLikeCount(uid int) (int, error)
}
