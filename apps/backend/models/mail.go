package models

type Mail struct {
	From    string
	To      []string
	Subject string
	Text    []byte
	HTML    []byte
}
