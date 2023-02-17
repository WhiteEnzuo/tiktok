package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB() {
	dsn := "root:root@tcp(8.130.28.213:3306)/like"
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	Db.Ping()
	//Db = mysql.NewMysql("root", "root", "8.130.28.213", "3306", "like").Server
}

// grom 框架下没有 Query 函数 Dao 层写崩溃了 QAQ
//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)

//func InitDB() {
//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//	dsn := "root:root@tcp(8.130.28.213:3306)/like?charset=utf8mb4&parseTime=True&loc=Local"
//	// 连接 MYSQL 数据库 https://gorm.cn/zh_CN/docs/connecting_to_the_database.html
//	var err error
//	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		// 提高性能
//		SkipDefaultTransaction: true, // 关闭默认事务
//		PrepareStmt:            true, // 缓存预编译语句
//	})
//	if err != nil {
//		panic("failed to connect database")
//	}
//}
