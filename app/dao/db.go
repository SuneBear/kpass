package dao

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/tidwall/buntdb"
)

var (
	DB     *buntdb.DB
	DBSalt = make([]byte, 128)
)

// Open open db
func Open(path string) (err error) {
	if path == "" {
		path = ":memory:"
	}
	if DB, err = buntdb.Open(path); err != nil {
		return
	}
	return DB.Update(func(tx *buntdb.Tx) error {
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

// InitIndex ...
func InitIndex() {
	DB.CreateIndex("user_by_created", UserKey("*"), buntdb.IndexJSON("created"))
	DB.CreateIndex("entry_by_owner", EntryKey("*"), buntdb.IndexJSON("ownerId"))
}
