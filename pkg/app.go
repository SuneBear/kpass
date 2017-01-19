package pkg

import (
	"time"

	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/dao"
	"github.com/seccom/kpass/pkg/logger"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/favicon"
	"github.com/teambition/gear/middleware/secure"
	"github.com/teambition/gear/middleware/static"
)

// Version is app version
const Version = "0.3.0"

// New returns a app instance
func New(dbPath string, env string) *gear.App {
	if env == "production" {
		logger.Init()
	}

	db, err := service.NewDB(dbPath)
	if err != nil {
		panic(err)
	}
	auth.Init(db.Salt, 10*time.Minute)

	app := gear.New()
	app.Use(favicon.NewWithIco(MustAsset("web/image/favicon.ico")))
	app.Use(secure.Default())

	if env == "development" {
		app.Use(static.New(static.Options{
			Root:        "./web",
			Prefix:      "/dev",
			StripPrefix: true,
		}))
	}
	app.UseHandler(logger.Default())
	app.UseHandler(newRouter(db))

	if env != "production" {
		initDemo(db)
	}
	return app
}

// initDemo creates demo user
func initDemo(db *service.DB) {
	userDao := dao.NewUser(db)
	entryDao := dao.NewEntry(db)
	secretDao := dao.NewSecret(db)

	if err := userDao.CheckID("demo"); err != nil {
		return
	}
	// create a demo user
	// client should make double sha256 hash.
	pass := util.SHA256Sum(util.SHA256Sum("demo"))
	pass = auth.EncryptUserPass("demo", pass)
	user, _ := userDao.Create("demo", pass)
	logger.Println(`create a demo user {id:"demo", pass:"demo"}:`)
	logger.Println(user)

	// create a demo entry
	entrySum, _ := entryDao.Create(user.ID, user.ID, "user", "demo entry", "")
	logger.Println(`create a demo entry:`)
	logger.Println(entrySum)

	// Add a secret to the entry
	token, _ := auth.NewToken(user.ID, pass, user.Pass)
	claims, _ := auth.JWT().Decode(token)
	key := claims.Get("key").(string)
	key, _ = auth.DecryptData(user.ID, key)

	secret := &schema.Secret{
		Name: "my secret",
		URL:  "https://demo.com",
		Pass: "6vRny_ZJwIlWqBH",
		Note: "Hello world!",
	}

	secretResult, _ := secretDao.Create(user.ID, key, entrySum.ID, secret)
	logger.Println(`Add a secret to demo entry:`)
	logger.Println(secretResult)

	logger.Printf("\nGet enties list:\nhttp://127.0.0.1:8088/entries?access_token=%s\n", token)
	logger.Printf("\nGet demo enty:\nhttp://127.0.0.1:8088/entries/%s?access_token=%s\n\n",
		entrySum.ID, token)
}
