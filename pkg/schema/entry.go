package schema

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Entry represents entry info
type Entry struct {
	OwnerID   string    `json:"ownerId"`
	OwnerType string    `json:"ownerType"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Priority  int       `json:"priority"`
	Secrets   []string  `json:"secrets"`
	Shares    []string  `json:"shares"`
	IsDeleted bool      `json:"isDeleted"`
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

// HasSecret returns whether the secret is in the Entry.Secrets
func (e *Entry) HasSecret(secretID string) bool {
	return StringSlice(e.Secrets).Has(secretID)
}

// RemoveSecret remove the secret from the Entry.Secrets
func (e *Entry) RemoveSecret(secretID string) bool {
	ok := false
	e.Secrets, ok = StringSlice(e.Secrets).Remove(secretID)
	return ok
}

// Result returns EntryResult intance
func (e *Entry) Result(ID uuid.UUID, secrets []*SecretResult, shares []*ShareResult) *EntryResult {
	if secrets == nil {
		secrets = []*SecretResult{}
	}
	if shares == nil {
		shares = []*ShareResult{}
	}
	return &EntryResult{
		ID:        ID,
		OwnerID:   e.OwnerID,
		OwnerType: e.OwnerType,
		Name:      e.Name,
		Category:  e.Category,
		Priority:  e.Priority,
		Secrets:   secrets,
		Shares:    shares,
		Created:   e.Created,
		Updated:   e.Updated,
	}
}

// Summary returns EntrySum intance
func (e *Entry) Summary(ID uuid.UUID) *EntrySum {
	return &EntrySum{
		ID:       ID,
		Name:     e.Name,
		Category: e.Category,
		Priority: e.Priority,
		Created:  e.Created,
		Updated:  e.Updated,
	}
}

// EntryResult represents desensitized entry
type EntryResult struct {
	ID        uuid.UUID       `json:"uuid"`
	OwnerID   string          `json:"ownerId"`
	OwnerType string          `json:"ownerType"`
	Name      string          `json:"name"`
	Category  string          `json:"category"`
	Priority  int             `json:"priority"`
	Secrets   []*SecretResult `json:"secrets"`
	Shares    []*ShareResult  `json:"shares"`
	Created   time.Time       `json:"created"`
	Updated   time.Time       `json:"updated"`
}

// String returns JSON string with desensitized entry info
func (e *EntryResult) String() string {
	return jsonMarshal(e)
}

// EntrySum represents desensitized entry
type EntrySum struct {
	ID       uuid.UUID `json:"uuid"`
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
