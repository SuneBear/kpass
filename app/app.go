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
	"github.com/teambition/gear/middleware/static"
)

// Version is app version
const Version = "0.1.0"

// New returns a app instance
func New(dbPath string, devMode bool) *gear.App {
	if devMode && dbPath == "./kpass.kdb" {
		dbPath = ""
	}
	pkg.InitLogger(os.Stdout)

	if err := dao.Open(dbPath); err != nil {
		panic(err)
	}
	dao.InitIndex()
	pkg.InitAuth(dao.DBSalt)
	pkg.InitJwt(10 * time.Minute)

	app := gear.New()
	app.Use(favicon.NewWithIco(MustAsset("web/image/favicon.ico")))
	app.Use(secure.Default())
	if devMode {
		app.Set("AppEnv", "development")
		app.Use(static.New(static.Options{
			Root:        "./web",
			Prefix:      "/dev",
			StripPrefix: true,
		}))
	} else {
		app.Set("AppEnv", "production")
	}
	app.UseHandler(pkg.Logger)
	initRouter()
	app.UseHandler(Router)

	userAPI.InitDemoUser()
	return app
}
