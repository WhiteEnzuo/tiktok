package RabbitMq

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"github.com/streadway/amqp"
	"go-micro.dev/v4/util/log"
)

type RabbitMq struct {
	Username string
	Password string
	Host     string
	Port     string
	Server   *amqp.Connection
}

func (r *RabbitMq) getConnect() error {
	url := "amqp://" + r.Username + ":" + r.Password + "@" + r.Host + ":" + r.Port + "/my_vhost"
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	r.Server = conn
	return nil
}

func NewRabbitMq(username, password, host, port string) *RabbitMq {
	r := new(RabbitMq)
	r.Username = username
	r.Password = password
	r.Host = host
	r.Port = port
	err := r.getConnect()
	if err != nil {
		log.Error("RabbitMq设置有问题")
		log.Error(err)
		return nil
	}
	return r
}
