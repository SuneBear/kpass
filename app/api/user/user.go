package userAPI

import (
	"github.com/seccom/kpass/app/crypto"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/dao/user"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
)

type tplJoin struct {
	ID   string `json:"id"`
	Pass string `json:"pass"` // should encrypt
}

func (t *tplJoin) Validate() error {
	if len(t.ID) < 3 {
		return &gear.Error{Code: 400, Msg: "invalid id, length of id should >= 3"}
	}
	if !crypto.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// Join ...
func Join(ctx *gear.Context) (err error) {
	body := new(tplJoin)
	if err = ctx.ParseBody(body); err == nil {
		if err = userDao.CheckID(body.ID); err != nil {
			return
		}

		var user *dao.User
		pass := pkg.Auth.EncryptUserPass(body.ID, body.Pass)
		if user, err = userDao.Create(body.ID, pass); err == nil {
			return ctx.JSON(200, user.Result())
		}
	}
	return
}

// Resource Owner Password Credentials Grant https://tools.ietf.org/html/rfc6749#page-37
type tplLogin struct {
	Type string `json:"grant_type"`
	ID   string `json:"username"`
	Pass string `json:"password"` // should encrypt
}

func (t *tplLogin) Validate() error {
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

// Login ...
func Login(ctx *gear.Context) (err error) {
	body := new(tplLogin)
	if err = ctx.ParseBody(body); err != nil {
		return
	}

	var user *dao.User
	if user, err = userDao.CheckLogin(body.ID, body.Pass); err != nil {
		return
	}

	token, err := pkg.Auth.NewToken(user.ID, body.Pass, user.Pass)
	if err != nil {
		return ctx.Error(err)
	}
	ctx.Set(gear.HeaderPragma, "no-cache")
	ctx.Set(gear.HeaderCacheControl, "no-store")
	return ctx.JSON(200, map[string]interface{}{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   pkg.Jwt.GetExpiresIn().Seconds(),
	})
}
