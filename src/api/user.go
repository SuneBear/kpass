package api

import (
	"strconv"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/dao"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
)

// User is API oject for users
//
// @Name User
// @Description User API
// @Accepts json
// @Produces json
type User struct {
	team *dao.Team
	user *dao.User
}

// NewUser returns a User API instance
func NewUser(db *service.DB) *User {
	return &User{dao.NewTeam(db), dao.NewUser(db)}
}

type tplUserJoin struct {
	ID   string `json:"id" swaggo:"true,user id,admin"`
	Pass string `json:"pass" swaggo:"true,user password hashed by sha256,15e2536def2490c115759ceabf012872fddbd7887fbe67e5074d1e66148d5d00"`
}

func (t *tplUserJoin) Validate() error {
	if len(t.ID) < 3 {
		return &gear.Error{Code: 400, Msg: "invalid id, length of id should >= 3"}
	}
	if !util.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// Join ...
//
// @Title Join
// @Summary Create a user
// @Description Create a user
// @Param body body tplUserJoin true "user info"
// @Success 200 schema.UserResult
// @Failure 400 string
// @Failure 401 string
// @Router POST /api/join
func (a *User) Join(ctx *gear.Context) error {
	body := new(tplUserJoin)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}
	if err := a.user.CheckID(body.ID); err != nil {
		return ctx.Error(err)
	}

	user, err := a.user.Create(body.ID, body.Pass)
	if err != nil {
		return ctx.Error(err)
	}
	// create a private team for the user.
	_, err = a.team.Create(body.ID, body.Pass, &schema.Team{
		Name:       body.ID,
		UserID:     body.ID,
		Visibility: "private",
		Members:    []string{body.ID},
	})
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.JSON(200, user.Result())
}

// Resource Owner Password Credentials Grant https://tools.ietf.org/html/rfc6749#page-37
type tplUserLogin struct {
	Type string `json:"grant_type" swaggo:"true,should always be \"password\",password"`
	ID   string `json:"username" swaggo:"true,user id,admin"`
	Pass string `json:"password" swaggo:"true,user password hashed by sha256,15e2536def2490c115759ceabf012872fddbd7887fbe67e5074d1e66148d5d00"`
}

func (t *tplUserLogin) Validate() error {
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
//
// @Title Login
// @Summary Login
// @Description Login with user id and pass, get the new access_token
// @Param body body tplUserLogin true "user auth info"
// @Success 200 AuthResult
// @Failure 400 string
// @Failure 401 string
// @Router POST /api/login
func (a *User) Login(ctx *gear.Context) (err error) {
	body := new(tplUserLogin)
	if err = ctx.ParseBody(body); err != nil {
		return
	}

	var user *schema.User
	if user, err = a.user.CheckLogin(body.ID, body.Pass); err != nil {
		return
	}

	token, err := auth.NewToken(user.ID)
	if err != nil {
		return ctx.Error(err)
	}
	ctx.Set(gear.HeaderPragma, "no-cache")
	ctx.Set(gear.HeaderCacheControl, "no-store")
	return ctx.JSON(200, &AuthResult{
		Token:  token,
		Type:   "Bearer",
		Expire: auth.JWT().GetExpiresIn().Seconds(),
		User:   user.Result(),
	})
}

// Find ...
//
// @Title Find
// @Summary get a user public info
// @Description get a user public info
// @Param userID path string true "user id"
// @Success 200 schema.UserResult
// @Failure 400 string
// @Failure 401 string
// @Router GET /api/user/{userID}
func (a *User) Find(ctx *gear.Context) (err error) {
	userID := ctx.Param("userID")
	if userID != "" {
		return ctx.ErrorStatus(400)
	}
	user, err := a.user.Find(userID)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, user.Result())
}

// PassResult ...
type PassResult struct {
	Pass string `json:"password" swaggo:"true,a random password,OG/O3QPm6Y)A"`
}

// Password generate a password
//
// @Title Password
// @Summary get a random password
// @Description get a random password by query options
// @Param len query uint false "password length" 12
// @Param num query uint false "numbers length that password include" 2
// @Param spec query uint false "special characters length that password include" 2
// @Success 200 PassResult
// @Router GET /api/password
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
	return ctx.JSON(200, &PassResult{util.RandPass(len, num, spec)})
}
