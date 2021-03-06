package api

import (
	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/dao"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
)

// Team is API oject for teams
type Team struct {
	team *dao.Team
}

// NewTeam returns a Team API instance
func NewTeam(db *service.DB) *Team {
	return &Team{dao.NewTeam(db)}
}

type tplCreate struct {
	Name string `json:"name"`
	Pass string `json:"pass"` // should encrypt
}

func (t *tplCreate) Validate() error {
	if t.Name == "" {
		return &gear.Error{Code: 400, Msg: "invalid team name"}
	}
	if !util.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid team pass, pass should be hashed by sha256"}
	}
	return nil
}

// Create ...
func (a *Team) Create(ctx *gear.Context) (err error) {
	body := new(tplCreate)
	if err = ctx.ParseBody(body); err == nil {
		claims, _ := auth.FromCtx(ctx)
		userID := claims.Get("id").(string)
		var res *schema.TeamResult
		if res, err = a.team.Create(userID, body.Name, body.Pass); err == nil {
			return ctx.JSON(200, res)
		}
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
				return &gear.Error{Code: 400, Msg: "invalid team name"}
			}
		case "isFrozen":
			_, ok := val.(bool)
			if !ok {
				return &gear.Error{Code: 400, Msg: "invalid team isFrozen"}
			}
		default:
			return &gear.Error{Code: 400, Msg: "invalid team property"}
		}
	}

	if empty {
		return &gear.Error{Code: 400, Msg: "no content"}
	}
	return nil
}

// Update ...
func (a *Team) Update(ctx *gear.Context) (err error) {
	TeamID, err := uuid.Parse(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	body := new(tplUpdate)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	team, err := a.team.Find(TeamID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if team.OwnerID != userID {
		return ctx.ErrorStatus(403)
	}

	changed := false
	for key, val := range *body {
		switch key {
		case "name":
			if name := val.(string); name != team.Name {
				changed = true
				team.Name = name
			}
		case "isFrozen":
			if isFrozen := val.(bool); isFrozen != team.IsFrozen {
				changed = true
				team.IsFrozen = isFrozen
			}
		}
	}

	if !changed {
		return ctx.End(204)
	}

	res, err := a.team.Update(TeamID, team)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

type tplMembers struct {
	Push []string `json:"$push"` // Removes members from team
	Pull []string `json:"$pull"` // Adds members to team
}

// Validate ...
func (t *tplMembers) Validate() error {
	if len(t.Push) == 0 && len(t.Pull) == 0 {
		return &gear.Error{Code: 400, Msg: "no content"}
	}
	if len(t.Push) > 100 || len(t.Pull) > 100 {
		return &gear.Error{Code: 400, Msg: "too many members"}
	}
	return nil
}

// Members ...
func (a *Team) Members(ctx *gear.Context) (err error) {
	TeamID, err := uuid.Parse(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	body := new(tplMembers)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	res, err := a.team.UpdateMembers(userID, TeamID, body.Pull, body.Push)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

// Delete ...
func (a *Team) Delete(ctx *gear.Context) (err error) {
	TeamID, err := uuid.Parse(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	team, err := a.team.Find(TeamID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if team.OwnerID != userID {
		return ctx.ErrorStatus(403)
	}

	team.IsDeleted = true
	if _, err = a.team.Update(TeamID, team); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}

// Restore ...
func (a *Team) Restore(ctx *gear.Context) (err error) {
	TeamID, err := uuid.Parse(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	team, err := a.team.Find(TeamID, true)
	if err != nil {
		return ctx.Error(err)
	}
	if team.OwnerID != userID {
		return ctx.ErrorStatus(403)
	}

	team.IsDeleted = false
	res, err := a.team.Update(TeamID, team)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

// FindByMember return teams for current user
func (a *Team) FindByMember(ctx *gear.Context) (err error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	team, err := a.team.FindByMemberID(userID)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, team)
}

type tplToken struct {
	Type string `json:"grant_type"`
	Pass string `json:"password"` // should encrypt
}

func (t *tplToken) Validate() error {
	if t.Type != "password" {
		return &gear.Error{Code: 400, Msg: "invalid_grant"}
	}
	if !util.IsHashString(t.Pass) {
		return &gear.Error{Code: 400, Msg: "invalid pass, pass should be hashed by sha256"}
	}
	return nil
}

// Token ...
func (a *Team) Token(ctx *gear.Context) (err error) {
	TeamID, err := uuid.Parse(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	body := new(tplToken)
	if err = ctx.ParseBody(body); err != nil {
		return
	}

	team, err := a.team.CheckToken(TeamID.String(), userID, body.Pass)
	if err != nil {
		return ctx.Error(err)
	}

	token, err := auth.AddTeamKey(ctx, TeamID, body.Pass, team.Pass)
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
