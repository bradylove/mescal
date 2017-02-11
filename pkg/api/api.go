package api

import (
	"github.com/bradylove/mescal/pkg/mescal"
	"golang.org/x/net/context"
)

type Store interface {
	Set(keyValue *mescal.KeyValue) (err error)
	Get(key string) (keyValue *mescal.KeyValue, err error)
}

type API struct {
	store Store
}

func New(s Store) *API {
	return &API{store: s}
}

func (a *API) Set(ctx context.Context, req *mescal.SetRequest) (*mescal.SetResponse, error) {
	if err := a.store.Set(req.GetKeyValue()); err != nil {
		return nil, err
	}

	return &mescal.SetResponse{KeyValue: req.GetKeyValue()}, nil
}

func (a *API) Get(ctx context.Context, req *mescal.GetRequest) (*mescal.GetResponse, error) {
	kv, err := a.store.Get(req.Key)
	if err != nil {
		return nil, err
	}

	return &mescal.GetResponse{KeyValue: kv}, nil
}
