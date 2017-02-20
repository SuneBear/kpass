package api

import (
	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/dao"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
)

// Secret is API oject for secrets
type Secret struct {
	entry  *dao.Entry
	secret *dao.Secret
}

// NewSecret returns a Secret API instance
func NewSecret(db *service.DB) *Secret {
	return &Secret{dao.NewEntry(db), dao.NewSecret(db)}
}

type tplSecretCreate struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Pass string `json:"password"`
	Note string `json:"note"`
}

func (t *tplSecretCreate) Validate() error {
	if (len(t.Name) + len(t.URL) + len(t.Pass) + len(t.Note)) == 0 {
		return &gear.Error{Code: 400, Msg: "content required"}
	}
	return nil
}

// Create ...
func (a *Secret) Create(ctx *gear.Context) error {
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	body := new(tplSecretCreate)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	entry, err := a.entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := auth.KeyFromCtx(ctx, entry.TeamID, "team")
	if err != nil {
		return ctx.Error(err)
	}
	secretResult, err := a.secret.Create(EntryID, userID, key, &schema.Secret{
		Name: body.Name,
		URL:  body.URL,
		Pass: body.Pass,
		Note: body.Note,
	})
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, secretResult)
}

type tplSecretUpdate map[string]interface{}

// Validate ...
func (t *tplSecretUpdate) Validate() error {
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
func (a *Secret) Update(ctx *gear.Context) error {
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	SecretID, err := util.ParseOID(ctx.Param("secretID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	body := new(tplSecretUpdate)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	entry, err := a.entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := auth.KeyFromCtx(ctx, entry.TeamID, "team")
	if err != nil {
		return ctx.Error(err)
	}

	res, err := a.secret.Update(EntryID, SecretID, userID, key, *body)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

// Delete ...
func (a *Secret) Delete(ctx *gear.Context) error {
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	SecretID, err := util.ParseOID(ctx.Param("secretID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}

	if err := a.secret.Delete(EntryID, SecretID, userID); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}
