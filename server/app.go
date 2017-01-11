package app

import (
	"os"

	"github.com/seccom/kpass/server/dao"
	"github.com/seccom/kpass/server/util"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/favicon"
)

// Version is app version
const Version = "0.0.1"

// New returns a app instance
func New() *gear.App {
	if err := dao.Open(""); err != nil {
		panic(err)
	}

	util.InitJwt(3600, "new jwt key", "old jwt key")
	util.InitLogger(os.Stdout)
	initRouter()

	app := gear.New()
	app.Use(favicon.NewWithIco(MustAsset("web/image/favicon.ico")))
	app.UseHandler(util.Logger)
	// app.UseHandler(util.Jwt)
	app.UseHandler(Router)
	return app
}
