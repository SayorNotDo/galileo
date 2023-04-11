package rabbitmq

import (
	"galileo/pkg/transport/broker"
	"github.com/streadway/amqp"
)

type publication struct {
	d   amqp.Delivery
	m   *broker.Message
	t   string
	err error
}

func (p *publication) Ack() error {
	return p.d.Ack(false)
}

func (p *publication) Error() error {
	return p.err
}

func (p *publication) Topic() string {
	return p.t
}

func (p *publication) Message() *broker.Message {
	return p.m
}
