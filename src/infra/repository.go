package infra

import (
	"broker/src/infra/amqp"
	"broker/src/model"
	appTypes "broker/src/types"
	"bytes"
	"errors"
)

type publishRepositoy struct {
	pub amqp.Publisher
}

func (pr *publishRepositoy) Save(p model.Publish, c *appTypes.Counter) (err error) {
	resCh := make(chan error)

	go func() {
		defer close(resCh)

		for _, d := range p.Data {
			resCh <- pr.pub.Publish(p.Topic, d)
			c.Inc()
		}
	}()

	var buff bytes.Buffer
	for e := range resCh {
		if e != nil {
			buff.WriteString(e.Error() + ";")
		}
	}

	errMsgs := buff.String()
	if len(errMsgs) > 0 {
		err = errors.New(errMsgs)
	}

	return err
}

func NewPublishRepository() PublishDataHandler {
	return &publishRepositoy{
		pub: amqp.NewPublisher(),
	}
}
