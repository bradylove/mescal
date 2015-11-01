package main

import (
	"crypto/tls"
	"crypto/x509"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
)

type Config struct {
	Port         string
	TLSCrtPath   string
	TLSKeyPath   string
	ClientCAPath string
}

func NewConfig() Config {
	port := kingpin.Flag("port", "Port to run the server on.").Short('p').Default("4567").String()

	// Flags needed: tls_crt, tls_key, root_ca, verify_client_crt
	tlsCrt := kingpin.Flag("tls_crt", "Path to server's TLS certificate in PEM format. Enables TLS and requires tls_key.").String()
	tlsKey := kingpin.Flag("tls_key", "Path to server's TLS key in PEM format. Enables TLS and requires tls_cert.").String()
	clientCa := kingpin.Flag("client_ca", "Path to server's client CA file in PEM format. Enables TLS client auth.").String()

	// TODO: If tls_crt or tls_key given. Validate that both are given.

	kingpin.Parse()

	return Config{
		Port:         *port,
		TLSCrtPath:   *tlsCrt,
		TLSKeyPath:   *tlsKey,
		ClientCAPath: *clientCa,
	}
}

func (c Config) TLSEnabled() bool {
	if c.TLSCrtPath != "" && c.TLSKeyPath != "" {
		return true
	}

	return false
}

func (c Config) TLSClientAuthEnabled() bool {
	return c.ClientCAPath != ""
}

func (c Config) ClientCAsPool() *x509.CertPool {
	certPool := x509.NewCertPool()
	caFile, err := ioutil.ReadFile(c.ClientCAPath)
	if err != nil {
		panic(err)
	}

	if ok := certPool.AppendCertsFromPEM(caFile); !ok {
		panic("Failed to add client CA to cert pool.")
	}

	return certPool
}

func (c Config) TLSCertificate() tls.Certificate {
	keyPair, err := tls.LoadX509KeyPair(c.TLSCrtPath, c.TLSKeyPath)
	if err != nil {
		panic(err)
	}

	return keyPair
}

func (c Config) TLSConfig() *tls.Config {
	certs := []tls.Certificate{c.TLSCertificate()}

	tlsConfig := tls.Config{
		Certificates: certs,
	}

	if c.TLSClientAuthEnabled() {
		tlsConfig.ClientCAs = c.ClientCAsPool()
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}

	return &tlsConfig
}
