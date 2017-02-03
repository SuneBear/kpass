package schema

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Share represents share info
type Share struct {
	EntryID uuid.UUID `json:"entryId"`
	Name    string    `json:"name"`
	Token   string    `json:"Token"`
	To      string    `json:"to"`
	TTL     int       `json:"ttl"`
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
		EntryID: s.EntryID,
		Name:    s.Name,
		To:      s.To,
		TTL:     s.TTL,
		Created: s.Created,
		Updated: s.Updated,
	}
}

// ShareResult represents desensitized share
type ShareResult struct {
	ID      uuid.UUID `json:"uuid"`
	EntryID uuid.UUID `json:"entryId"`
	Name    string    `json:"name"`
	To      string    `json:"to"`
	TTL     int       `json:"ttl"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
