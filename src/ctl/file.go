package ctl

import (
	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/dao"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
)

// File controller
type File struct {
	file  *dao.File
	entry *dao.Entry
	team  *dao.Team
	user  *dao.User
}

// NewFile ...
func NewFile(db *service.DB) *File {
	return &File{dao.NewFile(db), dao.NewEntry(db), dao.NewTeam(db), dao.NewUser(db)}
}

// Download ...
func (c *File) Download(ctx *gear.Context) (err error) {
	FileID, err := util.ParseOID(ctx.Param("fileID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	key := ""
	inline := true
	switch ctx.Query("refType") {
	case "user":
		userID := ctx.Query("refID")
		if userID == "" {
			return ctx.ErrorStatus(400)
		}
		user, err := c.user.Find(userID)
		if err != nil {
			return ctx.ErrorStatus(404)
		}
		if !FileID.Equal(user.Avatar) {
			return ctx.ErrorStatus(400)
		}
	case "team":
		TeamID, err := util.ParseOID(ctx.Query("refID"))
		if err != nil {
			return ctx.ErrorStatus(400)
		}
		team, err := c.team.Find(TeamID, false)
		if err != nil {
			return ctx.ErrorStatus(404)
		}
		if !FileID.Equal(team.Logo) {
			return ctx.ErrorStatus(400)
		}
	case "entry":
		inline = false
		key = ctx.Query("signed")
		EntryID, err := util.ParseOID(ctx.Query("refID"))
		if err != nil || key == "" {
			return ctx.ErrorStatus(400)
		}
		key, err = auth.FileKeyFromSigned(FileID, key)
		if err != nil {
			return ctx.ErrorStatus(401)
		}
		entry, err := c.entry.Find(EntryID, false)
		if err != nil {
			return ctx.ErrorStatus(404)
		}
		if !entry.HasFile(FileID.String()) {
			return ctx.ErrorStatus(400)
		}
	default:
		return ctx.ErrorStatus(400)
	}

	file, blob, err := c.file.FindFile(FileID, key)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Attachment(file.Name, file.Updated, blob.Reader(), inline)
}

// Upload ...
func (c *File) Upload(ctx *gear.Context) (err error) {
	return nil
}
