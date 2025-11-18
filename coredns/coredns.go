package main

//go:generate go run directives_generate.go

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
	ds "github.com/razin99/containers/coredns/core/dnsserver"
	_ "github.com/razin99/containers/coredns/core/plugin"
)

func init() {
	dnsserver.Directives = ds.Directives
}

func main() {
	coremain.Run()
}
