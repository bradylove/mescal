package main_test

import (
	. "github.com/bradylove/mescal/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// "bytes"
	"github.com/bradylove/mescal/msg"
	"net"
	"time"
)

var _ = Describe("Client", func() {
	var client net.Conn

	BeforeSuite(func() {
		cfg := NewConfig()
		go RunServer(cfg)

		time.Sleep(time.Millisecond * 100)

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
})
