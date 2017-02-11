package api_test

//go:generate hel

import (
	"testing"

	"github.com/bradylove/mescal/pkg/api"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
	mockStore *mockStore
	api       *api.API
}

func (s *APITestSuite) SetupTest() {
	s.mockStore = newMockStore()
	s.api = api.New(s.mockStore)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
