package msg_test

import (
	. "github.com/bradylove/mescal/msg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"time"
)

var _ = Describe("Command", func() {
	var buffer *bytes.Buffer

	BeforeEach(func() {
		var b []byte
		buffer = bytes.NewBuffer(b)
	})

	It("encodes a get command", func() {
		cmd := NewCommand("12345", NewGetCommand("foo"))
		err := cmd.Encode(buffer)

		Expect(err).To(BeNil())
		Expect(buffer.Bytes()).To(Equal(
			[]byte{0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 0, 0, 0, 0, 0, 13, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 34, 0, 0, 0, 102, 111, 111, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 0, 0, 0},
		))
	})

	It("decodes a get command", func() {
		encodedMsg := []byte{0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 0, 0, 0, 0, 0, 13, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 34, 0, 0, 0, 102, 111, 111, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 0, 0, 0}
		buf := bytes.NewBuffer(encodedMsg)

		cmd, err := DecodeCommand(buf)
		Expect(err).To(BeNil())
		Expect(cmd.Id).To(Equal("12345"))
		Expect(cmd.Action).To(Equal(ActionGet))

		subCmd := cmd.SubCommand.(GetCommand)
		Expect(subCmd.Key).To(Equal("foo"))
	})

	It("encodes and decodes a handshake command", func() {
		cmd := NewCommand("12345", NewHandshakeCommand("Go v0.1.1"))
		err := cmd.Encode(buffer)

		Expect(err).To(BeNil())

		decodedCmd, err := DecodeCommand(buffer)
		Expect(decodedCmd.Id).To(Equal("12345"))
		Expect(decodedCmd.Action).To(Equal(ActionHandshake))
	})

	It("encodes and decodes a set command", func() {
		expiry := time.Now().Unix()
		cmd := NewCommand("12345", NewSetCommand("foo", "bar", expiry))
		err := cmd.Encode(buffer)

		Expect(err).To(BeNil())

		decodedCmd, err := DecodeCommand(buffer)
		Expect(decodedCmd.Id).To(Equal("12345"))
		Expect(decodedCmd.Action).To(Equal(ActionSet))

		subCmd := cmd.SubCommand.(SetCommand)
		Expect(subCmd.Key).To(Equal("foo"))
		Expect(subCmd.Value).To(Equal("bar"))
		Expect(subCmd.Expiry).To(Equal(expiry))
	})
})
