package mail

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

//go:generate mockgen -destination=../../../mocks/email.go -package=mocks -source=./client.go Email
type Email interface {
	SetFrom(from string)
	SetTo(to []string)
	SetSubject(subject string)
	SetText(text []byte)
	SetHTML(html []byte)
	Send(addr string, a smtp.Auth) error
}

type EmailWrapper struct {
	*email.Email
}

func NewClient() Email {
	return &EmailWrapper{email.NewEmail()}
}

func (e *EmailWrapper) SetFrom(from string) {
	e.From = from
}

func (e *EmailWrapper) SetTo(to []string) {
	e.To = to
}

func (e *EmailWrapper) SetSubject(subject string) {
	e.Subject = subject
}

func (e *EmailWrapper) SetText(text []byte) {
	e.Text = text
}

func (e *EmailWrapper) SetHTML(html []byte) {
	e.HTML = html
}
