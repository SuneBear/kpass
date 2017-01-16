package teamDao

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
	"github.com/tidwall/gjson"
)

// Create ...
func Create(ownerID, name, pass string) (res *dao.TeamResult, err error) {
	TeamID := pkg.NewUUID(ownerID)
	team := &dao.Team{
		Name:    name,
		Pass:    pkg.Auth.EncryptUserPass(TeamID.String(), pass),
		OwnerID: ownerID,
		Members: []string{ownerID},
		Created: time.Now(),
	}
	team.Updated = team.Created
	res = team.Result(TeamID)
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		_, _, e := tx.Set(dao.TeamKey(TeamID.String()), team.String(), nil)
		return e
	})
	if err != nil {
		res = nil
		err = dao.DBError(err)
	}
	return
}

// Update ...
func Update(TeamID uuid.UUID, team *dao.Team) (res *dao.TeamResult, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		team.Updated = time.Now()
		_, _, e := tx.Set(dao.TeamKey(TeamID.String()), team.String(), nil)
		return e
	})
	if err != nil {
		return nil, dao.DBError(err)
	}
	return team.Result(TeamID), nil
}

// UpdateMembers ...
func UpdateMembers(userID string, TeamID uuid.UUID, pull, push []string) (res *dao.TeamResult, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		teamID := TeamID.String()
		value, e := tx.Get(dao.TeamKey(teamID))
		if e != nil {
			return e
		}
		team, e := dao.TeamFrom(value)
		if e != nil {
			return e
		}
		if team.IsDeleted {
			return buntdb.ErrNotFound
		}
		if !team.HasMember(userID) {
			return &gear.Error{Code: 403, Msg: "not team member"}
		}
		for _, user := range pull {
			if !team.RemoveMember(user) {
				return &gear.Error{Code: 400, Msg: "invalid team member to remove"}
			}
		}
		for _, user := range push {
			if _, e := tx.Get(dao.UserKey(user)); e != nil {
				// user not exists
				return &gear.Error{Code: 400, Msg: e.Error()}
			}
			team.AddMember(user)
		}
		res = team.Result(TeamID)
		_, _, e = tx.Set(dao.TeamKey(teamID), team.String(), nil)
		return e
	})
	if err != nil {
		res = nil
		err = dao.DBError(err)
	}
	return
}

// Find ...
func Find(TeamID uuid.UUID, IsDeleted bool) (team *dao.Team, err error) {
	err = dao.DB.View(func(tx *buntdb.Tx) (e error) {
		var res string
		if res, e = tx.Get(dao.TeamKey(TeamID.String())); e == nil {
			if team, e = dao.TeamFrom(res); e == nil {
				if team.IsDeleted != IsDeleted {
					e = buntdb.ErrNotFound
				}
			}
		}
		return e
	})
	if err != nil {
		team = nil
		err = dao.DBError(err)
	}
	return
}

// FindByOwnerID ...
func FindByOwnerID(ownerID string, IsDeleted bool) (teams []*dao.TeamResult, err error) {
	teams = make([]*dao.TeamResult, 0)
	cond := fmt.Sprintf(`{"ownerID":"%s"}`, ownerID)
	err = dao.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendGreaterOrEqual("team_by_owner", cond, func(key, value string) bool {

			team, e := dao.TeamFrom(value)
			if e != nil {
				e = fmt.Errorf("invalid team: %s, %s", key, value)
				return false
			}
			if team.OwnerID != ownerID {
				return false
			}
			if team.IsDeleted == IsDeleted {
				TeamID := dao.TeamIDFromKey(key)
				teams = append(teams, team.Result(TeamID))
			}
			return true
		})
		return nil
	})
	if err != nil {
		teams = nil
		err = dao.DBError(err)
	}
	return
}

// FindByMemberID ...
func FindByMemberID(memberID string) (teams []*dao.TeamResult, err error) {
	teams = make([]*dao.TeamResult, 0)
	conds := fmt.Sprintf(`members.#["%s"]#`, memberID)
	err = dao.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.Ascend("team_by_owner", func(key, value string) bool {
			if gjson.Get(value, conds).String() == memberID {
				team, e := dao.TeamFrom(value)
				if e != nil {
					e = fmt.Errorf("invalid team: %s, %s", key, value)
					return false
				}
				if team.IsDeleted == false {
					TeamID := dao.TeamIDFromKey(key)
					teams = append(teams, team.Result(TeamID))
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		teams = nil
		err = dao.DBError(err)
	}
	return
}

// CheckToken ...
func CheckToken(teamID, userID, pass string) (team *dao.Team, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		userKey := dao.UserKey(userID)
		value, e := tx.Get(userKey)
		if e != nil {
			return e
		}
		user, e := dao.UserFrom(value)
		if e != nil {
			return e
		}
		if user.IsBlocked || user.Attempt > 5 {
			return &gear.Error{Code: 403, Msg: "too many login attempts"}
		}
		value, e = tx.Get(dao.TeamKey(teamID))
		if e != nil {
			return e
		}
		team, e := dao.TeamFrom(value)
		if e != nil {
			return e
		}
		if team.IsDeleted {
			return buntdb.ErrNotFound
		}
		if !team.HasMember(userID) {
			return &gear.Error{Code: 403, Msg: "not team member"}
		}
		if !pkg.Auth.ValidateUserPass(teamID, pass, team.Pass) {
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
		team = nil
		err = dao.DBError(err)
	}
	return
}
