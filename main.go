package main

import (
	"flag"
	"os/exec"
	"runtime"
	"time"

	"github.com/seccom/kpass/app"
	"github.com/seccom/kpass/app/pkg"
)

var (
	address = flag.String("addr", "127.0.0.1:8088", `Auth service address to listen on.`)
	dbPath  = flag.String("dbpath", "./kpass.kdb", `KPass database pass.`)
)

func main() {
	flag.Parse()
	srv := app.New(*dbPath)

	s := srv.Start(*address)
	go func() {
		host := "http://" + s.Addr().String()
		pkg.Logger.Info("Start KPass: " + host)
		time.Sleep(time.Second)
		startBrowser(host)
	}()
	pkg.Logger.Fatal(s.Wait())
}

// startBrowser tries to open the URL in a browser
// and reports whether it succeeds.
func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
