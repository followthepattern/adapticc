package models

import validation "github.com/go-ozzo/ozzo-validation"

type Mail struct {
	From    string
	To      []string
	Subject string
	Text    []byte
	HTML    []byte
}

type EmailSignIn struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (e EmailSignIn) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required),
		validation.Field(&e.Email, validation.Required),
		validation.Field(&e.Message, validation.Required),
	)
}
