package app

import (
	"broker/src/infra"
	"broker/src/model"
)

type publishEntity struct {
	pub  *model.Publish
	repo infra.PublishDataHandler
}

func (p *publishEntity) Save() error {
	return p.repo.Save(*p.pub)
}

func newPublishEntity(p model.Publish) publishEntity {
	return publishEntity{
		pub:  &p,
		repo: infra.NewPublishRepository(),
	}
}
