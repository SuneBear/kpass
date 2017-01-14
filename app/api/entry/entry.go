package entryAPI

import (
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/dao/entry"
	"github.com/seccom/kpass/app/dao/secret"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
)

type tplCreate struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (t *tplCreate) Validate() error {
	if len(t.Name) == 0 {
		return &gear.Error{Code: 400, Msg: "entry name required"}
	}
	return nil
}

// Create ...
func Create(ctx *gear.Context) (err error) {
	body := new(tplCreate)
	if err = ctx.ParseBody(body); err == nil {
		claims, _ := pkg.Jwt.FromCtx(ctx)
		id := claims.Get("id").(string)
		var entry *dao.Entry
		if entry, err = entryDao.Create(id, "user", body.Name, body.Category); err == nil {
			return ctx.JSON(200, entry.Summary())
		}
	}
	return
}

// Find return the entry
func Find(ctx *gear.Context) error {
	id := ctx.Param("entryId")
	if !pkg.IsUUID(id) {
		return ctx.ErrorStatus(400)
	}

	entry, err := entryDao.Find(id)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := pkg.Auth.KeyFromCtx(ctx, entry.OwnerID)
	if err != nil {
		return ctx.Error(err)
	}

	var secrets []*dao.SecretResult
	if len(entry.Secrets) > 0 {
		if secrets, err = secretDao.FindSecrets(key, entry.Secrets...); err != nil {
			return ctx.Error(err)
		}
	}

	return ctx.JSON(200, entry.Result(secrets, nil))
}

// FindByUser return entries for current user
func FindByUser(ctx *gear.Context) (err error) {
	claims, _ := pkg.Jwt.FromCtx(ctx)
	id := claims.Get("id").(string)
	var entries []*dao.Entry
	if entries, err = entryDao.FindByOwnerID(id, false); err == nil {
		res := make([]*dao.EntrySum, 0, len(entries))
		for _, entry := range entries {
			res = append(res, entry.Summary())
		}
		return ctx.JSON(200, res)
	}
	return
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

// AddSecret ...
func AddSecret(ctx *gear.Context) error {
	id := ctx.Param("entryId")
	if !pkg.IsUUID(id) {
		return ctx.ErrorStatus(400)
	}

	body := new(tplSecretCreate)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	entry, err := entryDao.Find(id)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := pkg.Auth.KeyFromCtx(ctx, entry.OwnerID)
	if err != nil {
		return ctx.Error(err)
	}
	secret, err := secretDao.Create(key, &dao.Secret{
		Name: body.Name,
		URL:  body.URL,
		Pass: body.Pass,
		Note: body.Note,
	})
	if err != nil {
		return ctx.Error(err)
	}
	entry.Secrets = append(entry.Secrets, secret.ID)
	if err = entryDao.Update(entry); err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, secret)
}
