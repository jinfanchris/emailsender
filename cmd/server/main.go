package main

import (
	"mailsend/pkg/config"
	"mailsend/pkg/log"
	"mailsend/pkg/server"
	"net"

	pb "mailsend/pkg/grpc/mailer"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	logrus.Infof("config: %#v", c)

	creds, err := credentials.NewServerTLSFromFile(a.CertF, a.KeyF)
	if err != nil {
		logrus.Panic(err)
	}

	apiKey := c.ApiKey

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	mailServer := server.NewServer(
		apiKey,
		c.Sender.Account,
		c.Sender.Password,
		c.Sender.SMTPServer,
		c.Sender.Port,
	)

	pb.RegisterMailServiceServer(grpcServer, mailServer)

	lis, err := net.Listen("tcp", ":"+a.Port)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("Listening on port %s", a.Port)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Panic(err)
	}
}

