package schema

import (
	"encoding/json"
	"time"

	"github.com/seccom/kpass/src/util"
)

// Share represents share info
type Share struct {
	TeamID  util.OID  `json:"teamID"`
	EntryID util.OID  `json:"entryID"`
	Name    string    `json:"name"`
	Token   string    `json:"Token"`
	UserID  string    `json:"userID"`
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
func (s *Share) Result(ID util.OID) *ShareResult {
	return &ShareResult{
		ID:      ID,
		TeamID:  s.TeamID,
		EntryID: s.EntryID,
		Name:    s.Name,
		UserID:  s.UserID,
		Created: s.Created,
		Updated: s.Updated,
	}
}

// ShareResult represents desensitized share
type ShareResult struct {
	ID      util.OID  `json:"id"`
	TeamID  util.OID  `json:"teamID"`
	EntryID util.OID  `json:"entryID"`
	UserID  string    `json:"userID"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
