package config

import validation "github.com/go-ozzo/ozzo-validation"

type Cerbos struct {
	Address string `mapstructure:"address"`
}

func (cfg Cerbos) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Address, validation.Required),
	)
}
