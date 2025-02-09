package mq

import (
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/biz/mw/socket"
	amqp "github.com/rabbitmq/amqp091-go"
)

// 定义notify路由
const notifyRoutingKey = "notify"
const notifyExchangeName = "notifyExchange"

// Notify 点赞文章/评论，评论，回复评论，关注,系统通知
type Notify struct {
	nHashId string
}

func notifyInit() {
	err := ch.ExchangeDeclare(
		notifyExchangeName, // name
		"direct",           // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare notify exchange")
	//初始化消费者
	notifyConsumerInit()
	hlog.Infof("成功初始化notify mq")
}

// NotifyProduce 调用该函数前，需要检查用户是否在线
func NotifyProduce(notify *Notify) {
	// 将结构体转换为 JSON 字节切片
	data, err := json.Marshal(notify)
	if err != nil {
		failOnError(err, "Failed to publish a message")
		return
	}
	//创建上下文context
	ctx, cancel := newContext()
	defer cancel()
	err = ch.PublishWithContext(ctx,
		notifyExchangeName, // exchange
		notifyRoutingKey,   // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
	failOnError(err, "Failed to publish a message")
	hlog.Infof(" [x] Sent %s\n", string(data))
}
func notifyConsumerInit() {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,             // queue name
		notifyRoutingKey,   // routing key
		notifyExchangeName, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			notify := new(Notify)
			err := json.Unmarshal(d.Body, notify)
			if err != nil {
				hlog.Error(err)
				return
			}
			//推送消息到客户端
			socket.NotifyPushToClient(notify.nHashId)
		}
	}()
}
