package main

import (
	"os"

	"github.com/alecthomas/kingpin"

	"github.com/lalloni/grpcntt/client"
	"github.com/lalloni/grpcntt/server"
)

func main() {

	app := kingpin.New(os.Args[0], "A gRPC network testing tool")
	serverAddress := app.Flag("address", "Server address").Default(":80").String()
	useTLS := app.Flag("tls", "Use TLS (server mode generates a self-signed certificate, client mode skips certificate verification)").Bool()

	serve := app.Command("serve", "Start the testing server")
	san := serve.Flag("san", "List of host names / ip addresses to be added in self-signed certificate's SubjectAltName").Strings()
	org := serve.Flag("org", "Organization name to be set in self-signed certificate's DistinguishedName").String()

	ping := app.Command("ping", "Do ping test")
	count := ping.Flag("count", "Repeat count").Default("5").Uint64()
	size := ping.Flag("size", "Payload size").Default("1024").Uint64()
	sleep := ping.Flag("sleep", "Sleep duration between calls").Default("1s").Duration()
	timeout := ping.Flag("timeout", "Operation timeout").Default("5s").Duration()

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch cmd {
	case serve.FullCommand():
		app.FatalIfError(server.Serve(*org, *san, *serverAddress, *useTLS), "serve")
	case ping.FullCommand():
		app.FatalIfError(client.Ping(*serverAddress, *useTLS, *count, *size, *sleep, *timeout), "ping")
	}

}
