package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Client struct {
	Host    string
	channel *amqp.Channel
}

func NewClient(host string) *Client {
	return &Client{Host: host}
}

func (c *Client) ConnectToServer() error {
	conn, err := amqp.Dial(c.Host)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		return err
	}

	c.channel, err = conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (c *Client) CloseConnect() error  {
	return  c.channel.Close()
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
