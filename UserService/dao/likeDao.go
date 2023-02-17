package dao

import (
	"database/sql"
	"errors"
	//"gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Like 定义 gorm.Model 结构体，对应数据库中一张表
type Like struct {
	//gorm.Model
	Id  int //自增主键
	Uid int //点赞用户id
	Vid int //视频id
	Act int //是否点赞，1为点赞，0为取消赞
}

// TableName 为 model 定义表名
func (Like) TableName() string {
	return "likes"
}

// InsertLike uid 初次给 vid 点赞，插入数据库中
func InsertLike(uid int, vid int) error {
	_, err := Db.Query("INSERT INTO likes(uid,vid,act) VALUES(?,?,1)", uid, vid)
	if err != nil {
		log.Println(err.Error())
		return errors.New("插入点赞数据失败")
	}
	return nil
}

// GetLike 查询 uid 是否对 vid 视频点赞 0-不存在 1-点赞 2-没点赞 -1-错误
func GetLike(uid int, vid int) (int, error) {
	var act int
	err := Db.QueryRow("SELECT act FROM likes WHERE uid=? AND vid=?", uid, vid).Scan(&act)
	if err != nil {
		// 当前用户对该视频没点过赞
		if err == sql.ErrNoRows {
			return 0, nil
		} else {
			log.Println(err.Error())
			return -1, errors.New("查询是否点赞失败")
		}
	}
	if act == 1 {
		return 1, nil
	} else {
		return 2, nil
	}
}

// UpdateLike 根据 uid，vid, act 点赞或者取消赞
func UpdateLike(uid int, vid int, act int) error {
	_, err := Db.Query("UPDATE likes SET act=? WHERE uid=? AND vid=?", act, uid, vid)
	if err != nil {
		log.Println(err.Error())
		return errors.New("更新点赞失败")
	}
	return nil
}

// LikeList 查询 uid 点赞的全部视频 (？？按 updated_at 降序排序，先显示最新点赞)
func LikeList(uid int) ([]int, error) {
	var vidList []int
	rows, err := Db.Query("SELECT vid FROM likes WHERE uid=? AND act=?", uid, 1)
	if err != nil || rows.Err() != nil {
		log.Println(err.Error())
		return vidList, errors.New("查询点赞视频失败")
	}
	for rows.Next() {
		var vid int
		err = rows.Scan(&vid)
		if err != nil {
			log.Println(err.Error())
			return vidList, errors.New("查询点赞视频失败")
		}
		vidList = append(vidList, vid)
	}
	return vidList, nil
}

// LikeUserList 查询给 vid 点赞的全部 uid
func LikeUserList(vid int) ([]int, error) {
	var uidList []int
	rows, err := Db.Query("SELECT uid FROM likes WHERE vid=? AND act=?", vid, 1)
	if err != nil || rows.Err() != nil {
		log.Println(err.Error())
		return uidList, errors.New("查询点赞用户失败")
	}
	for rows.Next() {
		var uid int
		err = rows.Scan(&uid)
		if err != nil {
			log.Println(err.Error())
			return uidList, errors.New("查询点赞用户失败")
		}
		uidList = append(uidList, uid)
	}
	return uidList, nil
}
