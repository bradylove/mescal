package store_test

import "github.com/bradylove/mescal/pkg/mescal"

func (s *StoreTestSuite) TestSetAndGet() {
	inkv := &mescal.KeyValue{
		Key:   "a-key",
		Value: &mescal.KeyValue_Text{Text: "a-value"},
	}

	err := s.store.Set(inkv)
	s.Nil(err)

	outkv, err := s.store.Get("a-key")
	s.Nil(err)
	s.Equal(outkv.Key, "a-key")
	s.Equal(outkv.GetText(), "a-value")
}
