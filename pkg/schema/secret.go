package schema

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

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
