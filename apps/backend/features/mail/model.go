package mail

type MailModel struct {
	From    string
	To      []string
	Subject string
	Text    []byte
	HTML    []byte
}
