package service

import (
	"context"
	"crypto/tls"
	"fmt"
	pb "mailsend/pkg/grpc/mailer"
	"net/smtp"
	"sync"

	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
)

type state int

const (
	Sending state = iota
	Sent
	Failed
)

type Server struct {
	pb.UnimplementedMailServiceServer
	apiKey string
	smtp   smtpConfig

	states       map[string]state
	errors       map[string]error
	states_mutex sync.Mutex
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
		states:       make(map[string]state),
		errors:       make(map[string]error),
		states_mutex: sync.Mutex{},
	}
}

func (s *Server) State(ctx context.Context, u *pb.Uuid) (resp *pb.Status, err error) {

	uuid := u.GetUuid()
	state, ok := s.states[uuid]
	if !ok {
		err = fmt.Errorf("Unkown uuid: %s", uuid)
		return
	}

	var info string
	switch state {
	case Failed:
		info = fmt.Sprintf("Error sending mail: %v", s.errors[uuid])
	case Sent:
		info = "Mail Sent successfully"
	case Sending:
		info = "Mail Sending"
	}

	resp = &pb.Status{
		Status: int32(state),
		Info:   info,
	}
	return
}

func (s *Server) SendMail(ctx context.Context, req *pb.MailRequest) (resp *pb.Uuid, err error) {
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

	uuid := uuid.New().String()
	s.states_mutex.Lock()
	s.states[uuid] = Sending
	s.states_mutex.Unlock()
	go func(uuid string) {
		logrus.Infof("Sending Mail to %s", rcv)
		if err = sendMail(s.smtp.Account, rcv, sub, bdy, s.smtp.Server, s.smtp.Port, s.smtp.Password); err != nil {
			logrus.Errorf("Failed to send mail: %v", err)
			s.states_mutex.Lock()
			s.states[uuid] = Failed
			s.errors[uuid] = err
			s.states_mutex.Unlock()
			return
		}
		s.states_mutex.Lock()
		s.states[uuid] = Sent
		s.states_mutex.Unlock()
		logrus.Infof("Mail sent successfully to %s", rcv)
	}(uuid)

	resp = &pb.Uuid{
		Uuid: uuid,
	}

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
