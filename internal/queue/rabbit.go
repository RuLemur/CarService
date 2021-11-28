package queue

import (
	"context"
	"fmt"
	"github.com/RuLemur/CarService/internal/config"
	"github.com/RuLemur/CarService/internal/logger"
	"github.com/streadway/amqp"
)

var log = logger.NewDefaultLogger()

type Client struct {
	channel *amqp.Channel
	conn    *amqp.Connection
}

func (c *Client) Init(ctx context.Context) error {
	cfg := config.GetInstance().GetConfig()

	log.Infof("Connecting to RabbitMQ: %s", cfg.Queue.Host)
	var err error
	c.conn, err = amqp.Dial(cfg.Queue.Host)
	if err != nil {
		log.Errorf("Fail connect to RabbitMq: %s", err.Error())
		return err
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		log.Errorf("Fail to create RabbitMq channel: %s", err.Error())
		return err
	}
	log.Infof("Connected to RabbitMQ")
	return nil
}

func (c *Client) Ping(ctx context.Context) error {
	if c.conn.IsClosed() {
		return fmt.Errorf("rabbitMQ ebnulsa")
	}
	log.Debugf("Ping RabbitMQ")
	return nil
}

func (c *Client) Close() error {
	return c.CloseConnect()
}

func (c *Client) CloseConnect() error {
	log.Infof("Close RabbitMq connection")
	return c.channel.Close()
}

func (c *Client) SendMessageToQueue(message string) error {
	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := c.channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	// Handle any errors if we were unable to create the queue
	if err != nil {
		return err
	}

	err = c.channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return err
	}
	return nil
}
