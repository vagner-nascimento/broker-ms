package infra

import (
	"broker/src/model"
	appTypes "broker/src/types"
)

type PublishDataHandler interface {
	Save(p model.Publish, c *appTypes.Counter) error
}
