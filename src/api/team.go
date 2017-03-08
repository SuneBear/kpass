package api

import (
	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/bll"
	"github.com/seccom/kpass/src/model"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
)

// Team is API oject for teams
//
// @Name Team
// @Description Team API
// @Accepts json
// @Produces json
type Team struct {
	models  *model.All
	teamBll *bll.Team
}

// Init ...
func (a *Team) Init(blls *bll.All) *Team {
	a.models = blls.Models
	a.teamBll = blls.Team
	return a
}

type tplTeamCreate struct {
	Name string `json:"name" swaggo:"true,team name,Teambition"`
}

func (t *tplTeamCreate) Validate() error {
	if t.Name == "" {
		return &gear.Error{Code: 400, Msg: "invalid team name"}
	}
	return nil
}

// Create ...
//
// @Title Create
// @Summary Create a team
// @Description Create a team
// @Param Authorization header string true "access_token"
// @Param body body tplTeamCreate true "team body"
// @Success 200 schema.TeamResult
// @Failure 400 string
// @Failure 401 string
// @Router POST /api/teams
func (a *Team) Create(ctx *gear.Context) error {
	body := new(tplTeamCreate)
	if err := ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	key, err := auth.KeyFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	userID, _ := auth.UserIDFromCtx(ctx)
	res, err := a.teamBll.Create(userID, body.Name, key, "member")
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

type tplTeamUpdate map[string]interface{}

// Validate ...
func (t *tplTeamUpdate) Validate() error {
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
//
// @Title Update
// @Summary Update the team
// @Description only the team owner can update the team
// @Param Authorization header string true "access_token"
// @Param teamID path string true "team ID"
// @Param body body tplTeamUpdate true "team body"
// @Success 200 schema.TeamResult
// @Failure 400 string
// @Failure 401 string
// @Router PUT /api/teams/{teamID}
func (a *Team) Update(ctx *gear.Context) (err error) {
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	body := new(tplTeamUpdate)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	team, err := a.models.Team.Find(TeamID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if team.UserID != userID {
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

	res, err := a.models.Team.Update(TeamID, team)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

// RemoveMember ...
//
// @Title RemoveMember
// @Summary remove a team member
// @Description the team owner can remove other team member
// @Description team member can remove self
// @Param Authorization header string true "access_token"
// @Param teamID path string true "team ID"
// @Param userID path string true "team member ID"
// @Success 204
// @Failure 400 string
// @Failure 401 string
// @Router DELETE /api/teams/{teamID}/members/{userID}
func (a *Team) RemoveMember(ctx *gear.Context) (err error) {
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}
	memberID := ctx.Param("userID")
	if memberID == "" {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	if err = a.models.Team.RemoveMember(userID, memberID, TeamID); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}

// Delete ...
//
// @Title Delete
// @Summary Delete the team
// @Description only the team owner can delete the team
// @Param Authorization header string true "access_token"
// @Param teamID path string true "team ID"
// @Success 204
// @Failure 400 string
// @Failure 401 string
// @Router DELETE /api/entries/{teamID}
func (a *Team) Delete(ctx *gear.Context) (err error) {
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	team, err := a.models.Team.Find(TeamID, false)
	if err != nil {
		return ctx.Error(err)
	}
	if team.UserID != userID {
		return ctx.ErrorStatus(403)
	}
	if team.Visibility == "private" {
		return ctx.Error(&gear.Error{Code: 403, Msg: "private team can't be deleted"})
	}

	team.IsDeleted = true
	if _, err = a.models.Team.Update(TeamID, team); err != nil {
		return ctx.Error(err)
	}
	return ctx.End(204)
}

// Undelete ...
//
// @Title Undelete
// @Summary Undelete the team
// @Description only the team owner can undelete the team
// @Param Authorization header string true "access_token"
// @Param teamID path string true "entry ID"
// @Success 204
// @Failure 400 string
// @Failure 401 string
// @Router POST /api/teams/{teamID}:undelete
func (a *Team) Undelete(ctx *gear.Context) (err error) {
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	userID, _ := auth.UserIDFromCtx(ctx)
	team, err := a.models.Team.Find(TeamID, true)
	if err != nil {
		return ctx.Error(err)
	}
	if team.UserID != userID {
		return ctx.ErrorStatus(403)
	}

	team.IsDeleted = false
	res, err := a.models.Team.Update(TeamID, team)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}

// FindByMember return teams for current user
//
// @Title FindByMember
// @Summary Get teams for current user
// @Description Get teams for current user.
// @Param Authorization header string true "access_token"
// @Success 200 []schema.TeamResult
// @Failure 400 string
// @Failure 401 string
// @Router GET /api/teams
func (a *Team) FindByMember(ctx *gear.Context) (err error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	teams, err := a.models.Team.FindByMemberID(userID)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, teams)
}

type tplTeamInvite struct {
	UserID string `json:"userID" swaggo:"true,user to invite,jeo"`
}

func (t *tplTeamInvite) Validate() error {
	if t.UserID == "" {
		return &gear.Error{Code: 400, Msg: "invalid user id"}
	}
	return nil
}

// tplTeamInviteCode ...
type tplTeamInviteCode struct {
	Code string `json:"code"`
}

func (t *tplTeamInviteCode) Validate() error {
	if t.Code == "" {
		return &gear.Error{Code: 400, Msg: "invalid invite code"}
	}
	return nil
}

// Invite a user to the team
//
// @Title Invite
// @Summary Invite a user to the team
// @Description Invite a user to the team
// @Param Authorization header string true "access_token"
// @Param teamID path string true "team ID"
// @Param body body tplTeamInvite true "user to invite"
// @Success 200 tplTeamInviteCode
// @Failure 400 string
// @Failure 401 string
// @Router POST /api/teams/{teamID}/invite
func (a *Team) Invite(ctx *gear.Context) (err error) {
	TeamID, err := util.ParseOID(ctx.Param("teamID"))
	if err != nil {
		return ctx.ErrorStatus(400)
	}

	body := new(tplTeamInvite)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	key, err := auth.KeyFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	userID, _ := auth.UserIDFromCtx(ctx)

	code, err := a.teamBll.Invite(userID, key, body.UserID, TeamID)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, &tplTeamInviteCode{code})
}

// Join join a team by invite code
//
// @Title Join
// @Summary Join a team
// @Description Join a team by invite code
// @Param Authorization header string true "access_token"
// @Param body body tplTeamInviteCode true "invite code"
// @Success 200 schema.TeamResult
// @Failure 400 string
// @Failure 401 string
// @Router POST /api/teams/join
func (a *Team) Join(ctx *gear.Context) (err error) {
	body := new(tplTeamInviteCode)
	if err = ctx.ParseBody(body); err != nil {
		return ctx.Error(err)
	}

	key, err := auth.KeyFromCtx(ctx)
	if err != nil {
		return ctx.Error(err)
	}
	userID, _ := auth.UserIDFromCtx(ctx)
	res, err := a.teamBll.Join(userID, key, body.Code)
	if err != nil {
		return ctx.Error(err)
	}
	return ctx.JSON(200, res)
}
