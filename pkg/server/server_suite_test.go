package server_test

import (
	"testing"

	"github.com/bradylove/mescal/pkg/server"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	serverAddr string
}

func (s *ServerTestSuite) SetupTest() {
	var err error
	s.serverAddr, err = server.Start(":0")
	s.Nil(err)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
