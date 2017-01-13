package pkg

import (
	"io"
	"os"
	"time"

	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

// Logger middleware ...
var Logger = logging.New(os.Stdout)

// InitLogger ...
func InitLogger(w io.Writer) {
	Logger.Out = w
	Logger.SetLevel(logging.InfoLevel)

	Logger.SetLogInit(func(log logging.Log, ctx *gear.Context) {
		log["IP"] = ctx.IP()
		log["Method"] = ctx.Method
		log["URL"] = ctx.Req.URL.String()
		log["Start"] = time.Now()
		log["UserAgent"] = ctx.Get(gear.HeaderUserAgent)
	})

	Logger.SetLogConsume(func(log logging.Log, _ *gear.Context) {
		end := time.Now()
		log["Time"] = end.Sub(log["Start"].(time.Time)) / 1e6
		delete(log, "Start")

		if res, err := log.JSON(); err == nil {
			Logger.Output(end, logging.InfoLevel, res)
		} else {
			Logger.Output(end, logging.WarningLevel, log.String())
		}
	})
}
