package dao

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/tidwall/buntdb"
)

var db *buntdb.DB
var DBSalt = make([]byte, 128)

// Open open db
func Open(path string) (err error) {
	if path == "" {
		path = ":memory:"
	}
	if db, err = buntdb.Open(path); err != nil {
		return
	}
	return db.Update(func(tx *buntdb.Tx) error {
		salt, e := tx.Get(keyDBSalt)
		if e != nil {
			if _, e = rand.Read(DBSalt); e == nil {
				_, _, e = tx.Set(keyDBSalt, hex.EncodeToString(DBSalt), nil)
			}
			return e
		}

		if DBSalt, e = hex.DecodeString(salt); e == nil {
			if len(DBSalt) != 128 {
				return errors.New("invalid DBSalt")
			}
		}
		return e
	})
}
