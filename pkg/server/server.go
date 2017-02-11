package server

import (
	"log"
	"net"

	"github.com/bradylove/mescal/pkg/api"
	"github.com/bradylove/mescal/pkg/mescal"
	"github.com/bradylove/mescal/pkg/store"
	"google.golang.org/grpc"
)

func Start(addr string) (string, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return "", err
	}

	api := api.New(store.New())
	grpcServer := grpc.NewServer()
	mescal.RegisterMescalServer(grpcServer, api)

	go func() {
		log.Println("Server failed:", grpcServer.Serve(l))
	}()

	return l.Addr().String(), nil
}
