package api

import (
	"time"

	"github.com/seccom/kpass/server"
	"github.com/seccom/kpass/server/crypto"
	"github.com/seccom/kpass/server/dao"
	"github.com/teambition/gear"
)

type tplUserJoin struct {
	ID   string `json:"id"`
	Pass string `json:"pass"` // should encrypt
}

func (t *tplUserJoin) Validate() error {
	if len(t.ID) < 3 {
		return &gear.Error{Code: 400, Msg: "invalid id, length of id should >= 3"}
	}
	if !crypto.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// UserJoin ...
func UserJoin(ctx *gear.Context) (err error) {
	body := new(tplUserJoin)
	if err = ctx.ParseBody(body); err == nil {
		if err = dao.CheckUserID(body.ID); err != nil {
			return
		}

		var user *dao.User
		pass := crypto.Global().EncryptUserPass(body.ID, body.Pass)
		if user, err = dao.NewUser(body.ID, pass); err == nil {
			return ctx.JSON(200, user.Result())
		}
	}
	return
}

// Resource Owner Password Credentials Grant https://tools.ietf.org/html/rfc6749#page-37
type tplUserLogin struct {
	Type string `json:"grant_type"`
	ID   string `json:"username"`
	Pass string `json:"password"` // should encrypt
}

func (t *tplUserLogin) Validate() error {
	if t.Type != "password" {
		return &gear.Error{Code: 400, Msg: "invalid_grant"}
	}
	if len(t.ID) < 3 {
		return &gear.Error{Code: 400, Msg: "invalid id, length of id should >= 3"}
	}
	if !crypto.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// UserLogin ...
func UserLogin(ctx *gear.Context) (err error) {
	body := new(tplUserJoin)
	if err = ctx.ParseBody(body); err != nil {
		return
	}

	var user *dao.User
	if user, err = dao.CheckUserLogin(body.ID, body.Pass); err != nil {
		return
	}

	key := crypto.Global().AESKey(body.ID, body.Pass)
	// encrypt key
	if key, err = crypto.Global().EncryptData(body.ID, key); err != nil {
		return
	}
	if key, err = app.Jwt.Sign(map[string]interface{}{"id": body.ID, "key": key}); err != nil {
		return
	}

	ctx.Set(gear.HeaderPragma, "no-cache")
	ctx.Set(gear.HeaderCacheControl, "no-store")
	return ctx.JSON(200, map[string]interface{}{
		"access_token": key,
		"token_type":   "Bearer",
		"expires_in":   app.Jwt.GetExpiration() / time.Second,
	})
}
