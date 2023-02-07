package models

import (
	"common/config"
	"common/mysql"
	"gorm.io/gorm"
)

type mysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func mysqlConfigInit() (mysqlConfig, error) {
	var conf mysqlConfig
	err := config.ReadConfig("mysql", &conf)
	if conf.Username == "" {
		conf.Username = "root"
	}
	if conf.Password == "" {
		conf.Password = "root"
	}
	if conf.Host == "" {
		conf.Host = "8.130.28.213"
	}
	if conf.Port == "" {
		conf.Port = "3306"
	}
	if conf.DBName == "" {
		conf.DBName = "tiktok"
	}
	return conf, err
}

var DB *gorm.DB

func InitDB() {
	var err error

	mysqlCon, err := mysqlConfigInit()
	if err != nil {
		return
	}
	DB = mysql.NewMysql(mysqlCon.Username, mysqlCon.Password, mysqlCon.Host, mysqlCon.Port, mysqlCon.DBName).Server

	err = DB.AutoMigrate(&UserInfo{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
}
