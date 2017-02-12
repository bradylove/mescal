#!/usr/bin/env bash

# Production dependencies
go get github.com/golang/protobuf/{proto,protoc-gen-go}
go get golang.org/x/net/context

# Testing dependencies
go get github.com/satori/go.uuid
go get github.com/stretchr/testify
go get github.com/nelsam/hel
go get github.com/kyoh86/richgo
