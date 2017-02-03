package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/seccom/kpass/pkg/schema"
	"github.com/tidwall/buntdb"
)

const keyDBSalt = "DB_SALT"

// DB is buntdb.DB object
type DB struct {
	DB   *buntdb.DB
	Salt []byte
}

// NewDB returns a DB instance
func NewDB(path string) (db *DB, err error) {
	db = &DB{Salt: make([]byte, 64)}
	if path == "" {
		path = ":memory:"
	}
	if db.DB, err = buntdb.Open(path); err != nil {
		return
	}
	err = db.DB.Update(func(tx *buntdb.Tx) error {
		salt, e := tx.Get(keyDBSalt)
		if e != nil {
			if _, e = rand.Read(db.Salt); e == nil {
				_, _, e = tx.Set(keyDBSalt, hex.EncodeToString(db.Salt), nil)
			}
			return e
		}

		if db.Salt, e = hex.DecodeString(salt); e == nil {
			if len(db.Salt) != 64 {
				return errors.New("invalid DBSalt")
			}
		}
		return e
	})

	if err == nil {
		db.DB.CreateIndex("user_by_created", schema.UserKey("*"), buntdb.IndexJSON("created"))
		db.DB.CreateIndex("entry_by_owner", schema.EntryKey("*"), buntdb.IndexJSON("ownerId"))
		db.DB.CreateIndex("team_by_owner", schema.TeamKey("*"), buntdb.IndexJSON("ownerId"))
		db.DB.CreateIndex("share_by_user", schema.ShareKey("*"), buntdb.IndexJSON("to"))
		db.DB.CreateIndex("share_by_entry", schema.ShareKey("*"), buntdb.IndexJSON("entryId"))
	}
	return
}
