package config

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type DB struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

func (db DB) Validate() error {
	return validation.ValidateStruct(&db,
		validation.Field(&db.Host, validation.Required),
		validation.Field(&db.Port, validation.Required),
		validation.Field(&db.User, validation.Required),
		validation.Field(&db.Password, validation.Required),
		validation.Field(&db.DBName, validation.Required),
	)
}

func (db DB) ConnectionURL() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBName)
}
