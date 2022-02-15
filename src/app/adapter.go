package app

import (
	"broker/src/model"
)

type publishAdapter struct{}

func (p *publishAdapter) SaveAll(ps []model.Publish) *chan error {
	res := make(chan error)

	go func() {
		defer close(res)
		for _, p := range ps {
			res <- savePublishe(p)
		}
	}()

	return &res
}

func NewPublishAdapter() PublishHandler {
	return &publishAdapter{}
}
