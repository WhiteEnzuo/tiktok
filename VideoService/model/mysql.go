package model

import (
	"common/consul"
	"common/mysql"
	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4/util/log"
	"gorm.io/gorm"
)

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/13
 **/

var DB *gorm.DB

func init() {
	//配置中心
	config := consul.NewConfig("8.130.28.213", "8500")
	var mysqlMap map[string]interface{}
	err := config.GetConsulConfig("Video/mysql", &mysqlMap)
	if err != nil {
		log.Error("VideoMysql初始化失败", err)
	}
	var m mysql.Mysql
	err = mapstructure.Decode(mysqlMap, &m)
	if err != nil {
		log.Error("VideoMysql初始化失败", err)
	}
	err = m.GetConnect()

	if err != nil {
		log.Error("VideoMysql初始化失败", err)
	}
	DB = m.Server

}
