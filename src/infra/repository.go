package infra

import (
	"broker/src/model"
)

type publishRepositoy struct{}

func (pr *publishRepositoy) Save(p model.Publish) (err error) {
	return err
}

func NewPublishRepository() PublishDataHandler {
	return &publishRepositoy{}
}
