package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type Config struct {
	Port string
}

func NewConfig() Config {
	port := kingpin.Flag("port", "Port to run the server on.").Short('p').String()

	kingpin.Parse()

	return Config{
		Port: *port,
	}
}
