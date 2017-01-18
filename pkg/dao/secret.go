package dao

import (
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// Secret is database access oject for secrets
type Secret struct {
	db *service.DB
}

// NewSecret return a Secret intance
func NewSecret(db *service.DB) *Secret {
	return &Secret{db}
}

// Create ...
func (o *Secret) Create(userID, key string, EntryID uuid.UUID, secret *schema.Secret) (
	secretResult *schema.SecretResult, err error) {
	SecretID := util.NewUUID(EntryID.String())
	secret.Created = time.Now()
	secret.Updated = secret.Created
	secretID := SecretID.String()

	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := schema.EntryKey(EntryID.String())
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		// check user right for team
		if entry.OwnerType == "team" {
			value, e := tx.Get(schema.TeamKey(entry.OwnerID))
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
		}

		secretResult = secret.Result(SecretID)
		entry.Secrets = append(entry.Secrets, secretID)
		if value, e = auth.EncryptData(key, secret.String()); e == nil {
			if _, _, e = tx.Set(schema.SecretKey(secretID), value, nil); e == nil {
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
func (o *Secret) Update(userID, key string, EntryID, SecretID uuid.UUID, changes map[string]interface{}) (
	secretResult *schema.SecretResult, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		// transaction: one or more user(team members) may update the secret.
		value, e := tx.Get(schema.EntryKey(EntryID.String()))
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		// check user right for team
		if entry.OwnerType == "team" {
			value, e := tx.Get(schema.TeamKey(entry.OwnerID))
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
		} else if entry.OwnerID != userID {
			return &gear.Error{Code: 403, Msg: "no permission"}
		}

		if value, e = tx.Get(schema.SecretKey(SecretID.String())); e != nil {
			return e
		}
		if value, e = auth.DecryptData(key, value); e != nil {
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
			secret.Updated = time.Now()
			value, e = auth.EncryptData(key, secret.String())
			if e != nil {
				return e
			}
			_, _, e = tx.Set(schema.SecretKey(SecretID.String()), value, nil)
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
func (o *Secret) Delete(userID string, EntryID, SecretID uuid.UUID) error {
	secretID := SecretID.String()
	err := o.db.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := schema.EntryKey(EntryID.String())
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := schema.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return &gear.Error{Code: 404, Msg: "entry not found"}
		}
		if !entry.RemoveSecret(secretID) {
			return &gear.Error{Code: 404, Msg: "secret not found in the entry"}
		}
		// check user right for team
		if entry.OwnerType == "team" {
			value, e := tx.Get(schema.TeamKey(entry.OwnerID))
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
		}
		if _, _, e = tx.Set(entryKey, entry.String(), nil); e == nil {
			_, e = tx.Delete(schema.SecretKey(secretID))
		}
		return e
	})

	return dbError(err)
}

// Find ...
func (o *Secret) Find(key string, SecretID uuid.UUID) (secret *schema.Secret, err error) {
	err = o.db.DB.View(func(tx *buntdb.Tx) error {
		res, e := tx.Get(schema.SecretKey(SecretID.String()))
		if e != nil {
			return e
		}
		res, e = auth.DecryptData(key, res)
		if e != nil {
			return e
		}
		secret, e = schema.SecretFrom(res)
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// FindSecrets ...
func (o *Secret) FindSecrets(key string, ids ...string) (secrets []*schema.SecretResult, err error) {
	err = o.db.DB.View(func(tx *buntdb.Tx) error {
		for _, id := range ids {
			SecretID, _ := uuid.Parse(id)
			res, e := tx.Get(schema.SecretKey(id))
			if e != nil {
				return e
			}
			res, e = auth.DecryptData(key, res)
			if e != nil {
				return e
			}
			secret, e := schema.SecretFrom(res)
			if e != nil {
				return e
			}
			secrets = append(secrets, secret.Result(SecretID))
		}
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}
