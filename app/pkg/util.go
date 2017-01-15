package pkg

import (
	"github.com/google/uuid"
	"github.com/seccom/kpass/app/crypto"
)

// IsUUID ...
func IsUUID(id string) bool {
	if _, err := uuid.Parse(id); err == nil {
		return true
	}
	return false
}

// NewUUID return a UUID with given space
func NewUUID(space string) uuid.UUID {
	id, err := uuid.Parse(space)
	if err != nil {
		id = uuid.NewSHA1(uuid.NameSpaceOID, []byte(space))
	}
	return uuid.NewSHA1(id, crypto.RandBytes(16))
}
