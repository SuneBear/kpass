package bll

import (
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
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
	if err = b.Models.Team.CheckMember(entry.TeamID, userID, true); err != nil {
		return nil, err
	}
	if key, err = b.Models.Team.GetKey(entry.TeamID, userID, key); err != nil {
		return nil, err
	}

	return b.Models.Secret.Create(EntryID, userID, key, entry, secret)
}

// Update ...
func (b *Secret) Update(userID, key string, EntryID, SecretID util.OID, change map[string]interface{}) (*schema.SecretResult, error) {
	entry, err := b.Models.Entry.Find(EntryID, false)
	if err != nil {
		return nil, err
	}
	if !entry.HasSecret(SecretID.String()) {
		return nil, &gear.Error{Code: 403, Msg: "secret not found in the entry"}
	}
	if err = b.Models.Team.CheckMember(entry.TeamID, userID, true); err != nil {
		return nil, err
	}
	if key, err = b.Models.Team.GetKey(entry.TeamID, userID, key); err != nil {
		return nil, err
	}

	return b.Models.Secret.Update(EntryID, SecretID, userID, key, change)
}

// Delete ...
func (b *Secret) Delete(EntryID, SecretID util.OID, userID string) error {
	entry, err := b.Models.Entry.Find(EntryID, false)
	if err != nil {
		return err
	}
	if err = b.Models.Team.CheckMember(entry.TeamID, userID, true); err != nil {
		return err
	}
	if !entry.RemoveSecret(SecretID.String()) {
		return &gear.Error{Code: 404, Msg: "secret not found in the entry"}
	}
	return b.Models.Secret.Delete(EntryID, SecretID, userID, entry)
}
