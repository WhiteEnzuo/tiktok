package main

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
import (
	"VideoService/admin"
	"VideoService/model"
	"common/RabbitMq"
	"common/Redis"
	"common/Result"
	"common/consul"
	"common/mysql"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	if false {
		c := model.ContributeTask{
			//UserId:123,
			//VideoId: 123,
			//PictureId: 456,
			//VideoTitle: "123",

		}
		//c.ID=1
		a, err := c.QueryRandomId(15)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(a[0].VideoId)
		//fmt.Println(arr)

	}
	if false {
		f := model.File{
			VideoUrl: "/123",
			//Md5: "456",
		}
		err := f.QueryByUrl()
		fmt.Println(f)
		if err != nil {
			fmt.Println(err)
		}
	}
	if false {
		//配置中心
		config := consul.NewConfig("8.130.28.213", "8500")
		var test map[string]interface{}
		err := config.GetConsulConfig("Video/mysql", &test)
		if err != nil {
			fmt.Println(err)
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
		m := mysql.NewMysql("root", "root", "localhost", "3306", "java")
		server := m.Server
		server.Name()
	}
	if false {
		//Redis
		redis := Redis.NewRedis("127.0.0.1", "6379", "1")
		fmt.Println(redis.Keys("*"))

	}
	if false {
		//RabbitMq
		mq := RabbitMq.NewRabbitMq("admin", "admin", "8.130.28.213", "5672")
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
		//genToken, _ := token.GenToken(123, "456")
		//parseToken, _ := token.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImVuenVvIiwiZXhwIjoxNjc2ODkzMDA0LCJpc3MiOiJibHVlYmVsbCJ9.HQwV5j7zDq8lCg-2xXoKvocrkdU05mz8CPiMQnwtfp")
		//fmt.Println(parseToken)
	}

}
