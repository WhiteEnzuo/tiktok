package model

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/17
 **/
import (
	"errors"
	"gorm.io/gorm"
)

type ContributeTask struct {
	gorm.Model
	UserId     int64  `json:"userId" gorm:"user_id"`
	VideoId    string `json:"VideoId" gorm:"Column:video_id"`
	VideoTitle string `json:"videoTitle" gorm:"video_title"`
}

type ContributeTaskVo struct {
	gorm.Model
	UserId     int64  `json:"userId" gorm:"user_id"`
	VideoUrl   string `json:"videoUrl" gorm:"Column:video_url"`
	PictureUrl string `json:"pictureUrl" gorm:"Column:picture_url"`
	VideoTitle string `json:"videoTitle" gorm:"video_title"`
}

//Add 增
func (c *ContributeTask) Add() error {
	return DB.Table("contribute").Create(c).Error
}

//Delete 删
func (c *ContributeTask) Delete() error {
	err := DB.Table("contribute").Delete(c).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}

//UpData 改
func (c *ContributeTask) UpData() error {
	err := DB.Table("contribute").Updates(c).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}

//QueryById 通过Id查
func (c *ContributeTaskVo) QueryById() error {
	if c.ID <= 0 {
		return errors.New("请输入正确的id号")
	}
	err := DB.Select("contribute.video_title VideoTitle,contribute.user_id UserId,contribute.id ID,"+
		"contribute.created_at CreateAt,contribute.updated_at UpdatedAt,v.video_url video_url,v.image_url picture_url").
		Joins("inner join  file v on v.md5=contribute.video_id").
		Table("contribute").
		Where("id=?", c.ID).
		First(c).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}

//QueryByUserId 通过UserId查
func (c *ContributeTaskVo) QueryByUserId() ([]ContributeTaskVo, error) {
	if c.UserId == 0 {
		return nil, errors.New("请输入正确的id号")
	}
	var contributeTaskVos []ContributeTaskVo
	err := DB.Select("contribute.video_title VideoTitle,contribute.user_id UserId,contribute.id ID,"+
		"contribute.created_at CreateAt,contribute.updated_at UpdatedAt,v.video_url video_url,v.image_url picture_url").
		Joins("inner join  file v on v.md5=contribute.video_id").
		Table("contribute").
		Where("user_id=?", c.UserId).
		Find(&contributeTaskVos).Error

	return contributeTaskVos, err
}

//QueryRandomId 随机获取查
func (c *ContributeTaskVo) QueryRandomId(number int) ([]ContributeTaskVo, error) {
	if number < 1 {
		return nil, errors.New("请输入正确的数字大小")
	}
	if number > 30 {
		number = 30
	}
	var contributeTaskVo []ContributeTaskVo

	DB.Debug().Table("contribute").
		Select("contribute.video_title VideoTitle,contribute.user_id UserId,contribute.id ID," +
			"contribute.created_at CreateAt,contribute.updated_at UpdatedAt,v.video_url video_url,v.image_url picture_url").
		Joins("inner join  file v on v.md5=contribute.video_id").
		Order("RAND()").Limit(number).Find(&contributeTaskVo)

	return contributeTaskVo, nil
}

//QueryById 通过Id查
func (c *ContributeTask) QueryById() error {
	if c.ID <= 0 {
		return errors.New("请输入正确的id号")
	}
	err := DB.Select("*").
		Table("contribute").
		Where("id=?", c.ID).
		First(c).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}

//QueryByUserId 通过UserId查
func (c *ContributeTask) QueryByUserId() ([]ContributeTask, error) {
	if c.UserId == 0 {
		return nil, errors.New("请输入正确的id号")
	}
	var contributeTasks []ContributeTask
	err := DB.Select("*").
		Table("contribute").
		Where("user_id=?", c.UserId).
		Find(&contributeTasks).Error
	return contributeTasks, err
}

//QueryRandomId 随机获取查
func (c *ContributeTask) QueryRandomId(number int) ([]ContributeTask, error) {
	if number < 1 {
		return nil, errors.New("请输入正确的数字大小")
	}
	if number > 30 {
		number = 30
	}
	var contributeTasks []ContributeTask
	err := DB.Debug().Table("contribute").
		Select("*").
		Joins("inner join  file v on v.md5=contribute.video_id").
		Joins("Inner Join file p on p.md5=contribute.picture_id").Order("RAND()").
		Order("by contribute.create_time desc").
		Limit(number).
		Find(&contributeTasks).
		Error
	return contributeTasks, err
}
