package emailsender

import (
	"github.com/jinfanchris/emailsender/pkg/config"
	"github.com/jinfanchris/emailsender/pkg/emailer"
)

func LoadConfig(fn string) (*config.Config, error) {
	return config.LoadConfig(fn)
}

func NewEmailer(cfg *config.Config) *emailer.Emailer {
	return emailer.NewEmailer(cfg)
}
