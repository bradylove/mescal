package msg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMsg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Msg Suite")
}
