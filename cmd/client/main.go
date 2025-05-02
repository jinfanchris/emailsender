package main

import (
	"mailsend/pkg/client"
	"mailsend/pkg/config"
	"mailsend/pkg/log"

	"github.com/sirupsen/logrus"
)

func main() {

	a := ParseArgs()

	if err := log.Setup(true); err != nil {
		panic(err)
	}

	logrus.Infof("args: %#v", a)

	c, err := config.LoadConfig(a.ConfigF)
	if err != nil {
		logrus.Panic(err)
	}

	// logrus.Infof("config: %#v", c)

	cl, err := client.NewClient(a.Host+":"+a.Port, c.ApiKey, a.CertF, "hhost")
	if err != nil {
		logrus.Panic(err)
	}

	err = cl.SendEmail(
		c.Receiver.Account,
		"hhhhhhhhhh",
		"Test for fun - from lab",
	)

	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("Email sent successfully to %s", c.Receiver.Account)

}
