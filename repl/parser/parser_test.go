package parser_test

import (
	. "github.com/bradylove/mescal/repl/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("can parse a get command", func() {
		parser := NewParser("get some_key")
		statement, err := parser.Parse()
		Expect(err).To(BeNil())

		Expect(statement.Action).To(Equal(GET))
		Expect(statement.Key).To(Equal("some_key"))
	})

	It("can parse a set command", func() {
		parser := NewParser("set some_key 123456789 Hello World")
		statement, err := parser.Parse()
		Expect(err).To(BeNil())

		Expect(statement.Action).To(Equal(SET))
		Expect(statement.Key).To(Equal("some_key"))
		Expect(statement.Expiry).To(Equal(int64(123456789)))
		Expect(statement.Value).To(Equal("Hello World"))
	})
})
