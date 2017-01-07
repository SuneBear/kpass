package app

import "github.com/teambition/gear"

// Router middleware mounts the service routers
var Router = gear.NewRouter()

func noop(ctx *gear.Context) error {
	return ctx.ErrorStatus(404)
}

func init() {
	serviceInfo := map[string]string{
		"name":    "KPass",
		"version": Version,
	}

	Router.Get("/", func(ctx *gear.Context) error {
		return ctx.JSON(200, serviceInfo)
	})

	// Create a new user
	Router.Post("/join", noop)
	// login
	Router.Post("/login", noop)

	// Return current user info
	Router.Get("/user", noop)
	// Update current user info
	Router.Put("/user", noop)
	// Return the user info, for admin
	Router.Get("/users/:name", noop)
	// Update the user, block or unblock, for admin
	Router.Put("/users/:name", noop)

	// Create a new entry
	Router.Post("/entries", noop)
	// Return current user's entries list
	Router.Get("/entries", noop)
	// Update the entry
	Router.Put("/entries/:entryId", noop)
	// Delete the entry
	Router.Delete("/entries/:entryId", noop)
	// Add a secret to the entry
	Router.Post("/entries/:entryId/secrets", noop)
	// Update the secret
	Router.Put("/entries/:entryId/secrets/:secretId", noop)
	// Delete the secret
	Router.Delete("/entries/:entryId/secrets/:secretId", noop)
	// Add a share to the entry
	Router.Post("/entries/:entryId/shares", noop)
	// Update the share
	Router.Put("/entries/:entryId/shares/:shareId", noop)
	// Delete the share
	Router.Delete("/entries/:entryId/shares/:shareId", noop)

	// Create a team
	Router.Post("/teams", noop)
	// Return the team info
	Router.Get("/teams/:teamId", noop)
	// Return the team's entries list
	Router.Get("/teams/:teamId/entries", noop)
	// Update the team
	Router.Put("/teams/:teamId", noop)
	// Delete the team
	Router.Delete("/teams/:teamId", noop)

	// Return the shared entry
	Router.Get("/shares/:shareId", noop)
}
