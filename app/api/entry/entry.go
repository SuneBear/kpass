package entryAPI

import (
	"github.com/google/uuid"
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
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}
	userID, err := pkg.Auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}

	// POST /entries
	ownerID := userID
	ownerType := "user"
	// POST /teams/:teamID/entries
	if ctx.Param("teamID") != "" {
		TeamID, err := uuid.Parse(ctx.Param("teamID"))
		if err != nil {
			return ctx.ErrorStatus(400)
		}
		ownerID = TeamID.String()
		ownerType = "team"
	}

	var entry *dao.EntrySum
	if entry, err = entryDao.Create(userID, ownerID, ownerType, body.Name, body.Category); err == nil {
		return ctx.JSON(200, entry)
	}
	return
}

type tplUpdate map[string]interface{}

// Validate ...
func (t *tplUpdate) Validate() error {
	empty := true
	for key, val := range *t {
		empty = false

		switch key {
		case "name":
			v, ok := val.(string)
			if !ok || v == "" {
				return &gear.Error{Code: 400, Msg: "invalid entry name"}
			}
		case "category":
			_, ok := val.(string)
			if !ok {
				return &gear.Error{Code: 400, Msg: "invalid entry category"}
			}
		case "priority":
			v, ok := val.(float64)
			if !ok || v < 0 || v > 127 {
				return &gear.Error{Code: 400, Msg: "invalid entry priority"}
			}
		default:
			return &gear.Error{Code: 400, Msg: "invalid entry property"}
		}
	}

	if empty {
		return &gear.Error{Code: 400, Msg: "no content"}
	}
	return nil
}

// Update ...
func Update(ctx *gear.Context) (err error) {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := pkg.Auth.UserIDFromCtx(ctx)
	body := new(tplUpdate)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	entrySum, err := entryDao.Update(userID, EntryID, *body)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entrySum)
}

// Delete ...
func Delete(ctx *gear.Context) (err error) {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := pkg.Auth.UserIDFromCtx(ctx)
	entry, err := entryDao.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if entry.OwnerID != userID {
		return ctx.ErrorStatus(403)
	}

	if _, err = entryDao.UpdateDeleted(userID, EntryID, true); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}

// Restore ...
func Restore(ctx *gear.Context) (err error) {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := pkg.Auth.UserIDFromCtx(ctx)
	entry, err := entryDao.Find(EntryID, true)
	if err != nil {
		return ctx.Error(err)
	}
	if entry.OwnerID != userID {
		return ctx.ErrorStatus(403)
	}

	entrySum, err := entryDao.UpdateDeleted(userID, EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entrySum)
}

// Find return the entry
func Find(ctx *gear.Context) error {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	entry, err := entryDao.Find(EntryID, false)
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

	return ctx.JSON(200, entry.Result(EntryID, secrets, nil))
}

// FindByOwner return entries for current user
func FindByOwner(ctx *gear.Context) (err error) {
	userID, err := pkg.Auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	// GET /entries
	ownerID := userID
	ownerType := "user"
	// GET /teams/:teamID/entries
	if ctx.Param("teamID") != "" {
		TeamID, err := uuid.Parse(ctx.Param("teamID"))
		if err != nil {
			return ctx.ErrorStatus(400)
		}
		ownerID = TeamID.String()
		ownerType = "team"
	}

	entries, err := entryDao.FindByOwnerID(userID, ownerID, ownerType, false)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entries)
}
