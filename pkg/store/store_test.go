package store_test

import (
	"fmt"
	"math/rand"

	"github.com/bradylove/mescal/pkg/mescal"
)

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

func (s *StoreTestSuite) TestGettingAKeyThatDoesNotExist() {
	_, err := s.store.Get("a-key")
	s.NotNil(err)
}

func (s *StoreTestSuite) TestSetAndGetDoNotHaveDataRace() {
	inkv := &mescal.KeyValue{
		Key:   "a-key",
		Value: &mescal.KeyValue_Text{Text: "a-value"},
	}

	go func() {
		for i := 0; i < 10000; i++ {
			s.store.Set(inkv)
		}
	}()

	for i := 0; i < 10000; i++ {
		s.store.Get("a-key")
	}
}

func (s *StoreTestSuite) TestMinimizingColisions() {
	keys := make([]string, 1000000)

	for i := range keys {
		keys[i] = randString(20)
	}

	for _, k := range keys {
		inkv := &mescal.KeyValue{
			Key:   fmt.Sprint(k),
			Value: &mescal.KeyValue_Text{Text: "a-value"},
		}

		s.store.Set(inkv)
	}

	for i, k := range keys {
		resp, err := s.store.Get(k)
		s.Nil(err, fmt.Sprint("failed on attempt #", i))
		s.Equal(resp.Key, k)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
