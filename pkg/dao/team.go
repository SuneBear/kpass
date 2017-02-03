package dao

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
	"github.com/tidwall/gjson"
)

// Team is database access oject for teams
type Team struct {
	db *service.DB
}

// NewTeam return a Team intance
func NewTeam(db *service.DB) *Team {
	return &Team{db}
}

// Create ...
func (o *Team) Create(ownerID, name, pass string) (teamResult *schema.TeamResult, err error) {
	TeamID := util.NewUUID(ownerID)
	team := &schema.Team{
		Name:    name,
		Pass:    auth.SignPass(TeamID.String(), pass),
		OwnerID: ownerID,
		Members: []string{ownerID},
		Created: time.Now(),
	}
	team.Updated = team.Created
	teamResult = team.Result(TeamID)
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		_, _, e := tx.Set(schema.TeamKey(TeamID.String()), team.String(), nil)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Update ...
func (o *Team) Update(TeamID uuid.UUID, team *schema.Team) (teamResult *schema.TeamResult, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		team.Updated = time.Now()
		_, _, e := tx.Set(schema.TeamKey(TeamID.String()), team.String(), nil)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return team.Result(TeamID), nil
}

// UpdateMembers ...
func (o *Team) UpdateMembers(userID string, TeamID uuid.UUID, pull, push []string) (
	teamResult *schema.TeamResult, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		teamID := TeamID.String()
		value, e := tx.Get(schema.TeamKey(teamID))
		if e != nil {
			return e
		}
		team, e := schema.TeamFrom(value)
		if e != nil {
			return e
		}
		if team.IsDeleted {
			return &gear.Error{Code: 404, Msg: "team not found"}
		}
		if team.OwnerID != userID {
			return &gear.Error{Code: 403, Msg: "not team owner"}
		}
		for _, user := range pull {
			if !team.RemoveMember(user) {
				return &gear.Error{Code: 400, Msg: "invalid team member to remove"}
			}
		}
		for _, user := range push {
			if _, e := tx.Get(schema.UserKey(user)); e != nil {
				// user not exists
				return &gear.Error{Code: 400, Msg: e.Error()}
			}
			team.AddMember(user)
		}
		teamResult = team.Result(TeamID)
		_, _, e = tx.Set(schema.TeamKey(teamID), team.String(), nil)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Find ...
func (o *Team) Find(TeamID uuid.UUID, IsDeleted bool) (team *schema.Team, err error) {
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		var res string
		if res, e = tx.Get(schema.TeamKey(TeamID.String())); e == nil {
			if team, e = schema.TeamFrom(res); e == nil {
				if team.IsDeleted != IsDeleted {
					e = buntdb.ErrNotFound
				}
			}
		}
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// FindByOwnerID ...
func (o *Team) FindByOwnerID(ownerID string, IsDeleted bool) (teams []*schema.TeamResult, err error) {
	teams = make([]*schema.TeamResult, 0)
	cond := fmt.Sprintf(`{"ownerID":"%s"}`, ownerID)
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendGreaterOrEqual("team_by_owner", cond, func(key, value string) bool {

			team, e := schema.TeamFrom(value)
			if e != nil {
				e = fmt.Errorf("invalid team: %s, %s", key, value)
				return false
			}
			if team.OwnerID != ownerID {
				return false
			}
			if team.IsDeleted == IsDeleted {
				TeamID := schema.TeamIDFromKey(key)
				teams = append(teams, team.Result(TeamID))
			}
			return true
		})
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// FindByMemberID ...
func (o *Team) FindByMemberID(memberID string) (teams []*schema.TeamResult, err error) {
	teams = make([]*schema.TeamResult, 0)
	conds := fmt.Sprintf(`members.#["%s"]#`, memberID)
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.Ascend("team_by_owner", func(key, value string) bool {
			if gjson.Get(value, conds).String() == memberID {
				team, e := schema.TeamFrom(value)
				if e != nil {
					e = fmt.Errorf("invalid team: %s, %s", key, value)
					return false
				}
				if team.IsDeleted == false {
					TeamID := schema.TeamIDFromKey(key)
					teams = append(teams, team.Result(TeamID))
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// CheckToken ...
func (o *Team) CheckToken(teamID, userID, pass string) (team *schema.Team, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		userKey := schema.UserKey(userID)
		value, e := tx.Get(userKey)
		if e != nil {
			return e
		}
		user, e := schema.UserFrom(value)
		if e != nil {
			return e
		}
		if user.IsBlocked || user.Attempt > 5 {
			return &gear.Error{Code: 403, Msg: "too many login attempts"}
		}
		value, e = tx.Get(schema.TeamKey(teamID))
		if e != nil {
			return e
		}
		team, e := schema.TeamFrom(value)
		if e != nil {
			return e
		}
		if team.IsDeleted {
			return buntdb.ErrNotFound
		}
		if !team.HasMember(userID) {
			return &gear.Error{Code: 403, Msg: "not team member"}
		}
		if !auth.VerifyPass(teamID, pass, team.Pass) {
			user.Attempt++
			tx.Set(userKey, user.String(), nil)
			tx.Commit()
			return &gear.Error{Code: 400, Msg: "team password error"}
		}
		if user.Attempt > 0 {
			user.Attempt = 0
			tx.Set(userKey, user.String(), nil)
		}
		return nil
	})

	if err != nil {
		return nil, dbError(err)
	}
	return
}
