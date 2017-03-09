package model

import (
	"fmt"
	"time"

	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// Entry is database access oject for entries
type Entry struct {
	db *service.DB
}

// Init ...
func (m *Entry) Init(db *service.DB) *Entry {
	m.db = db
	return m
}

// Create ...
func (m *Entry) Create(userID string, entry *schema.Entry) (entrySum *schema.EntrySum, err error) {
	EntryID := util.NewOID()
	entry.Created = util.Time(time.Now())
	entry.Updated = entry.Created
	entrySum = entry.Summary(EntryID)
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		_, _, e := tx.Set(schema.EntryKey(EntryID), entry.String(), nil)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Update ...
func (m *Entry) Update(userID string, EntryID util.OID, entry *schema.Entry) (err error) {
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		entry.Updated = util.Time(time.Now())
		_, _, e := tx.Set(schema.EntryKey(EntryID), entry.String(), nil)
		return e
	})
	return dbError(err)
}

// UpdateDeleted ...
func (m *Entry) UpdateDeleted(userID string, EntryID util.OID, isDeleted bool) (
	entrySum *schema.EntrySum, err error) {
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
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
		entry.Updated = util.Time(time.Now())
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
func (m *Entry) Find(EntryID util.OID, IsDeleted bool) (entry *schema.Entry, err error) {
	err = m.db.DB.View(func(tx *buntdb.Tx) (e error) {
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
func (m *Entry) FindByTeam(TeamID util.OID, userID string, IsDeleted bool) (
	entries []*schema.EntrySum, err error) {
	entries = make([]*schema.EntrySum, 0)
	cond := fmt.Sprintf(`{"teamID":"%s"}`, TeamID.String())
	err = m.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendGreaterOrEqual("entry_by_team", cond, func(key, value string) bool {
			entry, e := schema.EntryFrom(value)
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

// AddFileByID ...
func (m *Entry) AddFileByID(EntryID, FileID util.OID, userID string) (err error) {
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := schema.EntryKey(EntryID)
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}

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

		entry.AddFile(FileID.String())
		_, _, e = tx.Set(entryKey, entry.String(), nil)
		return e
	})

	return dbError(err)
}

// RemoveFileByID ...
func (m *Entry) RemoveFileByID(EntryID, FileID util.OID, userID string) (err error) {
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := schema.EntryKey(EntryID)
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		if !entry.RemoveFile(FileID.String()) {
			return &gear.Error{Code: 404, Msg: "file not found in the entry"}
		}

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

		_, _, e = tx.Set(entryKey, entry.String(), nil)
		return e
	})

	return dbError(err)
}
