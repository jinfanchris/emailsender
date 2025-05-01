package main

import (
	"crypto/tls"
	"mailsend/pkg/config"
	"mailsend/pkg/log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := log.Setup(true); err != nil {
		panic(err)
	}

	logrus.Infof("Hello World")

	c, err := config.LoadDefaultConfig()
	if err != nil {
		logrus.Panic(err)
	}

	e := email.NewEmail()
	e.From = c.Sender.Account
	e.To = []string{c.Receiver.Account}
	e.Subject = "Test Email from Go (TLS)"
	e.Text = []byte("Hello,\n\nThis is a test email sent using Go with TLS.\n\nRegards,\nGo App")

	err = e.SendWithStartTLS(
		c.Sender.SMTPServer+":587",
		smtp.PlainAuth("", c.Sender.Account, c.Sender.Password, c.Sender.SMTPServer),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         c.Sender.SMTPServer,
		},
	)

	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("Email sent successfully to %s", c.Receiver.Account)

}
