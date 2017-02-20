package logger

import (
	"os"
	"time"

	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

// std is global logging instance ...
var std = logging.New(os.Stdout)

// Default returns the global logging instance
func Default() *logging.Logger {
	return std
}

// Init ...
func Init() {
	std.SetLevel(logging.InfoLevel)

	std.SetLogInit(func(log logging.Log, ctx *gear.Context) {
		log["IP"] = ctx.IP()
		log["Method"] = ctx.Method
		log["URL"] = ctx.Req.URL.String()
		log["Start"] = time.Now()
		log["UserAgent"] = ctx.Get(gear.HeaderUserAgent)
	})

	std.SetLogConsume(func(log logging.Log, _ *gear.Context) {
		end := time.Now()
		log["Time"] = end.Sub(log["Start"].(time.Time)) / 1e6
		delete(log, "Start")

		if res, err := log.JSON(); err == nil {
			std.Output(end, logging.InfoLevel, res)
		} else {
			std.Output(end, logging.WarningLevel, log.String())
		}
	})
}

// Info ...
func Info(v interface{}) {
	std.Info(v)
}

// Err ...
func Err(v interface{}) {
	std.Err(v)
}

// Fatal ...
func Fatal(v interface{}) {
	std.Fatal(v)
}

// Println ...
func Println(args ...interface{}) {
	std.Println(args...)
}

// Printf ...
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}
