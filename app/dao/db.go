package dao

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

var (
	DB     *buntdb.DB
	DBSalt = make([]byte, 64)
)

// Open open db
func Open(path string) error {
	var err error
	if path == "" {
		path = ":memory:"
	}
	if DB, err = buntdb.Open(path); err != nil {
		return DBError(err)
	}
	err = DB.Update(func(tx *buntdb.Tx) error {
		salt, e := tx.Get(keyDBSalt)
		if e != nil {
			if _, e = rand.Read(DBSalt); e == nil {
				_, _, e = tx.Set(keyDBSalt, hex.EncodeToString(DBSalt), nil)
			}
			return e
		}

		if DBSalt, e = hex.DecodeString(salt); e == nil {
			if len(DBSalt) != 64 {
				return errors.New("invalid DBSalt")
			}
		}
		return e
	})
	return DBError(err)
}

// InitIndex ...
func InitIndex() {
	DB.CreateIndex("user_by_created", UserKey("*"), buntdb.IndexJSON("created"))
	DB.CreateIndex("entry_by_owner", EntryKey("*"), buntdb.IndexJSON("ownerId"))
}

// DBError ...
func DBError(err error) error {
	if err == nil {
		return nil
	}
	if err == buntdb.ErrNotFound {
		return &gear.Error{Code: 404, Msg: err.Error()}
	}
	if _, ok := err.(*gear.Error); ok {
		return err
	}
	return &gear.Error{Code: 500, Msg: err.Error()}
}
