package mail

import (
	"os"
	"strconv"

	mail "github.com/xhit/go-simple-mail/v2"
)

// Sender .
type Sender struct {
	client *mail.SMTPClient
}

// NewSender .
func NewSender() *Sender {
	return &Sender{
		client: newClient(),
	}
}

// SendTextMail テキストメールを送信する
func (s *Sender) SendTextMail(toList []string, subject, body string) error {
	from := os.Getenv("MAIL_FROM_ADDRESS")
	if from == "" {
		panic("MAIL_FROM_ADDRESS is empty")
	}

	email := mail.NewMSG()
	email.SetFrom(from).
		AddTo(toList...).
		SetSubject(subject)
	email.SetBody(mail.TextPlain, body)

	if err := email.Send(s.client); err != nil {
		return err
	}
	return nil
}

func newClient() *mail.SMTPClient {
	host := os.Getenv("MAIL_HOST")
	if host == "" {
		panic("MAIL_HOST is empty")
	}
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic(err)
	}

	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	client, err := server.Connect()
	if err != nil {
		panic(err)
	}

	return client
}
