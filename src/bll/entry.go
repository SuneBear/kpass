package bll

import (
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
)

// Entry is Business Logic Layer for entry
type Entry struct {
	*Bll
}

// Find ...
func (b *Entry) Find(userID, key string, EntryID util.OID) (*schema.EntryResult, error) {
	entry, err := b.Models.Entry.Find(EntryID, false)
	if err != nil {
		return nil, err
	}
	if key, err = b.Models.Team.GetKey(entry.TeamID, userID, key); err != nil {
		return nil, err
	}

	var secrets []*schema.SecretResult
	var files []*schema.FileResult
	var shares []*schema.ShareResult
	if len(entry.Secrets) > 0 {
		if secrets, err = b.Models.Secret.FindSecrets(key, entry.Secrets...); err != nil {
			return nil, err
		}
	}
	if len(entry.Files) > 0 {
		if files, err = b.Models.File.FindFiles(EntryID, key, entry.Files...); err != nil {
			return nil, err
		}
	}
	return entry.Result(EntryID, secrets, files, shares), nil
}

// FindByTeam ...
func (b *Entry) FindByTeam(userID string, TeamID util.OID) ([]*schema.EntrySum, error) {
	if err := b.Models.Team.CheckUser(TeamID, userID); err != nil {
		return nil, err
	}
	return b.Models.Entry.FindByTeam(TeamID, userID, false)
}

// Delete ...
func (b *Entry) Delete(userID string, EntryID util.OID, isDelete bool) (*schema.EntrySum, error) {
	entry, err := b.Models.Entry.Find(EntryID, !isDelete)
	if err != nil {
		return nil, err
	}
	if err = b.Models.Team.CheckUser(entry.TeamID, userID); err != nil {
		return nil, err
	}
	return b.Models.Entry.UpdateDeleted(userID, EntryID, isDelete)
}

// DeleteFile ...
func (b *Entry) DeleteFile(userID string, EntryID, FileID util.OID) error {
	if err := b.Models.Entry.RemoveFileByID(EntryID, FileID, userID); err != nil {
		return err
	}
	return b.Models.File.Delete(FileID)
}
