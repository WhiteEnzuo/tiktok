package service

import (
	"UserService/dao"
	"common/RabbitMq"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

type LikeMQ struct {
	RabbitMQ  *RabbitMq.RabbitMq
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

// NewLikeRabbitMQ 获取likeMQ的对应队列。
func NewLikeRabbitMQ(queueName string) *LikeMQ {
	likeMQ := &LikeMQ{
		RabbitMQ:  RabbitMq.NewRabbitMq("admin", "admin", "8.130.28.213", "5672"),
		queueName: queueName,
	}
	cha, err := likeMQ.RabbitMQ.Server.Channel()
	likeMQ.channel = cha
	if err != nil {
		log.Fatalf("%s:%s\n", err, "获取通道失败")
	}
	return likeMQ
}

// Publish 点赞操作的发布配置
func (l *LikeMQ) Publish(message string) {

	_, err := l.channel.QueueDeclare(
		// 队列名字
		l.queueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否独占队列
		false,
		// 是否阻塞
		false,
		// 其他参数
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = l.channel.Publish(
		l.exchange,
		l.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		panic(err)
	}

}

// Consumer 消费者
func (l *LikeMQ) Consumer() {

	_, err := l.channel.QueueDeclare(l.queueName, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	messages, err1 := l.channel.Consume(
		l.queueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 消息队列是否阻塞
		false,
		nil,
	)
	if err1 != nil {
		panic(err1)
	}

	// 防止主进程过早结束子进程还没有运行完
	forever := make(chan bool)
	// 一直监听生产者
	switch l.queueName {
	case "likeAdd":
		go func() {
			// 点赞消费队列
			println(l.queueName)
			for d := range messages {
				// 参数解析
				params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
				uid, _ := strconv.ParseInt(params[0], 10, 64)
				vid, _ := strconv.ParseInt(params[1], 10, 64)
				// 尝试操作数据库的次数
				for i := 0; i < 3; i++ {
					flag := false
					like, err2 := dao.GetLike(int(uid), int(vid))
					if err2 != nil {
						flag = true
					} else {
						if like == 0 {
							if err2 = dao.InsertLike(int(uid), int(vid)); err2 != nil {
								flag = true
							}
						} else if like == 2 {
							if err2 = dao.UpdateLike(int(uid), int(vid), 1); err2 != nil {
								flag = true
							}
						}
						// 若没出错就结束
						if flag == false {
							break
						}
					}
				}
			}
		}()
	case "likeDel":
		go func() {
			// 取消赞消费队列
			println(l.queueName)
			for d := range messages {
				params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
				uid, _ := strconv.ParseInt(params[0], 10, 64)
				vid, _ := strconv.ParseInt(params[1], 10, 64)
				for i := 0; i < 3; i++ {
					flag := false
					like, err2 := dao.GetLike(int(uid), int(vid))
					if err2 != nil {
						flag = true
					} else {
						if like == 0 {
							flag = true
						} else if like == 1 {
							if err2 = dao.UpdateLike(int(uid), int(vid), 0); err2 != nil {
								flag = true
							}
						}
					}
					if flag == false {
						break
					}
				}
			}
		}()
	}
	<-forever
}

var RmqLikeAdd *LikeMQ
var RmqLikeDel *LikeMQ

func InitMQ() {
	RmqLikeAdd = NewLikeRabbitMQ("likeAdd")
	go RmqLikeAdd.Consumer()

	RmqLikeDel = NewLikeRabbitMQ("likeDel")
	go RmqLikeDel.Consumer()
}
