package main

import (
	"flag"

	"github.com/seccom/kpass/server"
	"github.com/seccom/kpass/server/util"
)

var (
	port = flag.String("port", ":8080", `Auth service port.`)
)

func main() {
	flag.Parse()
	srv := app.New()
	util.Logger.Info("Start KPass " + *port)
	srv.Listen(*port)
}
