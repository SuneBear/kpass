package ctl

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/bll"
	"github.com/seccom/kpass/src/model"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
)

// File controller
//
// @Name File
// @Description File controller
// @Accepts json
// @Produces json
type File struct {
	models *model.All
}

// Init ...
func (a *File) Init(blls *bll.All) *File {
	a.models = blls.Models
	return a
}

// Download ...
//
// @Title Download
// @Summary Download a file
// @Description Download a file by query.
// @Param fileID path string true "file ID"
// @Param refType query string true "refer object type: user | team | entry"
// @Param refID query string true "refer object ID"
// @Param signed query string false "signed string for verify, only need for entry type"
// @Success 200 []byte
// @Failure 400 string
// @Failure 401 string
// @Failure 404 string
// @Router GET /download/{fileID}?refType={refType}&refID={refID}&signed={signed}
func (c *File) Download(ctx *gear.Context) error {
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
		user, err := c.models.User.Find(userID)
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
		team, err := c.models.Team.Find(TeamID, false)
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
		entry, err := c.models.Entry.Find(EntryID, false)
		if err != nil {
			return ctx.ErrorStatus(404)
		}
		if !entry.HasFile(FileID.String()) {
			return ctx.ErrorStatus(404)
		}
	default:
		return ctx.ErrorStatus(400)
	}

	file, blob, err := c.models.File.FindFile(FileID, key)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Attachment(file.Name, file.Updated, blob.Reader(), inline)
}

// UploadAvatar :
//
// @Title UploadAvatar
// @Summary Upload a avatar
// @Description Upload a avatar and set it to the current user.
// @Success 200 schema.UserResult
// @Failure 400 string
// @Failure 401 string
// @Failure 404 string
// @Router POST /upload/avatar
func (c *File) UploadAvatar(ctx *gear.Context) error {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	user, err := c.models.User.Find(userID)
	if err != nil {
		return ctx.Error(err)
	}
	file, err := c.fileFromCtx(ctx, userID, "", true)
	if err != nil {
		return ctx.Error(err)
	}
	user.Avatar = file.ID
	if err = c.models.User.Update(user); err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, user.Result())
}

// UploadLogo :
//
// @Title UploadLogo
// @Summary Upload a logo
// @Description Upload a logo and set it to the team.
// @Param teamID path string true "team ID"
// @Success 200 schema.TeamResult
// @Failure 400 string
// @Failure 401 string
// @Failure 404 string
// @Router POST /upload/team/{teamID}/logo
func (c *File) UploadLogo(ctx *gear.Context) (err error) {
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	team, err := c.models.Team.Find(TeamID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if team.UserID != userID {
		return ctx.ErrorStatus(403)
	}

	file, err := c.fileFromCtx(ctx, userID, "", true)
	if err != nil {
		return ctx.Error(err)
	}
	team.Logo = file.ID
	teamResult, err := c.models.Team.Update(TeamID, team)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, teamResult)
}

// UploadFile :
//
// @Title UploadFile
// @Summary Upload a file
// @Description Upload a file to the entry.
// @Param refType query string true "refer object type: user | team | entry"
// @Param refID query string true "refer object ID"
// @Success 200 schema.FileResult
// @Failure 400 string
// @Failure 401 string
// @Failure 404 string
// @Router POST /upload/entry/{entryID}/file
func (c *File) UploadFile(ctx *gear.Context) (err error) {
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	entry, err := c.models.Entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := auth.KeyFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	userID, _ := auth.UserIDFromCtx(ctx)
	if key, err = c.models.File.GetTeamKey(entry.TeamID, userID, key); err != nil {
		return ctx.Error(err)
	}

	file, err := c.fileFromCtx(ctx, userID, key, false)
	if err != nil {
		return ctx.Error(err)
	}
	if err = c.models.Entry.AddFileByID(EntryID, file.ID, userID); err != nil {
		return ctx.Error(err)
	}
	file.SetDownloadURL("entry", EntryID.String())
	return ctx.JSON(200, file)
}

func (c *File) fileFromCtx(ctx *gear.Context, userID, key string, checkImage bool) (
	*schema.FileResult, error) {
	switch ctx.AcceptType(gear.MIMEMultipartForm) {
	case gear.MIMEMultipartForm:
		err := ctx.Req.ParseMultipartForm(1024 * 200)
		if err != nil {
			return nil, err
		}
		// only read the first file!
		for _, fileHeaders := range ctx.Req.MultipartForm.File {
			for _, fileHeader := range fileHeaders {
				if err := checkFileName(fileHeader.Filename, checkImage); err != nil {
					return nil, err
				}
				file, err := fileHeader.Open()
				if err != nil {
					return nil, err
				}
				return c.models.File.Create(userID, key, fileHeader.Filename, file)
			}
		}
	}
	return nil, &gear.Error{Code: 400, Msg: "invalid upload request"}
}

func checkFileName(filename string, checkImage bool) error {
	if filename == "" {
		return errors.New("invalid file name")
	}
	if !checkImage {
		return nil
	}
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".png", ".jpg", ".jpeg", ".gif":
		return nil
	default:
		return errors.New("invalid image file")
	}
}
