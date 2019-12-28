package main

import "os"

type Config struct {
	HttpPort string
	WsPort   string
}

func LoadConfig() Config {
	httpPort, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		httpPort = "2001"
	}
	wsPort, ok := os.LookupEnv("WS_PORT")
	if !ok {
		wsPort = "2002"
	}
	return Config{
		HttpPort: httpPort,
		WsPort:   wsPort,
	}
}
