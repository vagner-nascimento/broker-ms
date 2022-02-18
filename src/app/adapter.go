package app

import (
	"broker/src/model"
	appTypes "broker/src/types"
	"fmt"
)

type publishAdapter struct{}

func (p *publishAdapter) SaveAll(ps []model.Publish) <-chan error {
	res := make(chan error)

	go func() {
		defer close(res)

		var count appTypes.Counter
		for _, p := range ps {
			res <- savePublish(p, &count)
		}

		fmt.Println("publishAdapter.save ", count.Get(), " messages sent")
	}()

	return res
}

func NewPublishAdapter() PublishHandler {
	return &publishAdapter{}
}
