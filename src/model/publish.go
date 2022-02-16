package model

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Publish struct {
	Data  []interface{} `json:"data" validate:"required,min=1"`
	Topic string        `json:"topic" validate:"required,min=3"`
}

func (p *Publish) Validate() (valErrs validator.ValidationErrors) {
	v := validator.New()

	if err := v.Struct(*p); err != nil {
		valErrs = err.(validator.ValidationErrors)
	}

	return valErrs
}

func NewPublishesFromJsonBytes(bys []byte) (ps []Publish, err error) {
	err = json.Unmarshal(bys, &ps)

	return ps, err
}
