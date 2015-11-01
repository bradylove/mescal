#!/usr/bin/env bash
go get -u -t zombiezen.com/go/capnproto/...
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go get gopkg.in/alecthomas/kingpin.v2

ginkgo -r
