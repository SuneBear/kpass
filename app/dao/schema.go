package dao

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	// KeyPrefixUser ...
	keyPrefixUser   = "U:"
	keyPrefixTeam   = "T:"
	keyPrefixEntry  = "E:"
	keyPrefixSecret = "S:"
	keyPrefixShare  = "SH:"

	keyDBSalt = "DB_SALT"
)

// UserKey returns the user's db key
func UserKey(name string) string {
	return keyPrefixUser + name
}

// TeamKey returns the team's db key
func TeamKey(id string) string {
	return keyPrefixTeam + id
}

// TeamIDFromKey returns team' ID from key
func TeamIDFromKey(key string) uuid.UUID {
	val := key[len(keyPrefixTeam):]
	id, err := uuid.Parse(val)
	if err != nil {
		panic(err)
	}
	return id
}

// EntryKey returns the entry's db key
func EntryKey(id string) string {
	return keyPrefixEntry + id
}

// EntryIDFromKey returns entry' ID from key
func EntryIDFromKey(key string) uuid.UUID {
	val := key[len(keyPrefixEntry):]
	id, err := uuid.Parse(val)
	if err != nil {
		panic(err)
	}
	return id
}

// SecretKey returns the secret's db key
func SecretKey(id string) string {
	return keyPrefixSecret + id
}

// ShareKey returns the share's db key
func ShareKey(id string) string {
	return keyPrefixShare + id
}

// ShareIDFromKey returns share' ID from key
func ShareIDFromKey(key string) uuid.UUID {
	val := key[len(keyPrefixShare):]
	id, err := uuid.Parse(val)
	if err != nil {
		panic(err)
	}
	return id
}

// User represents user info
type User struct {
	ID        string    `json:"id"`
	Pass      string    `json:"pass"` // encrypt
	IsBlocked bool      `json:"isBlocked"`
	Attempt   int       `json:"attempt"` // login attempts
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
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
	Name      string    `json:"name"`
	Pass      string    `json:"pass"`
	IsFrozen  bool      `json:"isFrozen"`
	IsDeleted bool      `json:"isDeleted"`
	OwnerID   string    `json:"ownerID"`
	Members   []string  `json:"members"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
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

// HasMember returns whether the user is team's member
func (t *Team) HasMember(userID string) bool {
	return StringSlice(t.Members).Has(userID)
}

// AddMember adds the user to team's members
func (t *Team) AddMember(userID string) bool {
	ok := false
	t.Members, ok = StringSlice(t.Members).Add(userID)
	return ok
}

// RemoveMember removes the user from team's members
func (t *Team) RemoveMember(userID string) bool {
	ok := false
	t.Members, ok = StringSlice(t.Members).Remove(userID)
	return ok
}

// Result returns TeamResult intance
func (t *Team) Result(ID uuid.UUID) *TeamResult {
	return &TeamResult{
		ID:       ID,
		Name:     t.Name,
		IsFrozen: t.IsFrozen,
		OwnerID:  t.OwnerID,
		Members:  t.Members,
		Created:  t.Created,
		Updated:  t.Updated,
	}
}

// TeamResult represents desensitized team
type TeamResult struct {
	ID       uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	IsFrozen bool      `json:"isFrozen"`
	OwnerID  string    `json:"ownerID"`
	Members  []string  `json:"members"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// String returns JSON string with desensitized team info
func (t *TeamResult) String() string {
	return jsonMarshal(t)
}

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
	Name    string    `json:"name"`
	Salt    string    `json:"salt"`
	ToUser  string    `json:"toUser"`
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
func (s *Share) Result(ID uuid.UUID) *ShareResult {
	return &ShareResult{
		ID:      ID,
		Name:    s.Name,
		ToUser:  s.ToUser,
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
	ToUser  string    `json:"toUser"`
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

// StringSlice ...
type StringSlice []string

// Has returns whether the str is in the slice.
func (s StringSlice) Has(str string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == str {
			return true
		}
	}
	return false
}

// Add adds the str to the slice.
func (s StringSlice) Add(str string) ([]string, bool) {
	if s.Has(str) {
		return s, false
	}
	return append(s, str), true
}

// Remove remove the str from the slice.
func (s StringSlice) Remove(str string) ([]string, bool) {
	offset := 0
	for i := 0; i < len(s); i++ {
		if s[i] != str {
			s[offset] = s[i]
			offset++
		}
	}
	if offset < len(s) {
		return s[0:offset], true
	}
	return s, false
}
