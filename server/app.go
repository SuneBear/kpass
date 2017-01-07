package app

import (
	"github.com/teambition/gear"
)

// Version is app version
const Version = "0.0.1"

// New returns a app instance
func New() *gear.App {
	app := gear.New()
	app.Use(Favicon)
	app.UseHandler(Logger)
	app.UseHandler(Router)
	return app
}
