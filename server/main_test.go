package main_test

import (
	. "github.com/bradylove/mescal/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/bradylove/mescal/msg"
	"io/ioutil"
	"net"
	"time"
)

var _ = Describe("Client", func() {
	var client net.Conn

	BeforeSuite(func() {
		cfg := NewConfig()
		go RunServer(cfg)

		tlsCfg := Config{
			Port:       "4333",
			TLSCrtPath: "../test/cert.pem",
			TLSKeyPath: "../test/key.pem",
		}

		go RunServer(tlsCfg)

		time.Sleep(time.Millisecond * 1000)

		var err error
		client, err = net.Dial("tcp", "127.0.0.1:"+cfg.Port)
		if err != nil {
			panic(err)
		}
	})

	It("can handle a client handshake", func() {
		cmd := msg.NewCommand("12345", msg.NewHandshakeCommand("Go v0.1.1"))
		err := cmd.Encode(client)
		Expect(err).To(BeNil())

		result, err := msg.DecodeResult(client)
		Expect(err).To(BeNil())
		Expect(result.Id).To(Equal("12345"))
		Expect(result.Action).To(Equal(msg.ActionHandshake))

		subRes := result.SubResult.(msg.HandshakeResult)
		Expect(subRes.ClientId).To(BeAssignableToTypeOf("string"))
	})

	Context("With a client that has already performed handshake", func() {
		It("can handle a set command", func() {
			cmd := msg.NewCommand("12345", msg.NewSetCommand("doodle", "poodle", time.Now().Unix()))
			err := cmd.Encode(client)
			Expect(err).To(BeNil())

			result, err := msg.DecodeResult(client)
			Expect(err).To(BeNil())
			Expect(result.Id).To(Equal("12345"))
			Expect(result.Action).To(Equal(msg.ActionSet))
		})

		It("can handle a get command", func() {
			cmd := msg.NewCommand("12345", msg.NewGetCommand("doodle"))
			err := cmd.Encode(client)
			Expect(err).To(BeNil())

			result, err := msg.DecodeResult(client)
			Expect(err).To(BeNil())
			Expect(result.Id).To(Equal("12345"))
			Expect(result.Action).To(Equal(msg.ActionGet))

			subRes := result.SubResult.(msg.GetResult)
			Expect(subRes.Key).To(Equal("doodle"))
			Expect(subRes.Value).To(Equal("poodle"))
		})
	})

	Context("With TLS enabled", func() {
		It("fails to send a command with a client not using tls", func() {
			client, err := net.Dial("tcp", "127.0.0.1:4333")

			cmd := msg.NewCommand("12345", msg.NewHandshakeCommand("Go v0.1.1"))
			err = cmd.Encode(client)
			Expect(err).To(BeNil())

			var data []byte
			_, err = client.Read(data)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("EOF"))
		})

		It("allows a client to connect and send commands using tls", func() {
			certPool := x509.NewCertPool()
			caFile, err := ioutil.ReadFile("../test/ca.pem")
			Expect(err).To(BeNil())

			block, _ := pem.Decode(caFile)
			caCert, err := x509.ParseCertificate(block.Bytes)
			Expect(err).To(BeNil())

			certPool.AddCert(caCert)
			tlsCfg := tls.Config{
				RootCAs:    certPool,
				ServerName: "localhost",
			}

			client, err = tls.Dial("tcp", "127.0.0.1:4333", &tlsCfg)
			Expect(err).To(BeNil())

			cmd := msg.NewCommand("12345", msg.NewHandshakeCommand("Go v0.1.1"))
			err = cmd.Encode(client)
			Expect(err).To(BeNil())
		})
	})
})
