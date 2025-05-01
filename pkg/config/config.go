package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

// Config structure matches the TOML layout
type Config struct {
	ApiKey   string   `toml:"api_key"`
	Sender   Sender   `toml:"Sender"`
	Receiver Receiver `toml:"Receiver"`
}

type Sender struct {
	Account    string `toml:"account"`
	Password   string `toml:"password"`
	SMTPServer string `toml:"smtp_server"`
	Port       string `toml:"port"`
}

type Receiver struct {
	Account string `toml:"account"`
}

// LoadConfig loads the TOML config from a file path
func LoadConfig(path string) (*Config, error) {
	var cfg Config

	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

var DefaultFile = "runtime/config.toml"

func LoadDefaultConfig() (*Config, error) {
	return LoadConfig(DefaultFile)
}
