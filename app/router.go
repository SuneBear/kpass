package app

import (
	"github.com/seccom/kpass/app/api/entry"
	"github.com/seccom/kpass/app/api/user"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
)

// Router middleware mounts the service routers
var Router = gear.NewRouter()

func noop(ctx *gear.Context) error {
	return ctx.ErrorStatus(404)
}

func initRouter() {

	Router.Get("/", func(ctx *gear.Context) error {
		return ctx.HTML(200, string(MustAsset("web/index.html")))
	})

	// Create a new user
	Router.Post("/join", userAPI.Join)
	// login
	Router.Post("/login", userAPI.Login)

	// Return current user info
	Router.Get("/user", pkg.Jwt.Serve, noop)
	// Update current user info
	Router.Put("/user", pkg.Jwt.Serve, noop)
	// Return the user info, for admin
	Router.Get("/users/:id", pkg.Jwt.Serve, noop)
	// Update the user, block or unblock, for admin
	Router.Put("/users/:id", pkg.Jwt.Serve, noop)

	// Create a new entry
	Router.Post("/entries", pkg.Jwt.Serve, entryAPI.Create)
	// Return current user's entries list
	Router.Get("/entries", pkg.Jwt.Serve, entryAPI.FindByUser)
	// Update the entry
	Router.Put("/entries/:entryId", pkg.Jwt.Serve, entryAPI.Find)
	// Delete the entry
	Router.Delete("/entries/:entryId", pkg.Jwt.Serve, noop)
	// Add a secret to the entry
	Router.Post("/entries/:entryId/secrets", pkg.Jwt.Serve, entryAPI.CreateSecret)
	// Update the secret
	Router.Put("/entries/:entryId/secrets/:secretId", pkg.Jwt.Serve, noop)
	// Delete the secret
	Router.Delete("/entries/:entryId/secrets/:secretId", pkg.Jwt.Serve, noop)
	// Add a share to the entry
	Router.Post("/entries/:entryId/shares", pkg.Jwt.Serve, noop)
	// Update the share
	Router.Put("/entries/:entryId/shares/:shareId", pkg.Jwt.Serve, noop)
	// Delete the share
	Router.Delete("/entries/:entryId/shares/:shareId", pkg.Jwt.Serve, noop)

	// Create a team
	Router.Post("/teams", pkg.Jwt.Serve, noop)
	// Validate team token
	Router.Post("/teams/:teamId", pkg.Jwt.Serve, noop)
	// Return the team info
	Router.Get("/teams/:teamId", pkg.Jwt.Serve, noop)
	// Return the team's entries list
	Router.Get("/teams/:teamId/entries", pkg.Jwt.Serve, noop)
	// Create a new entry for team
	Router.Post("/teams/:teamId/entries", pkg.Jwt.Serve, noop)
	// Update the team
	Router.Put("/teams/:teamId", pkg.Jwt.Serve, noop)
	// Delete the team
	Router.Delete("/teams/:teamId", pkg.Jwt.Serve, noop)

	// Return the shared entry
	// Router.Get("/shares/:shareId", pkg.Jwt.Serve, noop)
}
