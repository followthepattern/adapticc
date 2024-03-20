package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Server struct {
	Host                  string
	Port                  string
	HmacSecret            string `mapstructure:"hmac_secret"`
	Ed25519PrivateKey     string `mapstructure:"ed25519_private_key"`
	Ed25519PublicKey      string `mapstructure:"ed25519_public_key"`
	LogLevel              int    `mapstructure:"log_level"`
	GraphqlSchemaFilepath string `mapstructure:"graphql_schema_filepath"`
}

func (cfg Server) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Host, validation.Required),
		validation.Field(&cfg.Port, validation.Required),
	)
}
