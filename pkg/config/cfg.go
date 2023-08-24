package config

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type Config struct {
	Version string `mapstructure:"version"`
	Server  Server `mapstructure:"server"`
	DB      DB     `mapstructure:"db"`
	Mail    Mail   `mapstructure:"mail"`
}

func (cfg *Config) Validate() error {
	return validation.ValidateStruct(cfg,
		validation.Field(&cfg.Server, validation.Required),
		validation.Field(&cfg.DB, validation.Required),
		validation.Field(&cfg.Mail, validation.Required),
	)
}

func Parse() (result *Config, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetEnvPrefix("ADAPTICC")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	err = viper.Unmarshal(&result)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into worker, %v", err)
	}

	err = result.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid configuration provided %v", err)
	}

	return
}
