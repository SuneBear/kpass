package api

import (
	"time"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/dao"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
)

// Share is API oject for teams
type Share struct {
	entry *dao.Entry
	share *dao.Share
	team  *dao.Team
	user  *dao.User
}

// NewShare returns a Share API instance
func NewShare(db *service.DB) *Share {
	return &Share{dao.NewEntry(db), dao.NewShare(db), dao.NewTeam(db), dao.NewUser(db)}
}

type tplShareCreate struct {
	Name   string `json:"name"`
	Pass   string `json:"pass"` // should encrypt
	UserID string `json:"userID"`
	Expire int    `json:"expire"` // seconds
}

func (t *tplShareCreate) Validate() error {
	if t.Name == "" {
		return &gear.Error{Code: 400, Msg: "invalid share name"}
	}
	if !util.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid share pass, pass should be hashed by sha256"}
	}
	if t.UserID == "" {
		return &gear.Error{Code: 400, Msg: "invalid user ID to share"}
	}
	if t.Expire < 10 {
		return &gear.Error{Code: 400, Msg: "invalid share expire time"}
	}
	return nil
}

// Create ...
func (a *Share) Create(ctx *gear.Context) (err error) {
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	body := new(tplShareCreate)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}
	if err = a.user.CheckID(body.UserID); err != nil {
		return ctx.Error(err)
	}

	entry, err := a.entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := auth.KeyFromCtx(ctx, entry.TeamID, "team")
	if err != nil {
		return ctx.Error(err)
	}
	userID, _ := auth.UserIDFromCtx(ctx)
	if err = a.team.CheckUser(entry.TeamID, userID); err != nil {
		return ctx.Error(err)
	}

	expire := time.Duration(body.Expire) * time.Second
	shareResult, err := a.share.Create(EntryID, key, body.Pass, expire, &schema.Share{
		EntryID: EntryID,
		TeamID:  entry.TeamID,
		Name:    body.Name,
		UserID:  body.UserID,
	})
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, shareResult)
}

// Delete ...
func (a *Share) Delete(ctx *gear.Context) (err error) {
	ShareID, err := util.ParseOID(ctx.Param("shareID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	if err := a.share.Delete(ShareID, userID); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}

type tplShareToken struct {
	Type string `json:"grant_type"`
	Pass string `json:"password"` // should encrypt
}

func (t *tplShareToken) Validate() error {
	if t.Type != "password" {
		return &gear.Error{Code: 400, Msg: "invalid_grant"}
	}
	if !util.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// Token ...
func (a *Share) Token(ctx *gear.Context) (err error) {
	ShareID, err := util.ParseOID(ctx.Param("shareID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	body := new(tplShareToken)
	if err = ctx.ParseBody(body); err != nil {
		return
	}

	share, err := a.share.Find(ShareID)
	if err != nil {
		return ctx.Error(err)
	}
	if share.UserID != userID {
		return ctx.ErrorStatus(403)
	}

	token, err := auth.AddShareKey(ctx, ShareID, body.Pass, share.Token)
	if err != nil {
		return ctx.Error(&gear.Error{Code: 401, Msg: err.Error()})
	}
	ctx.Set(gear.HeaderPragma, "no-cache")
	ctx.Set(gear.HeaderCacheControl, "no-store")
	return ctx.JSON(200, map[string]interface{}{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   auth.JWT().GetExpiresIn().Seconds(),
	})
}
