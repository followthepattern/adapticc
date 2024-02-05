package config

import validation "github.com/go-ozzo/ozzo-validation"

type Mail struct {
	Host     string
	Addr     string
	Username string
	Password string
}

func (m Mail) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Addr, validation.Required),
		validation.Field(&m.Username, validation.Required),
		validation.Field(&m.Password, validation.Required),
	)
}
