package schema

import (
	"encoding/json"
	"time"

	"github.com/seccom/kpass/pkg/util"
)

// Entry represents entry info
type Entry struct {
	TeamID    util.OID `json:"teamID"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Priority  int       `json:"priority"`
	IsDeleted bool      `json:"isDeleted"`
	Secrets   []string  `json:"secrets"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

// EntryFrom parse JSON string and returns a Entry intance.
func EntryFrom(str string) (*Entry, error) {
	entry := new(Entry)
	if err := json.Unmarshal([]byte(str), entry); err != nil {
		return nil, err
	}
	return entry, nil
}

// String returns JSON string with full entry info
func (e *Entry) String() string {
	return jsonMarshal(e)
}

// HasSecret returns whether the entry has the secret
func (e *Entry) HasSecret(secretID string) bool {
	return StringSlice(e.Secrets).Has(secretID)
}

// AddSecret adds the secret to the entry
func (e *Entry) AddSecret(secretID string) bool {
	ok := false
	e.Secrets, ok = StringSlice(e.Secrets).Add(secretID)
	return ok
}

// RemoveSecret removes the secret from the entry
func (e *Entry) RemoveSecret(secretID string) bool {
	ok := false
	e.Secrets, ok = StringSlice(e.Secrets).Remove(secretID)
	return ok
}

// Result returns EntryResult intance
func (e *Entry) Result(ID util.OID, secrets []*SecretResult, shares []*ShareResult) *EntryResult {
	if secrets == nil {
		secrets = []*SecretResult{}
	}
	if shares == nil {
		shares = []*ShareResult{}
	}
	return &EntryResult{
		ID:       ID,
		TeamID:   e.TeamID,
		Name:     e.Name,
		Category: e.Category,
		Priority: e.Priority,
		Secrets:  secrets,
		Shares:   shares,
		Created:  e.Created,
		Updated:  e.Updated,
	}
}

// Summary returns EntrySum intance
func (e *Entry) Summary(ID util.OID) *EntrySum {
	return &EntrySum{
		ID:       ID,
		TeamID:   e.TeamID,
		Name:     e.Name,
		Category: e.Category,
		Priority: e.Priority,
		Created:  e.Created,
		Updated:  e.Updated,
	}
}

// EntryResult represents desensitized entry
type EntryResult struct {
	ID       util.OID       `json:"id"`
	TeamID   util.OID       `json:"teamID"`
	Name     string          `json:"name"`
	Category string          `json:"category"`
	Priority int             `json:"priority"`
	Secrets  []*SecretResult `json:"secrets"`
	Shares   []*ShareResult  `json:"shares"`
	Created  time.Time       `json:"created"`
	Updated  time.Time       `json:"updated"`
}

// String returns JSON string with desensitized entry info
func (e *EntryResult) String() string {
	return jsonMarshal(e)
}

// EntrySum represents desensitized entry
type EntrySum struct {
	ID       util.OID `json:"id"`
	TeamID   util.OID `json:"teamID"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Priority int       `json:"priority"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// String returns JSON string with desensitized entry info
func (e *EntrySum) String() string {
	return jsonMarshal(e)
}
