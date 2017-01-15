package secretAPI

import (
	"github.com/google/uuid"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/dao/entry"
	"github.com/seccom/kpass/app/dao/secret"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
)

type tplCreate struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Pass string `json:"password"`
	Note string `json:"note"`
}

func (t *tplCreate) Validate() error {
	if (len(t.Name) + len(t.URL) + len(t.Pass) + len(t.Note)) == 0 {
		return &gear.Error{Code: 400, Msg: "content required"}
	}
	return nil
}

// Create ...
func Create(ctx *gear.Context) error {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	body := new(tplCreate)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	entry, err := entryDao.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := pkg.Auth.KeyFromCtx(ctx, entry.OwnerID)
	if err != nil {
		return ctx.Error(err)
	}
	secret, err := secretDao.Create(key, entry.ID, &dao.Secret{
		Name: body.Name,
		URL:  body.URL,
		Pass: body.Pass,
		Note: body.Note,
	})
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, secret)
}

type tplUpdate map[string]interface{}

// Validate ...
func (t *tplUpdate) Validate() error {
	empty := true
	for key, val := range *t {
		empty = false
		switch key {
		case "name", "url", "password", "note":
			if _, ok := val.(string); !ok {
				return &gear.Error{Code: 400, Msg: "invalid secret property"}
			}
		default:
			return &gear.Error{Code: 400, Msg: "invalid secret property"}
		}
	}

	if empty {
		return &gear.Error{Code: 400, Msg: "no content"}
	}
	return nil
}

// Update ...
func Update(ctx *gear.Context) error {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	SecretID, err := uuid.Parse(ctx.Param("secretID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	body := new(tplUpdate)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	entry, err := entryDao.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	// ensure the secret belong to the entry.
	if !entry.HasSecret(SecretID.String()) {
		return ctx.ErrorStatus(404)
	}
	// ensure current user has right.
	key, err := pkg.Auth.KeyFromCtx(ctx, entry.OwnerID)
	if err != nil {
		return ctx.Error(err)
	}

	secret, err := secretDao.Find(key, SecretID)
	if err != nil {
		return ctx.Error(err)
	}

	changed := false
	for key, val := range *body {
		switch key {
		case "name":
			if name := val.(string); name != secret.Name {
				changed = true
				secret.Name = name
			}
		case "url":
			if url := val.(string); url != secret.URL {
				changed = true
				secret.URL = url
			}
		case "password":
			if pass := val.(string); pass != secret.Pass {
				changed = true
				secret.Pass = pass
			}
		case "note":
			if note := val.(string); note != secret.Note {
				changed = true
				secret.Note = note
			}
		}
	}

	if !changed {
		return ctx.End(204)
	}
	res, err := secretDao.Update(key, SecretID, secret)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

// Delete ...
func Delete(ctx *gear.Context) error {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	SecretID, err := uuid.Parse(ctx.Param("secretID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	entry, err := entryDao.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	userID, err := pkg.Auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	if userID != entry.OwnerID {
		return ctx.ErrorStatus(403)
	}

	if err := secretDao.Delete(EntryID, SecretID); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}
