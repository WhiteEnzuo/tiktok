package model

import (
	"errors"
)

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/7
 **/

type File struct {
	VideoUrl string `json:"url" gorm:"video_url"`
	Md5      string `json:"md5" gorm:"md5"`
	FileName string `json:"fileName" gorm:"file_name"`
	ImageUrl string `json:"imageUrl" gorm:"image_url"`
}

// Add 增
func (f *File) Add() error {
	if f == nil || f.VideoUrl == "" || f.Md5 == "" || f.FileName == "" || f.ImageUrl == "" {
		return errors.New("不能传空值")
	}

	return DB.Table("file").Create(f).Error
}

// Delete 删
func (f *File) Delete() error {
	if f == nil {
		return errors.New("不能传空值")
	}
	err := DB.Table("file").Delete(f).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}

// QueryByMd5 通过MD5查
func (f *File) QueryByMd5() error {
	if f.Md5 == "" {
		return errors.New("MD5不能为空")
	}
	err := DB.Table("file").Where("md5=?", f.Md5).Debug().First(f).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}

// QueryByUrl 通过URL查
func (f *File) QueryByUrl() error {
	if f.VideoUrl == "" {
		return errors.New("URL不能为空")
	}
	err := DB.Table("file").Where("video_url=?", f.VideoUrl).First(f).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}
