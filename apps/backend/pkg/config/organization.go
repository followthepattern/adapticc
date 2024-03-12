package config

import validation "github.com/go-ozzo/ozzo-validation"

type Organization struct {
	Name  string
	Email string
	Url   string
}

func (o Organization) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Name, validation.Required),
		validation.Field(&o.Email, validation.Required),
		validation.Field(&o.Url, validation.Required),
	)
}
