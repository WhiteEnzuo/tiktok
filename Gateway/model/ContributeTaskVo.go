package model

import (
	"database/sql"
)

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/21
 **/
type DeletedAt sql.NullTime
type ContributeTaskVo struct {
	ID uint `gorm:"primarykey"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	DeletedAt  DeletedAt `gorm:"index"`
	UserId     int64     `json:"userId" gorm:"user_id"`
	VideoUrl   string    `json:"videoUrl" gorm:"Column:video_url"`
	PictureUrl string    `json:"pictureUrl" gorm:"Column:picture_url"`
	VideoTitle string    `json:"videoTitle" gorm:"video_title"`
}
