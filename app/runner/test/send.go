package main

import (
	"fmt"
	"galileo/app/runner/internal/biz"
	"galileo/pkg/transport/broker"
	"galileo/pkg/transport/broker/rabbitmq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//
//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Panicf("%s: %s", msg, err)
//	}
//}
//func main() {
//	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//	failOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//	ch, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	q, err := ch.QueueDeclare(
//		"hello",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	failOnError(err, "Failed to declare a queue")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	body := "Hello, world!"
//
//	err = ch.PublishWithContext(ctx,
//		"",
//		q.Name,
//		false,
//		false,
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        []byte(body),
//		})
//	failOnError(err, "Failed to publish a message")
//	log.Printf(" [x] Sent %s\n", body)
//}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := rabbitmq.NewBroker(
		broker.WithAddress("amqp://guest:guest@localhost:5672"),
		broker.WithCodec("json"),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		log.Panicf("can not connect to broker: %v", err)
	}

	var msg biz.Runner
	startTime := time.Now()
	msg.Id = 1
	msg.Timestamp = startTime.Unix()
	if err := b.Publish("logger.runner.ts", msg); err != nil {
		log.Panicf("can not publish to broker: %v", err)
	}
	fmt.Printf("Publish msg: %v\n", msg)
	<-interrupt
}
