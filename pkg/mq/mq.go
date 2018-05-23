package mq

import (
	"errors"
	"fmt"

	"github.com/joaofnfernandes/autoredeploy/pkg/mqconfig"
	"github.com/streadway/amqp"
)

const MAX_QUEUE_LEN = 10

type mq struct {
	config  mqconfig.MqConfig
	channel amqp.Channel
}

// Singleton
var instance *mq

func Read(mqConfig mqconfig.MqConfig) (string, error) {
	err := instantiateIfNotExists(mqConfig)
	if err != nil {
		return "", err
	}
	ch, err := instance.channel.Consume(
		instance.config.QueueName, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to read message from queue: %s", err))
	}
	msg := <-ch
	return string(msg.Body), nil
}

func Write(mqConfig mqconfig.MqConfig, message string) error {
	err := instantiateIfNotExists(mqConfig)
	if err != nil {
		return err
	}
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}
	err = instance.channel.Publish(
		"", // exhange
		instance.config.QueueName, // routing key
		false, // mandatory
		false, // immediate
		msg,   // message content
	)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to write to message queue: %s", err))
	}
	return nil
}

func instantiateIfNotExists(mqConfig mqconfig.MqConfig) error {
	if instance != nil {
		return nil
	}
	conn, err := amqp.Dial(mqConfig.ConnectionStr.String())
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to connect to message queue: %s", err))
	}
	ch, err := conn.Channel()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create channel in message queue: %s", err))
	}
	args := amqp.Table{}
	args["x-max-length"] = int16(MAX_QUEUE_LEN)

	// Make sure queue exists
	_, err = ch.QueueDeclare(
		mqConfig.QueueName, //name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		args,               // args
	)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create default queue: %s", err))
	}
	instance = &mq{mqConfig, *ch}
	return nil
}
