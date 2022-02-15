package app

import (
	"broker/src/infra"
	"broker/src/model"
	appTypes "broker/src/types"
)

type publishEntity struct {
	pub  *model.Publish
	repo infra.PublishDataHandler
}

func (p *publishEntity) Save(c *appTypes.Counter) error {
	return p.repo.Save(*p.pub, c)
}

func newPublishEntity(p model.Publish) publishEntity {
	return publishEntity{
		pub:  &p,
		repo: infra.NewPublishRepository(),
	}
}
