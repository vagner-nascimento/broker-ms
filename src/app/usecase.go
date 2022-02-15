package app

import (
	"broker/src/model"
	appTypes "broker/src/types"
)

func savePublish(p model.Publish, c *appTypes.Counter) error {
	pEnt := newPublishEntity(p)

	return pEnt.Save(c)
}
