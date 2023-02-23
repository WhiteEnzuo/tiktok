package main

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"UserService/admin"
	"common/RabbitMq"
	"common/Redis"
	"common/Result"
	"common/consul"
	"common/mysql"
	"common/token"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	if false {
		//配置中心
		config := consul.NewConfig("8.130.28.213", "8500")
		var test map[string]interface{}
		err := config.GetConsulConfig("Test", &test)
		if err != nil {
			return
		}
		fmt.Println(test)
	}

	if false {
		//Result用来传输
		result := Result.NewResult()
		result.OK().SetCode(201).SetDataKey("test", 123)
		fmt.Println(result.ToJsonString())
	}

	if true {
		server := admin.GetServer()
		err := server.Run()
		if err != nil {
			return
		}
	}
	if false {
		//gorm
		m := mysql.NewMysql("root", "root", "8.130.28.213", "3306", "like")
		server := m.Server
		server.Name()
	}
	if false {
		//Redis
		redis := Redis.NewRedis("8.130.28.213", "6379", "1")
		fmt.Println(redis.Keys("*"))

	}
	if false {
		//RabbitMq
		mq := RabbitMq.NewRabbitMq("admin", "admin", "8.130.28.213", "15672")
		server := mq.Server
		channel, _ := server.Channel()
		q, _ := channel.QueueDeclare(
			"hello",
			false,
			false,
			false,
			false,
			nil,
		)
		channel.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Test"),
			})
	}
	if false {
		genToken, _ := token.GenToken(123, "456")
		parseToken, _ := token.ParseToken(genToken)
		fmt.Println(parseToken)
	}

}
