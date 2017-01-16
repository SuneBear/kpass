package app

import (
	"time"

	"github.com/seccom/kpass/app/crypto"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/dao/entry"
	"github.com/seccom/kpass/app/dao/secret"
	"github.com/seccom/kpass/app/dao/user"
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
	if !devMode {
		pkg.InitLogger()
	}

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

	if devMode {
		initDemo()
	}
	return app
}

// initDemo creates demo user
func initDemo() {
	if err := userDao.CheckID("demo"); err != nil {
		return
	}
	// create a demo user
	// client should make double sha256 hash.
	pass := crypto.SHA256Sum(crypto.SHA256Sum("demo"))
	pass = pkg.Auth.EncryptUserPass("demo", pass)
	user, _ := userDao.Create("demo", pass)
	pkg.Logger.Println(`create a demo user {id:"demo", pass:"demo"}:`)
	pkg.Logger.Println(user)

	// create a demo entry
	entrySum, _ := entryDao.Create(user.ID, user.ID, "user", "demo entry", "")
	pkg.Logger.Println(`create a demo entry:`)
	pkg.Logger.Println(entrySum)

	// Add a secret to the entry
	token, _ := pkg.Auth.NewToken(user.ID, pass, user.Pass)
	claims, _ := pkg.Jwt.Decode(token)
	key := claims.Get("key").(string)
	key, _ = pkg.Auth.DecryptData(user.ID, key)

	secret := &dao.Secret{
		Name: "my secret",
		URL:  "https://demo.com",
		Pass: "6vRny_ZJwIlWqBH",
		Note: "Hello world!",
	}

	secretResult, _ := secretDao.Create(key, entrySum.ID, secret)
	pkg.Logger.Println(`Add a secret to demo entry:`)
	pkg.Logger.Println(secretResult)

	pkg.Logger.Printf("\nGet enties list:\nhttp://127.0.0.1:8088/entries?access_token=%s\n", token)
	pkg.Logger.Printf("\nGet demo enty:\nhttp://127.0.0.1:8088/entries/%s?access_token=%s\n\n", entrySum.ID, token)
}
