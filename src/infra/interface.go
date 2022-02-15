package infra

import "broker/src/model"

type PublishDataHandler interface {
	Save(p model.Publish) (err error)
}
