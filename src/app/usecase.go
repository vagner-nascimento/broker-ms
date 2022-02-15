package app

import "broker/src/model"

func savePublishe(p model.Publish) (err error) {
	pEnt := newPublishEntity(p)

	return pEnt.Save()
}
