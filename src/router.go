package src

import (
	"github.com/seccom/kpass/src/api"
	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/ctl"
	"github.com/seccom/kpass/src/service"
	"github.com/teambition/gear"
)

func noOp(ctx *gear.Context) error {
	return ctx.ErrorStatus(404)
}

func newRouter(db *service.DB) (Router *gear.Router) {
	Router = gear.NewRouter()

	fileCtl := ctl.NewFile(db)

	entryAPI := api.NewEntry(db)
	secretAPI := api.NewSecret(db)
	teamAPI := api.NewTeam(db)
	userAPI := api.NewUser(db)
	shareAPI := api.NewShare(db)

	// GET /download/fileID?refType=user&refID=userID
	// GET /download/fileID?refType=team&refID=teamID
	// GET /download/fileID?refType=entry&refID=entryID&signed=xxxx
	Router.Get("/download/:fileID", fileCtl.Download)

	Router.Post("/upload/avatar", auth.Middleware, fileCtl.UploadAvatar)
	Router.Post("/upload/team/:teamID/logo", auth.Middleware, fileCtl.UploadLogo)
	Router.Post("/upload/entry/:entryID/file", auth.Middleware, fileCtl.UploadFile)

	// generate a random password
	Router.Get("/api/password", userAPI.Password)

	// Create a new user
	Router.Post("/api/join", userAPI.Join)
	Router.Post("/api/login", userAPI.Login)
	// Return the user publicly info
	Router.Get("/api/user/:userID", auth.Middleware, userAPI.Find)

	// Create a team
	Router.Post("/api/teams", auth.Middleware, teamAPI.Create)
	// // Return current user's teams joined.
	Router.Get("/api/teams", auth.Middleware, teamAPI.FindByMember)
	// Undelete the entry
	Router.Post(`/api/teams/:teamID+:undelete`, auth.Middleware, teamAPI.Undelete)
	// Get team's token
	Router.Post("/api/teams/:teamID/token", auth.Middleware, teamAPI.Token)
	// Return the team's entries list
	Router.Get("/api/teams/:teamID/entries", auth.Middleware, entryAPI.FindByTeam)
	// Create a new entry for team
	Router.Post("/api/teams/:teamID/entries", auth.Middleware, entryAPI.Create)
	// Update the team
	Router.Put("/api/teams/:teamID", auth.Middleware, teamAPI.Update)
	// change the team's members
	Router.Put("/api/teams/:teamID/members", auth.Middleware, teamAPI.Members)
	// Delete the team
	Router.Delete("/api/teams/:teamID", auth.Middleware, teamAPI.Delete)
	// Return the team's shares list
	Router.Get("/api/teams/:teamID/shares", auth.Middleware, noOp)

	// Undelete the entry
	Router.Post("/api/entries/:entryID+:undelete", auth.Middleware, entryAPI.Undelete)
	// Get the full entry, with all secrets
	Router.Get("/api/entries/:entryID", auth.Middleware, entryAPI.Find)
	// Update the entry
	Router.Put("/api/entries/:entryID", auth.Middleware, entryAPI.Update)
	// Delete the entry
	Router.Delete("/api/entries/:entryID", auth.Middleware, entryAPI.Delete)
	// Add a secret to the entry
	Router.Post("/api/entries/:entryID/secrets", auth.Middleware, secretAPI.Create)
	// Update the secret
	Router.Put("/api/entries/:entryID/secrets/:secretID", auth.Middleware, secretAPI.Update)
	// Delete the secret
	Router.Delete("/api/entries/:entryID/secrets/:secretID", auth.Middleware, secretAPI.Delete)
	// Add a share to the entry
	Router.Post("/api/entries/:entryID/shares", auth.Middleware, shareAPI.Create)
	// Get shares list of the entry
	// Router.Get("/api/entries/:entryID/shares", auth.Middleware, shareAPI.FindByEntry)
	// Delete the share
	// Router.Delete("/api/entries/:entryID/shares/:shareID", auth.Middleware, entryAPI.DeleteShare)

	// Delete the file
	Router.Delete("/api/entries/:entryID/files/:fileID", auth.Middleware, entryAPI.DeleteFile)

	// Get shares list of the team
	// Router.Get("/api/teams/:teamID/shares", auth.Middleware, shareAPI.FindByTeam)
	// Get shares list to the current user
	// Router.Get("/api/shares/me", auth.Middleware, shareAPI.FindByUser)
	// Get the share
	// Router.Get("/api/shares/:shareID", auth.Middleware, shareAPI.ReadShare)
	return
}
