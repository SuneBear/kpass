package dao

import (
	"fmt"
	"time"

	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// Entry is database access oject for entries
type Entry struct {
	db *service.DB
}

// NewEntry return a Entry intance
func NewEntry(db *service.DB) *Entry {
	return &Entry{db}
}

// Create ...
func (o *Entry) Create(userID string, entry *schema.Entry) (entrySum *schema.EntrySum, err error) {
	EntryID := util.NewOID()
	entry.Created = time.Now()
	entry.Updated = entry.Created
	entrySum = entry.Summary(EntryID)
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		value, e := tx.Get(schema.TeamKey(entry.TeamID))
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
		if team.IsFrozen {
			return &gear.Error{Code: 403, Msg: "team is frozen"}
		}
		_, _, e = tx.Set(schema.EntryKey(EntryID), entry.String(), nil)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Update ...
func (o *Entry) Update(userID string, EntryID util.OID, changes map[string]interface{}) (
	entrySum *schema.EntrySum, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		// transaction: one or more user(team members) may update the entry.
		value, e := tx.Get(schema.EntryKey(EntryID))
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}

		// check user right for team
		value, e = tx.Get(schema.TeamKey(entry.TeamID))
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
		if team.IsFrozen {
			return &gear.Error{Code: 403, Msg: "team is frozen"}
		}

		changed := false
		for key, val := range changes {
			switch key {
			case "name":
				if name := val.(string); name != entry.Name {
					changed = true
					entry.Name = name
				}
			case "category":
				if category := val.(string); category != entry.Category {
					changed = true
					entry.Category = category
				}
			case "priority":
				if priority := int(val.(float64)); priority != entry.Priority {
					changed = true
					entry.Priority = priority
				}
			}
		}

		if changed {
			entry.Updated = time.Now()
			_, _, e = tx.Set(schema.EntryKey(EntryID), entry.String(), nil)
		}
		entrySum = entry.Summary(EntryID)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// UpdateDeleted ...
func (o *Entry) UpdateDeleted(userID string, EntryID util.OID, isDeleted bool) (
	entrySum *schema.EntrySum, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		// transaction: one or more user(team members) may update the entry.
		value, e := tx.Get(schema.EntryKey(EntryID))
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil {
			return e
		}

		entry.IsDeleted = isDeleted
		entry.Updated = time.Now()
		_, _, e = tx.Set(schema.EntryKey(EntryID), entry.String(), nil)
		entrySum = entry.Summary(EntryID)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Find ...
func (o *Entry) Find(EntryID util.OID, IsDeleted bool) (entry *schema.Entry, err error) {
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		var res string
		if res, e = tx.Get(schema.EntryKey(EntryID)); e == nil {
			if entry, e = schema.EntryFrom(res); e == nil {
				if entry.IsDeleted != IsDeleted {
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

// FindByTeam ...
func (o *Entry) FindByTeam(TeamID util.OID, userID string, IsDeleted bool) (
	entries []*schema.EntrySum, err error) {
	entries = make([]*schema.EntrySum, 0)
	cond := fmt.Sprintf(`{"teamID":"%s"}`, TeamID.String())
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendGreaterOrEqual("entry_by_team", cond, func(key, value string) bool {
			entry, e := schema.EntryFrom(value)
			fmt.Println(1212, key, value, entry)
			if e != nil {
				e = fmt.Errorf("invalid entry: %s, %s", key, value)
				return false
			}
			if entry.TeamID.String() != TeamID.String() {
				return false
			}
			if entry.IsDeleted == IsDeleted {
				EntryID := schema.EntryIDFromKey(key)
				entries = append(entries, entry.Summary(EntryID))
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
