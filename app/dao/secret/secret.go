package secretDao

import (
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/tidwall/buntdb"
)

// Create ...
func Create(key string, secret *dao.Secret) (res *dao.SecretResult, err error) {
	id := uuid.New()
	secret.Created = time.Now()
	return Update(key, id, secret)
}

// Update ...
func Update(key string, id uuid.UUID, secret *dao.Secret) (res *dao.SecretResult, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		secret.Updated = time.Now()
		res = secret.Result(id)
		s, e := pkg.Auth.EncryptData(key, secret.String())
		if e == nil {
			_, _, e = tx.Set(dao.SecretKey(id.String()), s, nil)
		}
		return e
	})
	if err != nil {
		res = nil
	}
	return
}

// FindSecrets ...
func FindSecrets(key string, ids ...uuid.UUID) (secrets []*dao.SecretResult, err error) {
	err = dao.DB.View(func(tx *buntdb.Tx) error {
		for _, id := range ids {
			res, e := tx.Get(dao.SecretKey(id.String()))
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
			secrets = append(secrets, secret.Result(id))
		}
		return nil
	})
	if err != nil {
		secrets = nil
	}
	return
}
