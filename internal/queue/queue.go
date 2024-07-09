package queue

import (
	"context"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/zeromicro/go-zero/core/logc"
)

type Queue struct {
	Producer *nsq.Producer
}

// NewNsqProducer 生产者
func NewNsqProducer(host string) *Queue {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(host, config)
	if err != nil {
		panic(err)
	}
	defer producer.Stop()

	return &Queue{
		Producer: producer,
	}
}

func (q *Queue) Send(topic string, body []byte) error {
	ctx := context.Background()
	err := q.Producer.Publish(topic, body)
	if err != nil {
		logc.Errorf(ctx, "send message to nsq error: %v", err)
		return err
	}

	return nil
}

type Msg struct {
	Title string
}

func (m *Msg) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return nil
}

func NewNsqConsumer(topic, channel, host string) error {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		logc.Errorf(context.Background(), "create nsq consumer error: %v", err)
		fmt.Printf("create nsq consumer error: %v\n", err)
		return err
	}

	msg := &Msg{Title: topic + " " + channel}

	consumer.AddHandler(msg)

	if err := consumer.ConnectToNSQD(host); err != nil {
		logc.Errorf(context.Background(), "connect to nsqd error: %v", err)
		fmt.Printf("connect to nsqd error: %v", err)
		return err
	}

	return nil
}
