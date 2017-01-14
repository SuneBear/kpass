package app

import (
	"os"

	"time"

	"github.com/seccom/kpass/app/api/user"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/favicon"
	"github.com/teambition/gear/middleware/secure"
)

// Version is app version
const Version = "0.0.1"

// New returns a app instance
func New() *gear.App {
	pkg.InitLogger(os.Stdout)

	if err := dao.Open(""); err != nil {
		panic(err)
	}
	pkg.InitAuth(dao.DBSalt)
	pkg.InitJwt(10 * time.Minute)

	app := gear.New()
	app.Use(favicon.NewWithIco(MustAsset("web/image/favicon.ico")))
	app.Use(secure.Default())
	app.UseHandler(pkg.Logger)
	initRouter()
	app.UseHandler(Router)

	userAPI.InitDemoUser()
	return app
}
