package admin

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"UserService/dao"
	"UserService/router"
	"common/config"
	"common/consul"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/web"
)

var service web.Service

type serverConfig struct {
	Host        string
	Port        string
	ServiceName string
}
type consulConfig struct {
	Host string
	Port string
}

func consulConfigInit() (consulConfig, error) {
	var conf consulConfig
	err := config.ReadConfig("consul", &conf)
	if err != nil {
		return conf, err
	}
	if conf.Host == "" {
		conf.Host = "8.130.28.213"
	}
	if conf.Port == "" {
		conf.Port = "8500"
	}
	return conf, nil
}
func serverConfigInit() (serverConfig, error) {
	var conf serverConfig
	err := config.ReadConfig("server", &conf)
	if err != nil {
		return conf, err
	}
	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}
	if conf.Port == "" {
		conf.Port = "8904"
	}
	if conf.ServiceName == "" {
		conf.ServiceName = "UserService"
	}
	return conf, nil
}
func routerInit() *gin.Engine {
	//创建gin
	server := gin.Default()
	//点赞接口
	router.Like(server)
	return server
}
func init() {
	//服务器配置
	serverConf, err := serverConfigInit()
	if err != nil {
		return
	}
	consulConf, err := consulConfigInit()
	if err != nil {
		return
	}
	if false {
		//var ConfigMap map[string]interface{}
		//获取consul的配置
		//err = consul.GetConsulConfig("Test", &ConfigMap)
		//if err != nil {
		//	return
		//}
	}
	// 数据库初始化
	dao.InitDB()
	server := routerInit()
	/**
		创建服务
	**/
	service = web.NewService(
		web.Name(serverConf.ServiceName),                                    //服务名
		web.Address(serverConf.Host+":"+serverConf.Port),                    //服务地址
		web.Handler(server),                                                 //gin服务
		web.Registry(consul.GetConsul(consulConf.Host+":"+consulConf.Port)), //注册中心
	)

}
func GetServer() web.Service {
	return service
}
