package app

import (
	"os"

	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
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

	pkg.InitJwt(3600, "new jwt key", "old jwt key")
	pkg.InitLogger(os.Stdout)
	initRouter()

	app := gear.New()
	app.Use(favicon.NewWithIco(MustAsset("web/image/favicon.ico")))
	app.UseHandler(pkg.Logger)
	// app.UseHandler(util.Jwt)
	app.UseHandler(Router)
	return app
}
