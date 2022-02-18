package app

import "broker/src/model"

type PublishHandler interface {
	SaveAll(pubs []model.Publish) <-chan error
}
