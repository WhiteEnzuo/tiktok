package model

import "errors"

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/7
 **/

type File struct {
	Url      string `json:"url" gorm:"url"`
	Md5      string `json:"md5" gorm:"md5"`
	FileName string `json:"fileName" gorm:"file_name"`
}

// Add 增
func (f *File) Add() error {
	if f == nil || f.Url == "" || f.Md5 == "" || f.FileName == "" {
		return errors.New("不能传空值")
	}

	return DB.Table("file").Create(f).Error
}

// Delete 删
func (f *File) Delete() error {
	if f == nil {
		return errors.New("不能传空值")
	}

	return DB.Table("file").Delete(f).Error
}

// QueryByMd5 通过MD5查
func (f *File) QueryByMd5() error {
	if f.Url == "" {
		return errors.New("MD5不能为空")
	}
	return DB.Table("file").Where("md5=?", f.Md5).First(f).Error
}

// QueryByUrl 通过URL查
func (f *File) QueryByUrl() error {
	if f.Url == "" {
		return errors.New("URL不能为空")
	}
	return DB.Table("file").Where("url=?", f.Url).First(f).Error
}
