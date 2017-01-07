package dao

import (
	"time"

	"github.com/google/uuid"
)

const (
	// KeyPrefixUser ...
	keyPrefixUser   string = "U:"
	keyPrefixTeam   string = "T:"
	keyPrefixEntry  string = "E:"
	keyPrefixSecret string = "S:"

	keyInitDBToken string = "INIT_TOKEN"
	keyUserAdmin   string = "ADMIN"
)

func UserKey(name string) string {
	return keyPrefixUser + name
}

func TeamKey(id uuid.UUID) string {
	return keyPrefixTeam + id.String()
}

func EntryKey(id uuid.UUID) string {
	return keyPrefixEntry + id.String()
}

func SecretKey(id uuid.UUID) string {
	return keyPrefixSecret + id.String()
}

type User struct {
	Id        string      `json:"id"`
	Pass      string      `json:"pass,omitempty"` // encrypt
	IsBlocked bool        `json:"isBlocked,omitempty"`
	Entries   []uuid.UUID `json:"entries,omitempty"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

type Team struct {
	Id        uuid.UUID   `json:"uuid"`
	Name      string      `json:"name"`
	Token     string      `json:"token,omitempty"`
	IsBlocked bool        `json:"isBlocked,omitempty"`
	IsDeleted bool        `json:"isDeleted,omitempty"`
	OwnerId   uuid.UUID   `json:"userId"`
	Members   []string    `json:"members"`
	Entries   []uuid.UUID `json:"entries"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

type Entry struct {
	Id        uuid.UUID     `json:"uuid"`
	OwnerId   string        `json:"ownerId"`
	OwnerType string        `json:"ownerType"`
	Name      string        `json:"name"`
	Category  string        `json:"category"`
	Priority  int           `json:"priority"`
	Secrets   []interface{} `json:"secrets"`
	Shares    []interface{} `json:"shares"`
	IsDeleted bool          `json:"isDeleted,omitempty"`
	Created   time.Time     `json:"created"`
	Updated   time.Time     `json:"updated"`
}

type Secret struct {
	Id      uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Url     string    `json:"url"`
	Pass    string    `json:"password"`          // encrypt
	Note    string    `json:"note"`              // encrypt
	RawData string    `json:"rawData,omitempty"` // encrypt
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Share struct {
	Id      uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Salt    string    `json:"salt,omitempty"`
	Allow   []string  `json:"allow"`
	TTL     int       `json:"ttl"`
	Expire  time.Time `json:"expire"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
