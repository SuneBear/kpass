package main

import (
	"flag"

	"github.com/seccom/kpass/app"
	"github.com/seccom/kpass/app/pkg"
)

var (
	port = flag.String("port", ":8080", `Auth service port.`)
)

func main() {
	flag.Parse()
	srv := app.New()
	pkg.Logger.Info("Start KPass " + *port)
	srv.Listen(*port)
}
