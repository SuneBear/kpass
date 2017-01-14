package app

import (
	"github.com/seccom/kpass/app/api/entry"
	"github.com/seccom/kpass/app/api/user"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
)

// Router middleware mounts the service routers
var Router = gear.NewRouter()

func noOp(ctx *gear.Context) error {
	return ctx.ErrorStatus(404)
}

func initRouter() {
	Router.Get("/", func(ctx *gear.Context) error {
		return ctx.HTML(200, string(MustAsset("web/index.html")))
	})

	// Create a new user
	// Request body:
	//  {
	// 		"id":"name",
	// 		"pass":"SHA256 hashed password"
	//  }
	// Return: user info object
	Router.Post("/join", userAPI.Join)
	// login
	// Request body:
	//  {
	// 		"grant_type":"password",
	//  	"username": "name",
	// 		"password":"SHA256 hashed password"
	//  }
	// Return: token object
	//  {
	// 		"access_token": "a very long JWT string",
	// 		"token_type": "Bearer",
	// 		"expires_in": 600,
	//  }
	//
	Router.Post("/login", userAPI.Login)

	// Return current user info
	Router.Get("/user", pkg.Jwt.Serve, noOp)
	// Update current user info
	Router.Put("/user", pkg.Jwt.Serve, noOp)
	// Return the user info, for admin
	Router.Get("/users/:id", pkg.Jwt.Serve, noOp)
	// Update the user, block or unblock, for admin
	Router.Put("/users/:id", pkg.Jwt.Serve, noOp)

	// Create a new entry
	// Request body:
	//  {
	// 		"name":"wechat",
	// 		"category":"登录信息"
	//  }
	// Return: entry info object
	Router.Post("/entries", pkg.Jwt.Serve, entryAPI.Create)
	// Return current user's entries list with summary info.
	Router.Get("/entries", pkg.Jwt.Serve, entryAPI.FindByUser)
	// Get the full entry, with all secrets
	Router.Get("/entries/:entryId", pkg.Jwt.Serve, entryAPI.Find)
	// Update the entry
	Router.Put("/entries/:entryId", pkg.Jwt.Serve, noOp)
	// Delete the entry
	Router.Delete("/entries/:entryId", pkg.Jwt.Serve, noOp)

	// Add a secret to the entry
	// Request body:
	//  {
	// 		"name":"kpass",
	// 		"url":"https://wechat.com/login",
	// 		"password":"123456",
	// 		"note":"other info",
	//  }
	// Return: secret info object
	Router.Post("/entries/:entryId/secrets", pkg.Jwt.Serve, entryAPI.AddSecret)
	// Update the secret
	Router.Put("/entries/:entryId/secrets/:secretId", pkg.Jwt.Serve, noOp)
	// Delete the secret
	Router.Delete("/entries/:entryId/secrets/:secretId", pkg.Jwt.Serve, noOp)
	// Add a share to the entry
	Router.Post("/entries/:entryId/shares", pkg.Jwt.Serve, noOp)
	// Update the share
	Router.Put("/entries/:entryId/shares/:shareId", pkg.Jwt.Serve, noOp)
	// Delete the share
	Router.Delete("/entries/:entryId/shares/:shareId", pkg.Jwt.Serve, noOp)

	// Create a team
	Router.Post("/teams", pkg.Jwt.Serve, noOp)
	// Validate team token
	Router.Post("/teams/:teamId", pkg.Jwt.Serve, noOp)
	// Return the team info
	Router.Get("/teams/:teamId", pkg.Jwt.Serve, noOp)
	// Return the team's entries list
	Router.Get("/teams/:teamId/entries", pkg.Jwt.Serve, noOp)
	// Create a new entry for team
	Router.Post("/teams/:teamId/entries", pkg.Jwt.Serve, noOp)
	// Update the team
	Router.Put("/teams/:teamId", pkg.Jwt.Serve, noOp)
	// Delete the team
	Router.Delete("/teams/:teamId", pkg.Jwt.Serve, noOp)

	// Return the shared entry
	// Router.Get("/shares/:shareId", pkg.Jwt.Serve, noOp)
}
