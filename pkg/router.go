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
	Router.Get("/user", auth.Jwt().Serve, noOp)
	// Update current user info
	Router.Put("/user", auth.Jwt().Serve, noOp)
	// Return the user info, for admin
	Router.Get("/users/:id", auth.Jwt().Serve, noOp)
	// Update the user, block or unblock, for admin
	Router.Put("/users/:id", auth.Jwt().Serve, noOp)

	// Create a new entry
	// Request body:
	//  {
	// 		"name":"wechat",
	// 		"category":"登录信息"
	//  }
	// Return: entry info object
	Router.Post("/entries", auth.Jwt().Serve, entryAPI.Create)
	// Return current user's entries list with summary info.
	Router.Get("/entries", auth.Jwt().Serve, entryAPI.FindByOwner)
	// Get the full entry, with all secrets
	Router.Get("/entries/:entryID", auth.Jwt().Serve, entryAPI.Find)
	// Update the entry
	Router.Put("/entries/:entryID", auth.Jwt().Serve, entryAPI.Update)
	// Delete the entry
	Router.Delete("/entries/:entryID", auth.Jwt().Serve, entryAPI.Delete)
	// Restore the entry
	Router.Put("/entries/:entryID/restore", auth.Jwt().Serve, entryAPI.Restore)

	// Add a secret to the entry
	// Request body:
	//  {
	// 		"name":"kpass",
	// 		"url":"https://wechat.com/login",
	// 		"password":"123456",
	// 		"note":"other info",
	//  }
	// Return: secret info object
	Router.Post("/entries/:entryID/secrets", auth.Jwt().Serve, secretAPI.Create)
	// Update the secret
	Router.Put("/entries/:entryID/secrets/:secretID", auth.Jwt().Serve, secretAPI.Update)
	// Delete the secret
	Router.Delete("/entries/:entryID/secrets/:secretID", auth.Jwt().Serve, secretAPI.Delete)
	// Add a share to the entry
	Router.Post("/entries/:entryID/shares", auth.Jwt().Serve, noOp)
	// Update the share
	Router.Put("/entries/:entryID/shares/:shareID", auth.Jwt().Serve, noOp)
	// Delete the share
	Router.Delete("/entries/:entryID/shares/:shareID", auth.Jwt().Serve, noOp)

	// Create a team
	Router.Post("/teams", auth.Jwt().Serve, teamAPI.Create)
	// // Return current user's teams joined.
	Router.Get("/teams", auth.Jwt().Serve, teamAPI.FindByMember)
	// Get team's token
	Router.Post("/teams/:teamID/token", auth.Jwt().Serve, teamAPI.Token)
	// Return the team's entries list
	Router.Get("/teams/:teamID/entries", auth.Jwt().Serve, entryAPI.FindByOwner)
	// Create a new entry for team
	Router.Post("/teams/:teamID/entries", auth.Jwt().Serve, entryAPI.Create)
	// Update the team
	Router.Put("/teams/:teamID", auth.Jwt().Serve, teamAPI.Update)
	// change the team's members
	Router.Put("/teams/:teamID/members", auth.Jwt().Serve, teamAPI.Members)
	// Delete the team
	Router.Delete("/teams/:teamID", auth.Jwt().Serve, teamAPI.Delete)

	// Return the shared entry
	// Router.Get("/shares/:shareID", auth.Jwt().Serve, noOp)
	return
}
