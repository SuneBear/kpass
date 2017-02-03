package dao

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/service"
	"github.com/seccom/kpass/pkg/util"
	"github.com/tidwall/buntdb"
)

// Share is database access oject for share
type Share struct {
	db *service.DB
}

// NewShare return a Share intance
func NewShare(db *service.DB) *Share {
	return &Share{db}
}

// Create ...
func (o *Share) Create(EntryID uuid.UUID, key, userID, pass, name string, ttl int, expire time.Duration) (
	shareResult *schema.ShareResult, err error) {
	ShareID := util.NewUUID(EntryID.String())
	token, err := auth.EncryptText(auth.SignPass(userID, pass), key)
	if err != nil {
		return nil, dbError(err)
	}
	share := &schema.Share{
		EntryID: EntryID,
		Name:    name,
		Token:   token,
		To:      userID,
		TTL:     ttl,
		Created: time.Now(),
	}
	share.Updated = share.Created
	shareResult = share.Result(ShareID)
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
		_, _, e := tx.Set(schema.ShareKey(ShareID.String()), share.String(), &buntdb.SetOptions{
			Expires: true,
			TTL:     expire,
		})
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// Find ...
func (o *Share) Find(ShareID uuid.UUID) (share *schema.Share, err error) {
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		var res string
		if res, e = tx.Get(schema.ShareKey(ShareID.String())); e == nil {
			share, e = schema.ShareFrom(res)
		}
		return e
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// FindByUserID ...
func (o *Share) FindByUserID(userID string) (shares []*schema.ShareResult, err error) {
	shares = make([]*schema.ShareResult, 0)
	cond := fmt.Sprintf(`{"to":"%s"}`, userID)
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendGreaterOrEqual("share_by_user", cond, func(key, value string) bool {
			share, e := schema.ShareFrom(value)
			if e != nil {
				e = fmt.Errorf("invalid share: %s, %s", key, value)
				return false
			}
			if share.To != userID {
				return false
			}
			ShareID := schema.ShareIDFromKey(key)
			shares = append(shares, share.Result(ShareID))
			return true
		})
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}

// FindByEntryID ...
func (o *Share) FindByEntryID(EntryID uuid.UUID) (shares []*schema.ShareResult, err error) {
	shares = make([]*schema.ShareResult, 0)
	conds := fmt.Sprintf(`{"entryId":"%s"}`, EntryID.String())
	err = o.db.DB.View(func(tx *buntdb.Tx) (e error) {
		tx.AscendGreaterOrEqual("share_by_entry", conds, func(key, value string) bool {
			share, e := schema.ShareFrom(value)
			if e != nil {
				e = fmt.Errorf("invalid share: %s, %s", key, value)
				return false
			}
			if share.EntryID.String() != EntryID.String() {
				return false
			}
			ShareID := schema.ShareIDFromKey(key)
			shares = append(shares, share.Result(ShareID))
			return true
		})
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}
