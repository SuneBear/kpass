package app

import (
	"crypto/rand"
	"os"

	"time"

	"github.com/seccom/kpass/app/api/user"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/favicon"
)

// Version is app version
const Version = "0.0.1"

// New returns a app instance
func New() *gear.App {
	pkg.InitLogger(os.Stdout)

	if err := dao.Open(""); err != nil {
		panic(err)
	}
	pkg.InitCrypto(dao.DBSalt)

	// We use rand key
	jwtKey := make([]byte, 64)
	rand.Read(jwtKey)
	pkg.InitJwt(10*time.Minute, jwtKey)

	app := gear.New()
	app.Use(favicon.NewWithIco(MustAsset("web/image/favicon.ico")))
	app.UseHandler(pkg.Logger)
	initRouter()
	app.UseHandler(Router)

	userAPI.InitDemo()
	return app
}
