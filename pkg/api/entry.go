package api

import (
	"fmt"

	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/dao"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
)

// Entry is API oject for entries
type Entry struct {
	entry  *dao.Entry
	secret *dao.Secret
	team   *dao.Team
}

// NewEntry returns a Entry API instance
func NewEntry(db *service.DB) *Entry {
	return &Entry{dao.NewEntry(db), dao.NewSecret(db), dao.NewTeam(db)}
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
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	body := new(tplEntryCreate)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}

	entry, err := a.entry.Create(userID, &schema.Entry{
		TeamID:   TeamID,
		Name:     body.Name,
		Category: body.Category,
		Secrets:  []string{},
	})
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entry)
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
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
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
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	entry, err := a.entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if err = a.team.CheckUser(entry.TeamID, userID); err != nil {
		return ctx.Error(err)
	}

	if _, err = a.entry.UpdateDeleted(userID, EntryID, true); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}

// Restore ...
func (a *Entry) Restore(ctx *gear.Context) (err error) {
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	entry, err := a.entry.Find(EntryID, true)
	if err != nil {
		return ctx.Error(err)
	}
	if err = a.team.CheckUser(entry.TeamID, userID); err != nil {
		return ctx.Error(err)
	}

	entrySum, err := a.entry.UpdateDeleted(userID, EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entrySum)
}

// Find return the entry
func (a *Entry) Find(ctx *gear.Context) error {
	EntryID, err := util.ParseOID(ctx.Param("entryID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	entry, err := a.entry.Find(EntryID, false)
	if err != nil {
		return ctx.Error(err)
	}
	key, err := auth.KeyFromCtx(ctx, entry.TeamID, "team")
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

// FindByTeam return entries for current user
func (a *Entry) FindByTeam(ctx *gear.Context) (err error) {
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	if err = a.team.CheckUser(TeamID, userID); err != nil {
		return ctx.Error(err)
	}

	entries, err := a.entry.FindByTeam(TeamID, userID, false)
	fmt.Println(111111, TeamID, userID, entries, err)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, entries)
}
