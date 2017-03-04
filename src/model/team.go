package model

import (
	"fmt"
	"time"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
	"github.com/tidwall/gjson"
)

// Team is database access oject for teams
type Team struct {
	db *service.DB
}

// Init ...
func (m *Team) Init(db *service.DB) *Team {
	m.db = db
	return m
}

// Create ...
func (m *Team) Create(userID, pass string, team *schema.Team) (teamResult *schema.TeamResult, err error) {
	TeamID := util.NewOID()
	team.Pass = auth.SignPass(TeamID.String(), pass)
	team.Created = util.Time(time.Now())
	team.Updated = team.Created
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		_, _, e := tx.Set(schema.TeamKey(TeamID), team.String(), nil)
		if e == nil {
			teamResult = team.Result(TeamID, IdsToUsers(tx, team.Members))
		}
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Update ...
func (m *Team) Update(TeamID util.OID, team *schema.Team) (teamResult *schema.TeamResult, err error) {
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		team.Updated = util.Time(time.Now())
		_, _, e := tx.Set(schema.TeamKey(TeamID), team.String(), nil)
		if e == nil {
			teamResult = team.Result(TeamID, IdsToUsers(tx, team.Members))
		}
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// AddMember ...
func (m *Team) AddMember(ownerID, userID string, TeamID util.OID) (
	teamResult *schema.TeamResult, err error) {
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		value, e := tx.Get(schema.TeamKey(TeamID))
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
		if team.UserID != ownerID {
			return &gear.Error{Code: 403, Msg: "not team owner"}
		}
		if team.Visibility == "private" {
			return &gear.Error{Code: 403, Msg: "can't change member in private team"}
		}

		if !team.AddMember(userID) {
			return &gear.Error{Code: 409, Msg: "already team member"}
		}
		_, _, e = tx.Set(schema.TeamKey(TeamID), team.String(), nil)
		if e == nil {
			teamResult = team.Result(TeamID, IdsToUsers(tx, team.Members))
		}
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// RemoveMember ...
func (m *Team) RemoveMember(ownerID, userID string, TeamID util.OID) error {
	err := m.db.DB.Update(func(tx *buntdb.Tx) error {
		value, e := tx.Get(schema.TeamKey(TeamID))
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
		if team.UserID != ownerID {
			return &gear.Error{Code: 403, Msg: "not team owner"}
		}
		if team.Visibility == "private" {
			return &gear.Error{Code: 403, Msg: "can't change member in private team"}
		}
		if !team.RemoveMember(userID) {
			return &gear.Error{Code: 400, Msg: "invalid team member to remove"}
		}
		_, _, e = tx.Set(schema.TeamKey(TeamID), team.String(), nil)
		return e
	})
	return dbError(err)
}

// Find ...
func (m *Team) Find(TeamID util.OID, IsDeleted bool) (team *schema.Team, err error) {
	err = m.db.DB.View(func(tx *buntdb.Tx) (e error) {
		var res string
		if res, e = tx.Get(schema.TeamKey(TeamID)); e == nil {
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

// FindByUserID ...
func (m *Team) FindByUserID(userID string, IsDeleted bool) (teams []*schema.TeamResult, err error) {
	teams = make([]*schema.TeamResult, 0)
	cond := fmt.Sprintf(`{"userID":"%s"}`, userID)
	err = m.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendGreaterOrEqual("team_by_user", cond, func(key, value string) bool {
			team, e := schema.TeamFrom(value)
			if e != nil {
				e = fmt.Errorf("invalid team: %s, %s", key, value)
				return false
			}
			if team.UserID != userID {
				return false
			}
			if team.IsDeleted == IsDeleted {
				TeamID := schema.TeamIDFromKey(key)
				teamResult := team.Result(TeamID, IdsToUsers(tx, team.Members))
				teams = append(teams, teamResult)
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
func (m *Team) FindByMemberID(memberID string) (teams []*schema.TeamResult, err error) {
	teams = make([]*schema.TeamResult, 0)
	err = m.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.Ascend("team_by_user", func(key, value string) bool {
			for _, r := range gjson.Get(value, "members").Array() {
				if r.String() == memberID {
					team, e := schema.TeamFrom(value)
					if e != nil {
						e = fmt.Errorf("invalid team: %s, %s", key, value)
						return false
					}
					if team.IsDeleted == false {
						TeamID := schema.TeamIDFromKey(key)
						teamResult := team.Result(TeamID, IdsToUsers(tx, team.Members))
						teams = append(teams, teamResult)
					}
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

// CheckUser check user' read right
func (m *Team) CheckUser(TeamID util.OID, userID string) error {
	err := m.db.DB.View(func(tx *buntdb.Tx) error {
		value, e := tx.Get(schema.TeamKey(TeamID))
		if e != nil {
			return e
		}
		team, e := schema.TeamFrom(value)
		if e != nil || team.IsDeleted {
			return &gear.Error{Code: 404, Msg: "team not found"}
		}
		if !team.HasMember(userID) {
			return &gear.Error{Code: 403, Msg: "not team member"}
		}
		if team.Visibility == "private" && team.UserID != userID {
			return &gear.Error{Code: 403, Msg: "private team"}
		}

		return nil
	})

	return dbError(err)
}

// SavePass ...
func (m *Team) SavePass(TeamID util.OID, userID, key, teamPass string) error {
	value, err := auth.EncryptText(key, teamPass)
	if err != nil {
		return dbError(err)
	}
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		_, _, e := tx.Set(schema.TeamPassKey(TeamID, userID), value, nil)
		return e
	})
	return dbError(err)
}

// GetPass ...
func (m *Team) GetPass(TeamID util.OID, userID, key string) (string, error) {
	teamPass := ""
	err := m.db.DB.View(func(tx *buntdb.Tx) error {
		val, e := tx.Get(schema.TeamPassKey(TeamID, userID))
		if e == nil {
			teamPass, e = auth.DecryptText(key, val)
		}
		return e
	})
	if err != nil {
		return "", dbError(err)
	}
	return teamPass, nil
}

// GetKey ...
func (m *Team) GetKey(TeamID util.OID, userID, key string) (string, error) {
	teamKey := ""
	err := m.db.DB.View(func(tx *buntdb.Tx) error {
		teamPass := ""
		val, e := tx.Get(schema.TeamPassKey(TeamID, userID))
		if e == nil {
			teamPass, e = auth.DecryptText(key, val)
		}
		if e == nil {
			if val, e = tx.Get(schema.TeamKey(TeamID)); e == nil {
				var team *schema.Team
				if team, e = schema.TeamFrom(val); e == nil {
					teamKey = auth.AESKey(teamPass, team.Pass)
				}
			}
		}
		return e
	})
	if err != nil {
		return "", dbError(err)
	}
	return teamKey, nil
}
