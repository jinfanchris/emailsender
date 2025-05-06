package main

import (
	"fmt"
	"mailsend/pkg/config"
	"mailsend/pkg/emailer"
	"mailsend/pkg/log"
	"time"

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

	emailer := emailer.NewEmailer(c)

	Subject := "Test Email from Go (TLS)"
	Text := "Hello,\n\nThis is a test email sent using Go with TLS.\n\nRegards,\nGo App"

	wait := make(chan error)
	emailer.SendEmail(Subject, Text, wait)

	tk := time.NewTicker(1 * time.Second)

	err = func() error {
		for {
			select {
			case <-tk.C:
				fmt.Print(".")
			case err := <-wait:
				return err

			}
		}
	}()

	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("Email sent successfully to %s", c.Receiver.Account)

}
