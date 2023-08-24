package config

import validation "github.com/go-ozzo/ozzo-validation"

type Server struct {
	Host       string
	Port       string
	HmacSecret string `mapstructure:"hmac_secret"`
	LogLevel   int    `mapstructure:"log_level"`
}

func (cfg *Server) Validate() error {
	return validation.ValidateStruct(cfg,
		validation.Field(&cfg.LogLevel, validation.Required),
		validation.Field(&cfg.Host, validation.Required),
		validation.Field(&cfg.Port, validation.Required),
	)
}
