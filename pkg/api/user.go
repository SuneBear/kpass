package api

import (
	"strconv"

	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/dao"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
)

// User is API oject for users
type User struct {
	user *dao.User
}

// NewUser returns a User API instance
func NewUser(db *service.DB) *User {
	return &User{dao.NewUser(db)}
}

type tplJoin struct {
	ID   string `json:"id"`
	Pass string `json:"pass"` // should encrypt
}

func (t *tplJoin) Validate() error {
	if len(t.ID) < 3 {
		return &gear.Error{Code: 400, Msg: "invalid id, length of id should >= 3"}
	}
	if !util.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// Join ...
func (a *User) Join(ctx *gear.Context) (err error) {
	body := new(tplJoin)
	if err = ctx.ParseBody(body); err == nil {
		if err = a.user.CheckID(body.ID); err != nil {
			return
		}

		var user *schema.User
		if user, err = a.user.Create(body.ID, body.Pass); err == nil {
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
	if !util.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// Login ...
func (a *User) Login(ctx *gear.Context) (err error) {
	body := new(tplLogin)
	if err = ctx.ParseBody(body); err != nil {
		return
	}

	var user *schema.User
	if user, err = a.user.CheckLogin(body.ID, body.Pass); err != nil {
		return
	}

	token, err := auth.NewToken(user.ID, body.Pass, user.Pass)
	if err != nil {
		return ctx.Error(err)
	}
	ctx.Set(gear.HeaderPragma, "no-cache")
	ctx.Set(gear.HeaderCacheControl, "no-store")
	return ctx.JSON(200, map[string]interface{}{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   auth.JWT().GetExpiresIn().Seconds(),
	})
}

// Password generate a password
func (a *User) Password(ctx *gear.Context) (err error) {
	len := 12
	num := 2
	spec := 2
	if val := ctx.Query("len"); val != "" {
		if len, err = strconv.Atoi(val); err != nil || len < 4 {
			return ctx.ErrorStatus(400)
		}
	}
	if val := ctx.Query("num"); val != "" {
		if num, err = strconv.Atoi(val); err != nil || num < 0 {
			return ctx.ErrorStatus(400)
		}
	}
	if val := ctx.Query("spec"); val != "" {
		if spec, err = strconv.Atoi(val); err != nil || spec < 0 {
			return ctx.ErrorStatus(400)
		}
	}
	return ctx.JSON(200, map[string]string{"password": util.RandPass(len, num, spec)})
}
