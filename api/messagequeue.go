package main

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

type mq struct {
	config  *config
	channel *amqp.Channel
}

// Singleton
var instance *mq

type config struct {
	user      string
	pass      string
	host      string
	port      string
	protocol  string
	queueName string
}

func (c *config) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/", c.protocol, c.user, c.pass, c.host, c.port)
}

func defaultConfig() *config {
	return &config{"admin", "password", "msgq", "5672", "amqp", "webhook"}
}

// todo: where should the connection be closed?
func defaultMessageQueue() (*mq, error) {
	config := defaultConfig()
	conn, err := amqp.Dial(config.String())
	if err != nil {
		errMsg := fmt.Sprintf("[API] Failed to create default message queue: %s", err)
		return nil, errors.New(errMsg)
	}

	ch, err := conn.Channel()
	if err != nil {
		errMsg := fmt.Sprintf("[API] Failed to create default message queue channel: %s", err)
		return nil, errors.New(errMsg)
	}

	// Make sure queue exists
	_, err = ch.QueueDeclare(
		config.queueName, //name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		errMsg := fmt.Sprintf("[API] Failed to create default queue: %s", err)
		return nil, errors.New(errMsg)
	}

	return &mq{config, ch}, nil
}

// todo: we should only accept json
func Write(msg string) error {
	if instance == nil {
		var err error
		instance, err = defaultMessageQueue()
		if err != nil {
			return err
		}
	}
	content := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	}

	err := instance.channel.Publish(
		"", // exhange
		instance.config.queueName, // routing key
		false,   // mandatory
		false,   // immediate
		content, // message content
	)
	if err != nil {
		return errors.New(fmt.Sprintf("[API] failed to write to message queue: %s", err))
	}
	return nil
}
