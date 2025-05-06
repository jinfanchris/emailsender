package mailsend

import (
	"github.com/jinfanchris/mailsend/pkg/config"
	"github.com/jinfanchris/mailsend/pkg/emailer"
)

func LoadConfig(fn string) (*config.Config, error) {
	return config.LoadConfig(fn)
}

func NewEmailer(cfg *config.Config) *emailer.Emailer {
	return emailer.NewEmailer(cfg)
}
