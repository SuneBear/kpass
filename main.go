package main

import (
	"flag"

	"github.com/seccom/kpass/server"
)

var (
	port = flag.String("port", ":8080", `Auth service port.`)
)

func main() {
	flag.Parse()
	srv := app.New()
	app.Logger.Info("Start KPass " + *port)
	srv.Listen(*port)
}
