package model

import (
	"github.com/seccom/kpass/src/schema"
	"github.com/teambition/gear"
	"github.com/tidwall/buntdb"
)

func dbError(err error) error {
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

// IdsToUsers ...
func IdsToUsers(tx *buntdb.Tx, ids []string) (users []*schema.UserResult) {
	for _, id := range ids {
		if res, e := tx.Get(schema.UserKey(id)); e == nil {
			if user, e := schema.UserFrom(res); e == nil {
				users = append(users, user.Result())
			}
		}
	}
	return
}
