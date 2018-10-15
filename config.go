package main

import "flag"

const (
	defaultListenAddress  = ":8080"
	defaultNamespace      = "local.test."
	defaultDatadogAddress = "127.0.0.1:8125"
)

type Config struct {
	ListenAddress          string
	StatsdCollectorAddress string
	Namesapce              string
	Verbose                bool
}

func LoadConfig() *Config {
	c := &Config{}

	flag.StringVar(&c.ListenAddress, "listen", defaultListenAddress, "Address and Port to bind health check, in host:port format")
	flag.StringVar(&c.StatsdCollectorAddress, "statsd", defaultDatadogAddress, "Address and Port to send statsd metrics, in host:port format")
	flag.StringVar(&c.Namesapce, "namespace", defaultNamespace, "Default statsd namespace")
	flag.BoolVar(&c.Verbose, "verbose", false, "Enable verbose logging")

	flag.Parse()

	return c
}
