package api_test

import (
	"context"

	"github.com/bradylove/mescal/pkg/mescal"
)

var (
	ctx = context.Background()
)

func (s *APITestSuite) TestSetWithText() {
	close(s.mockStore.SetOutput.Err)

	resp, err := s.api.Set(ctx, buildTextSetRequest("a-key", "a-value"))
	s.Nil(err)

	s.Equal(len(s.mockStore.SetCalled), 1)

	kv := resp.GetKeyValue()
	s.Equal(len(s.mockStore.SetCalled), 1)
	s.Equal(kv.Key, "a-key")
	s.Equal(kv.GetText(), "a-value")
}

func (s *APITestSuite) TestSetWithInteger() {
	close(s.mockStore.SetOutput.Err)

	resp, err := s.api.Set(ctx, buildIntegerSetRequest("a-key", 1234))
	s.Nil(err)

	kv := resp.GetKeyValue()
	s.Equal(len(s.mockStore.SetCalled), 1)
	s.Equal(kv.Key, "a-key")
	s.Equal(kv.GetInteger(), int64(1234))
}

func (s *APITestSuite) TestSetWithDecimal() {
	close(s.mockStore.SetOutput.Err)

	resp, err := s.api.Set(ctx, buildDecimalSetRequest("a-key", 1234.1234))
	s.Nil(err)

	kv := resp.GetKeyValue()
	s.Equal(len(s.mockStore.SetCalled), 1)
	s.Equal(kv.Key, "a-key")
	s.Equal(kv.GetDecimal(), float64(1234.1234))
}

func (s *APITestSuite) TestGet() {
	close(s.mockStore.GetOutput.Err)
	s.mockStore.GetOutput.KeyValue <- &mescal.KeyValue{
		Key:   "a-key",
		Value: &mescal.KeyValue_Text{Text: "a-value"},
	}

	resp, err := s.api.Get(ctx, &mescal.GetRequest{Key: "a-key"})
	s.Nil(err)

	kv := resp.GetKeyValue()
	s.Equal(kv.Key, "a-key")
	s.Equal(kv.GetText(), "a-value")
}

func buildTextSetRequest(key string, value string) *mescal.SetRequest {
	return &mescal.SetRequest{
		KeyValue: &mescal.KeyValue{
			Key:   key,
			Value: &mescal.KeyValue_Text{Text: value},
		},
	}
}

func buildIntegerSetRequest(key string, value int64) *mescal.SetRequest {
	return &mescal.SetRequest{
		KeyValue: &mescal.KeyValue{
			Key:   key,
			Value: &mescal.KeyValue_Integer{Integer: value},
		},
	}
}

func buildDecimalSetRequest(key string, value float64) *mescal.SetRequest {
	return &mescal.SetRequest{
		KeyValue: &mescal.KeyValue{
			Key:   key,
			Value: &mescal.KeyValue_Decimal{Decimal: value},
		},
	}
}
