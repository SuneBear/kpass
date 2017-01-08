package api

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/seccom/kpass/server/crypto"
	"github.com/seccom/kpass/server/dao"
	"github.com/teambition/gear"
)

type bodyTemplateJoin struct {
	ID   string `json:"id"`
	Pass string `json:"pass"` // encrypt
}

// Join ...
func Join(ctx *gear.Context) (err error) {
	body := &bodyTemplateJoin{}
	if err = parseRequestBody(ctx.Req.Body, body); err == nil {
		if body.ID == "" || body.Pass == "" {
			return ctx.ErrorStatus(400)
		}
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

func parseRequestBody(r io.ReadCloser, body interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err == nil {
		err = json.Unmarshal(buf, body)
	}
	return err
}
