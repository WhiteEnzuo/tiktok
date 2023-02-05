package admin

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"Gateway/gloabl"
	"Gateway/router"
	"common/config"
	"common/consul"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/registry"
)

var c registry.Registry
var ginServer *gin.Engine
var serverConf serverConfig

type serverConfig struct {
	Host string
	Port string
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
		conf.Port = "8900"
	}
	return conf, nil
}
func guiInit() {
	ginServer = gin.Default()
	router.Register(ginServer)
}
func init() {
	consulConf, err := consulConfigInit()
	if err != nil {
		return
	}
	serverConf, err = serverConfigInit()
	if err != nil {
		return
	}
	c = consul.GetConsul(consulConf.Host + ":" + consulConf.Port)
	gloabl.SetConsul(c)
	guiInit()

}
func ServerRun() {
	err := ginServer.Run(serverConf.Host + ":" + serverConf.Port)
	if err != nil {
		return
	}
}
