package service

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "mailsend/pkg/grpc/mailer"
	"os"

	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	pb.MailServiceClient
	ServerAddr string
	ApiKey     string
	CertF      string
}

func NewClient(serverAddr, apiKey, certF, serverName string) (c *Client, err error) {

	cert, err := os.ReadFile(certF)
	if err != nil {
		return
	}

	cp := x509.NewCertPool()
	if ok := cp.AppendCertsFromPEM(cert); !ok {
		return nil, fmt.Errorf("failed to append certificate")
	}

	creds := credentials.NewTLS(&tls.Config{
		RootCAs:    cp,
		ServerName: serverName,
	})

	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}

	c = &Client{
		MailServiceClient: pb.NewMailServiceClient(conn),
		ServerAddr:        serverAddr,
		ApiKey:            apiKey,
		CertF:             certF,
	}

	logrus.Infof("Client: %#v", c)

	return
}

func (c *Client) SendEmail(to, subject, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.MailRequest{
		ApiKey:   c.ApiKey,
		Receiver: to,
		Subject:  subject,
		Body:     body,
	}

	// Call the SendEmail method on the server
	Uuid, err := c.MailServiceClient.SendMail(ctx, req)
	if err != nil {
		return err
	}

	u := Uuid.GetUuid()
	logrus.Infof("Uuid: %s", u)
	tk := time.NewTicker(1 * time.Second)
	for range tk.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		req := &pb.Uuid{Uuid: u}
		s, err := c.MailServiceClient.State(ctx, req)
		if err != nil {
			return err
		}
		st := state(s.Status)
		switch st {
		case Failed:
			logrus.Error(s.Info)
			return fmt.Errorf("Error: %s", s.Info)
		case Sending:
			fmt.Print(".")
			continue
		case Sent:
			logrus.Infof("Sent")
			return nil
		}

	}

	return nil
}
