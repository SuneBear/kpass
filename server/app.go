package app

import (
	"github.com/seccom/kpass/server/dao"
	"github.com/teambition/gear"
)

// Version is app version
const Version = "0.0.1"

// New returns a app instance
func New() *gear.App {
	if err := dao.Open(""); err != nil {
		panic(err)
	}

	app := gear.New()
	app.Use(Favicon)
	app.UseHandler(Logger)

	InitJwt(3600, "new jwt key", "old jwt key")
	app.UseHandler(Jwt)
	app.UseHandler(Router)
	return app
}
