#!/usr/bin/env bash
go get -u -t zombiezen.com/go/capnproto/...
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega

ginkgo -r
