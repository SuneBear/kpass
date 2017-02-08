package schema

import (
	"encoding/json"
	"time"
)

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
	Avatar  string    `json:"avatar"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// String returns JSON string with desensitized user info
func (u *UserResult) String() string {
	return jsonMarshal(u)
}
