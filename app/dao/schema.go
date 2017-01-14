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
func TeamKey(id string) string {
	return keyPrefixTeam + id
}

// EntryKey returns the entry's db key
func EntryKey(id string) string {
	return keyPrefixEntry + id
}

// SecretKey returns the secret's db key
func SecretKey(id string) string {
	return keyPrefixSecret + id
}

// ShareKey returns the share's db key
func ShareKey(id string) string {
	return keyPrefixShare + id
}

// User represents user info
type User struct {
	ID        string      `json:"id"`
	Pass      string      `json:"pass"` // encrypt
	IsBlocked bool        `json:"isBlocked"`
	Attempt   int         `json:"attempt"` // login attempts
	Entries   []uuid.UUID `json:"entries"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

// UserFrom parse JSON string and returns a User intance.
func UserFrom(str string) (*User, error) {
	user := new(User)
	if err := json.Unmarshal([]byte(str), user); err != nil {
		return nil, err
	}
	return user, nil
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
	Pass      string      `json:"pass"`
	IsBlocked bool        `json:"isBlocked"`
	IsDeleted bool        `json:"isDeleted"`
	OwnerID   uuid.UUID   `json:"userId"`
	Members   []string    `json:"members"`
	Entries   []uuid.UUID `json:"entries"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

// TeamFrom parse JSON string and returns a Team intance.
func TeamFrom(str string) (*Team, error) {
	team := new(Team)
	if err := json.Unmarshal([]byte(str), team); err != nil {
		return nil, err
	}
	return team, nil
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

// Result returns EntryResult intance
func (e *Entry) Result(secrets []*SecretResult, shares []*ShareResult) *EntryResult {
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

// Summary returns EntrySum intance
func (e *Entry) Summary() *EntrySum {
	return &EntrySum{
		ID:       e.ID,
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

// Secret represents secret info
type Secret struct {
	Name    string    `json:"name"`
	URL     string    `json:"url"`
	Pass    string    `json:"password"`
	Note    string    `json:"note"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// SecretFrom parse JSON string and returns a Secret intance.
func SecretFrom(str string) (*Secret, error) {
	secret := new(Secret)
	if err := json.Unmarshal([]byte(str), secret); err != nil {
		return nil, err
	}
	return secret, nil
}

// String returns JSON string with full secret info
func (s *Secret) String() string {
	return jsonMarshal(s)
}

// Result returns EntryResult intance
func (s *Secret) Result(id uuid.UUID) *SecretResult {
	return &SecretResult{
		ID:      id,
		Name:    s.Name,
		URL:     s.URL,
		Pass:    s.Pass,
		Note:    s.Note,
		Created: s.Created,
		Updated: s.Updated,
	}
}

// SecretResult represents secret info
type SecretResult struct {
	ID      uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	URL     string    `json:"url"`
	Pass    string    `json:"password"`
	Note    string    `json:"note"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// String returns JSON string with full secret info
func (s *SecretResult) String() string {
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

// ShareFrom parse JSON string and returns a Share intance.
func ShareFrom(str string) (*Share, error) {
	share := new(Share)
	if err := json.Unmarshal([]byte(str), share); err != nil {
		return nil, err
	}
	return share, nil
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
