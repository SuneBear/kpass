package entryDao

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// Create ...
func Create(userID, ownerID, ownerType, name, category string) (entrySum *dao.EntrySum, err error) {
	EntryID := pkg.NewUUID(ownerID)
	entry := &dao.Entry{
		OwnerID:   ownerID,
		OwnerType: ownerType,
		Name:      name,
		Category:  category,
		Secrets:   []string{},
		Shares:    []string{},
		Created:   time.Now(),
	}
	entry.Updated = entry.Created
	entrySum = entry.Summary(EntryID)
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		// check right for team
		if ownerType == "team" {
			value, e := tx.Get(dao.TeamKey(ownerID))
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
		}
		_, _, e := tx.Set(dao.EntryKey(EntryID.String()), entry.String(), nil)
		return e
	})
	if err != nil {
		entrySum = nil
		err = dao.DBError(err)
	}
	return
}

// Update ...
func Update(EntryID uuid.UUID, entry *dao.Entry) (entrySum *dao.EntrySum, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		entry.Updated = time.Now()
		_, _, e := tx.Set(dao.EntryKey(EntryID.String()), entry.String(), nil)
		return e
	})
	if err != nil {
		return nil, dao.DBError(err)
	}
	return entry.Summary(EntryID), nil
}

// Find ...
func Find(EntryID uuid.UUID, IsDeleted bool) (entry *dao.Entry, err error) {
	err = dao.DB.View(func(tx *buntdb.Tx) (e error) {
		var res string
		if res, e = tx.Get(dao.EntryKey(EntryID.String())); e == nil {
			if entry, e = dao.EntryFrom(res); e == nil {
				if entry.IsDeleted != IsDeleted {
					e = buntdb.ErrNotFound
				}
			}
		}
		return e
	})
	if err != nil {
		entry = nil
		err = dao.DBError(err)
	}
	return
}

// FindByOwnerID ...
func FindByOwnerID(userID, ownerID, ownerType string, IsDeleted bool) (entries []*dao.EntrySum, err error) {
	entries = make([]*dao.EntrySum, 0)
	cond := fmt.Sprintf(`{"ownerId":"%s"}`, ownerID)
	err = dao.DB.View(func(tx *buntdb.Tx) (e error) {
		// check right for team
		if ownerType == "team" {
			value, e := tx.Get(dao.TeamKey(ownerID))
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
		}

		tx.AscendGreaterOrEqual("entry_by_owner", cond, func(key, value string) bool {
			entry, e := dao.EntryFrom(value)
			if e != nil {
				e = fmt.Errorf("invalid entry: %s, %s", key, value)
				return false
			}
			if entry.OwnerID != ownerID {
				return false
			}
			if entry.IsDeleted == IsDeleted {
				EntryID := dao.EntryIDFromKey(key)
				entries = append(entries, entry.Summary(EntryID))
			}
			return true
		})
		return nil
	})
	if err != nil {
		entries = nil
		err = dao.DBError(err)
	}
	return
}
