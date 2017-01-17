package secretDao

import (
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// Create ...
func Create(userID, key string, EntryID uuid.UUID, secret *dao.Secret) (
	secretResult *dao.SecretResult, err error) {
	SecretID := pkg.NewUUID(EntryID.String())
	secret.Created = time.Now()
	secret.Updated = secret.Created
	secretID := SecretID.String()

	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := dao.EntryKey(EntryID.String())
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := dao.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		// check user right for team
		if entry.OwnerType == "team" {
			value, e := tx.Get(dao.TeamKey(entry.OwnerID))
			if e != nil {
				return e
			}
			team, e := dao.TeamFrom(value)
			if e != nil || team.IsDeleted {
				return &gear.Error{Code: 404, Msg: "team not found"}
			}
			if !team.HasMember(userID) {
				return &gear.Error{Code: 403, Msg: "not team member"}
			}
			if team.IsFrozen {
				return &gear.Error{Code: 403, Msg: "team is frozen"}
			}
		}

		secretResult = secret.Result(SecretID)
		entry.Secrets = append(entry.Secrets, secretID)
		if value, e = pkg.Auth.EncryptData(key, secret.String()); e == nil {
			if _, _, e = tx.Set(dao.SecretKey(secretID), value, nil); e == nil {
				_, _, e = tx.Set(entryKey, entry.String(), nil)
			}
		}
		return e
	})

	if err != nil {
		return nil, dao.DBError(err)
	}
	return
}

// Update ...
func Update(userID, key string, EntryID, SecretID uuid.UUID, changes map[string]interface{}) (
	secretResult *dao.SecretResult, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		// transaction: one or more user(team members) may update the secret.
		value, e := tx.Get(dao.EntryKey(EntryID.String()))
		if e != nil {
			return e
		}
		entry, e := dao.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		// check user right for team
		if entry.OwnerType == "team" {
			value, e := tx.Get(dao.TeamKey(entry.OwnerID))
			if e != nil {
				return e
			}
			team, e := dao.TeamFrom(value)
			if e != nil || team.IsDeleted {
				return &gear.Error{Code: 404, Msg: "team not found"}
			}
			if !team.HasMember(userID) {
				return &gear.Error{Code: 403, Msg: "not team member"}
			}
			if team.IsFrozen {
				return &gear.Error{Code: 403, Msg: "team is frozen"}
			}
		} else if entry.OwnerID != userID {
			return &gear.Error{Code: 403, Msg: "no permission"}
		}

		if value, e = tx.Get(dao.SecretKey(SecretID.String())); e != nil {
			return e
		}
		if value, e = pkg.Auth.DecryptData(key, value); e != nil {
			return e
		}
		secret, e := dao.SecretFrom(value)
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
			secret.Updated = time.Now()
			value, e = pkg.Auth.EncryptData(key, secret.String())
			if e != nil {
				return e
			}
			_, _, e = tx.Set(dao.SecretKey(SecretID.String()), value, nil)
		}
		secretResult = secret.Result(SecretID)
		return e
	})
	if err != nil {
		return nil, dao.DBError(err)
	}
	return
}

// Delete ...
func Delete(userID string, EntryID, SecretID uuid.UUID) error {
	secretID := SecretID.String()
	err := dao.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := dao.EntryKey(EntryID.String())
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := dao.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		if !entry.RemoveSecret(secretID) {
			return &gear.Error{Code: 404, Msg: "secret not found in the entry"}
		}
		// check user right for team
		if entry.OwnerType == "team" {
			value, e := tx.Get(dao.TeamKey(entry.OwnerID))
			if e != nil {
				return e
			}
			team, e := dao.TeamFrom(value)
			if e != nil || team.IsDeleted {
				return &gear.Error{Code: 404, Msg: "team not found"}
			}
			if !team.HasMember(userID) {
				return &gear.Error{Code: 403, Msg: "not team member"}
			}
			if team.IsFrozen {
				return &gear.Error{Code: 403, Msg: "team is frozen"}
			}
		}
		if _, _, e = tx.Set(entryKey, entry.String(), nil); e == nil {
			_, e = tx.Delete(dao.SecretKey(secretID))
		}
		return e
	})

	return dao.DBError(err)
}

// Find ...
func Find(key string, SecretID uuid.UUID) (secret *dao.Secret, err error) {
	err = dao.DB.View(func(tx *buntdb.Tx) error {
		res, e := tx.Get(dao.SecretKey(SecretID.String()))
		if e != nil {
			return e
		}
		res, e = pkg.Auth.DecryptData(key, res)
		if e != nil {
			return e
		}
		secret, e = dao.SecretFrom(res)
		return e
	})
	if err != nil {
		return nil, dao.DBError(err)
	}
	return
}

// FindSecrets ...
func FindSecrets(key string, ids ...string) (secrets []*dao.SecretResult, err error) {
	err = dao.DB.View(func(tx *buntdb.Tx) error {
		for _, id := range ids {
			SecretID, _ := uuid.Parse(id)
			res, e := tx.Get(dao.SecretKey(id))
			if e != nil {
				return e
			}
			res, e = pkg.Auth.DecryptData(key, res)
			if e != nil {
				return e
			}
			secret, e := dao.SecretFrom(res)
			if e != nil {
				return e
			}
			secrets = append(secrets, secret.Result(SecretID))
		}
		return nil
	})
	if err != nil {
		return nil, dao.DBError(err)
	}
	return
}
