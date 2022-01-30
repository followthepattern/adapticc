package config

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

const (
	ModeDev  = "dev"
	ModeProd = "prod"
)

type Cfg struct {
	Api API `mapstructure:"api"`
	DB  DB  `mapstructure:"db"`
}

func (cfg *Cfg) Validate() error {
	return validation.ValidateStruct(cfg,
		validation.Field(&cfg.Api, validation.Required),
	)
}

func Parse() (result *Cfg, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
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

	if err != nil {
		return nil, fmt.Errorf("unable to init logger %v", err)
	}
	return
}
