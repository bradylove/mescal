package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type Config struct {
	Port            string
	TLSCrtPath      string
	TLSKeyPath      string
	RootCAPath      string
	VerifyClientCrt bool
}

func NewConfig() Config {
	port := kingpin.Flag("port", "Port to run the server on.").Short('p').String()

	// Flags needed: tls_crt, tls_key, root_ca, verify_client_crt
	tlsCrt := kingpin.Flag("tls_crt", "Path to server's TLS certificate in PEM format").String()
	tlsKey := kingpin.Flag("tls_key", "Path to server's TLS key in PEM format").String()
	rootCa := kingpin.Flag("root_ca", "Path to server's root CA file in PEM format").String()
	verifyClientCrt := kingpin.Flag("verify_client_crt", "Whether to validate the client certificate or not.").Bool()

	kingpin.Parse()

	return Config{
		Port:            *port,
		TLSCrtPath:      *tlsCrt,
		TLSKeyPath:      *tlsKey,
		RootCAPath:      *rootCa,
		VerifyClientCrt: *verifyClientCrt,
	}
}
