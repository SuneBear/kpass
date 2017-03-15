package src

import (
	"regexp"
	"strings"
	"time"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/logger"
	"github.com/seccom/kpass/src/service"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/cors"
	"github.com/teambition/gear/middleware/favicon"
	"github.com/teambition/gear/middleware/secure"
	"github.com/teambition/gear/middleware/static"
)

// Version is app version
const Version = "v1.0.0-alpha.1"

// New returns a app instance
func New(dbPath string, env string) *gear.App {
	if env == "production" {
		logger.Init()
	}

	db, err := service.NewDB(dbPath)
	if err != nil {
		panic(err)
	}
	auth.Init(db.Salt, 20*time.Minute)

	indexBody := "<h1>Kpass</h1>"
	faviconBin := []byte{}

	if env != "test" {
		indexBody = string(MustAsset("index.html"))
		faviconBin = MustAsset("favicon.ico")
	}

	staticOpts := static.Options{
		Root:        "",
		Prefix:      "/",
		StripPrefix: true,
		Files:       make(map[string][]byte),
	}
	for _, name := range AssetNames() {
		staticOpts.Files[name] = MustAsset(name)
	}
	if env == "development" {
		staticOpts.Root = "./web"
	}

	app := gear.New()
	app.Use(cors.New())
	app.Use(secure.Default)
	app.Use(favicon.NewWithIco(faviconBin))

	staticMiddleware := static.New(staticOpts)

	var routerPrefix = regexp.MustCompile(`^/(api|download|upload)/`)
	app.Use(func(ctx *gear.Context) (err error) {
		switch {
		case ctx.Path == "/logo.png" || ctx.Path == "/humans.txt" || ctx.Path == "/robots.txt" || strings.HasPrefix(ctx.Path, "/static/"):
			return staticMiddleware(ctx)
		case ctx.Path == "/" || !routerPrefix.MatchString(ctx.Path):
			return ctx.HTML(200, indexBody)
		}
		return nil
	})
	app.UseHandler(logger.Default())
	app.UseHandler(newRouter(db))
	return app
}
