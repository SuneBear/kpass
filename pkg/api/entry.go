package api

import (
	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/dao"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/teambition/gear"
)

// Entry is API oject for entries
type Entry struct {
	entry  *dao.Entry
	secret *dao.Secret
}

// NewEntry returns a Entry API instance
func NewEntry(db *service.DB) *Entry {
	return &Entry{dao.NewEntry(db), dao.NewSecret(db)}
}

type tplEntryCreate struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (t *tplEntryCreate) Validate() error {
	if len(t.Name) == 0 {
		return &gear.Error{Code: 400, Msg: "entry name required"}
	}
	return nil
}

// Create ...
func (a *Entry) Create(ctx *gear.Context) (err error) {
	body := new(tplEntryCreate)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}
	userID, err := auth.UserIDFromCtx(ctx)
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

	var entry *schema.EntrySum
	if entry, err = a.entry.Create(userID, ownerID, ownerType, body.Name, body.Category); err == nil {
		return ctx.JSON(200, entry)
	}
	return
}

type tplEntryUpdate map[string]interface{}

// Validate ...
func (t *tplEntryUpdate) Validate() error {
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
func (a *Entry) Update(ctx *gear.Context) (err error) {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	body := new(tplEntryUpdate)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	entrySum, err := a.entry.Update(userID, EntryID, *body)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entrySum)
}

// Delete ...
func (a *Entry) Delete(ctx *gear.Context) (err error) {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	entry, err := a.entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if entry.OwnerID != userID {
		return ctx.ErrorStatus(403)
	}

	if _, err = a.entry.UpdateDeleted(userID, EntryID, true); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}

// Restore ...
func (a *Entry) Restore(ctx *gear.Context) (err error) {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	entry, err := a.entry.Find(EntryID, true)
	if err != nil {
		return ctx.Error(err)
	}
	if entry.OwnerID != userID {
		return ctx.ErrorStatus(403)
	}

	entrySum, err := a.entry.UpdateDeleted(userID, EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entrySum)
}

// Find return the entry
func (a *Entry) Find(ctx *gear.Context) error {
	EntryID, err := uuid.Parse(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	entry, err := a.entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := auth.KeyFromCtx(ctx, entry.OwnerID)
	if err != nil {
		return ctx.Error(err)
	}

	var secrets []*schema.SecretResult
	if len(entry.Secrets) > 0 {
		if secrets, err = a.secret.FindSecrets(key, entry.Secrets...); err != nil {
			return ctx.Error(err)
		}
	}

	return ctx.JSON(200, entry.Result(EntryID, secrets, nil))
}

// FindByOwner return entries for current user
func (a *Entry) FindByOwner(ctx *gear.Context) (err error) {
	userID, err := auth.UserIDFromCtx(ctx)
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

	entries, err := a.entry.FindByOwnerID(userID, ownerID, ownerType, false)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entries)
}
