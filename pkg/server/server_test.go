package server_test

import (
	"context"

	"github.com/bradylove/mescal/pkg/mescal"
	"google.golang.org/grpc"
)

var (
	ctx = context.Background()
)

func (s *ServerTestSuite) TestSetAndGet() {
	cases := []struct {
		key, value string
	}{
		{"key-1", "value-1"},
		{"key-2", "value-2"},
		{"key-3", "value-3"},
	}

	client := buildMescalClient(s.serverAddr)

	for _, c := range cases {
		_, err := client.Set(ctx, buildSetRequest(c.key, c.value))
		s.Nil(err)

		resp, err := client.Get(ctx, &mescal.GetRequest{Key: c.key})
		s.Nil(err)

		kv := resp.GetKeyValue()
		s.Equal(kv.Key, c.key)
		s.Equal(kv.GetText(), c.value)
	}
}

func buildMescalClient(addr string) mescal.MescalClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return mescal.NewMescalClient(conn)
}

func buildSetRequest(key, value string) *mescal.SetRequest {
	return &mescal.SetRequest{
		KeyValue: &mescal.KeyValue{
			Key: key,
			Value: &mescal.KeyValue_Text{
				Text: value,
			},
		},
	}
}
