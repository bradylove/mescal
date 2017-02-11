package api_test

import (
	"context"
	"testing"

	"github.com/bradylove/mescal/pkg/api"
	"github.com/bradylove/mescal/pkg/mescal"
	"github.com/stretchr/testify/assert"
)

func TestSetWithText(t *testing.T) {
	resp, err := doSet(buildTextSetRequest("a-key", "a-value"))
	assert.Nil(t, err)

	kv := resp.GetKeyValue()
	assert.Equal(t, kv.Key, "a-key")
	assert.Equal(t, kv.GetText(), "a-value")
}

func TestSetWithInteger(t *testing.T) {
	resp, err := doSet(buildIntegerSetRequest("a-key", 1234))
	assert.Nil(t, err)

	kv := resp.GetKeyValue()
	assert.Equal(t, kv.Key, "a-key")
	assert.Equal(t, kv.GetInteger(), int64(1234))
}

func TestSetWithDecimal(t *testing.T) {
	resp, err := doSet(buildDecimalSetRequest("a-key", 1234.1234))
	assert.Nil(t, err)

	kv := resp.GetKeyValue()
	assert.Equal(t, kv.Key, "a-key")
	assert.Equal(t, kv.GetDecimal(), float64(1234.1234))
}

func doSet(req *mescal.SetRequest) (*mescal.SetResponse, error) {
	a := api.New()
	return a.Set(context.Background(), req)
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
