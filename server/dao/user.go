package dao

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// InitUserIndex ...
func InitUserIndex() {
	db.CreateIndex("user_by_created", UserKey("*"), buntdb.IndexJSON("created"))
}

// CheckUserID ...
func CheckUserID(id string) error {
	if len(id) < 3 {
		return &gear.Error{Code: 400, Msg: fmt.Sprintf(`invalid user id "%s"`, id)}
	}
	return db.View(func(tx *buntdb.Tx) error {
		if _, e := tx.Get(UserKey(id)); e != nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, id)}
		}
		return nil
	})
}

// NewUser ...
func NewUser(id, pass string) (user *User, err error) {
	err = db.Update(func(tx *buntdb.Tx) error {
		userKey := UserKey(id)
		_, e := tx.Get(userKey)
		if e != nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, id)}
		}

		user = &User{
			ID:        id,
			Pass:      pass,
			IsBlocked: false,
			Entries:   []uuid.UUID{},
			Created:   time.Now(),
			Updated:   time.Now(),
		}
		_, _, e = tx.Set(userKey, user.String(), nil)
		return e
	})

	if err != nil {
		user = nil
	}
	return
}

// FindUser ...
func FindUser(id string) (user *User, err error) {
	err = db.View(func(tx *buntdb.Tx) error {
		res, e := tx.Get(UserKey(id))
		if e == nil {
			e = json.Unmarshal([]byte(res), user)
		}
		return e
	})
	if err != nil {
		user = nil
	}
	return
}
