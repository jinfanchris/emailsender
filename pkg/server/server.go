package server

import (
	"context"
	"crypto/tls"
	pb "mailsend/pkg/grpc/mailer"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedMailServiceServer
	apiKey string
	smtp   smtpConfig
}

type smtpConfig struct {
	Account  string
	Password string
	Server   string
	Port     string
}

func NewServer(apiKey, smtpAccount, smtpPassword, smtpServer, smtpPort string) *Server {
	return &Server{
		apiKey: apiKey,
		smtp: smtpConfig{
			Account:  smtpAccount,
			Password: smtpPassword,
			Server:   smtpServer,
			Port:     smtpPort,
		},
	}
}

func (s *Server) SendMail(ctx context.Context, req *pb.MailRequest) (resp *pb.MailReply, err error) {
	apiKey := req.GetApiKey()
	rcv := req.GetReceiver()
	sub := req.GetSubject()
	bdy := req.GetBody()

	logrus.Infof("apiKey: %s", apiKey)
	logrus.Infof("rcv: %s", rcv)
	logrus.Infof("sub: %s", sub)
	logrus.Infof("bdy: %s", bdy)

	if apiKey != s.apiKey {
		err = ErrInvalidAPI
		logrus.Errorf("Invalid API key: %s v.s. %s", apiKey, s.apiKey)
		return
	}

	if err = sendMail(s.smtp.Account, rcv, sub, bdy, s.smtp.Server, s.smtp.Port, s.smtp.Password); err != nil {
		logrus.Errorf("Failed to send mail: %v", err)
		return
	}

	resp = &pb.MailReply{
		Status: "OK",
	}

	logrus.Infof("Mail sent successfully to %s", rcv)

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
