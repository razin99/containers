package main

//go:generate go run directives_generate.go

import (
	"os"
	"testing"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	ds "github.com/razin99/containers/coredns/core/dnsserver"
	_ "github.com/razin99/containers/coredns/core/plugin"
)

func TestCorefile(t *testing.T) {
	dnsserver.Directives = ds.Directives
	dnsserver.Quiet = true
	dnsserver.Port = "5353"
	caddy.Quiet = true

	contents, err := os.ReadFile("./Corefile")
	if err != nil {
		t.Fail()
	}

	ci, err := caddy.Start(caddy.CaddyfileInput{
		ServerTypeName: "dns",
		Filepath:       "Coredns",
		Contents:       contents,
	})
	if err != nil {
		t.Error(err)
	}
	defer ci.Stop()
}

func TestInvalidPlugin(t *testing.T) {
	dnsserver.Directives = ds.Directives
	dnsserver.Quiet = true
	dnsserver.Port = "5353"
	caddy.Quiet = true

	ci, err := caddy.Start(caddy.CaddyfileInput{
		ServerTypeName: "dns",
		Filepath:       "Coredns",
		Contents:       []byte("example.org:53 {\n    whoami\n}\n"),
	})
	if err == nil {
		defer ci.Stop()
		t.Errorf("Plugin whoami is not installed, expected coredns to fail")
	}
}
