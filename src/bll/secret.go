package bll

import (
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
)

// Secret is Business Logic Layer for secret
type Secret struct {
	*Bll
}

// Create ...
func (b *Secret) Create(userID, key string, EntryID util.OID, secret *schema.Secret) (*schema.SecretResult, error) {
	entry, err := b.Models.Entry.Find(EntryID, false)
	if err != nil {
		return nil, err
	}
	if key, err = b.Models.Team.GetKey(entry.TeamID, userID, key); err != nil {
		return nil, err
	}

	return b.Models.Secret.Create(EntryID, userID, key, secret)
}

// Update ...
func (b *Secret) Update(userID, key string, EntryID, SecretID util.OID, change map[string]interface{}) (*schema.SecretResult, error) {
	entry, err := b.Models.Entry.Find(EntryID, false)
	if err != nil {
		return nil, err
	}
	if key, err = b.Models.Team.GetKey(entry.TeamID, userID, key); err != nil {
		return nil, err
	}

	return b.Models.Secret.Update(EntryID, SecretID, userID, key, change)
}
