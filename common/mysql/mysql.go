package mysql

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"go-micro.dev/v4/util/log"
	my "gorm.io/driver/mysql" // gorm mysql 驱动包
	"gorm.io/gorm"            // gorm
)

type Mysql struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	Server   *gorm.DB
}

func (m *Mysql) GetConnect() error {
	// MySQL 配置信息
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Open 连接
	db, err := gorm.Open(my.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	m.Server = db
	return nil
}
func NewMysql(Username, Password, Host, Port, DBName string) *Mysql {
	m := new(Mysql)
	m.Username = Username
	m.Password = Password
	m.Host = Host
	m.Port = Port
	m.DBName = DBName
	err := m.GetConnect()
	if err != nil {
		log.Error("Mysql配置有问题")
		return nil
	}
	return m
}
