package main

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

const MAX_QUEUE_LEN = 10

type mq struct {
	config  *ApiServerConfig
	channel *amqp.Channel
}

// Singleton
var instance *mq

// TODO: where should the connection be closed?
func createMessageQueue(cfg ApiServerConfig) (*mq, error) {

	conn, err := amqp.Dial(cfg.mqConfig.connectionStr.String())
	if err != nil {
		errMsg := fmt.Sprintf("[API] Failed to connect to message queue: %s", err)
		return nil, errors.New(errMsg)
	}

	ch, err := conn.Channel()
	if err != nil {
		errMsg := fmt.Sprintf("[API] Failed to create channel in message queue: %s", err)
		return nil, errors.New(errMsg)
	}

	// Limit queue length. When queue is full, older messages are discarded
	args := amqp.Table{}
	args["x-max-length"] = int16(MAX_QUEUE_LEN)

	// Make sure queue exists
	_, err = ch.QueueDeclare(
		cfg.mqConfig.queueName, //name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // args
	)
	if err != nil {
		errMsg := fmt.Sprintf("[API] Failed to create default queue: %s", err)
		return nil, errors.New(errMsg)
	}

	return &mq{&cfg, ch}, nil
}

// TODO: we should only accept json
func Write(cfg ApiServerConfig, msg string) error {
	if instance == nil {
		var err error
		instance, err = createMessageQueue(cfg)
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
		cfg.mqConfig.queueName, // routing key
		false,   // mandatory
		false,   // immediate
		content, // message content
	)
	if err != nil {
		return errors.New(fmt.Sprintf("[API] failed to write to message queue: %s", err))
	}
	return nil
}
