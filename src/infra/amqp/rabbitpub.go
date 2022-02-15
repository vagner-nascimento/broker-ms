package amqp

import "fmt"

type rabbitPub struct{}

func (rp *rabbitPub) Publish(topic string, data interface{}) (err error) {
	//TODO impl rabbit pub
	fmt.Println("rabbitPub.Publish message sent to topic: "+topic, data)
	return nil
}

func NewPublisher() Publisher {
	return &rabbitPub{}
}
