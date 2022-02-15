package app

import "broker/src/model"

func savePublishes(ps []model.Publish) (err error) {
	// TODO: use channel to async proccess here
	for _, p := range ps {
		pEnt := newPublishEntity(p)

		if err = pEnt.Save(); err != nil {
			break
		}
	}

	return err
}
