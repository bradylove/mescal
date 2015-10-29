package msg_test

import (
	. "github.com/bradylove/mescal/msg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
)

var _ = Describe("Result", func() {
	var buffer *bytes.Buffer

	BeforeEach(func() {
		var b []byte
		buffer = bytes.NewBuffer(b)
	})

	It("can encode and decode a get result", func() {
		expiry := int64(1442537301)
		r := NewResult("12345", StatusSuccess, NewGetResult("foo", "bar", expiry))

		err := r.Encode(buffer)
		Expect(err).To(BeNil())
		Expect(buffer.Bytes()).To(Equal(
			[]byte{0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 0, 1, 0, 0, 0, 25, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 85, 95, 251, 85, 0, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 102, 111, 111, 0, 0, 0, 0, 0, 98, 97, 114, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 0, 0, 0},
		))

		decodedRes, err := DecodeResult(buffer)
		Expect(err).To(BeNil())
		Expect(decodedRes.Action).To(Equal(ActionGet))
		Expect(decodedRes.Id).To(Equal("12345"))

		subResult := decodedRes.SubResult.(GetResult)
		Expect(subResult.Key).To(Equal("foo"))
		Expect(subResult.Value).To(Equal("bar"))
		Expect(subResult.Expiry).To(Equal(expiry))
	})

	It("can encode and decode a handshake result", func() {
		r := NewResult("12345", StatusSuccess, NewHandshakeResult("123"))

		err := r.Encode(buffer)
		Expect(err).To(BeNil())

		decodedRes, err := DecodeResult(buffer)
		Expect(err).To(BeNil())
		Expect(decodedRes.Action).To(Equal(ActionHandshake))
		Expect(decodedRes.Id).To(Equal("12345"))

		subResult := decodedRes.SubResult.(HandshakeResult)
		Expect(subResult.ClientId).To(Equal("123"))
	})

	It("can encode and decode a set result", func() {
		r := NewResult("12345", StatusSuccess, NewSetResult())

		err := r.Encode(buffer)
		Expect(err).To(BeNil())

		decodedRes, err := DecodeResult(buffer)
		Expect(err).To(BeNil())
		Expect(decodedRes.Action).To(Equal(ActionSet))
		Expect(decodedRes.Id).To(Equal("12345"))
	})
})
