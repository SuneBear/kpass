package model

import (
	"time"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// Secret is database access oject for secrets
type Secret struct {
	db *service.DB
}

// Init ...
func (m *Secret) Init(db *service.DB) *Secret {
	m.db = db
	return m
}

// Create ...
func (m *Secret) Create(EntryID util.OID, userID, key string, secret *schema.Secret) (
	secretResult *schema.SecretResult, err error) {
	SecretID := util.NewOID()
	secret.Created = util.Time(time.Now())
	secret.Updated = secret.Created

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

		secretResult = secret.Result(SecretID)
		entry.AddSecret(SecretID.String())
		if value, e = auth.EncryptText(key, secret.String()); e == nil {
			if _, _, e = tx.Set(schema.SecretKey(SecretID), value, nil); e == nil {
				_, _, e = tx.Set(entryKey, entry.String(), nil)
			}
		}
		return e
	})

	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Update ...
func (m *Secret) Update(EntryID, SecretID util.OID, userID, key string, changes map[string]interface{}) (
	secretResult *schema.SecretResult, err error) {
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		// transaction: one or more user(team members) may update the secret.
		value, e := tx.Get(schema.EntryKey(EntryID))
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		if !entry.HasSecret(SecretID.String()) {
			return &gear.Error{Code: 403, Msg: "secret not found in the entry"}
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

		if value, e = tx.Get(schema.SecretKey(SecretID)); e != nil {
			return e
		}
		if value, e = auth.DecryptText(key, value); e != nil {
			return e
		}
		secret, e := schema.SecretFrom(value)
		if e != nil {
			return &gear.Error{Code: 404, Msg: "secret not found"}
		}

		changed := false
		for key, val := range changes {
			switch key {
			case "name":
				if name := val.(string); name != secret.Name {
					changed = true
					secret.Name = name
				}
			case "url":
				if url := val.(string); url != secret.URL {
					changed = true
					secret.URL = url
				}
			case "password":
				if pass := val.(string); pass != secret.Pass {
					changed = true
					secret.Pass = pass
				}
			case "note":
				if note := val.(string); note != secret.Note {
					changed = true
					secret.Note = note
				}
			}
		}

		if changed {
			secret.Updated = util.Time(time.Now())
			value, e = auth.EncryptText(key, secret.String())
			if e != nil {
				return e
			}
			_, _, e = tx.Set(schema.SecretKey(SecretID), value, nil)
		}
		secretResult = secret.Result(SecretID)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Delete ...
func (m *Secret) Delete(EntryID, SecretID util.OID, userID string) error {
	err := m.db.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := schema.EntryKey(EntryID)
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		if !entry.RemoveSecret(SecretID.String()) {
			return &gear.Error{Code: 404, Msg: "secret not found in the entry"}
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
		if _, _, e = tx.Set(entryKey, entry.String(), nil); e == nil {
			_, e = tx.Delete(schema.SecretKey(SecretID))
		}
		return e
	})

	return dbError(err)
}

// Find ...
// func (m *Secret) Find(key string, SecretID util.OID) (secret *schema.Secret, err error) {
// 	err = m.db.DB.View(func(tx *buntdb.Tx) error {
// 		res, e := tx.Get(schema.SecretKey(SecretID))
// 		if e != nil {
// 			return e
// 		}
// 		res, e = auth.DecryptText(key, res)
// 		if e != nil {
// 			return e
// 		}
// 		secret, e = schema.SecretFrom(res)
// 		return e
// 	})
// 	if err != nil {
// 		return nil, dbError(err)
// 	}
// 	return
// }

// FindSecrets ...
func (m *Secret) FindSecrets(key string, ids ...string) (secrets []*schema.SecretResult, err error) {
	secrets = make([]*schema.SecretResult, 0)
	err = m.db.DB.View(func(tx *buntdb.Tx) error {
		for _, id := range ids {
			ID, e := util.ParseOID(id)
			if e != nil {
				return e
			}
			res, e := tx.Get(schema.SecretKey(ID))
			if e != nil {
				return e
			}
			res, e = auth.DecryptText(key, res)
			if e != nil {
				return e
			}
			secret, e := schema.SecretFrom(res)
			if e != nil {
				return e
			}
			secrets = append(secrets, secret.Result(ID))
		}
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}
