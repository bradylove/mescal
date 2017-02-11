package store_test

//go:generate hel

import (
	"testing"

	"github.com/bradylove/mescal/pkg/store"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
	store *store.Store
}

func (s *StoreTestSuite) SetupTest() {
	s.store = store.New()
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
