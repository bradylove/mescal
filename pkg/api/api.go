package api

import (
	"context"

	"github.com/bradylove/mescal/pkg/mescal"
)

type API struct{}

func New() *API {
	return new(API)
}

func (a *API) Set(ctx context.Context, req *mescal.SetRequest) (*mescal.SetResponse, error) {
	return &mescal.SetResponse{KeyValue: req.GetKeyValue()}, nil
}
