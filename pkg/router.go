package pkg

import (
	"github.com/seccom/kpass/pkg/api"
	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/service"
	"github.com/teambition/gear"
)

func noOp(ctx *gear.Context) error {
	return ctx.ErrorStatus(404)
}

func newRouter(db *service.DB) (Router *gear.Router) {
	Router = gear.NewRouter()

	entryAPI := api.NewEntry(db)
	secretAPI := api.NewSecret(db)
	teamAPI := api.NewTeam(db)
	userAPI := api.NewUser(db)

	Router.Get("/", func(ctx *gear.Context) error {
		return ctx.HTML(200, string(MustAsset("web/index.html")))
	})

	Router.Get("/password", userAPI.Password)

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
	Router.Get("/user", auth.Middleware, noOp)
	// Update current user info
	Router.Put("/user", auth.Middleware, noOp)
	// Return the user info, for admin
	Router.Get("/users/:id", auth.Middleware, noOp)
	// Update the user, block or unblock, for admin
	Router.Put("/users/:id", auth.Middleware, noOp)

	// Create a new entry
	// Request body:
	//  {
	// 		"name":"wechat",
	// 		"category":"登录信息"
	//  }
	// Return: entry info object
	Router.Post("/entries", auth.Middleware, entryAPI.Create)
	// Return current user's entries list with summary info.
	Router.Get("/entries", auth.Middleware, entryAPI.FindByOwner)
	// Get the full entry, with all secrets
	Router.Get("/entries/:entryID", auth.Middleware, entryAPI.Find)
	// Update the entry
	Router.Put("/entries/:entryID", auth.Middleware, entryAPI.Update)
	// Delete the entry
	Router.Delete("/entries/:entryID", auth.Middleware, entryAPI.Delete)
	// Restore the entry
	Router.Put("/entries/:entryID/restore", auth.Middleware, entryAPI.Restore)

	// Add a secret to the entry
	// Request body:
	//  {
	// 		"name":"kpass",
	// 		"url":"https://wechat.com/login",
	// 		"password":"123456",
	// 		"note":"other info",
	//  }
	// Return: secret info object
	Router.Post("/entries/:entryID/secrets", auth.Middleware, secretAPI.Create)
	// Update the secret
	Router.Put("/entries/:entryID/secrets/:secretID", auth.Middleware, secretAPI.Update)
	// Delete the secret
	Router.Delete("/entries/:entryID/secrets/:secretID", auth.Middleware, secretAPI.Delete)
	// Add a share to the entry
	Router.Post("/entries/:entryID/shares", auth.Middleware, noOp)
	// Update the share
	Router.Put("/entries/:entryID/shares/:shareID", auth.Middleware, noOp)
	// Delete the share
	Router.Delete("/entries/:entryID/shares/:shareID", auth.Middleware, noOp)

	// Create a team
	Router.Post("/teams", auth.Middleware, teamAPI.Create)
	// // Return current user's teams joined.
	Router.Get("/teams", auth.Middleware, teamAPI.FindByMember)
	// Get team's token
	Router.Post("/teams/:teamID/token", auth.Middleware, teamAPI.Token)
	// Return the team's entries list
	Router.Get("/teams/:teamID/entries", auth.Middleware, entryAPI.FindByOwner)
	// Create a new entry for team
	Router.Post("/teams/:teamID/entries", auth.Middleware, entryAPI.Create)
	// Update the team
	Router.Put("/teams/:teamID", auth.Middleware, teamAPI.Update)
	// change the team's members
	Router.Put("/teams/:teamID/members", auth.Middleware, teamAPI.Members)
	// Delete the team
	Router.Delete("/teams/:teamID", auth.Middleware, teamAPI.Delete)

	// Return the shared entry
	// Router.Get("/shares/:shareID", auth.Middleware, noOp)
	return
}
