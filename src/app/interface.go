package app

import "broker/src/model"

type PublishHandler interface {
	SaveAll(pubs []model.Publish) (err error)
}
