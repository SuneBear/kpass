package secretDao

import (
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/tidwall/buntdb"
)

// Create ...
func Create(key string, EntryID uuid.UUID, secret *dao.Secret) (res *dao.SecretResult, err error) {
	SecretID := uuid.New()
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
			e = buntdb.ErrNotFound
			return e
		}

		res = secret.Result(SecretID)
		entry.Secrets = append(entry.Secrets, secretID)
		if value, e = pkg.Auth.EncryptData(key, secret.String()); e == nil {
			if _, _, e = tx.Set(dao.SecretKey(secretID), value, nil); e == nil {
				_, _, e = tx.Set(entryKey, entry.String(), nil)
			}
		}
		return e
	})

	if err != nil {
		res = nil
		err = dao.DBError(err)
	}
	return
}

// Update ...
func Update(key string, SecretID uuid.UUID, secret *dao.Secret) (res *dao.SecretResult, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		secret.Updated = time.Now()
		res = secret.Result(SecretID)
		s, e := pkg.Auth.EncryptData(key, secret.String())
		if e == nil {
			_, _, e = tx.Set(dao.SecretKey(SecretID.String()), s, nil)
		}
		return e
	})
	if err != nil {
		res = nil
		err = dao.DBError(err)
	}
	return
}

// Delete ...
func Delete(EntryID, SecretID uuid.UUID) error {
	secretID := SecretID.String()
	err := dao.DB.Update(func(tx *buntdb.Tx) error {
		entryKey := dao.EntryKey(EntryID.String())
		value, e := tx.Get(entryKey)
		if e != nil {
			return e
		}
		entry, e := dao.EntryFrom(value)
		if e != nil || entry.IsDeleted {
			return buntdb.ErrNotFound
		}

		if !entry.RemoveSecret(secretID) {
			return buntdb.ErrNotFound
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
		secret = nil
		err = dao.DBError(err)
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
		secrets = nil
		err = dao.DBError(err)
	}
	return
}
