package main

import (
	"github.com/spf13/pflag"
)

type Args struct {
	Version bool
	ConfigF string
	CertF   string
	KeyF    string
	Port    string
	Host    string
}

func ParseArgs() Args {
	args := Args{}

	pflag.BoolVarP(&args.Version, "version", "v", false, "Print version")
	pflag.StringVarP(&args.ConfigF, "config", "c", "runtime/config.toml", "Path to config file")
	pflag.StringVarP(&args.CertF, "cert", "C", "runtime/cert.pem", "Path to TLS certificate")
	pflag.StringVarP(&args.Port, "port", "p", "52333", "Port to listen on")
	pflag.StringVarP(&args.Host, "host", "H", "localhost", "Host")

	return args
}
