package userDao

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/seccom/kpass/app/dao"
	"github.com/seccom/kpass/app/pkg"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// CheckID ...
func CheckID(id string) error {
	if len(id) < 3 {
		return &gear.Error{Code: 400, Msg: fmt.Sprintf(`invalid user id "%s"`, id)}
	}
	err := dao.DB.View(func(tx *buntdb.Tx) error {
		if _, e := tx.Get(dao.UserKey(id)); e == nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, id)}
		}
		return nil
	})
	return dao.DBError(err)
}

// CheckLogin ...
func CheckLogin(id, pass string) (user *dao.User, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		userKey := dao.UserKey(id)
		value, e := tx.Get(userKey)
		if e != nil {
			return e
		}
		user, e = dao.UserFrom(value)
		if e != nil {
			return e
		}
		if user.IsBlocked || user.Attempt > 5 {
			return &gear.Error{Code: 403, Msg: "too many login attempts"}
		}
		if !pkg.Auth.ValidateUserPass(id, pass, user.Pass) {
			user.Attempt++
			tx.Set(userKey, user.String(), nil)
			tx.Commit()
			return &gear.Error{Code: 400, Msg: "user id or password error"}
		}
		if user.Attempt > 0 {
			user.Attempt = 0
			tx.Set(userKey, user.String(), nil)
		}
		return nil
	})

	if err != nil {
		return nil, dao.DBError(err)
	}
	return
}

// Create ...
func Create(userID, pass string) (user *dao.User, err error) {
	err = dao.DB.Update(func(tx *buntdb.Tx) error {
		userKey := dao.UserKey(userID)
		_, e := tx.Get(userKey)
		if e == nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, userID)}
		}

		user = &dao.User{
			ID:        userID,
			Pass:      pkg.Auth.EncryptUserPass(userID, pass),
			IsBlocked: false,
			Created:   time.Now(),
		}
		user.Updated = user.Created
		_, _, e = tx.Set(userKey, user.String(), nil)
		return e
	})

	if err != nil {
		return nil, dao.DBError(err)
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
		err = dao.DBError(err)
	}
	return
}

// Update ...
func Update(user *dao.User) error {
	err := dao.DB.Update(func(tx *buntdb.Tx) error {
		user.Updated = time.Now()
		_, _, e := tx.Set(dao.UserKey(user.ID), user.String(), nil)
		return e
	})
	return dao.DBError(err)
}
