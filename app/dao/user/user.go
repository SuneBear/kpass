package userDao

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/app/crypto"
	"github.com/seccom/kpass/app/dao"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// InitIndex ...
func InitIndex() {
	dao.DB.CreateIndex("user_by_created", dao.UserKey("*"), buntdb.IndexJSON("created"))
}

// CheckID ...
func CheckID(id string) error {
	if len(id) < 3 {
		return &gear.Error{Code: 400, Msg: fmt.Sprintf(`invalid user id "%s"`, id)}
	}
	return dao.DB.View(func(tx *buntdb.Tx) error {
		if _, e := tx.Get(dao.UserKey(id)); e == nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, id)}
		}
		return nil
	})
}

// CheckLogin ...
func CheckLogin(id, pass string) (user *dao.User, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		userKey := dao.UserKey(id)
		str, e := tx.Get(userKey)
		if e != nil {
			return &gear.Error{Code: 404, Msg: e.Error()}
		}
		user, e = dao.UserFrom(str)
		if e != nil {
			return e
		}
		if user.IsBlocked || user.Attempt > 5 {
			return &gear.Error{Code: 403, Msg: "too many login attempts"}
		}
		if !crypto.Global().ValidateUserPass(id, pass, user.Pass) {
			user.Attempt++
			tx.Set(userKey, user.String(), nil)
			return &gear.Error{Code: 400, Msg: "user id or password error"}
		}
		if user.Attempt > 0 {
			user.Attempt = 0
			tx.Set(userKey, user.String(), nil)
		}
		return nil
	})

	if err != nil {
		user = nil
	}
	return
}

// Create ...
func Create(id, pass string) (user *dao.User, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		userKey := dao.UserKey(id)
		_, e := tx.Get(userKey)
		if e == nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, id)}
		}

		user = &dao.User{
			ID:        id,
			Pass:      pass,
			IsBlocked: false,
			Entries:   []uuid.UUID{},
			Created:   time.Now(),
		}
		user.Updated = user.Created
		_, _, e = tx.Set(userKey, user.String(), nil)
		return e
	})

	if err != nil {
		user = nil
	}
	return
}

// Find ...
func Find(id string) (user *dao.User, err error) {
	err = dao.DB.View(func(tx *buntdb.Tx) error {
		res, e := tx.Get(dao.UserKey(id))
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

// Update ...
func Update(user *dao.User) error {
	return dao.DB.Update(func(tx *buntdb.Tx) error {
		user.Updated = time.Now()
		_, _, e := tx.Set(dao.UserKey(user.ID), user.String(), nil)
		return e
	})
}
