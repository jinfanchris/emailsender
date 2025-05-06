package emailer

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"mailsend/pkg/config"
	"net/smtp"
)

type Emailer struct {
	*config.Config
}

func NewEmailer(cfg *config.Config) *Emailer {
	return &Emailer{
		Config: cfg,
	}
}

func (e *Emailer) SendEmail(sub, bdy string, wait chan error) {
	go func() {
		wait <- sendMail(e.Sender.Account, e.Receiver.Account, sub, bdy, e.Sender.SMTPServer, e.Sender.Port, e.Sender.Password)
	}()

	return
}

func sendMail(sender, rcv, sub, bdy, smtpAddr, smtpPort, passWord string) error {
	e := email.Email{
		From:    sender,
		To:      []string{rcv},
		Subject: sub,
		Text:    []byte(bdy),
	}

	err := e.SendWithStartTLS(
		smtpAddr+":"+smtpPort,
		smtp.PlainAuth("", sender, passWord, smtpAddr),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         smtpAddr,
		},
	)

	return err
}
