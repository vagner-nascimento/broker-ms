package app

import (
	"broker/src/model"
)

type publishAdapter struct{}

func (p *publishAdapter) SaveAll(ps []model.Publish) (err error) {
	return savePublishes(ps)
}

func NewPublishAdapter() PublishHandler {
	return &publishAdapter{}
}
