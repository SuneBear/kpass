package dao

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

// User is database access oject for users
type User struct {
	db *service.DB
}

// NewUser return a User intance
func NewUser(db *service.DB) *User {
	return &User{db}
}

// CheckID ...
func (o *User) CheckID(id string) error {
	if len(id) < 3 {
		return &gear.Error{Code: 400, Msg: fmt.Sprintf(`invalid user id "%s"`, id)}
	}
	err := o.db.DB.View(func(tx *buntdb.Tx) error {
		if _, e := tx.Get(schema.UserKey(id)); e == nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, id)}
		}
		return nil
	})
	return dbError(err)
}

// CheckLogin ...
func (o *User) CheckLogin(id, pass string) (user *schema.User, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		userKey := schema.UserKey(id)
		value, e := tx.Get(userKey)
		if e != nil {
			return e
		}
		user, e = schema.UserFrom(value)
		if e != nil {
			return e
		}
		if user.IsBlocked || user.Attempt > 5 {
			return &gear.Error{Code: 403, Msg: "too many login attempts"}
		}
		if !auth.VerifyPass(id, pass, user.Pass) {
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
		return nil, dbError(err)
	}
	return
}

// Create ...
func (o *User) Create(userID, pass string) (user *schema.User, err error) {
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		userKey := schema.UserKey(userID)
		_, e := tx.Get(userKey)
		if e == nil {
			return &gear.Error{Code: 409, Msg: fmt.Sprintf(`user "%s" exists`, userID)}
		}

		user = &schema.User{
			ID:      userID,
			Pass:    auth.SignPass(userID, pass),
			Created: util.Time(time.Now()),
		}
		user.Updated = user.Created
		_, _, e = tx.Set(userKey, user.String(), nil)
		return e
	})

	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Find ...
func (o *User) Find(id string) (user *schema.User, err error) {
	err = o.db.DB.View(func(tx *buntdb.Tx) error {
		res, e := tx.Get(schema.UserKey(id))
		if e == nil {
			e = json.Unmarshal([]byte(res), user)
		}
		return e
	})
	if err != nil {
		user = nil
		err = dbError(err)
	}
	return
}

// Update ...
func (o *User) Update(user *schema.User) error {
	err := o.db.DB.Update(func(tx *buntdb.Tx) error {
		user.Updated = util.Time(time.Now())
		_, _, e := tx.Set(schema.UserKey(user.ID), user.String(), nil)
		return e
	})
	return dbError(err)
}
