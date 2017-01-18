package dao

import (
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
