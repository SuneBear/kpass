package entryDao

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/app/dao"
	"github.com/tidwall/buntdb"
)

// Create ...
func Create(ownerID, ownerType, name, category string) (entry *dao.Entry, err error) {
	entry = &dao.Entry{
		ID:        uuid.New(),
		OwnerID:   ownerID,
		OwnerType: ownerType,
		Name:      name,
		Category:  category,
		Secrets:   []uuid.UUID{},
		Shares:    []uuid.UUID{},
		Created:   time.Now(),
	}
	if err = Update(entry); err != nil {
		entry = nil
	}
	return
}

// Update ...
func Update(entry *dao.Entry) error {
	return dao.DB.Update(func(tx *buntdb.Tx) error {
		entry.Updated = time.Now()
		_, _, e := tx.Set(dao.EntryKey(entry.ID.String()), entry.String(), nil)
		return e
	})
}

// Find ...
func Find(id string) (entry *dao.Entry, err error) {
	err = dao.DB.View(func(tx *buntdb.Tx) error {
		res, e := tx.Get(dao.EntryKey(id))
		if e == nil {
			e = json.Unmarshal([]byte(res), entry)
		}
		return e
	})
	if err != nil {
		entry = nil
	}
	return
}

// FindByOwnerID ...
func FindByOwnerID(id string, IsDeleted bool) (entries []*dao.Entry, err error) {
	cond := fmt.Sprintf(`{"ownerId":"%s"}`, id)
	err = dao.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendRange("entry_by_owner", cond, cond, func(key, value string) bool {
			entry, e := dao.EntryFrom(value)
			if e != nil {
				e = fmt.Errorf("invalid entry: %s, %s", key, value)
				return false
			}
			if entry.IsDeleted == IsDeleted {
				entries = append(entries, entry)
			}
			return true
		})
		return nil
	})
	if err != nil {
		entries = nil
	}
	return
}
