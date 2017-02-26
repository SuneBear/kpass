package schema

import (
	"encoding/json"
	"time"

	"github.com/seccom/kpass/src/util"
)

// Team represents team info
type Team struct {
	Logo       util.OID  `json:"logo"`
	UserID     string    `json:"userID"`
	Name       string    `json:"name"`
	Pass       string    `json:"pass"`
	IsFrozen   bool      `json:"isFrozen"` // if true, member can't modify team's entires
	IsDeleted  bool      `json:"isDeleted"`
	Visibility string    `json:"visibility"` // enum: "private", "member"
	Members    []string  `json:"members"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
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
func (t *Team) Result(ID util.OID, Members []*UserResult) *TeamResult {
	return &TeamResult{
		ID:         ID,
		Logo:       DownloadURL(t.Logo, "team", ID.String(), ""),
		UserID:     t.UserID,
		Name:       t.Name,
		IsFrozen:   t.IsFrozen,
		Visibility: t.Visibility,
		Members:    Members,
		Created:    t.Created,
		Updated:    t.Updated,
	}
}

// TeamResult represents desensitized team
type TeamResult struct {
	ID         util.OID      `json:"id"`
	Logo       string        `json:"logo"`
	UserID     string        `json:"userID"`
	Name       string        `json:"name"`
	IsFrozen   bool          `json:"isFrozen"`
	Visibility string        `json:"visibility"`
	Members    []*UserResult `json:"members"`
	Created    time.Time     `json:"created"`
	Updated    time.Time     `json:"updated"`
}

// String returns JSON string with desensitized team info
func (t *TeamResult) String() string {
	return jsonMarshal(t)
}
