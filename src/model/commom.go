package model

import (
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
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

// All ....
type All struct {
	Entry  *Entry
	File   *File
	Secret *Secret
	Share  *Share
	Team   *Team
	User   *User
}

// Init ...
func (a *All) Init(db *service.DB) *All {
	a.Entry = new(Entry).Init(db)
	a.File = new(File).Init(db)
	a.Secret = new(Secret).Init(db)
	a.Share = new(Share).Init(db)
	a.Team = new(Team).Init(db)
	a.User = new(User).Init(db)
	return a
}
