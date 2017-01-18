package schema

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

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
