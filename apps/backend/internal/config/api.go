package config

import validation "github.com/go-ozzo/ozzo-validation"

type API struct {
	Host string
	Mode string
	Port string
}

func (a API) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Host, validation.Required),
		validation.Field(&a.Port, validation.Required),
		validation.Field(&a.Mode, validation.Required, validation.In(ModeProd, ModeDev)),
	)
}
