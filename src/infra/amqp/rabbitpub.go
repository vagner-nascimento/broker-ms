package amqp

import (
	"encoding/json"

	stAmq "github.com/streadway/amqp"
)

type rabbitPub struct{}

func (rp *rabbitPub) Publish(topic string, data interface{}) (err error) {
	var dBys []byte

	if dBys, err = json.Marshal(data); err == nil {
		var ch *stAmq.Channel

		if ch, err = getChannel(); err == nil {
			var q stAmq.Queue

			if q, err = ch.QueueDeclare(
				topic,
				false,
				false,
				false,
				false,
				nil,
			); err == nil {
				err = ch.Publish(
					"",
					q.Name,
					false,
					false,
					stAmq.Publishing{
						ContentType: "application/json",
						Body:        dBys,
					},
				)
			}
		}
	}

	return err
}

func NewPublisher() Publisher {
	return &rabbitPub{}
}
