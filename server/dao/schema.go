package dao

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	// KeyPrefixUser ...
	keyPrefixUser   string = "U:"
	keyPrefixTeam   string = "T:"
	keyPrefixEntry  string = "E:"
	keyPrefixSecret string = "S:"
	keyPrefixShare  string = "SH:"

	keyDBSalt    string = "DB_SALT"
	keyUserAdmin string = "ADMIN"
)

// UserKey returns the user's db key
func UserKey(name string) string {
	return keyPrefixUser + name
}

// TeamKey returns the team's db key
func TeamKey(id uuid.UUID) string {
	return keyPrefixTeam + id.String()
}

// EntryKey returns the entry's db key
func EntryKey(id uuid.UUID) string {
	return keyPrefixEntry + id.String()
}

// SecretKey returns the secret's db key
func SecretKey(id uuid.UUID) string {
	return keyPrefixSecret + id.String()
}

// ShareKey returns the share's db key
func ShareKey(id uuid.UUID) string {
	return keyPrefixShare + id.String()
}

// User represents user info
type User struct {
	ID        string      `json:"id"`
	Pass      string      `json:"pass"` // encrypt
	IsBlocked bool        `json:"isBlocked"`
	Entries   []uuid.UUID `json:"entries"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

// String returns JSON string with full user info
func (u *User) String() string {
	return jsonMarshal(u)
}

// Result returns UserResult intance
func (u *User) Result() *UserResult {
	return &UserResult{
		ID:      u.ID,
		Created: u.Created,
		Updated: u.Updated,
	}
}

// UserResult represents desensitized user
type UserResult struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// String returns JSON string with desensitized user info
func (u *UserResult) String() string {
	return jsonMarshal(u)
}

// Team represents team info
type Team struct {
	ID        uuid.UUID   `json:"uuid"`
	Name      string      `json:"name"`
	Token     string      `json:"token"`
	IsBlocked bool        `json:"isBlocked"`
	IsDeleted bool        `json:"isDeleted"`
	OwnerID   uuid.UUID   `json:"userId"`
	Members   []string    `json:"members"`
	Entries   []uuid.UUID `json:"entries"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

// String returns JSON string with full team info
func (t *Team) String() string {
	return jsonMarshal(t)
}

// Result returns TeamResult intance
func (t *Team) Result() *TeamResult {
	return &TeamResult{
		ID:      t.ID,
		Name:    t.Name,
		OwnerID: t.OwnerID,
		Members: t.Members,
		Created: t.Created,
		Updated: t.Updated,
	}
}

// TeamResult represents desensitized team
type TeamResult struct {
	ID      uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	OwnerID uuid.UUID `json:"userId"`
	Members []string  `json:"members"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// String returns JSON string with desensitized team info
func (t *TeamResult) String() string {
	return jsonMarshal(t)
}

// Entry represents entry info
type Entry struct {
	ID        uuid.UUID   `json:"uuid"`
	OwnerID   string      `json:"ownerId"`
	OwnerType string      `json:"ownerType"`
	Name      string      `json:"name"`
	Category  string      `json:"category"`
	Priority  int         `json:"priority"`
	Secrets   []uuid.UUID `json:"secrets"`
	Shares    []uuid.UUID `json:"shares"`
	IsDeleted bool        `json:"isDeleted"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

// String returns JSON string with full entry info
func (e *Entry) String() string {
	return jsonMarshal(e)
}

// Result returns EntryResult intance
func (e *Entry) Result(secrets []Secret, shares []ShareResult) *EntryResult {
	return &EntryResult{
		ID:        e.ID,
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

// EntryResult represents desensitized entry
type EntryResult struct {
	ID        uuid.UUID     `json:"uuid"`
	OwnerID   string        `json:"ownerId"`
	OwnerType string        `json:"ownerType"`
	Name      string        `json:"name"`
	Category  string        `json:"category"`
	Priority  int           `json:"priority"`
	Secrets   []Secret      `json:"secrets"`
	Shares    []ShareResult `json:"shares"`
	Created   time.Time     `json:"created"`
	Updated   time.Time     `json:"updated"`
}

// String returns JSON string with desensitized entry info
func (e *EntryResult) String() string {
	return jsonMarshal(e)
}

// Secret represents secret info
type Secret struct {
	ID      uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	URL     string    `json:"url"`
	Pass    string    `json:"password"` // encrypt
	Note    string    `json:"note"`     // encrypt
	RawData string    `json:"rawData"`  // encrypt
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// String returns JSON string with full secret info
func (s *Secret) String() string {
	return jsonMarshal(s)
}

// Share represents share info
type Share struct {
	ID      uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Salt    string    `json:"salt"`
	Members []string  `json:"members"`
	TTL     int       `json:"ttl"`
	Expire  time.Time `json:"expire"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// String returns JSON string with full share info
func (s *Share) String() string {
	return jsonMarshal(s)
}

// Result returns ShareResult intance
func (s *Share) Result() *ShareResult {
	return &ShareResult{
		ID:      s.ID,
		Name:    s.Name,
		Members: s.Members,
		TTL:     s.TTL,
		Expire:  s.Expire,
		Created: s.Created,
		Updated: s.Updated,
	}
}

// ShareResult represents desensitized share
type ShareResult struct {
	ID      uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Members []string  `json:"members"`
	TTL     int       `json:"ttl"`
	Expire  time.Time `json:"expire"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func jsonMarshal(v interface{}) (str string) {
	if res, err := json.Marshal(v); err == nil {
		str = string(res)
	}
	return
}
